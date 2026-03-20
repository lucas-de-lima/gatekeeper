# Changelog

Formato: [semver] e seções por data/versão em ordem decrescente.

## 📌 Important: Relationship with .github\instructions\project-memory.md

Este arquivo trabalha em **dupla com `project-memory.md`**:
- **CHANGELOG.md:** Registro técnico detalhado de todas as mudanças (commits, dependências, novos arquivos).
- **project-memory.md:** Síntese temporal com decisões, aprendizados, e referências cruzadas a este changelog.

**Prática recomendada:**
1. Adicione mudanças técnicas aqui com precisão (versões exatas, commits, arquivos modificados).
2. Para reflexões de design, decisões arquiteturais ou lições aprendidas, adicione correspondente em project-memory.md.
3. Cross-link: ao documentar em project-memory.md, referencie a versão/seção do CHANGELOG (ex: "Ver [Unreleased] para detalhes").
4. Esta dupla mantém rastreabilidade completa: CHANGELOG para o QUE, project-memory para o POR QUE.

## [Unreleased]

### Adicionado
- Documento de memória do projeto em `.github\instructions\project-memory.md`.
- Changelog inicial em `CHANGELOG.md`.
- Implementação do Épico 01 (fundação/infraestrutura):
  - `cmd/api/main.go`, `cmd/api/config/config.go`
  - middleware de request/health
  - logger estruturado (slog JSON) e configuração (Viper)
  - estrutura inicial de diretórios `internal/entity`, `internal/usecase`, `internal/infra`
  - Makefile, docker-compose, migrations base e sqlc placeholder
  - testes e lint totalmente rodados (`go test ./...`, `golangci-lint run`)
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
- Adicionado documento de estratégia de testes em `docs/Gatekeeper - Testing Strategy.md`.

### Planejado
- Implementar setup de Uber Dig e Viper em `cmd/api`.
- Configurar test suite com Testify + Mockery.
- Configurar Prometheus + Grafana na infraestrutura.
- Adicionar OpenTelemetry tracing (pós-MVP).
- Instalar CLI tools: sqlc, mockery, migrate, golangci-lint.

## [2026-03-19]

- Início da configuração do projeto e definição do roadmap MVP.
