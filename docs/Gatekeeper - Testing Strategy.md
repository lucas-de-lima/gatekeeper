# Gatekeeper - Testing Strategy & TDD

**Philosophy:** Pragmatic Testing over Dogmatic TDD.

**Focus:** High confidence, low maintenance (Test Churn).

## 1. O Veredito sobre TDD

Não exigiremos TDD estrito (Red-Green-Refactor) para 100% da base de código. O TDD será encorajado e focado onde ele realmente gera valor: no Core Domain (Regras de Negócio).

## 2. A Pirâmide de Testes do Gatekeeper

### 2.1. Camada de Use Cases (Regras de Negócio) – AQUI O TDD BRILHA

- Regra: Testes Unitários são OBRIGATÓRIOS.
- Recomendação: TDD é altamente recomendado.
- Como:
  - Escrever testes antes da lógica de `RegisterOrder`.
  - Casos de borda: tenant_id vazio, resident_id inexistente, notificação falhando.
- Técnica: Table-Driven Tests (padrão ouro em Go). Mockar repositórios (Ports/Interfaces).

### 2.2. Camada de Entidades (Validações) – TDD SIMPLES

- Regra: Testes unitários rápidos para invariantes (ex: `NewUser()` não aceita email inválido).

### 2.3. Camada de Repositório (Banco de Dados) – TEST-AFTER (Integração)

- Regra: TDD puro não funciona bem com SQL.
- Processo:
  1. Escrever query SQL.
  2. Gerar código com `sqlc`.
  3. Escrever teste de integração.
- Não usar `sqlmock` (não valida SQL real).
- Usar DB PostgreSQL real (Testcontainers ou Docker local).

### 2.4. Camada de HTTP (Handlers/Controllers) – TEST-AFTER (E2E/Integração)

- Regra: Validar caminho feliz (`Happy Path`).
- Ferramenta: `httptest.NewRecorder`.
- Verificar:
  - parsing JSON de requisição
  - status code esperado (ex: `201 Created`, `400 Bad Request`)

## 3. Cobertura de Código (Coverage)

- Não buscamos métrica de vaidade `100%`.
- Alvo para MVP: `70%+` nas pastas `internal/usecase` e `internal/entity`.
- Pastas de setup (`cmd/api`) podem ter coverage baixo.
