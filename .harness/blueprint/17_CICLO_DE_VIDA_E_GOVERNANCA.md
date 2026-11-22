# 17 — Ciclo de vida do projeto e governança de mudanças

**Responsável pelo produto:** Lucas de Lima  
**Estado:** decisões centrais confirmadas; máquina de estados proposta para orientar implementação  
**Termos:** GMUD é auditoria/autorização pós-homologação quando exigida pelo cliente; não é estratégia de branches.

## 1. Princípios

- Instalar e analisar um projeto é diferente de iniciar uma solicitação de mudança.
- A primeira solicitação do agente mestre pode ser a instalação/análise inicial, ou o Harness já pode estar instalado e receber uma solicitação de trabalho.
- Um toggle simples pode ajudar a indicar o modo, mas artefatos, manifestos, timestamps e evidências persistidos são a fonte de reconciliação. Toggle nunca substitui a leitura do estado.
- Arquitetura aprovada é baseline protegida. Agentes de camadas inferiores não a alteram por conta própria.
- O Harness entrega até a prontidão/homologação e, se necessário, o pacote para GMUD. O deploy de produção fica fora do escopo atual.
- Estado parcial precisa sobreviver ao fim da sessão; repetição deve retomar ou reconciliar, não apagar.

## 2. Modos de entrada

| Modo | Intenção | Comportamento inicial |
|---|---|---|
| `installation_analysis` | Instalar o Harness e compreender o projeto | Descoberta AI-first, classificar novo/legado, gerar estado inicial, registrar achados e próximos passos |
| `project_request` | Usar Harness instalado para uma demanda | Conferir instalação, estado, arquitetura e backlog antes de selecionar pipeline |
| `resume` | Continuar instalação, análise ou run incompleto | Localizar run aberto, validar referências e mudanças desde o checkpoint, pedir confirmação se a intenção não for recuperável |
| `reconcile` | Toggle sugere instalação mas artefatos já existem, ou vice-versa | Comparar `project_manifest`, `analysis_status`, timestamps e eventos; explicar conflito e adotar estado comprovado sem substituir dados |

O usuário pode explicitar o modo. Se houver ambiguidade, o mestre consulta os artefatos e faz uma pergunta única e focada somente quando a intenção ainda não puder ser inferida com segurança.

## 3. Estados de projeto propostos

```text
UNINITIALIZED → DISCOVERY_IN_PROGRESS → BASELINE_RECORDED
                                       ├─ sem solicitação → WAITING_FOR_REQUEST
                                       └─ com solicitação → REQUEST_READY
REQUEST_READY → ARCHITECTURE_IN_PROGRESS → AWAITING_ARCHITECTURE_APPROVAL
→ ARCHITECTURE_BASELINED → BACKLOG_READY → ENGINEERING_READY
→ DEVELOPMENT_IN_PROGRESS → VERIFICATION → HOMOLOGATION_READY
→ HOMOLOGATION_PENDING / HOMOLOGATED → GMUD_PENDING / GMUD_AUTHORIZED (se aplicável)
```

`FAILED`, `PAUSED`, `NEEDS_INPUT` e `CANCELLED` são estados de run que podem ocorrer em fases diferentes. Estado de projeto e estado de execução são separados. Homologação e GMUD exigem aprovação/registro da autoridade definida pelo projeto/cliente; o Harness só registra prova verificável, não presume aceite.

## 4. Instalação e análise inicial

1. Ler os marcadores do Harness, manifesto, versão de schema e registros de instalação.
2. Ler modo/toggle da solicitação e comparar com estado encontrado.
3. Criar `installation_run_id` e gravar checkpoint antes de operações de escrita.
4. Inspecionar progressivamente o repositório: fatos, hipóteses e lacunas com fonte.
5. Classificar: sistema novo sem arquitetura aprovada; legado com arquitetura observável; ou sistema misto/multirrepositório.
6. Para legado, registrar baseline `as-is`, componentes, dependências e limites de observação; avaliar prós, contras, riscos e comentários com proveniência. Não transformar a arquitetura sem uma demanda de mudança.
7. Produzir artefatos, memória e log de retomada; atualizar marcador com estado, data/hora, versão e resultado.
8. Se não houver problema/demanda de mudança, encerrar em `WAITING_FOR_REQUEST`. O usuário poderá instalar agora e fornecer trabalho em outra sessão.

Reexecução idempotente compara fingerprints e preserva personalizações. Diferenças de projeto após uma análise anterior geram nova observação e delta; não substituem baseline aprovado em silêncio.

## 5. Aprovação e bloqueio da arquitetura

O arquiteto registra as alternativas adequadas ao problema, critérios, prós, contras, alertas residuais, evidência e recomendação. Lucas seleciona ou pede revisão. A seleção cria uma versão de baseline aprovada que inclui escopo, data, aprovador e referências dos artefatos. Etapas de backlog, engenharia e desenvolvimento recebem essa referência.

Uma implementação que não caiba na arquitetura encaminha pergunta ao engenheiro responsável. O engenheiro tenta solução compatível e, se necessário, encaminha ao arquiteto. O arquiteto investiga causa e solução alternativa. Se houver motivo grave que exija reabrir a arquitetura, a proposta ao Lucas contém: problema reproduzível; impacto/evidências; tentativas nas camadas inferiores; por que não são suficientes; alternativa de arquitetura proposta; impactos/prós/contras; plano de migração/compatibilidade; tarefas e validação. A pipeline pausa. Só após autorização explícita nasce nova versão; histórico não é sobrescrito.

Falha não justificada como motivo grave não abre baseline. Se Lucas negar, o sistema registra a decisão e retorna aos caminhos compatíveis ou pausa/encerra a demanda.

## 6. Human-in-the-Loop e escalonamentos

Pedir a intervenção humana quando: intenção de produto/material de negócio faltar; decisão de arquitetura exigir autoridade de Lucas; aprovações de homologação/GMUD pertencerem ao cliente; efeitos externos excederem política; ou evidência conflitante não puder ser resolvida por inspeção. A pergunta inclui contexto mínimo, opções concretas, recomendação, consequência do adiamento e ação que o sistema retomará após resposta.

O mestre acumula bloqueios independentes para uma solicitação clara quando possível. Não converte silêncio em consentimento para arquitetura, homologação, GMUD ou produção.

## 7. Desenvolvimento e homologação

O escopo aprovado é concluído por unidade acordada (feature/épico ou outra unidade apropriada). O ciclo inclui, quando pertinente: teste que falha antes da implementação; implementação; unitários; integração; E2E (Playwright quando adequado); lint/formatação; build; análise de segurança/configuração; code review independente; QA e evidência reproduzível. Critérios obrigatórios vêm do projeto e da mudança; uma pipeline gerada não os remove.

Ao concluir, o Harness reporta diferenças, commits/identificadores quando disponíveis, testes e resultados, riscos conhecidos, requisitos de homologação e como validar. O artefato de prontidão destina-se ao ambiente/branch de homologação escolhido pelo projeto. O Harness não afirma que o cliente homologou sem evidência dessa aprovação.

## 8. Homologação e GMUD

Fluxo de referência configurável por cliente/projeto:

```text
verificação concluída → pronto para homologar → homologação registrada
→ (se política exigir GMUD) dossiê preparado → auditoria/aprovação GMUD registrada
→ autorizado para processo de deploy externo
```

GMUD é uma governança de mudança após a homologação; pode exigir auditoria, aprovadores, janela, risco, plano de mudança/retorno, evidências e ligações com o escopo/commits homologados. O template e as exigências variam por cliente. O Harness pode montar evidências e registrar a decisão, mas a autorização cabe ao processo/autoridade do cliente. A anotação “GMUD Épico XYZ” em um PR/registro e seus hashes de commits homologados são vínculo organizacional, não definição de workflow de branches.

Depois da aprovação aplicável, o Harness entrega o dossiê e encerra sua responsabilidade no ponto anterior ao deploy em produção. Nenhum agente do Harness inicia produção.

## 9. Registro de eventos e recuperação

Cada transição registra `event_id`, `project_id`, `run_id`, estado anterior/novo, ator (usuário/agente/helper), data/hora com timezone, referências, motivo, resultado e `correlation_id`. Aprovações incluem identidade explícita e objeto aprovado. Eventos são append-only no nível lógico; correção cria evento retificador em vez de reescrever a história.

Na retomada, validar que referências existem, schema é compatível, versão de executor e arquitetura continuam identificadas, e arquivos relevantes não mudaram desde evidência crítica. Se baseline, commit ou contexto estiver divergente, criar checkpoint de reconciliação.

## 10. Invariantes

1. Instalação sem solicitação de mudança pode concluir em espera.
2. Projeto novo não recebe baseline de arquitetura aprovada antes da seleção humana exigida.
3. Agente de desenvolvimento não altera decisão arquitetural aprovada.
4. Proposta de reabertura vem do arquiteto com evidências e solução e exige Lucas.
5. Toggle incorreto não sobrescreve estado persistido.
6. Homologação e GMUD são gates distintos e configuráveis; aprovação não pode ser inferida.
7. Produção permanece fora do ciclo automatizado do Harness.
