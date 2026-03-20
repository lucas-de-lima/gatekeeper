# Domain Overview

## Core Entities
- Tenant
- User (Admin, Gatekeeper)
- Unit
- Resident
- Order

## Order Lifecycle
- PENDING → DELIVERED | CANCELED | RETURNED

## Business Rules
- Every entity must belong to a tenant
- Orders cannot be deleted
- Residents may change units