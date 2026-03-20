# Gatekeeper Architecture

## Style
- Modular Monolith
- Clean Architecture

## Layers

1. Entity
- Core business rules
- Pure Go structs

2. Usecase
- Business logic
- Defines interfaces

3. Infra
- HTTP (Chi)
- Database (Postgres + sqlc)
- External services (WhatsApp)

## Key Patterns
- Repository Pattern
- Adapter Pattern (Notifier)
- Middleware for Tenant Context

## Multi-tenancy
- Shared DB
- Isolation via tenant_id