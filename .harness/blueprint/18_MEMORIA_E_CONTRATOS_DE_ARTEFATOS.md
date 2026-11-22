# 18 — Memória canônica, eventos e contratos de artefatos

**Responsável pelo produto:** Lucas de Lima  
**Estado:** princípios confirmados; layout e schemas abaixo são propostas de trabalho  
**Fonte canônica:** artefatos de projeto versionados, junto da arquitetura e decisões aprovadas.

## 1. Quatro categorias de estado

| Categoria | O que contém | Fonte de autoridade | Duração |
|---|---|---|---|
| Memória canônica | Baselines de arquitetura, decisões, engenharia acordada, contexto de produto aprovado, convenções do projeto | Repositório do projeto ou repositório central de documentação explicitamente vinculado | Persistente e versionada |
| Backlog e plano | Épicos, features, histórias, tarefas, dependências, critérios e estado | Artefatos do projeto; integração de ferramenta pode ser espelho | Persistente e versionada |
| Registro de execução | Runs, etapas, agentes, eventos, tentativas, decisões de retomada, evidências e próximas ações | Log do Harness com referências aos artefatos | Persistente conforme retenção definida |
| Memória auxiliar episódica | Handoff compacto e recuperação contextual entre sessões | Provider opcional como ai-memory, indexado por projeto/run | Auxiliar; reconstruível e não autoritativa |

Um agente sempre consulta os artefatos canônicos relevantes antes de usar lembrança episódica. Em divergência, a fonte canônica versionada prevalece; a divergência é registrada para revisão. Memória do modelo/conversa não é armazenamento durável.

## 2. Layout proposto

```text
.harness/
  manifest.yaml                 # ID, upstream/schema, executor e localização dos artefatos
  project-profile.yaml          # fatos confirmados, stack, escopo e políticas locais
  pipelines/                    # templates de projeto aprovados
  agents/                       # especializações de projeto (overrides identificados)
  skills/                       # procedimentos selecionados e de proveniência conhecida
  architecture/
    problem-brief.md
    alternatives/
    baseline.yaml
    decisions/
    as-is/                      # quando legado
  engineering/                  # contratos e decisões técnicas derivados
  backlog/                      # hierarquia e vínculos de escopo
  quality/                      # planos, critérios e evidências de validação
  governance/                   # homologação, GMUD, aprovações e pacotes
  runs/<run-id>/
    run.yaml
    events.jsonl
    pipeline-instance.yaml
    delegations/
    evidence-index.yaml
    handoff.md
  cache/                         # ignorado do Git: índices, sessões locais, temporários
```

O layout é exemplo lógico, não prescrição de arquivos finais. Em portfólio de muitos sistemas, cada repositório pode apontar para repositório de documentação central. O vínculo registra URI/ID, versão/commit, caminho, política de escrita e como validar correspondência. Não se deve duplicar e manter duas fontes canônicas concorrentes.

## 3. Versionamento e sensibilidade

Versionar o que explica o sistema, decisões, requisitos, backlog aprovado, definições de pipeline reutilizáveis, resultados necessários para auditoria e handoffs que orientam trabalho aberto. Ignorar caches, índice local, sockets, sessões efêmeras, logs de depuração com dados sensíveis e binários regeneráveis. Um manifesto declara cada classe como `tracked`, `local_only`, `external_reference` ou `generated_ephemeral`.

Logs evitam prompts inteiros e segredos por padrão. Guardam sumário, IDs/versões de modelo/executor, ferramentas invocadas, status, hashes, caminhos e resultados estruturados. Se texto bruto for necessário para auditoria, política do projeto define acesso, redação e retenção.

## 4. Contratos mínimos

### Project manifest

`project_id`, nome, `harness_version`, `schema_version`, `created_at`, perfil de plataforma, estado de instalação/análise, referências de memória canônica, política de sensibilidade, estado de arquitetura, localização de logs e integrações configuradas/não configuradas.

### Run record

`run_id`, `project_id`, `request_id`, `mode`, `created_at`, `updated_at`, `status`, `pipeline_definition_ref`, `pipeline_instance_id`, `architecture_baseline_ref`, executor e versão, escopo, approvers, próximo passo e bloqueios.

### Pipeline instance

Versão/resumo da definição usada, grafo de etapas selecionadas, justificativa de seleção/omissão, entradas por referência/hash, outputs, roteamento, limites, políticas e estado por etapa. Congelar instância no início; mudança posterior cria revisão vinculada.

### Event

Evento append-only com UUID, tempo RFC 3339, ator, run, etapa/delegação, tipo, payload estruturado, referências, causalidade (`caused_by`) e proveniência. Segredos e conteúdo não confiável têm classificação própria.

### Evidence item

`evidence_id`, tipo (arquivo, comando, teste, review, aprovação, fonte), localização, hash/commit quando aplicável, ferramenta e versão, instante, resultado, limitações e relação com requisito/critério de aceite.

### Handoff

Objetivo e estado atual; o que foi concluído e verificado; o que ficou em aberto; bloqueio/decisão aguardada; referências canônicas; alterações desde o checkpoint; próxima ação recomendada; agente/camada responsável; timestamp.

## 5. Proveniência epistêmica

Todo conhecimento relevante distingue:

- `observed`: obtido diretamente de arquivo, ferramenta, teste ou pessoa identificada;
- `inferred`: conclusão possível, com fontes e confiança/racional;
- `proposed`: desenho/recomendação ainda não aprovada;
- `human_confirmed`: resposta ou decisão explicitamente confirmada por Lucas/autoridade cadastrada.

O agente não rebaixa fato confirmado a hipótese nem eleva inferência a fato por repetição. Aprovação de uma baseline identifica exatamente quais artefatos e hashes/versões foram aprovados.

## 6. Retomada entre sessões/agentes

1. Resolver `project_id` e repositório canônico.
2. Ler manifesto, estado, run ativo mais recente e handoff.
3. Validar refs/hashes e detectar mudança do repositório/arquitetura.
4. Consultar apenas os artefatos relacionados ao próximo trabalho.
5. Consultar memória episódica para pista adicional e conferir cada pista com fonte canônica.
6. Registrar evento de retomada, agente novo e próximo passo.

Run concluído não some. Run interrompido mantém tentativa e artefatos; uma retomada ou novo pedido se liga ao anterior sem reescrevê-lo.

## 7. ai-memory e fonte canônica

ai-memory pode recuperar episódios, preferências persistentes ou handoffs entre sessões, desde que o mecanismo suporte escopo por projeto, proveniência, deleção/expurgo e política de privacidade necessária. Não deve ser necessário para reconstruir baseline, aprovação ou estado de pipeline. Se ficar indisponível, os artefatos canônicos e o handoff persistido mantêm a continuidade. A integração não deve armazenar indiscriminadamente todo conteúdo do repositório.

## 8. Validação dos artefatos

Cada schema tem identificador e versão; leitura aceita versões anteriores suportadas ou produz erro explicativo com migração proposta. Validar referências internas, timestamps, enumerações, evidência de gate, IDs e integridade de hash. Escritas devem ser atômicas sempre que possível e deixar backup/recuperação para atualização de estado. Gerar índice de artefatos para humano, evitar arquivo monolítico de memória.

## 9. Decisões abertas

Retenção temporal e imutabilidade física; assinatura/autoria dos registros; formato e tecnologia de schemas; escolha entre JSON/YAML/Markdown por classe; layout final; integração com documentação central em portfólios; política de privacidade por cliente. Esses itens não alteram a definição de memória canônica como artefatos persistentes do projeto.
