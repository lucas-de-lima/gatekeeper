# 19 — Integrações e arquitetura de conhecimento

**Responsável pelo produto:** Lucas de Lima  
**Estado:** intenções de integração confirmadas; adapters e compatibilidade precisam de validação por versão  
**Princípio:** conhecimento externo informa os agentes; decisões do produto e fatos do projeto continuam sob governança própria.

## 1. Mapa de capacidades

| Capacidade | Ferramenta/fonte | Papel desejado | Status no desenho |
|---|---|---|---|
| Delegação multiagente | Executor nativo suportado | Execução real dos especialistas e coordenação mestre | Obrigatório; runtime ainda a selecionar/validar |
| Estrutura de código | CodeGraph de Colby McHenry | Símbolos, arquivos e relações internas de código | Requerido no desenho integrado; prova de cobertura/frescor do projeto necessária |
| Grafo amplo de artefatos | Graphify | Relações entre código, documentação, esquemas SQL/config e outros formatos suportados | Requerido no desenho integrado; verificar sobreposição e proveniência |
| Conhecimento persistente auxiliar | ai-memory | Episódios/handoffs e busca entre sessões | Opcional e auxiliar à memória canônica |
| Melhoria do Harness | Harness² e seus métodos | Experimentos rastreáveis e propostas de melhoria | Fonte de métodos/adaptação; não deve ser adotado como executor sem análise de isolamento |
| Publicações Anthropic | Blogs/documentação/papers públicos relevantes | Biblioteca ampla e pesquisável para diferentes especialidades | Fonte externa curada continuamente, sem afiliação e sem autoridade normativa |
| Práticas de agentes e skills | ECC, Matt Pocock, Agent Skills | Procedimentos selecionados, com origem e licença registradas | Curadoria/adaptação; não espelhar coleções completas automaticamente |
| Spec-driven development | GitHub Spec Kit e conceitos correlatos | Artefatos de requisitos, plano, tarefas e aceitação | Conceitos compatíveis; pipeline canônica permanece no Harness |
| Observabilidade | OpenTelemetry GenAI conventions | Traces/eventos e contexto técnico | Base/compatibilidade a avaliar; redigir dados sensíveis |
| Análise de traces | Arize Phoenix | UI e experimentação opcional | Backend opcional; harness utilizável sem ele |
| Gateway de modelos | OmniRoute | Roteamento opcional | Nunca requisito; modelo/provedor direto deve funcionar |

Status de produto precisa usar valores honestos: `implemented`, `configured`, `optional_unconfigured`, `simulated`, `research_only`, `unsupported`. URL ou arquivo de configuração sozinho não é uma integração funcional.

## 2. CodeGraph + Graphify: hipótese de complementaridade

As descrições atuais dos repositórios oficiais sustentam divisão de ênfase: CodeGraph descreve extração local de símbolos/arestas de código e consultas via MCP/CLI/biblioteca; Graphify descreve um grafo que relaciona código com documentação, SQL, configuração e PDFs, com extração determinística em parte dos conteúdos. Ambos têm interseção em código. Assim, não assumir dois índices completos idênticos nem que a junção de respostas será correta por si só.

**Hipótese operacional a validar:**

1. CodeGraph é consultado para navegação de símbolos/calls/imports e trechos de código com referência de origem.
2. Graphify é consultado para relações interartefato e conteúdo amplo que CodeGraph não modela como foco.
3. Consultas de código podem ser comparadas durante a avaliação para medir cobertura, divergências e custo; o caminho normal roteia por intenção para evitar duplicar chamadas.
4. Cada resultado preserva nome do provider, versão/index, commit ou timestamp do repositório, arquivo/linhas, tipo da aresta e distinção extraída/inferida quando o provider a oferece.
5. Em desacordo ou indício de índice desatualizado, verificar o arquivo fonte diretamente. Grafo é navegação/prova auxiliar, não fonte autoritativa do requisito.

**Validações obrigatórias antes de adoção plena:** conjunto amostral de perguntas sobre código e relações documentais; precisão/recall contra evidência revisada; atualização após alteração; exclusões de dependências/gerados; comportamento com monorepo; custo local; integração com o executor escolhido; licenças e cadeia de instalação; funcionamento Windows, Linux e, especialmente, caminhos/índices compartilhados com WSL2. Repositório CodeGraph alerta que compartilhamento de índice/daemon entre Windows e WSL2 pode ter limitações; evitar assumir que um diretório de índice é seguro nos dois ambientes.

O requisito do Lucas é considerar ambas em conjunto porque no uso dele funcionam bem. A integração do Harness preserva essa intenção e mede as fronteiras na prática. A hipótese de roteamento por domínio é recomendação para tornar o uso complementar eficiente, não decisão de excluir uma ferramenta.

## 3. Biblioteca Anthropic

Lucas quer incorporar inteligência pública da Anthropic sobre os assuntos que publica, não só sobre Claude Code ou multiagentes. Interpretar isso como uma **biblioteca de conhecimento atualizável e curada**, não como captura ilimitada e automática de todo o site ou ingestão de texto integral sem avaliação.

Proposta do fluxo:

1. Manter registro de fontes (URL canônica, autor/publisher, título, data publicada/atualizada, acesso, assunto, formato, público, licença/termos conhecidos).
2. Descobrir materiais por áreas relevantes (agentes, contexto, avaliação, segurança, engenharia, programação, produto etc.); incluir artigos, guias, papers e repositórios públicos conforme necessidade.
3. Extrair metadados e resumo próprio com citações curtas/links; guardar texto integral somente se permitido por licença/termos e necessário.
4. Marcar validade temporal e escopo (geral, específico de produto/API, opinião, pesquisa, anúncio).
5. Recuperar no contexto da pergunta do especialista e citar origem/limites; não carregar tudo em todos os prompts.
6. Não deixar recomendação externa mudar decisão aprovada do Lucas ou política do Harness sem processo de decisão próprio.
7. Revarrer periodicamente e versionar adição/substituição; material indisponível não pode ser tratado como atual.

É tecnicamente inviável prometer cobertura completa de tudo o que Anthropic publica em todos os meios, idiomas e tópicos sem uma fonte/API/escopo fechado e atualização contínua. O sistema pode aproximar esse objetivo com catálogo rastreável, descoberta incremental, ranking de relevância e cobertura medida, mantendo a limitação explícita. Esse é um limite de escopo, não um motivo para reduzir a ambição da base de conhecimento.

## 4. Skills e proveniência

Agent Skills define diretórios com um `SKILL.md` e arquivos auxiliares opcionais; a especificação define metadados como `name` e `description`, estrutura e algumas restrições. Tratar isso como formato de interoperabilidade possível, não como garantia de que qualquer executor carrega a skill da mesma forma. Cada executor recebe adapter e teste de carregamento.

Catálogo interno deve registrar `skill_id`, capability, versão, idioma, fonte, licença, hash, executor compatível, ferramentas, riscos, exemplos de uso, conflitos conhecidos e substituições. Uma capacidade tem uma versão canônica escolhida por projeto; outras opções permanecem como referência comparativa, sem serem todas carregadas simultaneamente.

## 5. Memória, harness learning e observabilidade

- **ai-memory** recupera contexto episódico auxiliar; o dado canônico continua nos artefatos do projeto.
- **Harness²** inspira fluxo de hipótese, mudança, avaliação, comparação e adoção. Experimentos de harness são isolados do trabalho comum e não se autopromovem para alterar policy/arquitetura.
- **OpenTelemetry** descreve vocabulário e trace de execução. Instrumentar sem colocar prompt/segredo no span por padrão.
- **Phoenix** pode consultar traces e avaliações; indisponibilidade não bloqueia execução.
- **Testes do projeto/aceite** são a evidência da entrega. Observabilidade não prova correção.

Todos podem referenciar o mesmo `project_id`, `run_id`, `stage_id` e `evidence_id` sem duplicar a autoridade dos dados.

## 6. Modelos e OmniRoute

Separar seleção de modelo do grafo de pipeline: papel pode declarar capacidade requerida, orçamento ou classe de confidencialidade; o provider resolve rota permitida. OmniRoute é opcional. Recurso/modelo indisponível resulta em erro explícito ou rota alternativa aprovada, nunca em troca silenciosa que altere privacidade, custo ou qualidade.

## 7. Política de falha e fallback

| Falha | Comportamento esperado |
|---|---|
| Grafo offline/desatualizado | Inspecionar fonte diretamente; marcar cobertura limitada |
| ai-memory offline | Retomar de manifest/run/handoff canônicos |
| Phoenix/OTel backend ausente | Guardar relatório local estruturado, sem afirmar telemetria exportada |
| OmniRoute indisponível | Usar rota direta se permitida; registrar mudança |
| Fonte Anthropic não acessível | Usar somente material indexado com data/proveniência e registrar lacuna |
| Executor sem subagentes reais | Declarar incompatível; não usar persona sequencial como substituto |
| Skill sem licença/escopo claro | Referenciar ou redigir equivalente próprio até esclarecer uso |

## 8. Referências verificadas nesta pesquisa

Consulta em 2026-11-19 às fontes disponíveis; status e versões podem mudar antes da integração:

- [CodeGraph — repositório oficial](https://github.com/colbymchenry/codegraph): README atual descreve foco em estrutura de código, MCP/CLI, suporte Windows/Linux e nota de cuidado para índice compartilhado entre Windows/WSL2.
- [Graphify — repositório oficial](https://github.com/Graphify-Labs/graphify): descreve código mais documentação, SQL, configuração e PDF; documenta instalação/pré-requisitos, inclusive Python/uv.
- [Agent Skills — especificação](https://agentskills.io/specification): formato e campos de `SKILL.md`.
- [Harness² — setup](https://github.com/google-research/harness2/blob/main/docs/SETUP.md): fonte para inspecionar método, benchmark e executor; avaliar sandbox e dependências antes de reutilizar execução.
- [Anthropic — artigo sobre quando usar subagentes](https://claude.com/resources/articles/building-agents-when-to-fold-subagents-into-your-main-agent-loop): uma opinião metodológica de escopo próprio; não altera a estrutura de pilares do Harness.

Ver `references/SOURCE_CATALOG.md` para bibliografia completa. Licença e termos devem ser inspecionados no upstream específico antes de redistribuir código, conteúdo ou skills.
