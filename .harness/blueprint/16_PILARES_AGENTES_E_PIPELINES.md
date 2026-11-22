# 16 — Pilares, agentes especialistas e pipelines adaptáveis

**Responsável pelo produto:** Lucas de Lima  
**Estado:** estrutura dos quatro pilares confirmada; composição abaixo é um modelo inicial, ajustável ao problema  
**Decisões de autoridade:** `13_LUCAS_PRODUCT_VISION_AND_DECISIONS.md`

## 1. Conceitos que não devem ser confundidos

- **Pilar:** domínio de responsabilidade do ciclo principal: arquitetura, engenharia, gestão de backlog ou desenvolvimento.
- **Agente:** executor especialista com instruções e entregáveis próprios. Um pilar pode reunir vários agentes; a quantidade varia por problema.
- **Etapa:** unidade de trabalho e transição verificável. Pode invocar um ou mais agentes ou uma operação determinística.
- **Pipeline:** composição declarada de etapas, gates, dependências, rotas de exceção e resultados. Pode cruzar pilares.
- **Instância de pipeline:** execução concreta, congelada para rastreabilidade, que aponta para uma definição em determinada versão.

Os quatro pilares permanecem reconhecíveis, mas não exigem que toda solicitação passe por todos os agentes ou por todos os pilares. O mestre escolhe a composição proporcional ao problema, registrando por que etapas foram incluídas, reutilizadas ou dispensadas.

## 2. Hierarquia e direito de decisão

```text
Lucas (autoridade de produto e aprovação arquitetural)
└── Agente Mestre (gestão, composição e coordenação)
    ├── Pilar de Arquitetura
    ├── Pilar de Engenharia
    ├── Pilar de Gestão de Backlog
    └── Pilar de Desenvolvimento
```

Esta árvore descreve escalonamento e coordenação, não significa que todo trabalho precise ser executado em série. O fluxo de autoridade para dúvida técnica é desenvolvimento → engenharia → arquitetura. Uma dúvida sobre escopo/tarefa segue ao agente de backlog responsável. Se arquitetura detectar problema grave sem solução responsável, para e consulta Lucas. A alteração da arquitetura aprovada requer evidências e proposta de solução do arquiteto e autorização explícita de Lucas.

O mestre resolve ordem, capacidade e comunicação. Ele não substitui o parecer dos pilares nem se apropria das decisões reservadas a Lucas.

## 3. Catálogo inicial de responsabilidades

Os nomes abaixo são categorias de especialidade, não uma lista fechada de agentes que precisam ser criados para todo projeto.

| Pilar | Agentes/capacidades candidatas | Saída típica |
|---|---|---|
| Arquitetura | Analista do problema/domínio; arquiteto de solução; arquiteto de qualidade e segurança; revisor de alternativas; registrador de decisões | Problema e restrições, opções limitadas com critérios/prós/contras/alertas, arquitetura aprovada e baseline, decisões e riscos |
| Engenharia | Engenharia de plataforma/runtime; backend/API; frontend; dados/banco; segurança; integração; confiabilidade/observabilidade; infraestrutura/DevOps | Desenhos técnicos derivados, contratos, modelos, plano de execução, análise de impacto e questões escaladas |
| Gestão de backlog | Analista de produto/PO; gestor de backlog; decompositor e dependências; responsável por critérios de aceite | Épicos, features, histórias, tarefas, dependências, critérios e ordem proposta |
| Desenvolvimento | Implementador por área; especialista de testes; QA; revisor independente; integração/E2E; homologação; preparação de release | Alteração implementada, testes, revisão, evidência de homologação e pacote de prontidão para mudança |

O usuário pode aprovar especialidades e limites específicos por projeto. Agentes de revisão e QA devem, quando viável, receber evidência independente da execução original; aprovação não deve ser mera repetição do relatório do implementador.

## 4. Contrato de papel de agente

Cada papel tem manifesto com `role_id`, nome, pilar, versão, missão, responsabilidades, exclusões/autoridade, gatilhos de invocação, entradas obrigatórias, formato de saída, fontes autorizadas, ferramentas permitidas, escopo de escrita, regras de evidência, critérios de conclusão e caminhos de escalonamento.

Um agente deve declarar incerteza. Deve separar `observed`, `inferred`, `proposed` e `human_confirmed`; referenciar caminhos/linhas, comandos, testes, versões e datas quando aplicável. Se não consegue cumprir o contrato, retorna `needs_input` ou `blocked` com motivo, evidência coletada e pergunta concisa. Não completa por suposição silenciosa.

## 5. Contrato da etapa de pipeline

Campos mínimos propostos para uma etapa:

```yaml
stage_id: develop.feature
version: 1
pillar: development
purpose: Implementar a unidade de escopo aprovada com desenvolvimento guiado por testes.
depends_on: [engineering.design-approved, backlog.scope-approved]
agent_roles: [developer, test-specialist]
inputs: [approved_scope, engineering_contract, architecture_baseline_ref]
outputs: [change_set, test_report, evidence_index]
write_scope: [assigned_worktree]
completion_policy: all_required_outputs_present_and_verified
on_needs_input: pause_and_route
on_failure: preserve_attempt_and_escalate
retry_policy: bounded_with_new_attempt_record
```

O YAML é ilustrativo; schema formal deve ser versionado. Uma etapa inclui responsável pelo gate, evidência necessária, efeitos colaterais, timeout e política de repetição. Uma pipeline não pode apagar evidências nem enfraquecer invariantes de produto.

## 6. Pipelines estáveis, de projeto e temporárias

1. **Padrão upstream:** composição genérica revisada e versionada junto ao Harness.
2. **Reutilizável do projeto:** particularidade aprovada para domínio, cliente ou stack; vive na especialização do projeto.
3. **Temporária:** criada pelo mestre para uma solicitação/execução. Pode ser eliminada como definição ativa após o uso, mas a instância, a versão, os passos e os resultados permanecem auditáveis.
4. **Derivada:** cópia de um template com mudanças registradas; mantém origem e diferença.

Promover pipeline temporária exige evidência de reutilização, revisão de segurança/autoridade, entradas/saídas genéricas, ausência de dados sensíveis no template e aprovação conforme governança do projeto. O mestre pode criar e usar uma pipeline nova dentro das invariantes; não pode promover silenciosamente padrões específicos para o upstream.

## 7. Validação e invariantes da pipeline

Antes da execução, conferir: IDs únicos; referências e dependências existentes; ciclo no grafo de etapas; entradas produzidas antes do consumo; agente/adapter suportado; permissões suficientes e não excessivas; destinos de escrita sem disputa; gates de aprovação; tratamento de erro/retry; condição de parada; pontos de HITL; efeitos externos declarados; fronteira de homologação/GMUD/produção; proveniência da definição.

Uma definição com ciclo não intencional, autoridade incompatível, ausência de saída, etapa sem executor ou caminho de produção deve ser rejeitada. O mestre registra validação e o motivo. Uma etapa reprovada não é contornada removendo-a dinamicamente; é necessário criar uma nova versão ou uma rota segura autorizada.

## 8. Finalização e limites

- Arquitetura: pipeline de arquitetura encerra quando Lucas seleciona uma alternativa; os artefatos resultantes viram baseline versionado.
- Backlog: encerra com hierarquia e relações de escopo entregues.
- Desenvolvimento: encerra no escopo limitado acordado (épico, feature ou outra unidade), com QA, revisão, verificações e prontidão para homologação; o mestre espera Lucas selecionar o próximo escopo.
- Pipeline mestre: instala/análise, segue a sequência necessária para o problema e para no resultado solicitado. Um legado sem solicitação de mudança pode encerrar após baseline `as-is`.
- Todas: podem parar para HITL quando a intenção do produto não estiver clara ou quando uma decisão reservada não puder ser tomada pelo Harness.

## 9. Regra de custo e esforço proporcional

O arquiteto e o mestre primeiro caracterizam escala, qualidade, restrições, contexto e risco. Exploram um conjunto de opções plausíveis para o caso, sem uma obrigação de enumerar padrões irrelevantes. Devem declarar critérios de busca e explicar por que uma alternativa simples é suficiente ou por que maior sofisticação se justifica. Complexidade da análise cresce com incerteza, risco e impacto, não com a extensão de um catálogo de arquiteturas.

## 10. Questões abertas que não impedem este contrato

- Esquema executável final e tecnologia de validação.
- Nomes, cardinalidade e instruções concretas de cada especialista.
- Quantidade de concorrência e regras específicas de isolamento por executor.
- Critérios quantitativos de “alto risco” e de promoção para pipeline upstream.

Essas escolhas devem ser fechadas nas sessões de runtime, segurança e critérios de aceite, sem alterar a existência dos quatro pilares ou o requisito de agentes delegáveis.
