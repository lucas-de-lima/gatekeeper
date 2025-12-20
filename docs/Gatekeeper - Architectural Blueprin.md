# Gatekeeper - Architectural Blueprint

**Architecture Style:** Modular Monolith  
**Design Pattern:** Clean Architecture

## 1. Visão Geral

Adotaremos uma abordagem de Monólito Modular. Isso significa que o sistema será implantado como um único binário/serviço, mas internamente o código será organizado em módulos com fronteiras bem definidas, prevenindo acoplamento indevido. Isso facilita a futura extração para microsserviços, caso a escala exija (o que é improvável no curto prazo).

## 2. Diagrama de Camadas (Clean Architecture)

O fluxo de dependência deve ser sempre: **Externo -> Interno**.

```
[ Infraestrutura (Database, Web, External APIs) ]
        |
        v
[ Interface Adapters (Handlers, Controllers, Gateways) ]
        |
        v
[ Use Cases (Application Business Rules) ]
        |
        v
[ Entities (Enterprise Business Rules) ]
```

### 2.1. Camadas Detalhadas

**cmd/api:**

- Ponto de entrada (main.go).
- Carregamento de configurações (env).
- Injeção de dependência (Wiring).
- Inicialização do servidor HTTP.

**internal/entity:**

- Objetos de domínio puro (Go Structs).
- Regras de validação invariantes (ex: "Uma encomenda não pode ser criada sem destinatário").
- Sem tags de JSON ou DB (se possível).

**internal/usecase:**

- Orquestração do fluxo de negócio.
- Define interfaces para Repositórios e Serviços (Dependency Inversion).
- Exemplo: RegisterPackage, DeliverPackage, CreateTenant.

**internal/infra:**

- **web:** Handlers HTTP (Chi), Middlewares, DTOs (Data Transfer Objects) de entrada/saída.
- **repository:** Implementação concreta do acesso a dados (SQLc/Postgres).
- **service:** Implementações de serviços externos (ex: WhatsAppService).

## 3. Stack Tecnológica

| Componente | Tecnologia | Justificativa |
|------------|------------|---------------|
| Language | Go 1.23+ | Performance, tipagem forte, concorrência simples. |
| Database | PostgreSQL 16 | Relacional, robusto, suporte a JSONB se necessário. |
| Router | go-chi/chi | Leve, idiomatico, excelente suporte a middlewares (context). |
| ORM/Query | sqlc | Gera código Go type-safe a partir de SQL puro. Performance de raw SQL. |
| Logging | log/slog | Padrão da biblioteca (Go 1.21+), estruturado, zero alloc. |
| Config | godotenv | Simplicidade (12-factor app). |
| Migration | golang-migrate | Versionamento de schema agnóstico. |

## 4. Design Patterns Chave

### 4.1. Repository Pattern (Generic)

Para evitar repetição de código CRUD, utilizaremos Go Generics para criar um BaseRepository onde possível, mas sempre favorecendo queries específicas geradas pelo sqlc para performance.

### 4.2. Adapter Pattern (Notifications)

**Interface no Domínio:**

```go
type Notifier interface {
    Send(ctx context.Context, msg Message) error
}
```

**Implementação na Infra:**

```go
type WhatsAppAdapter struct { ... }
func (w *WhatsAppAdapter) Send(...) error { ... }
```

Isso permite trocar WhatsApp por Email ou SMS sem tocar no Core.

### 4.3. Middleware de Tenant Context

Todo request autenticado passará por um middleware que extrai o tenant_id do token JWT e o injeta no context.Context do Go. Esse contexto é passado até a camada de repositório, garantindo que nenhuma query seja executada sem o filtro de Tenant.
