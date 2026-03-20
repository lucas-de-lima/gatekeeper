# Gatekeeper Project Memory (Shared)

Este arquivo é a fonte canônica de estado e aprendizado para agentes de IA e humanos que trabalham no projeto.

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

## 2026-03-19: Lição aprendida
- A definição precisa de ferramentas (Logrus, Dig, Viper etc.) evita ambiguidade na especificação e facilita desenvolvimento rápido.
- Priorização: MVP lean antes de trocar para observabilidade completa.
