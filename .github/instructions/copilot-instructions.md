# Gatekeeper AI Instructions

You are working on the Gatekeeper project.

## Architecture
- Modular Monolith
- Clean Architecture
- Dependency flow: External -> Internal

Layers:
- entity: pure domain models (no frameworks)
- usecase: business logic and orchestration
- infra: web, repository, external services

## Rules

### Domain
- Entities must not depend on infra
- No database or JSON tags in entities (if possible)

### Usecases
- Must define interfaces for repositories/services
- Must not depend on concrete implementations

### Infra
- Implements interfaces defined in usecase
- Handles HTTP, DB, external APIs

## Multi-tenancy
- ALL queries must include tenant_id
- tenant_id must come from context.Context
- Never allow cross-tenant data access

## Error Handling
- Never ignore errors
- Always wrap errors with context:
  fmt.Errorf("context: %w", err)

## Logging
- Structured logs (slog)
- Always include:
  - tenant_id
  - request_id

## Code Style
- Simple, readable, maintainable
- Avoid over-engineering
- Prefer composition over inheritance

## Testing
- Use table-driven tests
- Mock dependencies via interfaces

## Performance
- Avoid loading unnecessary data into memory
- Prefer efficient SQL queries (sqlc)

## Security
- Never trust user input
- Always validate external data

## Shared Memory Reference
Before any code changes, read:
- `.github/instructions/project-memory.md`
- `CHANGELOG.md`

The memory file is the primary shared source for project state, decisions, and ongoing TODOS. Keep this instruction file minimal and point to it for detailed context.