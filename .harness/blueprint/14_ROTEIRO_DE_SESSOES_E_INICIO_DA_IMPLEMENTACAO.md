# 14 — Roteiro de sessões de concepção e início da implementação

**Responsável pelo produto:** Lucas de Lima  
**Estado:** roteiro recomendado; nenhuma implementação do harness começou  
**Base obrigatória:** ler `13_LUCAS_PRODUCT_VISION_AND_DECISIONS.md` e os artefatos já produzidos antes de cada sessão

## 1. Objetivo

Conduzir a concepção restante por assuntos coesos e decisões que afetam a arquitetura. Cada sessão deve produzir um artefato durável, registrar decisões e dependências, e atualizar o resumo de decisões. Não abrir uma sessão para cada ferramenta ou pergunta pequena.

O roteiro preserva a exigência de uma primeira versão integrada e de alta capacidade. A sequência organiza dependências; não autoriza reduzir os quatro pilares a uma demonstração ou excluir requisitos apenas para começar a codificar mais cedo.

## 2. O que já está feito

- A visão, decisões confirmadas, limites, correções e pendências estão registradas em `13_LUCAS_PRODUCT_VISION_AND_DECISIONS.md`.
- A autoridade do produto é Lucas de Lima.
- A estrutura do produto já está decidida em alto nível: agente mestre, especialistas delegáveis, quatro pilares, pipelines adaptáveis, memória durável e preparação para homologação.
- O código do harness ainda não foi iniciado.

## 3. Sessões recomendadas

### Sessão 1 — Runtime e execução multiagente

**Objetivo:** decidir como o agente mestre e os especialistas reais serão executados no Windows e Linux, e qual runtime inicial atende à exigência de delegação verdadeira.

**Cobrir:** agente mestre; criação e ciclo de vida de subagentes; troca de artefatos/contexto; execução paralela ou sequencial; limites do executor; adaptação Windows/Linux/WSL2; matriz de capacidades.

**Artefato de saída:** arquitetura de execução, matriz de runtimes suportados e contrato do adapter de executor.

**Condição para concluir:** existe pelo menos um caminho concreto para executar os agentes especialistas desde a primeira versão. Não usar troca de personas como fallback para subagentes indisponíveis.

### Sessão 2 — Agentes especialistas, quatro pilares e composição de pipelines

**Objetivo:** transformar a hierarquia organizacional em papéis, responsabilidades, interfaces e pipelines compostas.

**Cobrir:** responsabilidades do mestre; agentes candidatos de arquitetura, engenharia, backlog e desenvolvimento; escalonamento entre camadas; pipeline mestre; execução de escopo por feature/épico; geração, validação, versionamento e descarte de pipelines temporárias; promoção de pipeline reutilizável ao nível de projeto ou upstream.

**Artefato de saída:** mapa de agentes/responsabilidades e modelo de pipeline (estágios, entradas, saídas, gates, transições, erro, retry, conclusão e proveniência).

**Condição para concluir:** o modelo permite criar pipelines novas sem permitir que elas contornem contratos, limites de arquitetura ou gates aplicáveis. A execução de uma pipeline temporária continua rastreável depois que ela deixa de ser reutilizada.

### Sessão 3 — Ciclo de vida do projeto e autoridade das decisões

**Objetivo:** especificar os modos de instalação e operação, os checkpoints humanos e o fluxo de homologação.

**Cobrir:** toggle simples de instalação/análise versus operação normal; verificação contra manifest, artefatos e timestamps; retomada de instalação incompleta; onboarding de projeto novo e legado; modo de instalação de legado sem pedido de transformação; baseline as-is; aprovação e imutabilidade da arquitetura; escalonamento para reabrir decisão; HITL; preparação de homologação; GMUD opcional por política de cliente; produção fora do limite atual.

**Artefato de saída:** máquina de estados do projeto/run e política de autoridade, aprovações, escalonamento e homologação/GMUD.

**Condição para concluir:** instalação repetida não sobrescreve trabalho; análise de legado pode terminar sem modificar o sistema; nenhuma alteração arquitetural é feita sem autorização de Lucas; nenhum deploy de produção é atribuído ao harness.

### Sessão 4 — Artefatos, memória, logs e retomada

**Objetivo:** definir como o sistema sabe o estado do projeto e continua uma tarefa entre sessões sem depender da conversa completa.

**Cobrir:** memória canônica do projeto; registros de arquitetura e engenharia; backlog; logs de tarefa/iteração com data, resultado, evidência e próxima ação; escopo de leitura/escrita dos agentes; handoff; arquivos versionados versus estado local; grandes portfólios e eventual repositório central de documentação; integração de ai-memory sem duplicar a fonte canônica.

**Artefato de saída:** modelo de dados e layout de armazenamento, schemas dos principais artefatos, fingerprints de retomada, regras de atualização e retenção.

**Condição para concluir:** outra sessão/agente consegue retomar usando artefatos e memória referenciada; um arquivo `memory.md` sem limites não é a única fonte de continuidade.

### Sessão 5 — Conhecimento e integrações

**Objetivo:** fechar o papel, as fronteiras e os contratos das ferramentas e coleções de conhecimento.

**Cobrir:** CodeGraph de Colby McHenry; Graphify; uso complementar e controle de indexação duplicada; proveniência e atualização; ai-memory; Harness²; biblioteca consultável de publicações Anthropic; curadoria de ECC e Matt Pocock; ideias/artefatos de Spec Kit; OpenTelemetry e Phoenix; provedor de modelo direto e OmniRoute opcional.

**Artefato de saída:** mapa de integrações e fluxo de contexto, contrato por provider, matriz de estado real (implementado/configurado/opcional/não suportado) e registro versionado das fontes externas.

**Condição para concluir:** cada integração tem uma função definida, dados de entrada/saída, dependências, tratamento de falha e limites de confiança. Conteúdo de fontes externas não se torna política do produto e conteúdo inferido não se torna fato confirmado.

### Sessão 6 — Segurança, dependências e portabilidade

**Objetivo:** transformar a autonomia desejada em política operacional e definir distribuição/atualização com risco controlado.

**Cobrir:** acesso a arquivos, rede e credenciais; risco de prompt injection; política de dependência confiável; proveniência, licença e integridade de releases; retenção/redação de logs; escopo de escrita dos agentes; comandos e efeitos externos; Windows/Linux/WSL2; instalação, atualização upstream, preservação de personalização e recuperação.

**Artefato de saída:** política de confiança/dependências, capability profile do executor e procedimento de distribuição e upgrade.

**Condição para concluir:** a automação permanece alta para trabalho comum, mas não confunde popularidade de pacote com segurança, não expõe segredos e não consegue mudar a arquitetura aprovada sem a autorização estabelecida.

### Sessão 7 — Arquitetura integrada, critérios de aceite e plano de implementação

**Objetivo:** integrar os artefatos anteriores, fechar dependências críticas e autorizar o início do desenvolvimento do harness.

**Cobrir:** arquitetura de componentes e deployment; interação mestre/pilares/agents/providers; fluxos novos e legados; pipeline generator; contratos e APIs locais; estratégia de observabilidade; plano de implementação integrado; cenário sintético ou anonimizado de validação; critérios de aceite por pilar e para o E2E completo.

**Artefatos de saída:** arquitetura integrada aprovada, ADRs das decisões críticas, backlog de implementação rastreável aos critérios de aceite, plano de validação e lista explícita de limitações conhecidas.

**Gate de início do código:** os artefatos determinam um runtime multiagente real, contrato de pipeline adaptável, autoridade arquitetural, estado/memória recuperável, fronteira de produção e critérios de aceite. Questões opcionais podem seguir como trabalho de integração sem bloquear a base, mas não se deve fingir que uma integração foi implementada se ainda está em aberto.

## 4. Quando começa a implementação

**A implementação de software do harness começa após a Sessão 7**, quando a arquitetura integrada e seus contratos estiverem aprovados por Lucas. A conversa e a documentação de concepção já começaram; o desenvolvimento do produto ainda não.

O primeiro trabalho de código deve ser uma fatia integrada que comprova a arquitetura escolhida: invocação mestre, delegação real a especialistas, composição/validação de pipeline, registro de artefatos e retomada, com os quatro pilares representados no fluxo. Isso não é um protótipo que reduz os requisitos; é a fundação executável do sistema completo. O trabalho prossegue em incrementos internos até entregar os critérios de aceite acordados, sem transformar cada incremento em uma nova rodada de aprovação do usuário.

Integrações precisam ter estado real e declarado. Se uma fonte estiver em pesquisa ou depender de credencial/serviço não configurado, marcar essa condição precisamente e continuar os itens independentes; não representar configuração de placeholder como integração concluída.

## 5. Regra de saída de cada sessão

Cada sessão deve:

1. Começar lendo `13_LUCAS_PRODUCT_VISION_AND_DECISIONS.md` e os artefatos anteriores relevantes.
2. Tratar como fixas as decisões já confirmadas; abrir apenas conflitos reais ou informação que altere materialmente a solução.
3. Entregar ao menos um artefato durável, adequado ao assunto (ADR, modelo, schema, matriz ou plano).
4. Registrar decisões, justificativas, dependências e perguntas realmente abertas.
5. Atualizar o registro de visão/decisões e a manifest/readme do pacote quando houver mudança durável.
6. Não iniciar implementação de código antes do gate da Sessão 7.

As sessões não são reuniões de confirmação de cada componente. Cada uma aborda uma área inteira, produz resultado revisável e deixa as próximas com contexto persistido.

## 6. Progresso documental em 2026-11-19

Foram criados os documentos 15–21 como drafts operacionais para avançar as sessões 1–7 sem obrigar Lucas a repetir a visão. A pesquisa atual encontrou documentação oficial de multiagentes (OpenAI Responses API, indicada como beta na página consultada), subagentes Claude Code, CodeGraph, Graphify, Agent Skills e Harness². Isso embasa os contratos, não encerra a validação prática de adapters.

Estado por tema:

- **Sessão 1:** contrato de execução e matriz em documento 15; executor inicial e prova Windows/Linux continuam abertos.
- **Sessão 2:** mapa de pilares e contratos de composição em documento 16; roster final/schema continuam abertos.
- **Sessão 3:** draft de estados e autoridade em documento 17; política específica de cliente para homologação/GMUD segue configurável.
- **Sessão 4:** modelo proposto de artefatos/memória em documento 18; formatos, retenção e ambiente de portfólio seguem abertos.
- **Sessão 5:** mapa de integrações e fonte primária atualizados em documento 19 e `references/`; complementaridade dos grafos ainda precisa de ensaio com dados representativos.
- **Sessão 6:** baseline adaptativo proposto em documento 20; fornecedores, política de dados e limites por executor requerem validação.
- **Sessão 7:** plano integrado e critérios em documento 21; thresholds, ambiente real de validação e escolhas do runtime requerem decisão/teste.
- **Amarração das integrações:** documento 22 define capability catalog, acesso de subagentes, proveniência, ondas de instalação e cenários de colaboração. Ainda falta selecionar/validar executor e executar os smoke tests; clonar fontes sem isso não conclui a integração.

Portanto, a documentação avançou substancialmente, mas as sessões 1–7 não devem ser marcadas como aprovadas apenas porque há um draft. O primeiro código ainda não começou.

## 7. Identidade do build e plano até release

Lucas confirmou explicitamente que o objetivo atual é construir o Harness, não usar um Harness já pronto para desenvolver o sistema privado ou outro sistema-alvo. Para essa distinção foram adicionados `23_IDENTIDADE_DO_WORKSPACE_E_MODOS_DE_TRABALHO.md`, `24_PLANO_MESTRE_DE_CONSTRUCAO_DO_HARNESS.md` e `BLUEPRINT_STATUS.yaml`.

O status atual é `workspace_kind=blueprint_package`, `product_status=under_construction`, `implementation_status=not_started`, `target_system_active=false`. O primeiro work item de código, depois de escolher destino e runtime, é `H0-01`: criar o repositório executável do Harness separado deste pacote e ligá-lo ao blueprint. O agente de cada sessão deve ler o estado, o plano e o run/handoff e continuar o próximo item desbloqueado. Release só pode ser marcada pronta após H0–H8, AC-01–AC-15, evidências globais e decisão explícita de Lucas.

