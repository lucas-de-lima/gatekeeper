# Agent Start Here — Build the Portable Engineering Harness

You are the local coding agent responsible for turning this blueprint into a working, integrated product when asked to implement it. Read the package-root `AGENTS.md` for persistent repository guidance, then use this file as the implementation brief. Treat the other files in this package as the product specification and design record.

The current package state is recorded in `BLUEPRINT_STATUS.yaml`: this is the Harness blueprint, implementation has not started, and there is no active target system. Before choosing the next implementation task, read `23_IDENTIDADE_DO_WORKSPACE_E_MODOS_DE_TRABALHO.md` and continue from `24_PLANO_MESTRE_DE_CONSTRUCAO_DO_HARNESS.md`. The target of the current build is the Harness itself.

## 1. Mission

Build a **portable, project-agnostic, AI-first Engineering Harness**. The harness is a reusable engineering toolbox that can be placed in or referenced from any software repository. When introduced to a project, an AI agent must lead the onboarding from the first interaction, understand the project using targeted evidence, and create a project-specific harness instance that is versioned with that project.

The result is not a prompt collection. It is a workflow system with an E2E pipeline, invocable stages, skills, tools, context assembly, project knowledge, persistent memory integration, execution policies, observability, verification, and controlled learning.

## 2. Non-negotiable instruction: AI-first onboarding

**Do not implement onboarding as a large scanner or opaque installer that prepares the repository before the AI agent starts.**

The first meaningful actor in onboarding is the AI agent already operating in the target project. The agent must:

1. Read this harness's bootstrap instructions and establish the onboarding goal with the user.
2. Inspect the project progressively using its native file, search, Git, and terminal capabilities.
3. Decide which evidence is needed next and explain why when an operation is material or potentially risky.
4. Invoke small deterministic helpers only when they provide a clear advantage (for example, listing tracked files, reading package metadata, running an existing test command, or building/updating a structural graph).
5. Interpret raw results itself; helpers must not secretly make architectural decisions or write a complete project profile without AI review.
6. Distinguish observed facts, inferred hypotheses, and human-confirmed facts.
7. Create project-specific files and native agent entrypoints through AI-directed operations.
8. Validate the generated instance and report what is configured, what is missing, what is inferred, and what needs user confirmation.

A small helper script is allowed. A hidden, monolithic, pre-AI scan is not. No mandatory `harness install ./project` command may be required to start onboarding. A CLI can later expose deterministic operations or improve convenience, but AI remains the orchestrator and owner of the onboarding process.

The AI may use a concise bootstrap prompt/skill, an optional global agent skill, a project copy of the harness, or a path to the harness repository. Choose the simplest reliable entry mechanism for the supported executors. Document its usage. If a native agent needs an adapter-specific entrypoint, generate it through the AI-led process.

## 3. Product decisions already approved

- The product is universal and not specialized in a language, framework, cloud, IDE, model vendor, or architecture style.
- It is maintained in a central, version-controlled upstream repository.
- It can be cloned anywhere or carried on a USB drive.
- Installing it into a project creates a local instance whose project-specific information and customizations are committed to that project's Git repository.
- Prefer a managed physical copy/vendor model for project files over requiring a Git submodule. Subtree may be evaluated as an optional alternative, not the default.
- Each installed instance records upstream version, source revision, schema version, integration providers and project profile version.
- Upgrades must be explicit, preserve local changes, detect conflicts, and produce a report. Never silently overwrite project customizations.
- A Master Agent coordinates the work and is primarily a management specialist. Real specialist agents are organized across architecture, engineering, backlog management, and development. A pipeline stage is not necessarily one agent, but an executor without real delegated-agent execution is unsupported.
- The Master may compose new task-specific pipelines or reuse project patterns. Every definition is validated against the authority, safety, evidence, and production-boundary invariants.
- The system includes Spec-Driven Development, curated skills, structural knowledge, persistent memory, model abstraction, observability/evaluation, and evidence-based harness improvement.
- Do not build a bespoke sandbox platform. Reuse the executor's native sandbox/approval controls where available and provide conservative project-level policies. Be transparent where a runtime lacks equivalent controls.
- Harness² is the reference for harness-level experiments and recursive improvement. Do not add a separate coding-benchmark platform as a mandatory dependency in the first implementation.
- Use CodeGraph and Graphify as complementary knowledge providers, with distinct query intent, explicit overlap/freshness checks, and source verification. Their exact adapter/query behavior is a design validation item.
- ai-memory handles persistent episodic/project memory and handoffs. Harness²'s improvement harness is not a replacement for a general memory service.
- ECC and Matt Pocock are sources for curated skills and engineering practices. Do not install both complete collections or duplicate overlapping skills blindly.
- OmniRoute is optional. The core must work with a native/direct model provider and not require a gateway.
- OpenTelemetry is the instrumentation foundation; Phoenix is an optional UI/analysis backend, not a required runtime dependency.
- User-facing work should not be split into endless rounds of tiny prototypes and approvals. Build a coherent integrated implementation, validate it end-to-end, and present one consolidated result. Ask only for genuinely blocking information, permissions, or irreversible decisions.

## 4. Required reading order

Read these files before implementing:

1. `01_PRODUCT_VISION_AND_REQUIREMENTS.md`
2. `02_ARCHITECTURE.md`
3. `03_E2E_PIPELINE.md`
4. `04_AI_FIRST_ONBOARDING.md`
5. `05_COMPONENT_INTEGRATION_CATALOG.md`
6. `06_CONTRACTS_AND_ARTIFACTS.md`
7. `07_PORTABILITY_AND_VERSIONING.md`
8. `08_EXECUTION_POLICIES.md`
9. `09_OBSERVABILITY_EVALUATION_AND_LEARNING.md`
10. `10_IMPLEMENTATION_AND_ACCEPTANCE.md`
11. `11_DECISIONS_AND_OPEN_ITEMS.md`
12. `12_GLOSSARY.md`
13. `13_LUCAS_PRODUCT_VISION_AND_DECISIONS.md` — authoritative addendum for decisions Lucas clarified after the original blueprint
14. `14_ROTEIRO_DE_SESSOES_E_INICIO_DA_IMPLEMENTACAO.md` — design-session sequence and implementation gate
15. `references/SOURCE_CATALOG.md` and `references/INTEGRATION_NOTES.md`
16. `15_MODELO_DE_EXECUCAO_E_ADAPTERS.md` through `20_POLITICA_DE_SEGURANCA_PORTABILIDADE_E_DEPENDENCIAS.md`
17. `21_ARQUITETURA_INTEGRADA_E_PLANO_DE_IMPLEMENTACAO.md`
18. `22_PLANO_DE_AMARRACAO_E_INTEGRACAO_FUNCIONAL.md` before implementing or installing integrations
19. `23_IDENTIDADE_DO_WORKSPACE_E_MODOS_DE_TRABALHO.md` for build-vs-target classification
20. `24_PLANO_MESTRE_DE_CONSTRUCAO_DO_HARNESS.md` for all build phases, work items, dependencies, and readiness gates

If any earlier section in this brief conflicts with Lucas's decisions recorded in `13_LUCAS_PRODUCT_VISION_AND_DECISIONS.md` or their detailed operational contracts in `15`–`22`, follow the later clarified decision and record any still-open implementation choice. These artifacts are design work; do not infer from them that coding is authorized or has started.

## 5. Required behavior before implementation

Perform a focused assessment of your own environment and the requested target location. Do not start with a giant scan. Identify:

- Which coding-agent runtime is available and which native capabilities it supports.
- Whether the harness source directory is writable and whether the target project is a Git repository.
- Available language runtimes and package managers only as needed by implementation.
- Constraints on network access, credentials, local models, and external services.

Inspect current official upstream documentation for the referenced projects before adopting dependencies. Record the date and evidence for compatibility, licenses, install procedures, platform support, and command/API contracts. The reference links are starting points, not permission to assume current details.

Do not copy entire upstream repositories into the product merely because they are in the reference catalog. Prefer narrow integrations, documented adapters, formats, or methods. Respect licenses and attribution.

## 6. Implementation approach

Build a complete, internally coherent system. You may organize the coding work into logical engineering tasks and run unit/integration checks as you go, but do not turn those internal checks into a repeated user-facing approval loop. Do not ask the user to approve every component separately.

Use good engineering practice: versioned schemas, tests, explicit errors, logging, platform-aware paths, safe file writes, idempotent onboarding, rollback/recovery where feasible, and no hidden destructive behavior.

Resolve low-risk implementation details autonomously. If a dependency is optional, provide a graceful fallback. If a requested upstream service is unavailable, document the limitation and keep the core usable.

## 7. Final delivery from the local agent

Deliver:

- A working harness source repository with versioning.
- An AI-first bootstrap/onboarding entrypoint and clear operator instructions.
- At least one complete E2E pipeline and independently invocable stages.
- A project instance created in a representative target repository.
- A manifest and Git ignore policy separating committed project knowledge from local runtime data.
- Adapters for the executor(s) actually supported, with a declared capability matrix.
- Working model interface with optional gateway support.
- Structural knowledge and memory integrations or documented graceful adapters.
- Traces, deterministic verification, evaluation reports, and a controlled learning proposal workflow.
- Unit/integration tests and a consolidated E2E validation report.
- A list of deviations, unavailable integrations, known limitations, and commands/examples to use the product.

Do not claim success for integrations that are only represented by placeholder configuration. Mark each integration as implemented, optional/unconfigured, simulated, or not supported.
