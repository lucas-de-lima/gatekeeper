# Changelog

Formato: [semver] e seções por data/versão em ordem decrescente.

## [Unreleased]

### Adicionado
- Documento de memória do projeto em `.vscode/instructions/project-memory.md`.
- Changelog inicial em `CHANGELOG.md`.
- Atualização da stack em `docs/Gatekeeper - Architectural Blueprin.md` para Go 1.26, Logrus, Viper, Dig, Validator, Testify/Mockery.
- Atualização de padrões em `docs/Gatekeeper - Engineering Standards.md` com logging e testes novos.
- Instalação de dependências do MVP:
  - Logrus (v1.9.4) para logging estruturado.
  - Viper (v1.21.0) para configuração avançada.
  - Uber Dig (v1.19.0) para injeção de dependências.
  - go-chi/chi (v5.2.5) + jwtauth (v5.4.0) para routing e JWT.
  - lib/pq (v1.12.0) para conexão com PostgreSQL.
  - go-playground/validator (v10.30.1) para validação de structs.
  - Testify (v1.11.1) para testes e assertions.
  - google/uuid (v1.6.0) para geração de UUIDs.

### Planejado
- Implementar setup de Uber Dig e Viper em `cmd/api`.
- Configurar test suite com Testify + Mockery.
- Configurar Prometheus + Grafana na infraestrutura.
- Adicionar OpenTelemetry tracing (pós-MVP).
- Instalar CLI tools: sqlc, mockery, migrate, golangci-lint.

## [2026-03-19]

- Início da configuração do projeto e definição do roadmap MVP.
