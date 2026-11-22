# 13 — Lucas's Product Vision and Decisions

**Record type:** conversation-derived product decision record and design addendum  
**Author / product owner:** Lucas de Lima  
**Captured:** 2026-11-19  
**Status:** vision and decisions captured; implementation has not started

## 1. Authority and purpose

This document preserves product decisions and intent clarified by Lucas after review of the original blueprint. It is part of the specification bundle and must be read before implementation. It supplements the original blueprint; where it makes an explicit decision that narrows or changes an earlier open item, this record is authoritative. It does not convert proposals or unresolved questions below into approved requirements.

The goal is to prevent implementation from depending on chat history. Future sessions should use this record and the other blueprint files as the source of truth, update this record when Lucas makes a durable decision, and preserve the distinction between confirmed decisions, proposals, and open questions.

## 2. Product vision

Build a high-capability, portable AI-first software engineering harness that can take a software problem from early ideation through system design, engineering definition, backlog decomposition, implementation, verification, and preparation for customer or project homologation.

The harness should perform substantial architecture and engineering analysis itself, including selecting which relevant design alternatives deserve consideration, explaining trade-offs and residual risks, and producing artifacts that allow implementation to proceed. Lucas should evaluate the important decisions and deliverables instead of manually revisiting all architecture and engineering concepts throughout the project.

The intended product behaves like a software delivery organization serving a client: it includes a coordinating leader and specialists for the work performed by architecture, engineering, product/backlog, development, testing, quality, review, and homologation roles. The role model is represented by actual specialist agents, not by a single master agent merely changing personas.

Lucas is the first user and will decide when the harness is ready to distribute. The initial system must already provide a coherent, capable end-to-end lifecycle; the four pillars below are not a reduced demonstration scope. The long-range ambition is very broad, including solving complex problems at the scale of major technology companies. That ambition is a north star, not a verifiable first-release success percentage.

## 3. Confirmed operating model

### 3.1 The four pillars

The lifecycle has exactly four primary pillars:

1. **Architecture** — understand the problem and relevant context, assess suitable architectural and system-design alternatives, recommend a small set of strong candidates, and produce the architecture decision and supporting artifacts.
2. **Software engineering** — translate the selected architecture into implementable technical decisions, interfaces, constraints, quality requirements, and engineering artifacts.
3. **Backlog management** — translate the approved engineering and product outcomes into epics, features, user stories, and tasks, with acceptance criteria and dependencies.
4. **Development** — implement the selected epic or feature end to end, including test-first work, verification, independent review, quality assurance, integration and end-to-end testing, release preparation, and homologation.

Each pillar is a pipeline and may contain multiple specialist agents. A pillar is not synonymous with a single agent. Its composition depends on the problem and can include, for example, database, backend, frontend, data, security, reliability, infrastructure, testing, product, or review specialists.

### 3.2 Master agent and delegation

- A **Master Agent** coordinates specialist agents and pipelines. Its primary strength is leadership, planning, delegation, state management, and resolving where work goes next; it is not expected to replace deep architecture or engineering expertise.
- The Master Agent may coordinate a master pipeline that composes the four pillar pipelines and any needed project-specific or temporary pipelines.
- Specialists must be real delegated agents with their own responsibility and scope. A host that cannot support the required subagent model is not a supported execution environment for the intended system. The Master Agent must not silently substitute role-playing personas for unavailable specialists.
- The main agent should consult the harness's agents, project artifacts, and memory before bringing a question to Lucas. Human-in-the-loop checkpoints remain necessary when product intent or another material fact cannot be inferred responsibly.
- The Master Agent chooses the next stage within the pipeline that Lucas invoked. When Lucas asks for a bounded feature or epic, the run ends at that scope and waits for him to select the next one; it does not continue into unrelated backlog items automatically.

### 3.3 Architecture alternatives and lock

- Analysis is problem-sensitive. The architect should consider the patterns and architectures relevant to the actual problem, rather than enumerate irrelevant options. For example, do not add messaging or distributed architecture when a simpler design adequately solves the problem.
- Present a limited set of strong alternatives with selection criteria, benefits, disadvantages, residual risks, mitigations, and points Lucas should monitor. The recommendation should minimize disadvantages where practical, not merely list them.
- Lucas chooses the architecture. The accepted design becomes a versioned, canonical architecture baseline and governance constraint for all later agents.
- Developers and engineers first try to solve implementation issues within the approved architecture. Escalation flows from development to engineering, then to architecture when the lower layer cannot resolve the problem.
- Architecture is not changed silently. If the architect cannot resolve the issue within the baseline, the architect must assemble concrete evidence and a viable proposed solution, then ask Lucas to authorize reopening the decision. Only Lucas can authorize that reopening.

### 3.4 Development quality

Development means more than writing application code. The intended development pipeline can include:

- Test-first development: establish a failing test before the implementation, then implement and make it pass.
- Relevant unit, integration, contract, browser/end-to-end, lint, type, build, and other project-appropriate checks.
- Frontend browser verification such as Playwright when appropriate, beyond unit tests alone.
- Backend integration verification and end-to-end validation where the system warrants it.
- Independent code review and QA; do not make passing checks equivalent to independent acceptance.
- DevOps/CI and release-readiness work needed to prepare an artifact for homologation.
- Customer or project homologation before any configured production change-control process.

The specific agents and gates are selected for the system; this list does not require every project to use every testing technology.

### 3.5 Homologation, GMUD, and production boundary

- The harness may prepare work for a homologation environment or release candidate. It does **not** deploy to production under the current product boundary.
- **GMUD** is a post-homologation audit/change-authorization process that, when required by a client or project, confirms the change may proceed to production deployment. It is not a Git branching strategy.
- A PR or record named, for example, `GMUD Épico XYZ` may organize the evidence and link the exact commit hashes homologated for that epic. This naming convention is traceability/organization, not a branch architecture.
- GMUD is conditional on the client's/project's process; Lucas does not require it for every one of his own systems. A configured GMUD workflow must record the relevant homologation evidence and wait for the required authorization. Production deployment remains outside the harness's current responsibility.

## 4. Installation, project state, and legacy onboarding

### 4.1 Install mode versus normal operation

The Master Agent must recognize whether a request is:

- installing/onboarding the harness in a repository that has not yet been analyzed;
- operating a repository whose harness installation and initial analysis already exist; or
- resuming or explicitly refreshing an existing installation.

The user-facing signal can remain a simple toggle. The Master Agent cross-checks it against durable evidence such as the harness manifest, project-analysis artifacts, and their timestamps. If the toggle indicates first installation but a valid installation already exists, it treats the signal as mistaken and avoids overwriting or repeating the completed onboarding. If evidence shows partial onboarding, it resumes from the recorded state instead of starting blindly.

### 4.2 Legacy project install-only request

Installing the harness into a legacy system does not itself authorize a modernization project. When the user only asks to install/analyze:

1. Discover the existing system progressively and record the **as-is** architecture as the baseline.
2. Analyze and document its benefits, liabilities, uncertainties, and comments with supporting evidence.
3. Produce the relevant artifacts, project memory, logs, and next-step handoff.
4. Stop and wait for a concrete problem or transformation request.

If a problem is supplied as part of the request, the pipelines may proceed to design alternatives and a target architecture for Lucas to decide. The as-is baseline remains historical evidence; it is not silently replaced by a proposed target.

## 5. Adaptive pipeline system

Pipelines must be composable and extensible. The harness must not hard-code a closed set of flows or prevent new pipelines from existing.

- The four pillars and any built-in end-to-end workflows are reusable patterns/templates, not mandatory routes for every problem.
- The Master Agent may select an existing pipeline, compose several pipelines, create a project-reusable pipeline, or create a temporary pipeline for one task.
- A generated pipeline records its purpose, scope, stages, responsible agents, inputs, outputs, prerequisites, acceptance/completion conditions, failure/escalation paths, applicable gates, and provenance/version.
- Generated pipeline definitions are validated before execution against harness contracts, including bounded retries/loops and architecture/security/approval invariants.
- “Temporary/discardable” means it need not remain an active reusable workflow. The version actually executed and its run history remain recorded so work can be audited and resumed.
- Project-reusable pipelines belong to that project's specialization. A generally reusable improvement to the upstream harness is a candidate for evaluation and controlled promotion, not an automatic global change.

The enduring invariants belong to the harness; the exact route through stages adapts to the task, project, and client process.

## 6. Memory, logs, and continuity

- Each project has canonical, persistent project memory alongside its accepted architecture and project artifacts. The architectural baseline, durable decisions, engineering constraints, product terms, and approved backlog must not exist only in chat or in an episodic memory service.
- Large portfolios may keep shared documentation in a separate repository when project scale warrants it. A separate repository is optional, not the default for every project.
- Task, feature, app, and project work needs scoped memory and handoff data. Do not make a single ever-growing `memory.md` per agent the only continuity mechanism.
- Record task/iteration logs with timestamps, what was done, outcomes/evidence, decisions or blockers, and the next concrete task. Agents should use these records to resume without replaying the full prior conversation.
- Every agent may read the shared project context. Write access is scoped to the agent's responsibility; changes to canonical architecture or another protected artifact follow its owner/approval rules.
- Episodic memory, canonical project decisions, code structure, telemetry, and harness-learning evidence are separate concerns that may link to one another.

## 7. Knowledge and integrations: intended roles

The desired stack uses each source/tool for its strongest role, with a narrow integration boundary and explicit provenance.

| Source/capability | Intended contribution | Boundary |
|---|---|---|
| CodeGraph — Colby McHenry (`https://github.com/colbymchenry/codegraph`) | Local code structure, symbols, call/import relationships, context retrieval, and change impact | Code intelligence; graph results do not replace source verification where freshness or unsupported constructs matter. |
| Graphify (`https://github.com/Graphify-Labs/graphify`) | Relate code with documentation, SQL schemas, configurations, and other supported artifacts | Preserve extracted/inferred provenance. Business rules remain supported by project-owned evidence; a graph is not authority by itself. |
| ai-memory | Episodic/project-session memory and handoffs | Does not replace canonical project architecture and decisions. |
| Harness² | Evidence-based experiments and improvement of harness agents/skills/policies | Not the project's general or episodic memory. |
| Anthropic public material | Broad knowledge source for agents about agent design, orchestration, tools, context, evaluation, and other software/agent-engineering topics Anthropic publishes | Attribute sources, preserve dates and product-specific scope, and do not let a vendor publication override Lucas's product decisions. |
| ECC and Matt Pocock skills | Curated procedures and engineering practices | Select/adapt useful material; avoid blindly duplicating complete overlapping catalogs. |
| GitHub Spec Kit | Selected SDD concepts and artifacts | The harness owns the pipeline; do not create a competing parallel workflow. |
| OpenTelemetry / optional Phoenix | Execution observability and analysis | Not project memory or proof of task correctness. OpenTelemetry adoption remains under evaluation. |
| OmniRoute | Gateway obrigatório para chamadas de modelo no caminho inicial de release | Todas as chamadas de modelo passam pelo gateway local; a rota upstream exata é allowlist local por provider/modelo e classe de dados. Adapter direto não faz parte do caminho inicial. |

### Anthropic knowledge collection

Lucas wants the harness's knowledge library to use the intelligence Anthropic publishes publicly across relevant topics, not merely the single subagent article and not only material about Claude Code. This is a knowledge source for the harness's specialists. Relevant material should be findable at the point of use and carry a source URL, author/publisher, date/version where available, topic, and notes about general versus vendor-specific applicability. The harness is not affiliated with Anthropic.

Anthropic's article about folding some subagents into a main loop is one input to orchestration design, not an instruction to remove Lucas's specialist-agent organization. Lucas's approved product structure takes precedence over opinions in any external source.

## 8. Portability, dependencies, and autonomy

- The upstream harness is maintained as a repository and may be cloned or carried on removable media. A managed physical/vendor-style project instance is preferred over requiring submodules.
- Windows is Lucas's primary environment; Linux must be well supported. WSL2 is available as a compatibility option. Cloud and IDE-specific integrations are not current priorities; terminal support is lower priority.
- The initial integrated system must include its required specialist agents and a genuine multi-agent execution environment. An executor without subagent support is not considered adequate for this product.
- Keep projects vanilla where possible. Do not add a dependency without a real architectural/engineering need. Avoid obscure or untrusted packages; assess source, maintenance, licensing, release integrity, ecosystem use, and dependency risk. Familiarity or popularity alone is not proof of safety.
- Lucas delegates the concrete security/privacy policy to the implementation design. His product intent allows broad agent autonomy for normal engineering work, including useful network access, memory writes, and dependency operations, subject to the project's dependency trust controls and protected architecture decision. Production deploy is outside the harness boundary.
- The system should minimize repetitive clarification. Agents first use project context, approved knowledge, other specialists, and available evidence. Ask Lucas when unresolved product intent or an architectural reopening decision materially blocks responsible progress.

## 9. Distribution and project specialization

- Lucas maintains the upstream harness and decides when it is ready for wider distribution.
- Downstream users may clone it, install it in a new project, or place it over a legacy project.
- Upstream harness updates can be identified and incorporated into an installed project instance without silently destroying its project-specific specialization.
- Project-specific agent/skill customizations and artifacts are maintained in the project specialization. A downstream maintainer is responsible for adapting those specializations when upstream changes; they are not blindly overwritten as if they were generic defaults.

## 10. First validation target and success

Lucas has a private, early-ideation system he intends to use, but will not disclose its details in this discussion so the harness is not tailored to a single undisclosed system. It currently has a problem/opportunity and possible business value but no selected architecture.

Validation should demonstrate the four-pillar flow on a representative problem while keeping the system's details private or using a sanitized/synthetic case until Lucas chooses otherwise. The strongest early signal is that architecture/design alternatives and engineering artifacts are sufficiently useful to save substantial time, followed by correct backlog decomposition, tested development, and delivery to homologation.

The long-term “95% of Microsoft/Google-scale problems” statement is an ambition, not a first-version acceptance metric. Evaluation requires representative tasks and evidence across the product lifecycle.

## 11. Proposed decomposition — not yet approved as a final agent roster

These roles are candidates to help decompose each pillar. Lucas approved specialist agents and multi-agent pillars, but has not yet approved every role name, boundary, or exact count.

- **Master:** program/delivery coordinator, state and pipeline manager, delegation and escalation.
- **Architecture:** problem/domain analyst, solution/system architect, quality-attribute/security architect as needed, alternative reviewer, decision-record author.
- **Engineering:** specialists in platform, backend, frontend, data/database, security, integration, reliability, and other needed engineering disciplines.
- **Backlog:** product owner/product analyst, backlog decomposer, dependency/acceptance manager.
- **Development:** implementers by relevant technical area, test-first/test specialist, QA, independent code reviewer, integration/E2E specialist (for example Playwright where appropriate), DevOps/release-readiness specialist, homologation/client perspective when needed.

Pillar membership should be selected by task; this list is not a requirement that every project use every specialist.

## 12. Open decisions and follow-ups

These remain unresolved; do not silently assume an answer:

1. Exact initial executor(s), agent launch mechanisms, and adapter matrix for Windows and Linux while preserving mandatory specialist delegation.
2. Final specialist roster, expertise boundaries, handoff schemas, and which specialists can run in parallel within each pillar.
3. Concrete artifact/schema set for architecture alternatives, approved architecture baseline, engineering design, backlog hierarchy, development stages, QA, homologation, GMUD, and pipeline definitions.
4. Exact project memory/log directory layout, log retention, and how large multi-system portfolios link project repositories to a central documentation repository when used.
5. How CodeGraph and Graphify indexes are separately queried/combined, how overlapping code indexing is controlled, and how freshness is verified in Windows/WSL2.
6. Versioned collection/update workflow for Anthropic publications and other external knowledge, including copyright/licensing constraints and the retrieval/index strategy.
7. Dependency trust rubric and which operations may be automatic versus blocked/escalated.
8. How client-specific homologation and GMUD requirements are represented, what evidence a GMUD package needs, and how its approval is verified. No generic branch model is implied.
9. Pipeline generation/validation schema, ephemeral artifact location, lifecycle, and criteria for promoting a useful project pipeline into an upstream reusable pattern.
10. Representative evaluation suite and objective acceptance thresholds for the initial high-capability lifecycle.

## 13. Explicit corrections to prior interpretations

- **GMUD is a post-homologation change audit/authorization gate when a client/project requires it. It is not Gitflow, Git branching strategy, or an environment branch model.** `GMUD Épico XYZ` is an organizational PR/record label linking homologated commit hashes.
- **No fixed, closed set of pipelines.** The harness provides pillars, primitives, contracts, templates, and invariant controls; the Master Agent may compose temporary or project-reusable pipelines.
- **No persona fallback for mandatory specialists.** Supported executors must provide genuine subagent execution.
- **Legacy installation can be analysis-only.** Record and assess the as-is architecture, create artifacts and handoff, then stop if no software problem has been requested.
- **Gateway e política de dados iniciais.** As chamadas de modelo do caminho inicial passam pelo OmniRoute local/loopback; não há chamada direta de provider nesse caminho. Material `synthetic`, `public`, `internal` e `confidential` pode ser enviado quando necessário, mas o context assembler deve minimizar escopo/custo e cada rota upstream provider/modelo deve estar explicitamente allowlisted para aquela classe. Segredos, credenciais, material de autenticação e dados regulados continuam proibidos por padrão. Se não houver rota elegível, a chamada falha sem fallback silencioso.

## 14. Operational design artifacts produced 2026-11-19

The following documents turn the confirmed vision into proposed operational contracts without treating proposals as Lucas-approved implementation choices:

- `15_MODELO_DE_EXECUCAO_E_ADAPTERS.md` defines genuine delegation and a runtime capability test. Exact initial executor(s) remain open.
- `16_PILARES_AGENTES_E_PIPELINES.md` defines the four-pillar vocabulary, candidate role families, dynamic-pipeline invariants and proportional selection. Exact agent roster and executable schema remain open.
- `17_CICLO_DE_VIDA_E_GOVERNANCA.md` describes install/reconcile/resume, legacy baseline, architectural reopening, homologation and GMUD.
- `18_MEMORIA_E_CONTRATOS_DE_ARTEFATOS.md` proposes project-canonical memory, append-only logical events, run/evidence/handoff contracts and optional episodic memory.
- `19_INTEGRACOES_E_ARQUITETURA_DE_CONHECIMENTO.md` records CodeGraph + Graphify as intended complementary providers, with overlap/freshness evaluation; and Anthropic material as a broad, attributed, curated knowledge source.
- `20_POLITICA_DE_SEGURANCA_PORTABILIDADE_E_DEPENDENCIAS.md` proposes adaptive risk controls for high autonomy, dependency trust and Windows/Linux/WSL2.
- `21_ARQUITETURA_INTEGRADA_E_PLANO_DE_IMPLEMENTACAO.md` integrates the components and identifies feasibility boundaries, acceptance criteria and implementation phases.
- `22_PLANO_DE_AMARRACAO_E_INTEGRACAO_FUNCIONAL.md` defines a shared capability/provider/context layer, per-source installation modes, integration waves and collaboration smoke scenarios. It distinguishes tools, skills, knowledge sources, observability and evaluation.

Official primary documentation was checked on 2026-11-19 for current multi-agent API patterns, Claude Code subagents, CodeGraph, Graphify, Agent Skills, ai-memory, ECC, Matt Pocock Skills, Harness², Spec Kit, OmniRoute and Phoenix. See `references/SOURCE_CATALOG.md`. These checks are not proof that the Harness adapters are implemented or that every platform/version is compatible.

## 15. Current status distinction

**Confirmed by Lucas:** the product vision and corrections in sections 1–13.  
**Proposed by the design documents:** schemas/layouts, initial capability policy, query routing between graphs, and phased technical decomposition.  
**Still needs a concrete choice or test before coding:** instalação e smoke test do gateway OmniRoute e de uma rota upstream localmente configurada, final schemas, representative validation environment/thresholds, and support scope for multiple executors in the first release.  
**Not started:** implementation code for the Harness.

## 16. Integration clarification

Lucas wants the tools, skills and knowledge sources to collaborate and be available across the specialist agents. The proposed bridge is a shared capability/provider registry plus executor-aware access, context/provenance assembly and shared project/run artifacts. This does not assume that upstream tools communicate directly with one another. “Available to all agents” means discoverable and callable for relevant work through each supported executor; it does not mean every integration is loaded into every prompt or that every role receives every credential.

The design assumes “YA Memory” refers to the catalog's `ai-memory`. If Lucas intended a different project, this name must be corrected before integrating that dependency. Current upstream documentation marks native Windows as experimental; WSL2/Linux-hosted operation should be evaluated for Lucas's Windows setup before choosing a deployment mode.

## 17. Agent entrypoint clarification

A short package-root `AGENTS.md` now routes coding agents to this decision record, the start brief, and the integration plan. It is guidance for the blueprint package. The development model should create the separate root `AGENTS.md` for the executable Harness repository and adapter-specific entrypoints as part of implementation; the product cannot rely on Codex's file convention as its only cross-executor interface.

## 18. Build-versus-target distinction

Lucas clarified that current work is constructing the Harness itself—the system that will later help develop other systems. Until release readiness is approved and recorded, a code session in the Harness source repository must use `work_mode=harness_build` and take its next work item from `24_PLANO_MESTRE_DE_CONSTRUCAO_DO_HARNESS.md`. It must not assume it is developing Lucas's undisclosed product or another target system. The separate readiness lifecycle and workspace identities are specified in documents 23 and 24; `BLUEPRINT_STATUS.yaml` records the current package as blueprint-only, with implementation not started and no target active.

