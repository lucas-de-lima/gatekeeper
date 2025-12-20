# Gatekeeper - MVP Backlog & Roadmap

**Status:** Planning  
**Context:** MVP Development Phase

Este documento quebra o escopo do produto em entregáveis técnicos organizados por ordem de dependência.

## ÉPICO 01: Fundação & Infraestrutura (The Bedrock)

**Objetivo:** Estabelecer a base tecnológica, ambiente de desenvolvimento e estrutura do banco de dados multi-tenant.

### US-01: Setup do Projeto e Arquitetura Base

Como Engenheiro, quero uma estrutura de pastas organizada seguindo Clean Arch para iniciar o desenvolvimento.

- [ ] Inicializar módulo Go (go mod init).
- [ ] Criar estrutura de pastas (cmd, internal/entity, internal/usecase, internal/infra).
- [ ] Configurar Makefile (run, build, test, lint).
- [ ] Configurar Linter (golangci-lint) com regras estritas.

### US-02: Infraestrutura de Dados e Docker

Como Engenheiro, quero rodar o banco de dados localmente para desenvolver as queries.

- [ ] Criar docker-compose.yml com PostgreSQL 16.
- [ ] Configurar golang-migrate para versionamento de schema.
- [ ] Escrever Migration 001: Tabelas tenants e users.
- [ ] Configurar sqlc (sqlc.yaml) e gerar primeiras structs.

### US-03: Logging e Configuração

Como Operador, quero logs estruturados para debuggar problemas em produção.

- [ ] Implementar carregamento de variáveis de ambiente (.env).
- [ ] Configurar Logger (slog) com output JSON.
- [ ] Criar middleware básico de Logging de Request HTTP (Method, Path, Latency).

## ÉPICO 02: Gestão de Acesso e Multi-tenancy (The Gate)

**Objetivo:** Garantir que cada requisição seja autenticada e isolada ao Tenant correto.

### US-04: Login e Autenticação (JWT)

Como Admin/Porteiro, quero me logar para acessar o sistema.

- [ ] Implementar Hash de senha (Argon2 ou Bcrypt).
- [ ] Criar UseCase AuthenticateUser (busca user, valida senha).
- [ ] Gerar Token JWT assinado contendo user_id, role e tenant_id.
- [ ] Criar Endpoint POST /api/v1/login.

### US-05: Middleware de Tenant Context

Como Sistema, quero garantir que dados de um condomínio não vazem para outro.

- [ ] Criar Middleware que intercepta o JWT.
- [ ] Validar assinatura do Token.
- [ ] Extrair tenant_id e injetar no context.Context do Go.
- [ ] Rejeitar requisições sem token (exceto login/healthcheck).

### US-06: Gestão de Usuários (Backoffice Admin)

Como Admin, quero cadastrar meus porteiros.

- [ ] Criar Migration 002: Refinar tabela users.
- [ ] Criar UseCase CreateUser (apenas Admin pode criar).
- [ ] Validar unicidade de email dentro do Tenant.

## ÉPICO 03: Cadastro de Moradores e Unidades (The Inventory)

**Objetivo:** Mapear quem mora onde para poder endereçar as encomendas.

### US-07: Gestão de Unidades

Como Admin, quero cadastrar os apartamentos/blocos do condomínio.

- [ ] Criar Migration 003: Tabelas units e residents.
- [ ] Implementar CRUD de Unidades (CreateUnit, ListUnits).
- [ ] Garantir constraint única (tenant_id, block, number).

### US-08: Gestão de Moradores

Como Admin, quero cadastrar moradores e vinculá-los a uma unidade.

- [ ] Implementar CRUD de Moradores.
- [ ] Validar formato de CPF e Telefone (apenas validação básica/regex).
- [ ] Implementar Endpoint GET /residents com filtro por Unidade e Nome (busca parcial).

## ÉPICO 04: Processo de Recebimento (Inbound)

**Objetivo:** O coração do sistema. Registrar a chegada de pacotes.

### US-09: Registro de Encomenda

Como Porteiro, quero registrar rapidamente uma encomenda para liberar o entregador.

- [ ] Criar Migration 004: Tabela orders.
- [ ] Implementar UseCase RegisterOrder.
- [ ] Input: Selecionar Unidade/Morador, Digitar Rastreio (opcional), Origem.
- [ ] Gerar ID único e Status inicial PENDING.

### US-10: Listagem de Encomendas (Dashboard)

Como Porteiro, quero ver o que está na portaria aguardando retirada.

- [ ] Implementar Endpoint GET /orders.
- [ ] Filtros obrigatórios: status (default: PENDING), unit_id.
- [ ] Ordenação: Mais antigas primeiro.

## ÉPICO 05: Notificações (The Signal)

**Objetivo:** Avisar o morador que a encomenda chegou (Design Patterns).

### US-11: Arquitetura de Notificação (Adapter)

Como Arquiteto, quero desacoplar a regra de negócio do envio de mensagem.

- [ ] Definir Interface Notifier no domínio (internal/entity/notifier.go).
- [ ] Implementar LogNotifier (Mock que apenas loga no stdout) para desenvolvimento.
- [ ] Integrar chamada do Notifier.Send() dentro do UseCase RegisterOrder (de preferência em goroutine para não bloquear).

### US-12: Implementação WhatsApp (MVP)

Como Morador, quero receber um Zap quando minha encomenda chegar.

- [ ] Escolher driver simples (ex: API externa mockada ou lib Go).
- [ ] Implementar WhatsAppAdapter que satisfaz a interface Notifier.
- [ ] Formatador de mensagem: "Olá {Nome}, chegou uma encomenda da {Origem} na portaria."

## ÉPICO 06: Processo de Retirada (Outbound)

**Objetivo:** Baixar a encomenda e garantir a rastreabilidade da entrega.

### US-13: Baixa de Encomenda

Como Porteiro, quero registrar que entreguei o pacote ao morador.

- [ ] Implementar UseCase DeliverOrder.
- [ ] Input: order_id, receiver_name, receiver_doc.
- [ ] Validação: Encomenda deve estar PENDING.
- [ ] Ação: Mudar status para DELIVERED, gravar delivered_at (NOW).

### US-14: Histórico/Relatório Simples

Como Admin, quero ver o histórico de entregas do mês.

- [ ] Ajustar GET /orders para suportar filtro de data (start_date, end_date) e status DELIVERED.
