# Gatekeeper Project Memory (Shared)

Este arquivo é a fonte canônica de estado e aprendizado para agentes de IA e humanos que trabalham no projeto.

## 📌 Important: Relationship with CHANGELOG.md

Este arquivo trabalha em **dupla com `CHANGELOG.md`**:
- **CHANGELOG.md:** Registro detalhado de TODAS as mudanças de código/docs (commits, adições, correções).
- **project-memory.md:** Síntese temporal de **decisões, aprendizados e estado** do projeto, com referências ao CHANGELOG para detalhes técnicos.

**Prática recomendada:**
1. Quando adicionar uma entrada aqui, referencie o commit/seção no CHANGELOG (ex: "Ver CHANGELOG [Unreleased] para detalhes de implementação").
2. O CHANGELOG mantém a verdade sobre O QUE foi feito; esta memória mantém POR QUE foi feito e LIÇÕES aprendidas.
3. Mantenha esta memória concisa; detalhes técnicos vão para CHANGELOG.
4. Atualize ambos em conjunto após sprints ou marcos significativos.

## 2026-03-19: Início e definição do escopo
- Projeto: Gatekeeper, SaaS de gestão de encomendas multi-tenant para portarias.
- Arquitetura definida: monólito modular (Clean Architecture).
- Stack inicial: Go 1.23+, PostgreSQL 16, go-chi, sqlc, log/slog, godotenv, golang-migrate.
- Roadmap MVP definido em docs (módulos, auth, tenants, unidades, moradores, orders).

## 2026-03-19: Atualização da stack para MVP
- Mudança para Go 1.26.0.
- Logger: Logrus (sirupsen/logrus).
- DI: Uber Dig (uber-go/dig).
- Config: Viper (spf13/viper).
- Validação: go-playground/validator.
- Testes: Testify (stretchr/testify) + Mockery (vektra/mockery).
- Observabilidade: Prometheus + Grafana (métricas), OpenTelemetry-Go (tracing). OAuth2 para Google como plano futuro.

## 2026-03-19: Organização de docs
- `docs/Gatekeeper - Architectural Blueprin.md` atualizado com nova stack e fases do MVP.
- `docs/Gatekeeper - Engineering Standards.md` atualizado com novas diretrizes de logging, testes e ferramental.

## 2026-03-19: Status do roadmap
- Épico 01: fundação / infraestrutura em andamento (setup do projeto, DI, config, logger, migrations, sqlc).
- Épico 02: gestão de acesso planejado (autenticação JWT, middleware tenant_id).
- Épico 03: cadastro de moradores/unidades planejado (CRUD e validação).
- Épico 04: processo de recebimento em planejamento (orders, notificações).

## 2026-03-19: Instalação de dependências do MVP
- Instaladas todas as dependências Go do Épico 01 (fundação):
  - **Logging:** sirupsen/logrus v1.9.4
  - **Config:** spf13/viper v1.21.0
  - **DI:** go.uber.org/dig v1.19.0
  - **Router:** go-chi/chi/v5 v5.2.5 + go-chi/jwtauth/v5 v5.4.0
  - **DB Driver:** lib/pq v1.12.0
  - **UUID:** google/uuid v1.6.0
  - **Validation:** go-playground/validator/v10 v10.30.1
  - **Testing:** stretchr/testify v1.11.1
- Commit: `build: install MVP dependencies`
- go.mod e go.sum commitados com sucesso.

## 2026-03-19: Correção de go.mod e go.sum
- Identificado: go.sum estava vazio após primeira tentativa de instalação.
- Causa: Erro durante `go install` de CLI tools interferiu no processo.
- Solução: Re-executar `go get -u` com todas as dependências do MVP em um comando único.
- Resultado: go.mod e go.sum agora com 44 dependências totalmente resolvidas e lockadas.
- Commit: `fix: properly install and lock MVP dependencies in go.mod and go.sum`
- Verificação: `go mod verify` passou com sucesso.

## 2026-03-19: Lições aprendidas
- A definição precisa de ferramentas (Logrus, Dig, Viper etc.) evita ambiguidade na especificação e facilita desenvolvimento rápido.
- Priorização: MVP lean antes de trocar para observabilidade completa.
- Prática de Documentação: Manter project-memory.md e CHANGELOG.md sincronizados e com cross-references garante rastreabilidade completa sem redundância.
- CLI tools (sqlc, mockery, migrate, golangci-lint) devem ser instaladas via `go install` DEPOIS de estabelecer go.mod com dependências principais.
- **Importante:** Executar instalação de dependências do módulo em um único comando `go get` é mais confiável que múltiplos `go get` em paralelo.
  - Ver [CHANGELOG Unreleased](../../../CHANGELOG.md) para lista exata de versões instaladas.
