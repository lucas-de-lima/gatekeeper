# Gatekeeper - Database Schema

**Engine:** PostgreSQL 16+  
**Strategy:** Single Database, Shared Schema, Row-Level Separation (Tenant ID).

## 1. Diretrizes de Modelagem

- **Multi-tenancy:** Todas as tabelas de negócio DEVEM ter a coluna tenant_id (UUID).
- **Primary Keys:** Usaremos UUID v7 (Time-ordered) ou BIGINT. Vamos priorizar UUID v7 para evitar enumeração de recursos e facilitar sharding futuro, além de manter ordenação temporal.
- **Audit:** Colunas created_at e updated_at são obrigatórias.
- **Soft Delete:** Utilizar coluna deleted_at apenas em entidades críticas (Users, Residents). Encomendas (Orders) não devem ser deletadas, apenas canceladas.

## 2. Tabelas Core

### tenants

Representa o condomínio ou empresa cliente.

- id (UUID, PK)
- name (VARCHAR)
- document (CNPJ - VARCHAR)
- status (ENUM: ACTIVE, SUSPENDED)
- created_at, updated_at

### users

Porteiros e Administradores.

- id (UUID, PK)
- tenant_id (UUID, FK -> tenants.id)
- name (VARCHAR)
- email (VARCHAR, Unique per Tenant)
- password_hash (VARCHAR)
- role (ENUM: ADMIN, GATEKEEPER)
- active (BOOLEAN)

### units

Apartamentos, casas ou setores.

- id (UUID, PK)
- tenant_id (UUID, FK -> tenants.id)
- block (VARCHAR) - Bloco/Quadra
- number (VARCHAR) - Número/Apto
- Index composto: (tenant_id, block, number) UNIQUE.

### residents

Moradores/Destinatários.

- id (UUID, PK)
- tenant_id (UUID, FK -> tenants.id)
- unit_id (UUID, FK -> units.id, Nullable - pode mudar de apto)
- name (VARCHAR)
- cpf (VARCHAR)
- phone (VARCHAR) - Formato E.164 para WhatsApp
- email (VARCHAR)
- active (BOOLEAN)

### orders

As encomendas.

- id (UUID, PK)
- tenant_id (UUID, FK -> tenants.id)
- resident_id (UUID, FK -> residents.id)
- unit_id (UUID, FK -> units.id) - Snapshot do local na hora da entrega
- registered_by (UUID, FK -> users.id) - Porteiro que recebeu
- tracking_code (VARCHAR, Indexed)
- sender (VARCHAR) - Ex: Amazon, Correios
- status (ENUM: PENDING, DELIVERED, CANCELED, RETURNED)
- delivered_at (TIMESTAMP)
- delivered_to_name (VARCHAR) - Nome de quem retirou (se diferente do resident)
- delivered_to_doc (VARCHAR) - Documento de quem retirou
- created_at

## 3. Segurança Row-Level

Todas as consultas SQL (via sqlc) devem obrigatoriamente incluir a cláusula WHERE tenant_id = $1.

Exemplo:

```sql
-- name: ListOrders :many
SELECT * FROM orders
WHERE tenant_id = $1 AND status = 'PENDING'
ORDER BY created_at DESC;
```
