# 24 — Plano mestre de construção do Harness até prontidão

**Produto em construção:** Portable Engineering Harness  
**Autoridade:** Lucas de Lima  
**Estado inicial:** plano detalhado; implementação do produto ainda não iniciada  
**Modo de trabalho enquanto no repositório-fonte:** `harness_build`  
**Identidade e seleção de modo:** documento `23_IDENTIDADE_DO_WORKSPACE_E_MODOS_DE_TRABALHO.md`  
**Especificações técnicas:** documentos `13` e `15`–`22`; instrução persistente em `AGENTS.md`.

## 1. O que este plano controla

Este é o backlog mestre para construir **o Harness**. Os itens abaixo não são histórias de usuário de um sistema-alvo. Todo item pertence ao `product_id=portable-engineering-harness`. Alterações neste blueprint, na implementação do Harness e nas fixtures de validação devem ser classificadas como manutenção de especificação, construção ou verificação do Harness.

Nenhuma sessão, subagente ou compactação de contexto deve escolher “o próximo passo” apenas pela memória da conversa. Consultar este plano e o estado persistido do build. Um item só passa a `done` quando suas saídas e evidências forem registradas; texto afirmando “feito” não é evidência.

## 2. Resultado final pretendido

O Harness está pronto para release quando uma pessoa puder instalá-lo em um projeto novo ou legado, invocar um Master Agent que identifique corretamente o modo, use delegação real para especialistas em quatro pilares, componha pipelines apropriadas, preserve estado/decisões entre sessões, use as integrações previstas por providers funcionais, conduza um ciclo de desenvolvimento completo e verificável até homologação/GMUD configurável, e pare antes de produção. O produto precisa demonstrar tudo isso em fixtures sem depender do projeto privado do Lucas.

O primeiro release deve ser integrado e de alta capacidade. As fases ordenam dependências e evidência; não autorizam remover pilares, substituir agentes por personas ou omitir ferramentas exigidas. Componentes opcionais por decisão registrada (por exemplo, Phoenix como backend e OmniRoute como gateway escolhido) precisam ter adapter e estado real; não é necessário que cada instalação de projeto os configure.

## 3. Estado e regra de seleção do próximo item

### Estados de work item

`proposed` → `ready` → `in_progress` → `verification` → `done`  
Alternativas: `blocked`, `needs_decision`, `deferred_by_lucas`, `rejected_with_reason`.

`done` requer saída durável, critérios do item verificados, referências de commit/run/build quando existirem, decisão/dependência atualizada e próximo item desbloqueado. `blocked` requer razão/evidência e ação necessária. Dependência externa, não ausência de vontade, determina o bloqueio.

### Algoritmo de roteamento para a sessão seguinte

1. Confirmar `workspace_kind=harness_source` e `work_mode=harness_build` conforme documento 23.
2. Ler o arquivo de estado do build (a ser criado no início da implementação), o último run/handoff e este plano.
3. Verificar se algum `in_progress` tem retomada segura; continuar esse item antes de abrir trabalho paralelo.
4. Selecionar item `ready` da fase ativa cujas dependências estejam `done`.
5. Se não houver item `ready`, resolver dependência pendente ou enviar decisão material a Lucas; não iniciar fase adiante por conveniência.
6. Registrar no run o `work_item_id`, objetivo, saída esperada, limites, acceptance criteria e modo.
7. Ao fechar, gravar resultados, evidências, riscos e `next_ready_item`.

O agente pode propor ajuste de sequência quando evidência mostrar dependência ausente; registra motivo, impacto e itens afetados. Não apaga item nem altera meta de prontidão sem decisão de Lucas.

## 4. Convenção de registro de cada tarefa

Este documento define ID, descrição, dependências, saídas e critério de saída. A implementação converte cada item para o backlog/schema escolhido com campos:

```yaml
work_item_id: H-RT-04
workspace_kind: harness_source
work_mode: harness_build
title: Verificar dispatch real de especialistas no executor inicial
phase: H2
status: ready
depends_on: [H0-RT-02, H1-CT-01]
owner_role: runtime-engineer
outputs: [adapter-report, smoke-run-reference]
acceptance: [AC-RT-01, AC-RT-02]
evidence_refs: []
decision_refs: []
blocked_by: []
```

IDs deste plano são estáveis e devem permanecer rastreáveis caso sejam decompostos. A decomposição pode criar filhos, mas não perde a relação com o ID principal.

## 5. Fases, tarefas e gates

### H0 — Controle do build, identidade e decisões que destravam implementação

**Objetivo:** criar destino inequívoco para o produto e impedir mistura entre Harness-fonte e produto-alvo.  
**Gate H0:** identidade, runtime de referência, plano de estado e critérios de release aprovados/registrados.

| ID | Trabalho | Depende de | Entrega/critério de saída |
|---|---|---|---|
| H0-01 | Criar repositório executável do Harness separado do pacote blueprint e registrar a relação entre ambos | Nenhum | Repositório-fonte com root limpo, LICENSE/versão escolhidas e referência ao blueprint; não misturar código na pasta do pacote ZIP |
| H0-02 | Definir `harness-build.yaml`, estado de release, work item ativo e handoff de build | H0-01 | Manifesto com `workspace_kind=harness_source`, `product_status=under_construction`, `work_mode_default=harness_build` e esquema versionado |
| H0-03 | Identificar executor inicial e validar execução real de subagentes no ambiente alvo | H0-01 | Relatório com executor/versão, Windows, dispatch, retorno, erro, tools, isolamento, retomada e limites; Lucas escolhe caso haja alternativas materiais |
| H0-04 | Definir plataforma primária de build/execução e matriz Windows nativo, Linux e WSL2 | H0-03 | Matriz declara `supported`, `supported-via-WSL2`, `degraded` ou `unsupported`, com justificativa e plano de prova |
| H0-05 | Definir gateway/model policy inicial e classificação de dados permitida | H0-03 | Perfil local sem segredo commitado; diretiva de quais dados de projeto podem chegar à rota upstream |
| H0-06 | Confirmar destino de logs, memória auxiliar e backup; preservar memória canônica em projeto/documentação vinculada | H0-01 | Registro da decisão de armazenamento, scope e recuperação; nenhum serviço é declarado ativo por placeholder |
| H0-07 | Registrar decisão sobre número de executores suportados no release 1 | H0-03 | Declara um adapter de referência mais extensibilidade ou vários adapters completos; decisão de produto documentada |
| H0-08 | Congelar backlog/AC versionados deste plano e gerar estado de build inicial | H0-02–H0-07 | Estado versionado com fases, owners, status iniciais e `next_ready_item`; backlog mantém todos os requisitos confirmados |

### H1 — Especificações executáveis e arquitetura do produto

**Objetivo:** transformar os contratos de design em interfaces, schemas e regras validáveis antes de espalhar integração em código.  
**Gate H1:** adapters, estados, dados e limites de autoridade têm contratos versionados e uma arquitetura aprovada para o Harness.

| ID | Trabalho | Depende de | Entrega/critério de saída |
|---|---|---|---|
| H1-01 | Definir arquitetura de componentes, limites de processo, arquivos/config e direção de dependências | H0 | Architecture Decision Records para executor, estado, plugins/providers e dados locais/remotos |
| H1-02 | Formalizar schemas de workspace, build-state, project manifest e instalação | H1-01 | Schema/version/migration policy e exemplos válidos/inválidos; diferencia `harness_source` de `managed_target` |
| H1-03 | Formalizar schemas de agente, pipeline, instância, etapa, delegação, run, evento, handoff, evidence e provider | H1-01 | Schemas versionados ligados aos contratos 15,16,18,22; validação falha com erro explicativo |
| H1-04 | Definir state machine de projeto e de run e transições idempotentes | H1-02/H1-03 | Diagramas/transições testáveis para bootstrap, discovery, pipeline, pause, failure, resume, homolog e close |
| H1-05 | Definir autoridade, aprovações e enforcement, incluindo baseline de arquitetura e produção | H1-04 | Policy conformance map: quem pode propor, aprovar, escrever, promover pipeline, homologar, solicitar GMUD |
| H1-06 | Especificar logs/eventos append-only, correlation IDs, segredos e retenção | H1-03/H1-05 | Schema e redaction rules compatíveis com privacidade e retomada entre agentes |
| H1-07 | Definir interfaces de capability/provider e envelope comum de resultado/falha | H1-03 | Contratos `ProviderAdapter`, `AgentExecutorAdapter`, `SkillCatalog`, `MemoryProvider`, `TraceExporter` |
| H1-08 | Definir versão do produto, compatibilidade de config/artefatos e política de upgrade | H1-02/H1-03 | Versão semântica ou regra equivalente, upgrade report, migrations, rollback e suporte de N-1 decidido |

### H2 — Bootstrap, executor, agente mestre e identidade da sessão

**Objetivo:** uma invocação real reconhece que está desenvolvendo o Harness, estabelece estado e coordena especialistas sem confundir seu papel com o sistema-alvo.  
**Gate H2:** delegação real e seleção de modo/build plan funcionam em ambiente de desenvolvimento.

| ID | Trabalho | Depende de | Entrega/critério de saída |
|---|---|---|---|
| H2-01 | Criar root `AGENTS.md` do Harness-fonte e entrypoints do executor | H1/H0 | Agente recebe mensagem inequívoca de `harness_build`; instruções curtas apontam ao build state e plano |
| H2-02 | Implementar bootstrap AI-first e descoberta mínima de projeto sem scanner monolítico | H1-02/H2-01 | O agente começa antes de helpers e decide o que inspecionar; fatos têm proveniência |
| H2-03 | Implementar resolver de `workspace_kind`/`work_mode` com prompt, manifest, Git root, timestamps e verificação | H1-02/H2-02 | Invocações deliberadas distinguem blueprint/source/fixture/target; conflito resulta em reconcile ou HITL |
| H2-04 | Implementar carregador de instruções e leitura contextual de planos/artefatos | H2-01/H2-03 | Primeiro run lê o necessário; depois consulta artefatos pela tarefa, sem depender de transcript completo |
| H2-05 | Implementar adapter de executor real e relatório de capabilities | H0-03/H1-07 | Master inicia agentes com identidade/contexto separado e retorno/falha identificável |
| H2-06 | Implementar registry de papéis e dispatch de delegações | H2-05/H1-03 | Agente recebe tarefa, objetivo, context refs, limites, saída e escopo de escrita próprios |
| H2-07 | Implementar wait/result/error/timeout/cancel/retry sem perda de tentativa | H2-05/H1-04 | Todo resultado termina em estado explícito; tentativa reexecutada se liga à anterior |
| H2-08 | Garantir visibilidade de tools/providers a todos os agentes compatíveis | H2-05/H1-07 | Capability smoke matrix cobre Master e todos os papéis; sem herança de tools presumida |
| H2-09 | Implementar síntese gerencial do Master e escalonamento entre camadas | H2-06/H1-05 | Master coordena, não se declara especialista nem altera decisão de autoridade reservada |

### H3 — Persistência, memória canônica, pipeline dinâmica e retomada

**Objetivo:** converter decisão em execuções recuperáveis que o Master pode compor, iniciar, parar e continuar.  
**Gate H3:** pipeline simples e pipeline gerada passam de criação a encerramento com estado íntegro.

| ID | Trabalho | Depende de | Entrega/critério de saída |
|---|---|---|---|
| H3-01 | Criar armazenamento de projeto, configuração, eventos, artefatos e caches separados | H1-06/H2-03 | Manifesto identifica o que é versionado, local, externo e regenerável |
| H3-02 | Implementar memória canônica e índice de artefatos | H3-01 | Baseline, backlog, decisions, engineering, QA e governance têm owners e referências estáveis |
| H3-03 | Implementar run/event log, evidence index, checkpoint e handoff | H3-01/H1-06 | Outra sessão retoma com fontes, status, evidência, mudanças e próximo passo |
| H3-04 | Implementar recuperação de build state e seleção de próximo work item | H3-03/H0-08 | Nova sessão em `harness_build` escolhe item desbloqueado do plano 24, não inventa backlog-alvo |
| H3-05 | Implementar schema/validador de grafo de pipeline (dependências, ciclo, inputs/outputs, papéis) | H1-03/H2-06 | Pipeline inválida, sem agente ou sem saída é rejeitada antes de executar |
| H3-06 | Implementar gerador/compositor de pipeline proporcional à solicitação | H3-05/H2-09 | Cria definição temporária/reutilizável com justificativa de escolha e snapshot por run |
| H3-07 | Implementar estado por etapa, gates, pause/resume, retries e roteamento de exceção | H3-05/H3-03 | Instância reproduzível e auditável; mudanças geram revisão ligada ao original |
| H3-08 | Implementar HITL, perguntas agrupadas e escalonamento com contexto/evidência | H1-05/H3-07 | Silêncio nunca vira aprovação de arquitetura, homologação, GMUD ou produção |
| H3-09 | Implementar instalação/reconciliação idempotente via AI-first workflow | H2-03/H3-01 | Toggle incorreto não reinstala/sobrescreve; artefatos/timestamps resolvem estado com log |
| H3-10 | Implementar promoção explícita de pipeline temporária a projeto/upstream | H3-06/H1-05 | Requer provenance, generalização, revisão e autoridade declarada; não é automática |

### H4 — Agentes e quatro pilares end-to-end

**Objetivo:** entregar a estrutura especializada da organização e seus contratos por pilar; sem “um agente genérico por pilar”.  
**Gate H4:** execução sintética atravessa todas as etapas de cada pilar com responsabilidades e escalação observáveis.

| ID | Trabalho | Depende de | Entrega/critério de saída |
|---|---|---|---|
| H4-01 | Definir catálogo inicial de especialistas com foco, fronteira, inputs/outputs e tools | H1-03/H2-06 | Roles para Master; arquitetura; engenharia; backlog; dev, teste, QA, review e homologação/release readiness |
| H4-02 | Implementar pipeline de arquitetura e método proporcional de alternativas | H4-01/H3-07 | Descoberta do problema, critérios, alternativas/prós/contras/alertas e escolha explícita de Lucas |
| H4-03 | Implementar baseline `as-is` e modo legado analysis-only | H4-02/H3-02 | Arquitetura observável, limitações, riscos, evidências e parada à espera sem mudança requisitada |
| H4-04 | Implementar engenharia especializada e controle de compatibilidade arquitetural | H4-01/H4-02/H1-05 | Contratos técnicos derivam da baseline; conflito escala desenvolvimento→engenharia→arquitetura |
| H4-05 | Implementar reabertura arquitetural com pacote do arquiteto | H4-04/H1-05 | Proposta contém problema grave, provas, tentativas e solução; só Lucas autoriza nova baseline |
| H4-06 | Implementar gestão de backlog para épico/feature/história/tarefa/dependências | H4-01/H4-02/H3-02 | Hierarquia ligada a necessidade, arquitetura, aceite e ordem proposta |
| H4-07 | Implementar pipeline de desenvolvimento por unidade de escopo e TDD | H4-01/H4-04/H4-06 | Teste que falha primeiro quando aplicável, alteração, testes e rastreio de requisito |
| H4-08 | Implementar QA, review independente, integração e agente E2E apropriado | H4-07 | Evidência separada do implementador; Playwright quando apropriado; falha volta à correção sem apagar prova |
| H4-09 | Implementar gates de homologação e GMUD pós-homologação configurável | H4-08/H1-05 | Estados “pronto/enviado/aprovado”, evidências, autoridade e stop antes de produção |
| H4-10 | Implementar escolha do próximo escopo por Lucas após unidade concluída | H3-08/H4-09 | Desenvolvimento fecha feature/épico definido e não inicia o próximo por decisão própria |

### H5 — Providers e integração colaborativa das ferramentas

**Objetivo:** instalar/configurar ferramentas e dar aos agentes acesso funcional compartilhado por capability, com fontes, pins, falhas e estado real.  
**Gate H5:** cada integração exigida tem status `smoke_verified` nos agentes/plataformas declarados, ou uma lacuna aprovada que não invalide readiness.

| ID | Trabalho | Depende de | Entrega/critério de saída |
|---|---|---|---|
| H5-01 | Criar lockfile/manifesto de providers, verificação de licença/integridade e install report | H1-07/H1-08 | Sources oficiais, version/tag/hash, runtime, license, config e estado; idempotência e remoção |
| H5-02 | Implementar adapter CodeGraph e ciclo build/sync/status/query | H5-01/H2-08 | Agentes aprovados consultam símbolos/relacionamentos com arquivo/linhas e revision fresca |
| H5-03 | Implementar adapter Graphify local e consulta de código/documentos/SQL/config | H5-01/H2-08 | Grafo local e/ou MCP selecionado com origem, proveniência e cobertura declaradas |
| H5-04 | Implementar roteamento combinado e avaliação de overlap CodeGraph/Graphify | H5-02/H5-03/H3-03 | Query direcionada por intenção, comparação em modo avaliação, divergência checada na fonte |
| H5-05 | Validar indexes/daemon no Windows nativo, Linux e WSL2; separar cache quando necessário | H5-02/H5-03/H0-04 | Sem compartilhamento inseguro de DB/socket; run registra runtime/index revision |
| H5-06 | Instalar/configurar ai-memory e integrar MCP, scope por projeto e lifecycle/handoff | H5-01/H1-07/H3-03 | Mestre e especialistas recebem/recuperam handoff; dado canônico continua no repo |
| H5-07 | Resolver execução do ai-memory em Windows: nativo se suportado ou WSL2/host controlado | H5-06/H0-04 | Matriz real de suporte e round-trip no executor; credenciais e rede fora de arquivos commitados |
| H5-08 | Criar catálogo de skills e loader on-demand com metadados e permissões por executor | H1-07/H2-08 | Cada especialista descobre skill aplicável sem receber coleção inteira em toda chamada |
| H5-09 | Baixar/clonar ECC em pin e curar agent/rules/hooks/skills utilizáveis | H5-08/H5-01 | Inventário, licença, dependências e conflitos; adaptar apenas capacidades aprovadas; não rodar installer amplo sem diff |
| H5-10 | Baixar/clonar skills Matt Pocock em pin e curar conjunto compatível | H5-08/H5-01 | Skills estáveis/beta diferenciadas; atribuição e versão; disponibilidade em todos hosts declarados |
| H5-11 | Implementar índice/biblioteca consultável de publicações Anthropic | H5-08/H3-03 | Fontes por tópico, atualização, data, citation/provenance, escopo e política de conteúdo licitamente armazenado |
| H5-12 | Integrar conceitos/artefatos Spec Kit sem pipeline concorrente | H5-08/H3-05 | Mapping/import-export dos artefatos que agreguem valor; nenhuma duplicação de status/authority |
| H5-13 | Preparar Harness² em ambiente de avaliação isolado e com benchmarks licenciados | H5-01/H0-04/H0-05 | Setup reproduzível, credenciais/redes segregadas, benchmark amostral e resultado importável; sem runtime de produto |
| H5-14 | Instrumentar traces/evals via OpenTelemetry e adicionar export local/backend | H1-06/H3-03 | IDs correlacionados, payload sensível redigido, fallback local honesto |
| H5-15 | Implementar adapter Phoenix opcional e estado de configuração verificável | H5-14 | Trace visualizável em Phoenix local/permitido; runtime continua funcional sem Phoenix |
| H5-16 | Implementar OmniRoute local como gateway obrigatório e adapter direto adiado | H0-05/H1-07 | Gateway funcional; rota upstream autorizada, custo/uso e fallback sem mudança de privacidade silenciosa |
| H5-17 | Implementar context assembler federado para skills, grafos, memória e fontes | H5-02–H5-16/H3-03 | Relevância, limite, dedup, classificação de fato/inferência, freshness e referências |
| H5-18 | Validar acesso/resultado de cada capability por cada papel de agente | H2-08/H5-02–H5-17 | Matriz cobertura Master+4 pilares; falhas reportam `unavailable/denied/stale`, nunca fabricam sucesso |

### H6 — Segurança, controle de efeitos e operação segura

**Objetivo:** transformar autonomia em controles executáveis e tornar instalação, memória e upgrades recuperáveis.  
**Gate H6:** efeitos de ferramentas são classificados e protegidos pelo executor/policy; limitações não são disfarçadas por prompt.

| ID | Trabalho | Depende de | Entrega/critério de saída |
|---|---|---|---|
| H6-01 | Implementar capability profiles por executor e modelo | H0/H2-05 | Ferramentas, subagentes, rede, sandbox, filesystem, timeout, memória, secrets e trace declarados |
| H6-02 | Aplicar escopos de escrita e isolamento/serialização para agentes paralelos | H6-01/H2-06 | Sem writers concorrentes não isolados; cada diff atribuído a run/agente |
| H6-03 | Tratar conteúdo do repo, skills, web, MCP e outros agentes como entrada não confiável | H6-01/H5-17 | Injeção não altera hierarquia/permissions; saída ofensiva registrada/rejeitada |
| H6-04 | Centralizar secrets fora de artefatos/logs; redigir dados em traces e memória | H6-01/H5-06/H5-14 | Scans/inspection de outputs demonstram que segredo não vaza por padrão |
| H6-05 | Implementar policy de rede/depêndencias e relatório de confiança | H6-01/H5-01 | Fontes, scripts, licença, cadeia de supply, lockfile e ações externas passam por política configurada |
| H6-06 | Implementar instalação/reversão de providers com backup, diff e sem sobrescrita | H5-01/H6-01 | Install/uninstall preserves user config and project data, reports partial failure and recovery |
| H6-07 | Implementar upgrade/migração de schema e Harness upstream | H1-08/H3-01/H6-06 | Detecta customization conflict, gera preview/backup e mantém instância utilizável ou rollback |
| H6-08 | Formalizar retenção/expurgo e política de portfólio/documentação central | H0-06/H3-01/H5-06 | IDs e links por sistema; nenhuma duplicação concorrente de canonical memory |
| H6-09 | Checar license/attribution e remover componentes não redistribuíveis | H5-09–H5-13/H5-15 | Third-party notices e release SBOM/inventory; benchmark data separada de software |

### H7 — Portabilidade, qualidade e upgrade em sistemas suportados

**Objetivo:** provar que o produto funciona além da máquina de desenvolvimento, nas combinações efetivamente declaradas.  
**Gate H7:** perfis de sistema suportados têm instalação, uso multiagente, providers e recuperação verificados.

| ID | Trabalho | Depende de | Entrega/critério de saída |
|---|---|---|---|
| H7-01 | Adaptar paths, shells, subprocessos, permissões, atomic writes e locks | H0-04/H6-01 | Suíte de portabilidade passa em Windows e Linux declarados; diferenças documentadas |
| H7-02 | Testar caminho WSL2 específico sem pressupor índice compartilhado com Windows | H5-05/H5-07/H7-01 | Instruções e caches distintos ou prova suportada de compartilhamento |
| H7-03 | Validar instalação/removibilidade em repo limpo e repo com customização/alterações | H3-09/H6-06 | Nenhuma perda de trabalho ou config externa; manifest e report coerentes |
| H7-04 | Validar atualização de versão e migrações entre schemas suportados | H6-07/H7-03 | Cenários N-1 e conflito com customização têm resultado explícito e recuperação |
| H7-05 | Avaliar precisão/frescor dos grafos em fixture pequena, monorepo e multi-repo | H5-04/H5-05 | Metrics/ground truth, falhas e limites documentados; thresholds definidos por Lucas antes de release |
| H7-06 | Avaliar retomada após encerramento/compactação/interrupção/falha de provider | H3-03/H3-07/H5-06 | Continua do checkpoint correto; sem duplicar efeito e sem perder decisão/evidência |
| H7-07 | Avaliar custo/latência/confiabilidade por pilar e limites de concorrência | H2-05/H5-18/H5-14 | Baseline com settings/versões e tradeoffs, sem otimização que rebaixe qualidade do requisito |

### H8 — Validação E2E e prontidão de release

**Objetivo:** demonstrar critério completo do produto, decidir readiness com evidência e entregar artefatos de operação.  
**Gate H8:** todos os critérios obrigatórios passam, limitações aceitas por Lucas estão registradas e release checklist aprovado.

| ID | Trabalho | Depende de | Entrega/critério de saída |
|---|---|---|---|
| H8-01 | Implementar fixture de sistema novo e validar instalação→ideação→arquitetura→aprovação | H2–H7 | Run persistido e artefatos; alternativas proporcionais; sem detalhe do sistema privado |
| H8-02 | Implementar fixture legado e validar analysis-only→baseline as-is→waiting | H2–H7 | Evidência as-is/prós/contras; nenhum arquivo do sistema alterado sem solicitação |
| H8-03 | Validar solicitação sintética completa pelos quatro pilares até homologação | H2–H7 | Story/epic, design, TDD, unit/integration/E2E, review, QA e pacote de homologação |
| H8-04 | Validar falha de arquitetura, escalada e autorização de reabertura | H4-05/H8-03 | Sem autorização, baseline permanece; com autorização, nova versão retém proveniência |
| H8-05 | Validar GMUD condicional e parada pré-produção | H4-09/H8-03 | GMUD só pós-homologação, vínculo de evidência/commits e nenhum deploy executado |
| H8-06 | Validar dynamic pipeline temporary/project/reusable e após descarte de definição | H3-05–H3-10/H8-03 | Snapshot e histórico auditáveis; nenhuma rota contorna gate |
| H8-07 | Validar instalação e uso de cada integração; medir colaboração entre providers/agentes | H5-18/H7-05 | Status real por provider e papel; cenários do doc 22 pass/fail com evidência |
| H8-08 | Executar regressões e avaliação Harness² na amostra permitida | H5-13/H7-07/H8-01–H8-07 | Relatórios reproduzíveis, custo/limitações e sem vazamento de respostas/regras do benchmark |
| H8-09 | Fechar AC-01–AC-15 e gates adicionais deste documento | H8-01–H8-08 | Traceability matrix liga cada requisito, tarefa, teste/evidência e resultado |
| H8-10 | Preparar documentação operacional, versão, release notes, setup/uninstall/upgrade, support matrix | H6-09/H8-09 | Uma pessoa instala, verifica, atualiza, recupera e remove sem inferir passos da conversa |
| H8-11 | Produzir relatório final de readiness, limitações, dependências e decisão recomendada | H8-10 | Evidência por plataforma/integração, itens não atendidos, risco residual e recomendação explícita |
| H8-12 | Registrar decisão de Lucas sobre `ready_for_target_use` e versão de release | H8-11 | Release passa a `ready_for_target_use` só após aprovação; sem aceite permanece `release_candidate` |

## 6. Dependências e caminho crítico

```text
H0 identidade/runtime
  → H1 contratos/esquemas/autoridade
    → H2 execução multiagente e identidade da sessão
      → H3 estado/pipeline/retomada
        → H4 quatro pilares e fluxo de produto
        → H5 providers/skills/conhecimento
          → H6 segurança/instalação/upgrade
            → H7 Windows/Linux/qualidade/portabilidade
              → H8 validação integrada e decisão de release
```

H4 e H5 podem ter tarefas paralelas após seus contratos e interfaces estarem estáveis, mas H8 depende das duas. H6 inicia controles em paralelo quando interfaces de H1/H2 existem, mas a aprovação de H6 precisa cobrir efeitos das integrações H5. Se o executor não satisfaz multiagente real, H0/H2 bloqueiam mais implementação; não construa um harness sem a capacidade que é fundamental.

## 7. Definição de “Harness em construção”, “release candidate” e “pronto”

| Estado | Evidência mínima | O que o Master/AGENTS deve dizer |
|---|---|---|
| `blueprint_only` | Só especificação/artefatos | “Estamos preparando a construção; o Harness não está instalado nem operacional.” |
| `under_construction` | Source repo e build plan ativos, H0–H7 em andamento | “Este workspace está construindo o Harness. Não é uma sessão de produto-alvo.” |
| `release_candidate` | H8 executado; relatório e defeitos/limitações pendentes | “Harness está em avaliação; pode ser testado em fixtures, não tratado como liberado para produto privado/cliente.” |
| `ready_for_target_use` | Todos os gates obrigatórios aprovados por Lucas; versão publicada/pinada e support matrix clara | “Harness está liberado nesta versão e combinação executor/plataforma; abra managed_target e informe objetivo.” |
| `maintenance` | Harness liberado e work item altera/testa o próprio Harness | “Estamos mantendo o Harness, não entregando a feature de um sistema-alvo.” |

`ready_for_target_use` é versionado e compatível por plataforma/executor. Release não elimina estado de manutenção; um defeito na ferramenta retorna `harness_source`/`harness_build` ou `harness_verification` sem confundir com a feature em um sistema-alvo.

## 8. Gates globais de conclusão

O Harness não termina quando o último arquivo é criado. O produto pode ser marcado `ready_for_target_use` somente se:

1. H0–H8 estão `done`, ou exceções são `deferred_by_lucas` que não invalidam requisito obrigatório, com aceitação explícita e versionada.
2. AC-01–AC-15 do documento 21 têm evidência e nenhuma afirmação de integração usa apenas placeholder.
3. Há agente mestre e specialists reais dos quatro pilares; todas as responsabilidades essenciais têm contratos e execução verificável.
4. O resolver de identidade/modalidade diferencia o Harness-fonte, fixture e projeto-alvo antes da primeira alteração.
5. O plano 24 permite retomar build por ID e escolher automaticamente item desbloqueado correto em nova sessão.
6. Pipeline dinâmica preserva gates e tem snapshot auditável após definição temporária expirar.
7. CodeGraph, Graphify e ai-memory operam por providers funcionais; especialistas elegíveis têm acesso verificado e resultado com scope/proveniência/freshness.
8. ECC e Matt Pocock têm fontes pinadas/catalogadas, skills aprovadas disponíveis aos papéis compatíveis, licença e divergências registradas. Não é requisito carregar todas simultaneamente.
9. Biblioteca Anthropic tem escopo inicial explícito, retrieval funcional, URLs/datas e limites de cobertura/termos; cobertura “tudo que foi publicado” não é falsa promessa.
10. Harness² roda em bancada de avaliação isolada e relata método, benchmark, versões, custo e limitações; não altera produção sozinho.
11. Gateway OmniRoute, logs/redação, OTel e portabilidade têm estado coerente com o perfil de release; Phoenix não vira hard dependency e a rota upstream não muda silenciosamente.
12. Instalação e remoção não perdem alterações e distinguem setup local, instalação no alvo e configuração global.
13. Novo, legado sem demanda, legado com demanda, entrega até homolog, GMUD se aplicável, pausa, retomada e falha de provider foram demonstrados.
14. Harness nunca promove nem executa deploy de produção; a fronteira aparece no executor, pipeline e smoke tests.
15. Uma verificação independente avalia critérios, rastreabilidade, segurança, portabilidade e documentação antes do gate de Lucas.
16. Lucas aceita a recomendação de release com limitations, support matrix e versão específica apresentados.

## 9. O que o agente de código informa em cada handoff

```text
workspace_kind / work_mode:
product_status / release:
phase_id / work_item_id:
work items completed and evidence:
verification run references:
files/commits changed:
open blockers or Lucas decisions:
risks/limitations discovered:
next ready work item and why:
current Harness readiness state:
explicit statement: building Harness or using it on target system:
```

Se o último campo não puder ser respondido com manifesto + solicitação + plano, o agente deve reconciliar a identidade antes de prosseguir.

## 10. Relação deste plano com o plano técnico anterior

`21_ARQUITETURA_INTEGRADA_E_PLANO_DE_IMPLEMENTACAO.md` explica arquitetura, viabilidade, fases e critérios. Este documento decompõe essa direção em dependências e work items rastreáveis até o gate final. `23_IDENTIDADE_DO_WORKSPACE_E_MODOS_DE_TRABALHO.md` controla que tipo de produto e objetivo cada run está tratando. Os três se complementam e não substituem decisões confirmadas no documento 13.

## 11. Questões abertas sem perda do objetivo

As únicas escolhas que não podem ser inventadas pelo agente implementador são: repositório/destino do código executável; executor/runtime inicial; política de rota upstream/gateway e de dados; thresholds de avaliação após baseline; e limitações que Lucas autoriza no release. O agente deve apresentar evidências e recomendação concreta para cada escolha. Estas questões não mudam o objetivo: construir o Harness completo e diferenciar o build de seu uso em sistemas-alvo.
