# Gatekeeper

SaaS backend for package management in residential and commercial buildings.

## Architecture
- Modular Monolith
- Clean Architecture
- Multi-tenant (tenant_id isolation)

## Stack
- Go
- PostgreSQL
- Chi
- sqlc

## Structure
- cmd/ → entrypoint
- internal/ → core logic
- infra/ → external layers

## Principles
- Security first
- Strong tenant isolation
- Simple and maintainable code