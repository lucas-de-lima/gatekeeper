# Gatekeeper - Engineering Standards & Guidelines

**Language:** Go 1.23+  
**Philosophy:** "Simple, Readable, Maintainable."

## 1. Go Style Guide

Seguiremos o Uber Go Style Guide e o Effective Go, com as seguintes ênfases:

### 1.1. Tratamento de Erros

- Nunca ignore erros. Use `_` apenas se o erro for realmente irrelevante.
- **Wrap Errors:** Ao retornar um erro de uma camada inferior, adicione contexto.

```go
// Ruim
return err

// Bom
return fmt.Errorf("failed to create order: %w", err)
```

- **Sentinel Errors:** Defina erros de domínio em variáveis exportadas na camada entity ou usecase para checagem com errors.Is().

```go
var ErrOrderNotFound = errors.New("order not found")
```

### 1.2. Naming Conventions

- **Interfaces:** Devem ser definidas onde são usadas (Consumer side), não onde são implementadas, exceto interfaces centrais de domínio.
- **Variáveis:** Curtas quando o escopo é pequeno (i, err, ctx). Descritivas quando o escopo é grande ou exportado.
- **Acrônimos:** Mantenha o case consistente (ServeHTTP, ID, URL). Nada de Id, Url.

## 2. Estrutura de Logs

Utilizaremos log/slog para logs estruturados em JSON (em produção) e Texto (em dev).

**Regra:** Todo log de erro ou operação crítica deve conter o tenant_id e o request_id.

```go
logger.Info("order created",
    slog.String("tenant_id", tenantID),
    slog.String("order_id", orderID),
    slog.String("user_id", userID),
)
```

## 3. Testes

- **Unitários:** Obrigatórios para toda a camada de usecase e helpers utilitários. Devemos mockar as interfaces de repositório.
- **Integração:** Obrigatórios para a camada de repository (testar queries reais num DB Dockerizado) e handler (testar rotas HTTP).
- **Tabela de Testes:** Use Table Driven Tests para cobrir múltiplos cenários.

## 4. Git Flow & Commits

- **Branchs:** main (produção), develop (integração), feat/nome-da-feature, fix/nome-do-bug.
- **Conventional Commits:** Obrigatório.

```
feat: add resident registration

fix: resolve whatsapp notification timeout

chore: update go.mod dependencies

docs: update readme

refactor: simplify database connection logic
```

## 5. Ferramental

- **Linter:** golangci-lint deve rodar em todo PR. Configuração estrita.
- **Makefiles:** Use um Makefile para comandos comuns (make run, make test, make migrate).
