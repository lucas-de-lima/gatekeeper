# 23 — Identidade do workspace e modos de trabalho

**Autoridade do produto:** Lucas de Lima  
**Estado:** contrato operacional obrigatório para qualquer agente que trabalhe no Harness  
**Propósito:** impedir que uma sessão de construção do Harness seja confundida com uma sessão na qual o Harness desenvolve um sistema de usuário.

## 1. Regra central

Até que o Harness cumpra o critério de prontidão deste pacote e Lucas registre uma decisão de release, o objetivo ativo do repositório-fonte é **construir, integrar, validar e preparar o próprio Harness**. O Harness ainda não existe como produto operacional confiável; portanto, seus futuros agentes, pipelines e políticas não podem ser tratados como ferramentas já disponíveis para construir o Harness.

Um agente de desenvolvimento chamado para trabalhar neste código está atuando como **agente implementador do Harness**, orientado por este plano. Não está executando o fluxo de projeto-alvo do Harness, mesmo que a tarefa envolva arquitetura, backlog, código, testes ou integrações.

O Harness poderá ser usado para desenvolver outros sistemas somente após o gate de prontidão e uma invocação explícita em um workspace-alvo. Até então, nenhuma sessão deve assumir que o Harness está “pronto” porque existe uma documentação ou um executável parcial.

## 2. Identidades de workspace

Cada workspace recebe uma identidade explícita, persistida em manifesto no seu root. Não inferir apenas pelo nome de uma pasta.

| `workspace_kind` | Conteúdo | Identificador recomendado | Agente está construindo o quê? |
|---|---|---|---|
| `blueprint_package` | Esta especificação e seus manifestos/documentos | `engineering-harness-blueprint` | Documentação/especificação do Harness |
| `harness_source` | Repositório executável do Harness, seus adapters, skills e ferramentas de runtime | `portable-engineering-harness` | O produto Harness |
| `harness_fixture` | Projeto sintético/sanitizado para testar o Harness | ID de fixture vinculado ao plano de avaliação | Fixture de validação; não é produto privado do Lucas |
| `managed_target` | Sistema novo ou legado no qual o Harness foi instalado para uma demanda real | `project_id` do projeto-alvo | O sistema-alvo, somente se modo/solicitação confirmarem isso |
| `portfolio_docs` | Repositório central de documentação que agrega sistemas de um portfólio | ID do portfólio e IDs vinculados | Memória/artefatos documentais dos sistemas relacionados |

Não confundir `harness_source` com `managed_target`: ambos podem conter configuração do Harness, mas têm objetivos, arquivos de controle e gates diferentes.

## 3. Modos de trabalho

`workspace_kind` identifica o tipo de repositório; `work_mode` identifica a intenção da invocação atual. São campos separados.

| `work_mode` | Onde é válido | Objetivo |
|---|---|---|
| `blueprint_maintenance` | `blueprint_package` | Avaliar ou alterar documentação do produto sem criar código executável |
| `harness_build` | `harness_source` | Construir features, providers, agents, pipelines, instaladores e distribuição do Harness |
| `harness_verification` | `harness_source` ou `harness_fixture` | Validar comportamento/aceite do Harness em fixture controlada; correções precisam ser explicitamente incluídas no run |
| `harness_release` | `harness_source` | Executar checklist e gerar release candidata após critérios de prontidão |
| `harness_installation` | `managed_target` | Instalar/configurar o Harness no projeto, analisar estrutura e persistir estado |
| `target_discovery` | `managed_target` | Compreender sistema novo/legado, gerar baseline e aguardar solicitação quando ausente |
| `target_delivery` | `managed_target` | Resolver uma solicitação de produto/software usando o Harness já instalado e liberado |
| `portfolio_maintenance` | `portfolio_docs` | Atualizar índices e vínculos entre documentação e repositórios de um portfólio |

Modos podem ter runs filhos explícitos, mas uma sessão não migra implicitamente de `harness_build` para `target_delivery` nem de `harness_verification` para execução de trabalho em sistema real.

## 4. Manifestos de identidade e estado

O repositório-fonte do Harness deve conter um manifesto próprio, distinto do manifesto de pacote `.harness/` gerado em sistemas-alvo:

```yaml
schema_version: 1
workspace_kind: harness_source
product_id: portable-engineering-harness
product_status: under_construction
work_mode_default: harness_build
current_release: null
active_build_plan: 24_PLANO_MESTRE_DE_CONSTRUCAO_DO_HARNESS.md
active_work_item: null
readiness_decision: pending
last_updated_at: "RFC3339"
```

Campos de estado mudam por evento/versionamento. `product_status` só pode transicionar a `release_candidate` ou `ready_for_target_use` com os gates do documento 24 registrados. A existência de configuração de agentes, MCP ou executável não altera o status por si só.

Uma instância em projeto-alvo terá manifesto separado:

```yaml
schema_version: 1
workspace_kind: managed_target
project_id: project-id
harness_version: pinned-release
installation_state: installed | analysis_in_progress | needs_input
target_work_mode: target_discovery | target_delivery
architecture_baseline_ref: null
active_request_id: null
```

O schema é ilustrativo e deve ser formalizado/validado na implementação. O Harness não escreve um manifesto `managed_target` no próprio código-fonte para encobrir a identidade `harness_source`.

## 5. Identificação de uma nova invocação

O agente começa identificando o workspace e o objetivo nesta ordem:

1. Ler instruções aplicáveis `AGENTS.md` desde a raiz efetiva até a pasta de trabalho.
2. Ler manifesto de identidade/estado do workspace e localizar o plano ativo.
3. Ler solicitação atual do usuário; comando explícito e atual prevalece sobre o modo default, salvo conflito com autoridade/gates.
4. Conferir Git root/remotes, estrutura e arquivos-canônicos para validar que o manifesto corresponde ao diretório.
5. Comparar o modo pedido, a identidade persistida, o estado de release e o estado do plano.
6. Registrar um `run_id` com `workspace_kind`, `work_mode`, `plan_ref`/`request_id`, fase e objetivo antes de escrever.

Se prompt e marcadores divergem, não escolher silenciosamente. Reconciliar por evidências e, se a intenção continuar ambígua, perguntar: “Você quer continuar construindo/validando o Harness ou executar o Harness para trabalhar em um sistema-alvo?” Não iniciar implementação de produto-alvo com base em nome de arquivo ou inferência.

## 6. Roteamento de trabalho durante construção

Em `harness_build`, o agente deve:

- consultar `24_PLANO_MESTRE_DE_CONSTRUCAO_DO_HARNESS.md` e o estado atual do item/fase;
- selecionar o próximo item `ready`, respeitando dependências e critério de saída;
- usar a especificação da capacidade e integração relevante (`15`–`23`);
- registrar a alteração como implementação do Harness, incluindo link a work item e evidências;
- manter `workspace_kind=harness_source` e `work_mode=harness_build` durante todo o run;
- entregar status, item concluído, evidência, bloqueios e próximo item pronto.

O agente não precisa inventar uma pipeline de software do Harness para si mesmo; a pipeline futura do produto está sendo implementada. Use o contrato do executor atual e o plano mestre. Se o Harness parcialmente construído já for utilizado em um teste de auto-hospedagem, executar em `harness_verification`, workspace/branch isolado, dados sintéticos, sem permitir que o produto escreva sua própria autoridade ou marque seu próprio aceite.

## 7. Roteamento para produto-alvo após prontidão

Em `target_delivery`, exigir todos estes sinais:

1. Workspace é `managed_target`, identificado pelo repositório/manifesto do projeto.
2. Harness está numa versão liberada e compatível, com capability matrix válida no executor atual.
3. Manifesto da instância confirma instalação/análise concluída ou operação de retomada identificável.
4. Existe solicitação de produto/software ativa ou uma ação explícita de instalação/descoberta.
5. Arquitetura e gates são lidos a partir da memória canônica do projeto-alvo.

Sem solicitação, a instalação/análise pode criar baseline de legado, registrar evidências e terminar em `waiting_for_request`; não inventar feature. Instalação e operação não são iguais a desenvolvimento do Harness.

## 8. Códigos de bloqueio e exceções

- `workspace_identity_conflict`: pasta aberta contradiz manifesto/remotes/solicitação.
- `work_mode_ambiguous`: não é possível distinguir construção do Harness de entrega ao produto-alvo.
- `build_plan_missing`: o código-fonte do Harness não tem plano/estado suficiente para escolher próximo item.
- `release_gate_not_met`: alguém tentou usar Harness em um projeto-alvo antes da decisão de prontidão.
- `fixture_to_target_boundary`: uma fixture de validação foi confundida com um produto real.

Bloqueio não equivale a encerrar a tarefa inteira. O agente pode continuar inspeção/documentação segura e reportar evidências; não pode executar trabalho no domínio errado.

## 9. Texto obrigatório nas instruções de construção

O `AGENTS.md` do repositório-fonte e o prompt de invocação devem informar claramente:

> Este repositório é o código-fonte do Portable Engineering Harness. Enquanto o estado de prontidão não for aprovado e registrado, o trabalho aqui constrói e valida o Harness; não desenvolve um sistema-alvo. Use o plano mestre, registre cada mudança num item e num run, e não marque o Harness como pronto sem cumprir o gate completo. Para desenvolver outro sistema, abra um workspace-alvo e use uma versão liberada do Harness.

## 10. Critério de aceite do contrato

Uma validação do sistema deve apresentar quatro invocações controladas e resultados corretos: manutenção do blueprint; construção do Harness; instalação/análise de legado sem demanda (após release); entrega de uma feature sintética (após release). Cada run identifica `workspace_kind` e `work_mode`, seleciona artefatos corretos, não cruza dados/objetivos e termina com a próxima ação adequada.
