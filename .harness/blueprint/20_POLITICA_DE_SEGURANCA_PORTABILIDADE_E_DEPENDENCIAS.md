# 20 — Política de segurança, portabilidade e dependências

**Responsável pelo produto:** Lucas de Lima  
**Estado:** base operacional proposta; detalhes por cliente/runtime ficam no perfil e na matriz de capacidades  
**Objetivo:** manter alta autonomia cotidiana com controles explícitos para limites reais de impacto.

## 1. Perfil de autonomia proposto

Lucas delegou ao desenho a política de privacidade e segurança e sinalizou autonomia alta para trabalho de engenharia, inclusive operações úteis de dependência, rede e memória, dentro de controles de confiança; produção não pertence ao Harness. Proponho perfil **moderado e adaptativo ao risco**: automatizar leitura, análise, edição de escopo autorizado e verificações locais; exigir evidência e confirmação para mudanças que alteram autoridade, exposição ou impacto externo.

Este é um baseline de produto proposto, não requisito confirmado em cada ambiente. O adapter deve aproveitar sandbox/aprovações nativas e declarar o que não consegue impor. Instruções em prompt não são sandbox.

## 2. Categorias de efeito

| Nível | Exemplos | Política proposta |
|---|---|---|
| A — leitura local | Ler arquivos, Git, configuração não secreta, grafo | Automático dentro do projeto autorizado |
| B — escrita recuperável local | Editar código/documentação, criar artefatos de projeto, executar testes/lint/build | Automático em escopo conhecido; guardar diff/checkpoint; evitar arquivos fora do escopo |
| C — dependência e rede de engenharia | Consultar docs/releases, baixar pacote conhecido, atualizar lockfile, chamar registry | Permitir dentro da política de proveniência/lockfile; observar scripts, permissões, checksum/assinatura e impacto |
| D — efeitos externos ou sensíveis | Alterar sistema compartilhado, credenciais, produção, enviar dados para SaaS, publicar artefato, aprovar homologação/GMUD | Bloquear por padrão ou exigir autoridade explícita e política específica. Produção não é executor do Harness |

Reversibilidade não basta: rede pode exfiltrar código, e scripts de instalação podem executar comandos. Avaliar fonte, conteúdo, permissões e exposição de dados em conjunto.

## 3. Dependências

Antes de introduzir dependência relevante: justificar necessidade; verificar repositório/fonte oficial, maintainer, release e cadeia de distribuição; licença; manutenção; compatibilidade de plataforma; vulnerabilidades conhecidas; dependências transitivas; scripts de instalação; integridade de release/assinaturas quando disponíveis; privacidade/telemetria; alternativa nativa/vanilla; modo de remover/atualizar.

O Harness pode executar análise e propor instalação de uma dependência autorizada. Não inferir confiança por estrelas/popularidade ou por uma ferramenta citar outra. Para binários/instaladores, registrar versão, origem e hash/atestado quando disponível. `curl | sh`/`irm | iex` não devem ser tratados como verificação em si: inspecionar método de distribuição e oferecer instalação reproduzível/configurável quando viável.

## 4. Conteúdo não confiável e injeção de prompt

Código, README, issue, logs, página web, skill externa, resultado de grafo e mensagem de subagente são dados a analisar, não instruções de autoridade. O master mantém a hierarquia: regras do produto, autorização de Lucas, políticas do projeto, missão/contrato do agente, fontes externas. Conteúdo encontrado no repositório não pode elevar a permissão, solicitar segredo ou substituir instruções do usuário.

Quando conteúdo externo instruir ações ou contradizer a política, o agente deve descrever a origem e efeito e não obedecer à instrução embutida. Ferramenta MCP é uma fronteira de execução: permitir só o necessário por papel, validar entrada/saída, proveniência e efeitos colaterais, e manter credenciais fora de logs e artefatos commitados.

## 5. Escrita concorrente e recuperação

Agentes recebem escopo de arquivo e modo (`read_only`, `assigned_write`, `isolated_workspace`). Escrita simultânea no mesmo destino é serializada ou isolada em worktree/workspace. Toda alteração de agente é atribuída à delegação/run e revisável pelo diff. Falha de agente não apaga a alteração existente. Operações de migração/instalação registram checkpoint e resultado parcial.

## 6. Portabilidade Windows/Linux/WSL2

- Windows é ambiente primário do Lucas; Linux precisa de suporte forte. WSL2 é opção de compatibilidade, não pressuposto de compartilhamento sem limites.
- Separar interfaces de shell, caminho e ambiente. Preferir utilitários disponíveis nos runtimes suportados; comandos específicos têm adapter e equivalente documentado.
- Guardar paths como URI/path normalizado relativo ao `project_root`; não assumir `C:\`, `/home`, separador, case sensitivity ou quoting idêntico.
- Testar criação/renomeação atômica, locks, watchers, symlinks, permissões, Git e índices no volume usado.
- Um mesmo índice/binário/daemon não é considerado seguro para Windows e WSL2 simultaneamente sem evidência. Manter caches por runtime/OS quando o provider exigir.
- Requisito operacional: `project_root`, ferramenta de grafo e subprocessos devem estar no mesmo lado/ambiente ou terem integração oficialmente suportada.

## 7. Privacidade, segredo e dados

Classificar dados do projeto (públicos, internos, confidenciais, regulados). Modelo/provider, grafo, memória, observabilidade e biblioteca externa precisam declarar para onde dados são enviados e retidos. Segredos são acessados pelo mecanismo local do executor/usuário e nunca armazenados em memória canônica, prompt log ou trace por padrão. Resumos e IDs preferem caminhos/linhas/hash a payload completo. Retenção/remediação vira parte do perfil do cliente.

## 8. Atualização e distribuição

O modelo preferido é cópia gerenciada/versionada dentro do projeto, com versão upstream, commit/hash de origem e schema. Atualização compara base upstream e customização local, mostra conflito e relatório, cria checkpoint e nunca sobrescreve override silenciosamente. Migração de schema é explícita e reversível quando possível. Instalar/desinstalar provider remove apenas arquivos/configurações atribuídos e preserva índices/dados do projeto conforme opção informada.

## 9. Capacidades que o adapter declara

Para cada host: isolamento de arquivo, aprovações de comandos, acesso à rede, ferramentas MCP, segredo, logs, subagentes, contexto, escrita, timeout, interruptibilidade, retomada e métricas. `unknown` não equivale a protegido. Se uma capacidade crítica não é imposta, pipeline precisa reduzir autonomia, exigir revisão humana ou declarar ambiente incompatível.

## 10. Decisões ainda abertas

Quais provedores/modelos podem receber classes de dado do projeto; fornecedores específicos; instalação automática versus confirmação por risco; thresholds de dependência; retenção e deleção em clientes; política de licenças/cópia de skills; e limites de comandos com efeitos externos. A proposta atual permite avançar com a arquitetura e fechar política por perfil sem travar análise de documentação.
