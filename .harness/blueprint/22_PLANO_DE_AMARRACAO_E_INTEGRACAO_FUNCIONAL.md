# 22 — Plano de amarração e integração funcional

**Responsável pelo produto:** Lucas de Lima  
**Estado:** contrato de integração proposto para orientar implementação e validação  
**Objetivo:** fazer componentes cooperarem no trabalho real dos agentes, e não apenas baixar/clonar repositórios.

O modo de construção e a distinção entre o código-fonte do Harness e um projeto-alvo são obrigatórios conforme `23_IDENTIDADE_DO_WORKSPACE_E_MODOS_DE_TRABALHO.md`; a sequência de implementação dos providers está rastreada por H5-01–H5-18 no plano 24.

## 1. Resposta de arquitetura

Uma pasta com cada repositório clonado não é um Harness integrado. Para haver colaboração, falta uma camada comum que conecte **identidade do projeto/run**, **registro de agentes**, **catálogo de capacidades**, **adapters de provider**, **contexto/proveniência** e **estado compartilhado recuperável**.

O fluxo que cada agente usa deve ser:

```text
agente identifica a necessidade
  → consulta capability catalog com o project_id/run_id
  → adapter seleciona ferramenta/fonte permitida
  → provider devolve resposta estruturada + fonte + versão/frescor
  → context assembler limita e rotula o contexto
  → agente decide e produz artefato/evidência
  → estado e handoff ficam recuperáveis pelos outros agentes
```

Esse contrato torna as ferramentas disponíveis aos especialistas sem colocar o conteúdo inteiro de todos os repos em cada prompt. “Disponível para todos os agentes” significa que todos podem descobrir e invocar capacidades compatíveis e permitidas; não significa carregar todas as skills e índices em todo contexto, nem dar credenciais e efeitos externos idênticos a todo papel.

## 2. Camada comum de integração

| Peça | Responsabilidade |
|---|---|
| Agent Registry | Mapeia `agent_role_id` para pilar, versão de instruções, capabilities que pode solicitar, ferramentas e escopo de escrita |
| Capability Catalog | Declara capacidades sem acoplar agente a fornecedor: `code_structure.query`, `cross_artifact.query`, `project_memory.recall`, `project_memory.handoff`, `skill.retrieve`, `knowledge.search`, `evaluation.run`, `trace.export`, `model.complete` |
| Provider Registry | Liga cada capability a um ou mais providers configurados, com versão, plataforma, saúde, política de dados e fallback permitido |
| Provider Adapter | Implementa chamadas/configuração de provider, traduz o protocolo e normaliza resposta/erro; não decide produto nem arquitetura |
| Context Assembler | Recupera material pertinente, aplica limite, desduplicação, classificação epistemológica, origem, data e confiança, e fornece links de verificação |
| Project/Run Context | Acrescenta `project_id`, `request_id`, `run_id`, `stage_id`, baseline aprovada, commit/fingerprint e classificação de sensibilidade |
| Shared Artifact/Event Store | Persiste artefatos canônicos, delegações, eventos, aprovações e handoffs para todo agente retomar |
| Executor Capability Adapter | Torna tools/MCP, arquivos, skills, eventos e subagentes disponíveis conforme o executor, sem presumir herança de ferramentas |

Cada provider retorna um envelope mínimo:

```json
{
  "capability": "code_structure.query",
  "provider": "codegraph",
  "provider_version": "pinned-version",
  "project_id": "project-id",
  "indexed_revision": "git-commit-or-fingerprint",
  "status": "ok",
  "items": [],
  "sources": [{"path": "relative/path", "lines": "12-30", "hash": "..."}],
  "limitations": [],
  "observed_at": "RFC3339 timestamp"
}
```

O envelope é modelo lógico; schemas concretos devem ser versionados. O adapter deve preservar resposta nativa quando normalização perder semântica. `ok` sem `sources` pode ser útil para ferramenta de execução, mas não basta para evidência factual do projeto.

## 3. Disponibilidade para todos os especialistas

O instalador configura a capability por projeto/runtime uma vez; cada chamada de agente recebe acesso segundo o perfil do executor. A validação deve cobrir mestre e especialistas de todos os quatro pilares, não apenas a sessão principal.

1. Se o executor propaga MCP/tools compartilhados aos subagentes, verificar por teste documentado que cada agente vê a capability configurada.
2. Se não propaga, adapter registra provider para cada agente no início, ou declara incompatibilidade se a instalação por agente não for confiável/portável.
3. Skills e conhecimento são descobertos pelo catálogo e carregados sob demanda; instruções comuns pequenas explicam como buscar.
4. `project_id`, `run_id`, `stage_id` e `agent_role_id` acompanham eventos de ferramenta quando o protocolo permite.
5. Um agente pode devolver erro `capability_unavailable`/`scope_denied`, sem inventar resultado ou dizer que ferramenta foi usada.
6. O mestre pode delegar explicitamente consulta a agente especializado; o resultado compartilhado fica no run, com proveniência, em vez de depender da memória privada do worker.

O catálogo oferece capacidade; o manifesto do agente determina quando e como usá-la. Especialidade não é uma ACL global. Operações de leitura de conhecimento podem ser comuns; gravação em memória, alteração de pipeline ou uso de segredo têm política própria.

## 4. Classes de componentes: como integrar de verdade

| Fonte | O que ela é | Como entra na colaboração | O que não deve ser presumido |
|---|---|---|---|
| CodeGraph | Ferramenta executável de indexação/query de código, local-first, com interface MCP/CLI declarada | Instalar versão fixada; inicializar índice por projeto/runtime; expor capability de código a agentes; testar atualização e referências de fonte | Que configuração MCP do mestre automaticamente aparece em todo subagente; que índice Windows/WSL2 é compartilhável |
| Graphify | Ferramenta/skill local para construir e consultar grafo de código e artefatos; há também opção MCP com topologia própria | Fixar pacote; gerar grafo local por projeto; expor query local/stdio MCP; registrar inputs cobertos e arestas/proveniência | Que CLI/skill local e Graphify Hosted MCP são o mesmo provider ou têm a mesma privacidade; não conectar serviço hosted sem aprovação de dados |
| ai-memory (assumi que “YA Memory” significa este projeto) | Serviço de memória persistente com MCP, API e integrações de ciclo de vida | Escolher serviço local ou hospedagem controlada; criar escopo/namespace por projeto; registrar recall/handoff e integrar lifecycle do executor; validar acesso pelo mestre e especialistas | Que memória auxiliar substitui artefatos canônicos; que Windows nativo tem paridade com Linux (matriz atual marca native Windows experimental) |
| ECC / Everything Claude Code | Coleção de skills, agentes, rules, comandos, hooks e exemplos de configuração | Fixar clone/cache; catalogar arquivos/skills; selecionar e adaptar conteúdo útil para o formato/capability registry; preservar licença; rodar só módulos de que precisamos | Que o instalador da coleção combina com o Harness ou é neutro; ele pode gravar configurações e regras no executor |
| Matt Pocock Skills | Coleção versionada de skills, com plugin gerenciado para Claude Code e instalação de skills escolhidas por outras ferramentas | Fixar origem; avaliar skill por skill, formato, estabilidade, licença e sobreposição; disponibilizar skills aprovadas no catálogo compartilhado | Que plugin Claude é instalação universal para Codex ou que conteúdo beta/in-progress tem estabilidade da coleção promovida |
| Harness² | Ferramenta experimental de avaliação/recursão e conjunto de benchmarks para harnesses | Instalar em ambiente de avaliação separado; selecionar benchmark/licença/modelos; executar amostra controlada; importar métricas, configs e artefatos de experimento via contrato de resultado | Que seja serviço MCP para agentes de produto ou motor da pipeline do dia a dia |
| Anthropic publicações | Conteúdo de pesquisa, guias, posts e documentação em múltiplos canais | Registro/versionamento de fontes + busca/retrieval contextual; apresentar URL, data, escopo e distinção entre orientação e política | Que exista corpus completo/estático ou direito para baixar e redistribuir tudo |
| OpenTelemetry | Convenções e bibliotecas de instrumentação/tracing | Instrumentar o mestre, workers e adapters no ponto que gera eventos; compartilhar correlation IDs; mandar traces para destino configurado | Que traces sejam memória canônica, nem que prompt/conteúdo completo deva ser capturado |
| Phoenix | Backend/UI opcional para traces e análise | Configurar exporter apontando ao Phoenix local/permitido; vincular trace e run; manter export/local report fallback | Que Phoenix seja dependência para runtime ou prova de correção |
| OmniRoute | Gateway obrigatório de modelos no caminho inicial, com endpoint compatível e painel | Expor como provider de `model.complete`; rodar local/loopback, guardar credenciais em local seguro e registrar rota/uso; a rota upstream é allowlist por provider/modelo e classe de dados | Que fallback automático preserve custo, privacidade, modelo e qualidade sem policy explícita |
| Spec Kit | CLI/processo e templates de SDD com comandos/workflows próprios | Mapear seus artefatos e métodos para os contratos do Harness; importar/exportar ou adotar extensões isoladas se justificadas | Que devemos executar outro pipeline de implementação competindo com o pipeline mestre |
| Agent Skills specification | Formato aberto de empacotamento de instruções/procedimentos | Usar como formato de skill quando executor aceitar; adaptar ou gerar índice por executor | Que padrão de diretório garanta descoberta/execução igual em todos os hosts |

## 5. Um ponto importante sobre “clonar todos”

Clonar pode ser útil para auditoria, pins, busca offline e curadoria; não é requisito para cada capacidade. Para executáveis, preferir release oficial/pacote versionado e integridade verificável à cópia de todo o repositório. Para coleções de skills, manter checkout/cache pinado pode ser a fonte de curadoria; copiar para o projeto apenas arquivos escolhidos e permitidos. Para conteúdo que só é referência, catálogo linkado e snapshots permitidos bastam. Código upstream inteiro não deve ser copiado nem instalado dentro do Harness só porque foi citado.

Um lockfile de providers proposto registra `provider_id`, URL oficial, release/tag/commit, pacote/artefato, licença, hash/atestado disponível, origem da instalação, runtimes suportados, data de validação, configuração não secreta, capacidades, status e procedimento de atualização/remoção.

## 6. Plano de instalação funcional por ondas

### Onda 0 — Ambiente e baseline

- Confirmar executor real e versionamento do host, Windows nativo ou WSL2, runtime, diretório do projeto e identidade do projeto.
- Registrar dependências já disponíveis, política de dados e ferramentas permitidas.
- Criar o manifesto de providers e pins antes de instalar qualquer coisa.
- Usar `dry-run`/checagem de versão e registrar impacto dos instaladores. Configuração global não é padrão: preferir escopo de projeto ou configuração gerenciada reversível.

### Onda 1 — Acesso compartilhado e memória

- Provar delegação real do mestre e especialistas.
- Configurar serviço/escopo de ai-memory, MCP e handoff/lifecycle para runtime escolhido; decidir se memória roda localmente ou em servidor controlado.
- Garantir que artefatos canônicos do Harness funcionem mesmo com ai-memory desligado.

### Onda 2 — Contexto de código e artefatos

- Instalar CodeGraph e Graphify em versões fixadas e pelo caminho oficial adequado ao Windows/Linux.
- Construir dois índices por projeto (com cache isolado por OS/runtime quando necessário).
- Expor os dois aos especialistas por adapters MCP/CLI; configurar instruções de roteamento por intenção e manter Graphify local se o código for privado.
- Confirmar que pelo menos mestre, arquiteto, engenheiro e desenvolvedor conseguem consultar ambos; QA/revisor conseguem reproduzir consulta e fonte.

### Onda 3 — Skills e biblioteca

- Baixar/clonear ECC e Matt Pocock em cache de referência controlado, sem executar instalador global.
- Inventariar arquivos, licença, estabilidade, dependencies, platform assumptions e conflitos; selecionar skills e manter lineage.
- Normalizar o índice de skills para todos os agentes suportados, carregando cada skill sob demanda.
- Criar o registro Anthropic: fontes por tópico, ingestão legal/permitida, atualização incremental, busca e links de citação.

### Onda 4 — Medição e autoaperfeiçoamento controlado

- Instrumentar eventos do Harness com OpenTelemetry e exportar traces para arquivo/backend permitido.
- Instalar Phoenix apenas se necessário para visualizar/avaliar traces; manter modo local/headless funcional.
- Criar sandbox/ambiente separado do produto para Harness². Conferir requisitos da versão (docs atuais citam Linux e Python 3.12, Node 20+, `uv`, credenciais Google Cloud, e Docker para certos benchmarks), custo, acesso de rede e licenças dos benchmarks.
- Rodar benchmark pequeno com tarefas e saída de dados revisadas; importar relatório de avaliação, não executar o Harness² como agente permanente em todo run.

### Onda 5 — Gateway de modelos obrigatório no caminho inicial

- Instalar OmniRoute com versão fixada e bind local/loopback; configurar provider/API key em secret store e allowlist de provider/modelo por perfil e classe de dados.
- Validar `/v1/models`, uma chamada de completions/responses compatível com o executor e log de rota/fallback usando contexto synthetic/public.
- Desabilitar fallback silencioso: a chamada falha se não houver rota explicitamente allowlisted que aceite a classe de dados. O caminho direto de provider não integra o release inicial.

## 7. Critérios de colaboração (smoke scenarios)

Estes testes de aceitação serão pedidos na implementação; a lista não declara que já foram executados:

1. O mestre inicia arquiteto e engenheiro em tarefas diferentes; ambos retornam com IDs e artefatos que o mestre combina.
2. Um especialista pesquisa símbolo/chamada no CodeGraph, recebe caminho/linhas/commit, e outro especialista consegue obter o mesmo resultado sob a mesma versão de índice.
3. Um arquiteto pergunta relação entre requisito documental, tabela SQL e componente; Graphify recupera o grafo e a fonte original é aberta para verificação.
4. A mesma consulta de código é avaliada nos dois grafos para identificar sobreposição; o roteador só duplica consulta no modo avaliação.
5. Um agente salva handoff em ai-memory com escopo do projeto, e outra sessão/agente o recupera; a baseline arquitetural é lida do repositório e prevalece se divergir.
6. Um agente recebe skill de TDD aprovada sob demanda; outro pode consultar a mesma skill, mas nenhuma regra específica de Claude é tratada como universal sem adapter.
7. Alteração em arquivo invalida/atualiza índice conforme contrato; consulta stale é detectada e instrução de verificar fonte é acionada.
8. Harness² executa fora da pipeline normal e emite avaliação versionada; seu resultado pode abrir proposta de mudança mas não altera skills/arquitetura automaticamente.
9. Todos os agentes elegíveis geram eventos ligados pelo mesmo run/correlation IDs; o relatório de traces omite segredo e conteúdo sensível por padrão.
10. Desabilitar cada provider deixa um estado degradado honesto e comportamento definido, sem declarar que a integração continua ativa.

## 8. Requisitos de reversibilidade e controle

Instaladores devem suportar escopo de projeto/usuário, inventário de alterações, versão, logs sem segredo e remoção segura. Não usar configuração `latest` como pin reprodutível. Antes de executar scripts upstream, inspecionar o conteúdo/efeito e preferir releases verificáveis. Configurações dos executores devem ser geradas por adapter com cópia/backup/diff e removidas sem apagar dados/índices do projeto inadvertidamente.

Cada integração mantém status independente: `planned`, `downloaded`, `installed`, `configured`, `connected`, `smoke_verified`, `optional_unconfigured`, `degraded`, `failed`, `removed`. Só `smoke_verified` significa que o agente chamou a capability e recebeu resultado válido.

## 9. Decisões de execução que ainda são necessárias

1. Qual executor hospeda a primeira instalação real e seus subagentes? Recomendação técnica: provar primeiro o Codex que Lucas usa no Windows, depois testar o caminho Linux; isso é uma recomendação pelo ambiente, não decisão aprovada.
2. ai-memory será serviço local de usuário, processo em container/host local, ou serviço em LAN? Isso define privacidade, backup, rede e configuração.
3. Os índices e artefatos ficam apenas no projeto local, em um repositório de docs central, ou combinação por manifest? Base canônica continua no projeto ou docs repo vinculado.
4. Qual perfil OmniRoute e qual rota upstream provider/modelo explicitamente allowlisted serão usados no primeiro smoke? O gateway permanece local/loopback até decisão de deployment distinta.
5. Qual amostra de projeto/fixture sanitizada pode validar os dois grafos e agentes sem revelar o projeto privado?

Essas decisões determinam configuração real, não obrigam outra longa rodada de entrevista. O executor deve ser a única decisão exigida antes de um smoke test integrado local; as demais podem ter defaults locais explicitamente registrados, sem segredo commitado.

## 10. Limites claros

- Os agentes podem compartilhar acesso às capacidades, mas ferramenta não deve rodar em todo chamado sem necessidade, pois isso aumenta contexto, custo, permissões e ruído.
- Harness² não é um “agente extra” que todos consultam para desenvolver produto; ele valida alterações no próprio Harness em experimentos separados.
- ECC, Matt Pocock e Anthropic não são plugins intercambiáveis. São fontes distintas de skills, configuração e conhecimento que precisam de curadoria e origem registrada.
- CodeGraph e Graphify podem colaborar através do Harness. Não presumir que um fala diretamente com o outro ou compartilha banco de dados; a ponte é o registry/adapters/context assembler.
- Não há “download único que deixa tudo pronto”: credenciais, autenticação, inicialização de índices, instalação MCP, hooks, configuração de política e validação por subagente compõem o trabalho funcional.
