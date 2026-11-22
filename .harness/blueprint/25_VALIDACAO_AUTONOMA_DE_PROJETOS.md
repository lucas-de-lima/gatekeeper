# 25 - Validacao autonoma de projetos

## 1. Objetivo

Definir um procedimento executavel e agnostico para que um agente externo
valide qualquer projeto novo ou legado sem contaminar o blueprint, o repositorio
fonte do Harness ou outro projeto-alvo.

Cada stack entra apenas no runner/adapters locais do projeto. O blueprint nao
codifica ferramentas, testes ou evidencias de uma tecnologia especifica.

## 2. Identidade obrigatoria

A nova sessao abre o repositorio-alvo como workspace primario e le:

1. `.harness/AGENTS.md`
2. `.harness/project-manifest.yaml`
3. `.harness/TARGET_CONTEXT.yaml`, quando existir
4. `.harness/MEMORY.md`
5. `.harness/LOG.md`
6. `.harness/blueprint/`
7. arquivos canonicos da stack do projeto

Antes de executar, confirme Git root, remote, branch, estado do worktree,
`workspace_kind=managed_target`, work mode e versao pinned do Harness.

Sem identificacao suficiente, produza `needs_input` ou `blocked`. Nao inferir
apenas pelo nome da pasta.

## 3. Contrato de isolamento

A copia em `.harness/blueprint/` pertence ao projeto. Ela pode registrar
baseline local, evidencias do projeto e adapters locais.

Ela nao pode alterar diretamente:

- o repositorio central `portable-engineering-harness`
- o blueprint oficial
- outro projeto-alvo
- produtos ou subsistemas nao solicitados

Promover uma evidencia ou padrao ao blueprint oficial exige generalizacao,
remocao de dados/nomes/paths, fixture sintetica, revisao e aprovacao.

## 4. Descoberta de ferramentas

O agente nao assume uma stack ou runner.

Ele verifica de forma read-only: sistema operacional, shell, runtimes,
gerenciadores de dependencias, manifestos de build/test, ferramentas instaladas,
servidor/DB/emulador quando declarados e variaveis por nome sem publicar
segredos.

Os comandos sao definidos pelo projeto. Nao referenciar tecnologia, ferramenta ou comando no documento generico.
Toda familia de stack deve ser declarada pelo projeto local e aprovada pelo
runner da instancia.
- web: build/test/e2e declarado pelo projeto

Se nao houver runner, o agente cria uma proposta local, pede aprovacao para
instalar ferramentas e nao executa acoes irreversiveis.

## 5. Runner deterministic

O runner local deve:

1. criar um `run_id`
2. registrar workspace, work mode, commit e objetivo
3. descobrir ferramentas sem instalar implicitamente
4. registrar versoes
5. executar somente comandos aprovados
6. capturar exit code, stdout, stderr e artefatos
7. preservar evidencias antigas
8. calcular SHA-256
9. persistir checkpoint/handoff
10. terminar em estado explicito

O runner coleta fatos. O agente interpreta resultados e separa fatos observados
de hipoteses.

## 6. Evidencias persistentes

Cada execucao gera no alvo:

- `.harness/runs/<run-id>/run.json`
- `.harness/runs/<run-id>/events.jsonl`
- `.harness/runs/<run-id>/checkpoint.json`
- `.harness/runs/<run-id>/evidence-index.json`
- `.harness/runs/<run-id>/handoff.md`

Artefatos adicionais ficam em `artifacts/` e precisam ser indexados.
Atualizacoes posteriores preservam `.harness/MEMORY.md`, `.harness/LOG.md`,
baseline local e historico de runs.

Uma conversa ou resumo do agente nao substitui esses arquivos.

## 7. Classificacao do resultado

Usar:

- `passed`
- `failed`
- `observed`
- `blocked`
- `needs_input`
- `not_applicable`
- `unknown`
- `completed`

Regras:

- ausencia de runtime = `blocked`, nao falha da aplicacao
- teste nao executado = `not_run`/`blocked`, nao `passed`
- erro do projeto = `failed`
- erro de infraestrutura = `blocked`
- resultado ambíguo = `unknown`
- onboarding concluido sem implementacao = `completed`

Nunca omitir stderr, exit code, versao ou limitacao.

## 8. Revisao independente

Para mudanca de codigo:

- especialista implementa
- revisor sem write scope verifica diff e evidencias
- QA verifica testes
- Master coordena sem substituir especialistas
- findings bloqueantes impedem encerramento

Somente quando aplicavel: arquitetura autoriza fronteiras, Lucas autoriza
reopening reservado, homologacao externa autoriza release e producao permanece
fora do scope normal do agente.

## 9. Estados finais

Um run termina em `completed`, `failed`, `blocked`, `needs_input` ou
`waiting_for_request`.

Nenhuma feature nova deve ser criada por inferencia. Entrega exige pedido
explicito, baseline aprovado, arquitetura compativel, test-first, testes,
review independente e gate apropriado.

## 10. Criterio de operacionalizacao

O sistema esta apto quando possui identificacao explicita do workspace, runner
local persistente, schemas de run/checkpoint/evidence, discovery agnostico,
instalacoes opt-in e pinned, write scope isolado, memoria local duravel,
review independente e promocao explicitamente revisada.
