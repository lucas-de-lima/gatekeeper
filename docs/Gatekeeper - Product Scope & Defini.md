# Gatekeeper - Product Scope & Definition

**Document Status:** Draft v1.0  
**Owner:** Lucas (Tech Lead) / AI Architect

## 1. Visão do Produto

O Gatekeeper é um sistema backend SaaS (Software as a Service) projetado para modernizar a gestão de encomendas em portarias residenciais e comerciais. O objetivo primário é substituir o tradicional "Livro de Ocorrências/Protocolo" de papel por um fluxo digital seguro, auditável e eficiente.

**Mantra do Produto:** "Segurança do registro em papel, velocidade do digital."

## 2. Modelo de Negócio

- **Modelo:** SaaS B2B Multi-tenant.
- **Infraestrutura:** Instância única compartilhada, isolamento lógico de dados por tenant_id.
- **Estratégia de Distribuição:** Assinatura mensal por condomínio/empresa (Tenant).

## 3. Atores e Papéis

### 3.1. Tenant (O Cliente)

A entidade legal (Condomínio, Prédio Comercial, Empresa) que contrata o serviço.

- Possui configurações globais (ex: habilitar notificação WhatsApp, tempo de retenção de logs).

### 3.2. Admin (O Gestor)

Geralmente o Síndico, Zelador Chefe ou Gerente de Facilities.

- **Poderes:** Criar/Remover Porteiros, Cadastrar Moradores, Visualizar Relatórios, Configurar Integrações.
- Acesso total aos dados apenas do seu Tenant.

### 3.3. Gatekeeper (O Porteiro/Recepcionista)

O usuário operacional diário.

- **Poderes:** Registrar entrada de encomendas (Check-in), Registrar saída/baixa (Check-out), Consultar histórico recente.
- Interface simplificada e de alta velocidade.

### 3.4. Resident (O Destinatário)

O usuário final passivo (no MVP).

- **Interação:** Recebe notificações (WhatsApp/Email) e fornece código/documento para retirada.
- Não possui login no sistema neste MVP.

## 4. Escopo do MVP (Minimum Viable Product)

### Funcionalidades CORE (Obrigatórias)

**Autenticação & Gestão de Acesso:**

- Login seguro para Admin e Gatekeeper.
- Isolamento estrito de dados entre Tenants (Middleware de Contexto).

**Gestão de Moradores/Unidades:**

- CRUD simples de Unidades (Apto 101, Sala 304).
- Vínculo de Moradores às Unidades (Nome, CPF, Telefone Celular).

**Fluxo de Recebimento (Inbound):**

- Registro rápido: Quem recebeu (Porteiro), Para quem é (Morador/Unidade), Origem (Correios/Amazon/Mercado Livre), Código de Rastreio (Opcional).
- Geração de ID único da encomenda.

**Notificação (Outbound):**

- Disparo assíncrono de mensagem via WhatsApp (via Adapter Pattern) informando a chegada.
- Fallback: Se falhar, apenas logar o erro, não bloquear o registro.

**Fluxo de Retirada (Outbound):**

- Baixa da encomenda.
- Registro de quem retirou (Próprio morador ou autorizado).
- Timestamp imutável da entrega.

### Funcionalidades FORA DO MVP (Roadmap Futuro)

- Aplicativo Mobile para o Morador.
- Reconhecimento de texto (OCR) para ler etiquetas de encomendas via câmera.
- Assinatura digital na tela (Touch).
- Integração com Lockers Inteligentes.
- Módulo de Correspondências (Cartas simples).

## 5. Requisitos Não-Funcionais (SLA)

- **Disponibilidade:** 99.5% (Horário comercial é crítico).
- **Latência:** APIs de registro e busca < 200ms.
- **Segurança:** Senhas com hash (Argon2/Bcrypt), HTTPS obrigatório, Sanitização de inputs.
