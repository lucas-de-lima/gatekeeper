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

## 4. Abordagem Oficial: Table-Driven Tests

Para todos os testes unitários (especialmente em usecases e entities), adotaremos estritamente o padrão Table-Driven Tests.

### 4.1. Estrutura Padrão

Todo teste deve definir um slice de structs anônimas. Cada struct representa um cenário de teste completo, contendo o nome do cenário, os dados de entrada, a configuração dos mocks e o resultado esperado.

### 4.2. Exemplo de Template (Obrigatório seguir a estrutura)

```go
func TestRegisterOrderUseCase_Execute(t *testing.T) {
    // 1. Definição da "Tabela" de casos de teste
    tests := []struct {
        name           string
        inputDTO       dto.RegisterOrderInput
        mockBehavior   func(mockRepo *MockOrderRepository, mockNotifier *MockNotifier)
        expectedError  error
        expectedStatus string
    }{
        {
            name: "success: valid order registration",
            inputDTO: dto.RegisterOrderInput{ TenantID: "t-1", ResidentID: "r-1" },
            mockBehavior: func(repo *MockOrderRepository, notifier *MockNotifier) {
                repo.On("Save", mock.Anything).Return(nil)
                notifier.On("Send", mock.Anything).Return(nil)
            },
            expectedError:  nil,
            expectedStatus: "PENDING",
        },
        {
            name: "error: missing resident id",
            inputDTO: dto.RegisterOrderInput{ TenantID: "t-1", ResidentID: "" },
            mockBehavior: func(repo *MockOrderRepository, notifier *MockNotifier) {
                // Nenhum mock é chamado pois falha na validação
            },
            expectedError:  entity.ErrInvalidResident,
            expectedStatus: "",
        },
    }

    // 2. Execução iterativa (O "Drive" da tabela)
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            // Setup local dos mocks para este cenário
            mockRepo := new(MockOrderRepository)
            mockNotifier := new(MockNotifier)
            tc.mockBehavior(mockRepo, mockNotifier)
            
            usecase := NewRegisterOrderUseCase(mockRepo, mockNotifier)
            
            // Execução
            output, err := usecase.Execute(context.Background(), tc.inputDTO)
            
            // Asserts (usando testify/assert ou stdlib)
            if tc.expectedError != nil {
                assert.ErrorIs(t, err, tc.expectedError)
            } else {
                assert.NoError(t, err)
                assert.Equal(t, tc.expectedStatus, output.Status)
            }
            
            // Verifica se as expectativas dos mocks foram atendidas
            mockRepo.AssertExpectations(t)
            mockNotifier.AssertExpectations(t)
        })
    }
}
```

### 4.3. Benefícios Esperados

- **Clareza:** Fica óbvio o que o código faz apenas lendo a lista de `name` e `expectedError`.
- **Escalabilidade:** Para testar um novo erro de negócio, basta adicionar de 5 a 10 linhas na declaração da tabela, sem duplicar a lógica de setup, execução e assertions.

