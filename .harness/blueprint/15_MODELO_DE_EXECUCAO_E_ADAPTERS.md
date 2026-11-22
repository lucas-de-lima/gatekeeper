# 15 — Modelo de execução e adapters de agentes

**Responsável pelo produto:** Lucas de Lima  
**Estado:** desenho técnico proposto; selecionar e validar executores concretos ainda é uma decisão aberta  
**Precedência:** este documento aplica as decisões confirmadas em `13_LUCAS_PRODUCT_VISION_AND_DECISIONS.md` e substitui recomendações incompatíveis dos documentos anteriores.

## 1. Decisão de produto que orienta este desenho

O Harness precisa de um agente mestre com responsabilidade de gestão e coordenação e de agentes especialistas reais, invocáveis como execuções delegadas. O mestre não substitui especialidades com perfis, personas ou mudança de prompt. Um executor que não permita delegar trabalho a agentes separados não satisfaz o requisito de produto.

“Agente especialista” é um papel operacional com instruções, contexto, ferramentas, limites, entradas e entregas explícitos. “Pilar” é uma área funcional que pode conter muitos especialistas e muitas etapas. A separação exata em processos, sessões, threads ou workers depende do executor, desde que a fronteira de contexto e a delegação sejam reais e observáveis.

## 2. Camadas do modelo

1. **Entrada do executor:** recebe a solicitação, identifica o projeto e inicia o agente mestre.
2. **Agente mestre:** reconhece instalação/análise inicial ou operação, lê o estado persistido, escolhe ou compõe uma pipeline e coordena os pilares.
3. **Agentes especialistas:** executam trabalhos delimitados em suas especialidades e devolvem resultado com evidências, artefatos e estado.
4. **Estado durável do Harness:** guarda definições versionadas, instâncias de pipeline, eventos de execução, artefatos, aprovações e referências para retomada.
5. **Adapters de capacidade:** traduzem invocação, espera, mensagens, ferramentas disponíveis e resultados para as capacidades nativas de cada executor.
6. **Ferramentas externas:** grafo, memória auxiliar, modelo, observabilidade e verificação. São acessadas por contratos e não são o orquestrador.

Um helper determinístico pode validar schemas, criar diretórios, calcular hashes, consultar estado ou validar transições. Não deve escolher arquitetura, distribuir autoridade, inventar resultado de agente ou substituir o mestre.

## 3. Contrato mínimo do executor

Um executor é suportado somente se seu adapter comprovar, por documentação atual e um ensaio de integração, todos os requisitos obrigatórios:

| Capacidade | Obrigatória | Evidência exigida |
|---|---:|---|
| Iniciar agente mestre com instruções próprias | Sim | Entrada nativa e identificável |
| Iniciar especialista como execução delegada distinta | Sim | Identidade/ID separado e instrução de papel própria |
| Retornar, aguardar e associar saída à delegação | Sim | Resultado terminal ou falha observável |
| Compartilhar referências de artefatos do projeto | Sim | Caminho/ID estável, sem depender de conversa anterior |
| Delimitar ferramentas/contexto quando oferecido pelo host | Sim para declarar suporte; não para requisito universal de produto | Matriz explícita de capacidade e limitações |
| Tratar erro, timeout e cancelamento sem declarar sucesso | Sim | Estado de falha e possibilidade de retomada |
| Operar no sistema operacional declarado | Sim | Versão, comandos e limitações testados |
| Agentes com chamadas aninhadas | Desejável; não presumir | Medir profundidade e limite de concorrência |
| Paralelismo | Desejável; execução serial continua válida | Limites do host e de edição compartilhada |

Se a ferramenta do host não oferece uma capacidade, o adapter deve reportar `unsupported` ou `degraded` e o master deve parar quando a capacidade for requisito. Nunca simular delegação com uma segunda persona dentro da mesma execução.

## 4. Ciclo de delegação

Para cada chamada, o mestre grava uma tarefa de delegação antes de iniciar o especialista. A tarefa contém `run_id`, `pipeline_instance_id`, `stage_id`, `agent_role_id`, objetivo, contexto mínimo, referências de entrada, limites, entregáveis, prazo/timeout e modo de escrita. O adapter devolve `executor_session_id` e estado inicial.

O mestre aguarda ou recebe notificação, valida o resultado contra o contrato da etapa e grava uma conclusão de delegação. Resultado textual sem evidência/entregáveis não satisfaz automaticamente a etapa. Estados mínimos: `queued`, `running`, `succeeded`, `needs_input`, `failed`, `timed_out`, `cancelled`, `rejected`. Repetição deve gerar nova tentativa vinculada à anterior, sem apagar a falha original.

Agentes podem trabalhar em paralelo quando as tarefas forem independentes e os destinos de escrita forem isolados. Escritores concorrentes no mesmo artefato precisam de lock, branch/worktree ou execução serial. Paralelismo é mecanismo de execução, não requisito de todas as pipelines.

## 5. Perfil do agente mestre

O mestre é gestor de entrega: identifica o objetivo, o estado e o próximo trabalho; compõe a sequência; delega aos especialistas; monitora dependências e limites; sintetiza evidências; coordena Human-in-the-Loop; mantém o estado recuperável; e faz escalonamentos conforme a autoridade definida.

O mestre não se apresenta como especialista técnico universal. Ele pode compreender os relatórios para coordenar e apontar inconsistências, mas encaminha julgamento especializado aos agentes pertinentes. Não aprova uma reabertura arquitetural em nome de Lucas nem realiza deploy de produção.

## 6. Candidatos de runtime e método de escolha

Documentação oficial consultada em 2026-11-19 mostra dois padrões distintos, ainda não uma escolha aprovada para o produto: a API multiagente da OpenAI oferece delegação coordenada e concorrência configurável, mas a documentação consultada a classifica como beta; Claude Code documenta agentes personalizados com contexto, ferramentas e permissões próprios. Essas referências demonstram que há modelos reais de delegação, mas não comprovam por si só suporte operacional completo do Harness em Windows/Linux, disponibilidade de credenciais, custo, privacidade, durabilidade de sessão ou compatibilidade com todos os requisitos.

O primeiro adapter deve ser escolhido por uma matriz prática: execução em Windows nativo; execução Linux; suporte a agentes com identidade/contexto próprios; política de ferramentas e aprovações; transporte de artefatos; retomada após reinício; execução paralela; configuração portátil; licença e custo; comportamento com projetos privados; manutenção da API. Não escolher somente pelo nome “multi-agent”.

**Proposta:** adotar arquitetura de adapters e fazer uma verificação comparativa curta entre os executores que Lucas pretende usar, antes de iniciar a implementação do runtime. Um adapter inicial não reduz o número de agentes nem fecha a porta aos demais. A lista de executores prioritários permanece aberta, pois o produto deve ser validado em ambiente real antes dessa decisão.

## 7. Matriz de capacidade por adapter

Cada executor/versionamento deve declarar:

| Campo | Valores/descrição |
|---|---|
| `executor_id`, versão e plataforma | Produto, release, sistema operacional e arquitetura |
| `native_subagents` | `yes`, `no`, `partial`, com prova/documentação |
| `subagent_context_isolation` | Tipo e limites conhecidos |
| `nested_delegation` | Suporte, limite ou não documentado |
| `parallelism` | Máximo configurado/observado e limitações de escrita |
| `tool_inheritance` | Herança, restrição e diferenças por agente |
| `filesystem_scope` | Acesso comum, isolamento, permissões |
| `session_resume` | Como recuperar execuções e IDs |
| `network_and_approval` | Controles do host e diferenças por plataforma |
| `known_gaps`, `last_verified_at` | Lacunas, versão e data da validação |

O manifesto registra capacidade observada, não uma promessa futura. Revalidar quando o executor atualizar APIs, modelos, políticas ou sistema operacional.

## 8. Critérios de aceite do adapter

O ensaio deve demonstrar: o mestre inicia pelo menos dois papéis distintos; cada especialista recebe apenas a instrução/contexto destinado a ele; um resultado de falha não é tratado como sucesso; um evento fica registrado; o mestre retoma a coordenação após receber os resultados; limites de escrita funcionam como declarados; e o caminho opera nas plataformas reivindicadas. Registrar custo e latência como medições, não como garantias.

## 9. Referências de base

- OpenAI, [Multi-agent (Responses API)](https://developers.openai.com/api/docs/guides/responses-multi-agent) — coordenação, árvore de agentes, concorrência, eventos e condições de uso; API em beta segundo a página consultada.
- Anthropic, [Create custom subagents (Claude Code)](https://code.claude.com/docs/en/sub-agents) — contexto próprio, prompts, ferramentas, permissões e execução/retomada.
- Ver também os links e ressalvas registrados em `references/SOURCE_CATALOG.md`. Compatibilidade precisa ser revalidada antes da integração.
