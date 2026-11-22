# Portable Engineering Harness â€” Blueprint

**Status:** Architecture and implementation brief for a local coding agent  
**Purpose:** Build a portable, project-agnostic, AI-first software engineering harness.  
**Package type:** Documentation/specification bundle; this ZIP intentionally contains no implementation code.

## Start here

1. Read [`AGENTS.md`](AGENTS.md) for persistent repository guidance, then [`AGENT_START_HERE.md`](AGENT_START_HERE.md) for the implementation brief.
2. Read [`01_PRODUCT_VISION_AND_REQUIREMENTS.md`](01_PRODUCT_VISION_AND_REQUIREMENTS.md) and [`02_ARCHITECTURE.md`](02_ARCHITECTURE.md).
3. Read [`03_E2E_PIPELINE.md`](03_E2E_PIPELINE.md) and [`04_AI_FIRST_ONBOARDING.md`](04_AI_FIRST_ONBOARDING.md) carefully. The AI-first onboarding requirement is non-negotiable.
4. Use [`05_COMPONENT_INTEGRATION_CATALOG.md`](05_COMPONENT_INTEGRATION_CATALOG.md), [`06_CONTRACTS_AND_ARTIFACTS.md`](06_CONTRACTS_AND_ARTIFACTS.md), and [`07_PORTABILITY_AND_VERSIONING.md`](07_PORTABILITY_AND_VERSIONING.md) as implementation guidance.
5. Follow [`08_EXECUTION_POLICIES.md`](08_EXECUTION_POLICIES.md), [`09_OBSERVABILITY_EVALUATION_AND_LEARNING.md`](09_OBSERVABILITY_EVALUATION_AND_LEARNING.md), and [`10_IMPLEMENTATION_AND_ACCEPTANCE.md`](10_IMPLEMENTATION_AND_ACCEPTANCE.md).
6. Read [`13_LUCAS_PRODUCT_VISION_AND_DECISIONS.md`](13_LUCAS_PRODUCT_VISION_AND_DECISIONS.md) as the authoritative addendum for product decisions clarified after this blueprint was drafted.
7. Use [`14_ROTEIRO_DE_SESSOES_E_INICIO_DA_IMPLEMENTACAO.md`](14_ROTEIRO_DE_SESSOES_E_INICIO_DA_IMPLEMENTACAO.md) to sequence remaining design sessions and identify the implementation start gate.
8. Consult [`references/SOURCE_CATALOG.md`](references/SOURCE_CATALOG.md) and [`references/INTEGRATION_NOTES.md`](references/INTEGRATION_NOTES.md) for upstream sources and caveats.
9. Read the operational design contracts [`15_MODELO_DE_EXECUCAO_E_ADAPTERS.md`](15_MODELO_DE_EXECUCAO_E_ADAPTERS.md) through [`20_POLITICA_DE_SEGURANCA_PORTABILIDADE_E_DEPENDENCIAS.md`](20_POLITICA_DE_SEGURANCA_PORTABILIDADE_E_DEPENDENCIAS.md), then the integrated plan in [`21_ARQUITETURA_INTEGRADA_E_PLANO_DE_IMPLEMENTACAO.md`](21_ARQUITETURA_INTEGRADA_E_PLANO_DE_IMPLEMENTACAO.md).
10. Before implementing providers or running upstream installers, use [`22_PLANO_DE_AMARRACAO_E_INTEGRACAO_FUNCIONAL.md`](22_PLANO_DE_AMARRACAO_E_INTEGRACAO_FUNCIONAL.md) as the shared-capability and functional-installation contract.
11. Read [`BLUEPRINT_STATUS.yaml`](BLUEPRINT_STATUS.yaml), [`23_IDENTIDADE_DO_WORKSPACE_E_MODOS_DE_TRABALHO.md`](23_IDENTIDADE_DO_WORKSPACE_E_MODOS_DE_TRABALHO.md), and [`24_PLANO_MESTRE_DE_CONSTRUCAO_DO_HARNESS.md`](24_PLANO_MESTRE_DE_CONSTRUCAO_DO_HARNESS.md) to determine whether work is on the Harness itself and which task comes next.
12. For autonomous testing of an existing target, use [`25_VALIDACAO_AUTONOMA_DE_PROJETOS.md`](25_VALIDACAO_AUTONOMA_DE_PROJETOS.md).

**Precedence:** where original requirements conflict with Lucas's clarified decisions in document 13 or its operational specifications in documents 15â€“21, the clarified decisions and newer specifications govern. Document 13 marks confirmed choices versus proposals; documents 15â€“21 do not imply implementation has started.

## What this blueprint specifies

The target is a **portable Engineering Harness** that can be cloned or carried on removable storage, then introduced into any new or legacy repository. Its core remains generic. An AI agent, from the first moment, conducts discovery and onboarding, creates a project-specific instance, and then operates an end-to-end (E2E) task pipeline.

The system combines these capability areas:

- Orchestration and selective delegation.
- Spec-driven engineering workflows.
- Curated engineering skills and reusable procedures.
- Structural code intelligence / knowledge graph.
- Persistent cross-session and cross-agent memory.
- Provider-neutral model access, with optional gateway routing.
- Execution traces, verification and quality evaluation.
- Controlled, evidence-based improvement of the harness itself.
- Portable project installation, versioning and per-project specialization.

## Critical product decisions

- **AI-first from moment zero.** No mandatory pre-AI scanner, wizard, installer, or large opaque script. The AI agent leads discovery, decides what to inspect, invokes bounded deterministic helpers when useful, interprets evidence, creates project artifacts, and validates the result.
- A CLI may exist later as an optional convenience for deterministic operations. A mandatory `harness install <path>` command is **not** the onboarding architecture.
- The Master Agent coordinates the E2E pipeline and is a management specialist, not a persona for the other disciplines. Real specialist agents form four pillars: architecture, engineering, backlog management, and development.
- Stages and agents do not map one-to-one, but an executor without real subagent delegation is unsupported. Pipeline composition and the number of specialists invoked remain problem-dependent.
- Keep the harness generic. Project expertise comes from AI-led onboarding, project-local configuration, verified project knowledge, and project-owned artifactsâ€”not hard-coded stacks.
- Keep the upstream harness repository separate from the copied/versioned project instance.
- Do not install every referenced project wholesale. Integrate capabilities through narrow adapters and explicit contracts.
- CodeGraph and Graphify are intended to work together as complementary knowledge providers. Their query boundaries, overlap, freshness, and platform behavior require validation.
- Pipelines are composable and may be generated for an individual request, retained for a project, or promoted as reviewed upstream patterns; generated definitions cannot bypass authority or safety invariants.
- Avoid an endless user-facing prototyping loop. The local agent should implement a coherent integrated system from this specification, use normal internal engineering checks, and deliver a consolidated E2E validation report. Ask the user only for genuinely blocking decisions or permissions.

## Package contents

| File | Purpose |
|---|---|
| `AGENT_START_HERE.md` | Execution brief and primary instructions for the local coding agent |
| `AGENTS.md` | Short persistent root guidance and reading map for agents working on this blueprint package |
| `01_PRODUCT_VISION_AND_REQUIREMENTS.md` | Vision, goals, non-goals, requirements and product principles |
| `02_ARCHITECTURE.md` | Logical architecture and component boundaries |
| `03_E2E_PIPELINE.md` | Task lifecycle, stage contracts, transitions, resume and direct invocation |
| `04_AI_FIRST_ONBOARDING.md` | Mandatory AI-led installation and project discovery behavior |
| `05_COMPONENT_INTEGRATION_CATALOG.md` | Role and integration approach for every selected reference |
| `06_CONTRACTS_AND_ARTIFACTS.md` | Inter-module contracts and durable artifact models |
| `07_PORTABILITY_AND_VERSIONING.md` | Portable distribution, project instance, Git, upgrade and USB behavior |
| `08_EXECUTION_POLICIES.md` | Practical security and execution boundaries without building a sandbox platform |
| `09_OBSERVABILITY_EVALUATION_AND_LEARNING.md` | Tracing, quality gates, HarnessÂ²-inspired improvement loop |
| `10_IMPLEMENTATION_AND_ACCEPTANCE.md` | Build directives, delivery expectations and E2E acceptance criteria |
| `11_DECISIONS_AND_OPEN_ITEMS.md` | Approved decisions and items the local agent must resolve transparently |
| `12_GLOSSARY.md` | Shared terminology |
| `13_LUCAS_PRODUCT_VISION_AND_DECISIONS.md` | Product-owner vision, clarified decisions, operating model, and open questions |
| `14_ROTEIRO_DE_SESSOES_E_INICIO_DA_IMPLEMENTACAO.md` | Recommended design-session sequence and gate for starting implementation |
| `15_MODELO_DE_EXECUCAO_E_ADAPTERS.md` | Real multi-agent execution contract and executor capability matrix |
| `16_PILARES_AGENTES_E_PIPELINES.md` | Four-pillar roles, specialist contracts, and dynamic pipeline validation |
| `17_CICLO_DE_VIDA_E_GOVERNANCA.md` | Installation/resume modes, architecture authority, homologation and GMUD |
| `18_MEMORIA_E_CONTRATOS_DE_ARTEFATOS.md` | Canonical memory, event/run/evidence/handoff models |
| `19_INTEGRACOES_E_ARQUITETURA_DE_CONHECIMENTO.md` | Graph tools, external knowledge, skills, memory and optional providers |
| `20_POLITICA_DE_SEGURANCA_PORTABILIDADE_E_DEPENDENCIAS.md` | Risk-based autonomy and Windows/Linux/dependency boundaries |
| `21_ARQUITETURA_INTEGRADA_E_PLANO_DE_IMPLEMENTACAO.md` | Consolidated architecture, feasibility limits, phases and acceptance criteria |
| `22_PLANO_DE_AMARRACAO_E_INTEGRACAO_FUNCIONAL.md` | Shared provider/capability layer, per-source integration modes, installation waves and collaboration acceptance scenarios |
| `23_IDENTIDADE_DO_WORKSPACE_E_MODOS_DE_TRABALHO.md` | Workspace identities and explicit modes distinguishing Harness construction from target-system delivery |
| `24_PLANO_MESTRE_DE_CONSTRUCAO_DO_HARNESS.md` | Ordered build backlog, phase dependencies/gates, handoff contract, and complete readiness definition |
| `BLUEPRINT_STATUS.yaml` | Current package/product status, active target flag, and first build work item |
| `references/SOURCE_CATALOG.md` | All supplied and added links |
| `references/INTEGRATION_NOTES.md` | Repository-level notes, overlaps, caveats and selection guidance |
| `diagrams/*.mmd` | Mermaid architecture and pipeline diagrams |

## Important caveat

This is a design and implementation brief, not a claim that every upstream repository is production-ready, mutually compatible, or suitable as a runtime dependency. Before integrating a project, the local agent must inspect its current official repository, current documentation, license, platform support, install method, dependencies, and API/CLI contract. Preserve the distinction between **using a project's ideas**, **adopting its file format**, **integrating its executable**, and **vendoring its source code**.
