# Changelog

Formato: [semver] e seções por data/versão em ordem decrescente.

## [Unreleased]

### Adicionado
- Documento de memória do projeto em `.vscode/instructions/project-memory.md`.
- Changelog inicial em `CHANGELOG.md`.
- Atualização da stack em `docs/Gatekeeper - Architectural Blueprin.md` para Go 1.26, Logrus, Viper, Dig, Validator, Testify/Mockery.
- Atualização de padrões em `docs/Gatekeeper - Engineering Standards.md` com logging e testes novos.

### Planejado
- Implementar setup de Uber Dig e Viper em `cmd/api`.
- Configurar test suite com Testify + Mockery.
- Configurar Prometheus + Grafana na infraestrutura.
- Adicionar OpenTelemetry tracing (pós-MVP).

## [2026-03-19]

- Início da configuração do projeto e definição do roadmap MVP.
