# 01 — Product Vision and Requirements

## 1. Product statement

Build a **Portable Engineering Harness**: a version-controlled, project-agnostic toolbox that lets an AI coding agent conduct disciplined software engineering in any repository. The agent leads onboarding from the beginning, learns the target project's actual conventions and structure, creates a project-specific configuration, and executes tasks through a traceable E2E pipeline.

The harness should make a coding agent operate like a careful engineering team: clarify intent, inspect the system, make a plan, implement in controlled steps, verify the result, preserve useful knowledge, and improve its methods through evidence.

## 2. Goals

1. **Portable:** one upstream repository can be cloned or carried on removable media and installed into many unrelated projects.
2. **Universal:** the core does not assume Go, JavaScript, Python, a specific architecture, framework, IDE, LLM, or cloud.
3. **AI-first:** the AI agent is in charge from the first onboarding interaction. Deterministic scripts are subordinate helpers, never a hidden precondition.
4. **End-to-end:** a task can move from intent to verified deliverable with linked artifacts and clear status.
5. **Composable:** a pipeline can be run end-to-end or individual stages can be invoked directly.
6. **Portable across executors:** native adapters expose the same logical capabilities through each agent's conventions.
7. **Context-aware:** relevant project structure, docs, decisions, history, and procedures are assembled only when useful.
8. **Traceable:** requirements, plans, changes, tests, reviews, and decisions are linked.
9. **Safe by default:** use existing executor controls and least privilege; avoid unnecessary custom security infrastructure.
10. **Measurable:** record execution traces and deterministic/semantic evaluation results.
11. **Continuously improvable:** experiment with changes to skills, prompts, routing, or workflows using controlled comparisons and approval.
12. **Maintainable:** avoid needless services, agents, databases, vector stores, message brokers, and overlapping frameworks.

## 3. Non-goals

- Building a new foundation model or fine-tuning service.
- Building an independent IDE or coding agent that replaces Codex/Claude Code/OpenCode/Kilo.
- Creating a bespoke remote sandbox or container orchestration platform.
- Automatically installing every upstream repository in the catalog.
- Forcing one architecture such as hexagonal, DDD, CQRS, microservices, or a particular testing strategy onto every project.
- Scanning the entire repository with a large, opaque script before an AI agent participates.
- Requiring a one-to-one mapping between pipeline stages and agents; the product does require real specialists grouped in its four pillars and a master coordinator.
- Automatically adopting self-generated improvements without evaluation and approval policy.
- Requiring a vector database or hosted memory service.
- Requiring OmniRoute, Phoenix, Graphify, or ai-memory to be operational for the basic pipeline to run; integrations should degrade gracefully where practical.
- A long user-facing cycle of approving isolated prototypes before the integrated system exists.

## 4. Core product principles

### 4.1 AI leads, deterministic tools assist

The AI agent performs the reasoning and orchestration: what to inspect, what to ask, which helper to invoke, how to interpret evidence, what project artifacts to create, and whether onboarding is complete. Deterministic helpers may collect factual data or execute operations, but they must be small, transparent, scoped, and invoked by the agent for a clear reason.

### 4.2 Progressive discovery, not blind full scans

Start from cheap, high-signal evidence: repository root, Git status, README, top-level file list, package/build manifests, CI workflows, and existing agent instructions. Follow references and inspect relevant files. Build a broader index or knowledge graph only when the agent has a reason to do so and understands the operation.

### 4.3 Evidence and provenance

Every important project-profile claim should be classified:

- `OBSERVED`: directly read from a file, command result, CI configuration, or source structure.
- `INFERRED`: a reasoned hypothesis based on evidence; it must not be presented as certainty.
- `CONFIRMED`: explicitly validated by the user or project owner.

Important claims should carry source paths, symbol references, command evidence, timestamps, or revision identifiers where applicable.

### 4.4 One source of truth per concern

Avoid duplicate systems. A project fact should have a canonical owner. A skill should have one canonical selected version per capability. A requirement should have one canonical specification. A pipeline stage should have one declared contract. Graphify and CodeGraph are both intended in the integrated design; assign complementary query domains and preserve distinct provenance rather than treating either as the sole index.

### 4.5 Main agent first, delegation when justified

The Master Agent owns coordination, task state, and convergence. The Harness uses genuine specialist-agent executions organized into the architecture, engineering, backlog-management, and development pillars. Which specialists run for a stage depends on the problem; stage and agent are not one-to-one. An executor without real delegated-agent support is unsupported.

### 4.6 Artifacts over conversational assumptions

The pipeline must create durable artifacts that allow a different session or model to continue. Conversation history is useful but never the only source of truth.

### 4.7 Bounded autonomy

Set budgets, deadlines, retry limits, tool permissions, and escalation conditions. When the agent is stuck, it should stop and report evidence rather than loop indefinitely.

## 5. Functional requirements

### FR-01 — AI-first bootstrap

A human can open a target project in a supported coding agent and tell the agent to use the harness. The AI agent must lead onboarding without a mandatory pre-AI CLI installer or monolithic scanner.

### FR-02 — Project discovery

The agent performs progressive discovery and can use native file/search/Git/terminal tools plus optional bounded helpers. It identifies languages, build/test commands, frameworks, directory structure, architecture evidence, CI, conventions, and known gaps.

### FR-03 — Project profile

Generate a project-local profile with claims, status, provenance, and unresolved questions. Never silently turn an inference into a confirmed rule.

### FR-04 — Pipeline engine

Define versioned pipelines made of stages, gates, transitions, budgets, and output contracts. Support end-to-end invocation, direct stage invocation, resumption, and controlled retry.

### FR-05 — Stage independence

Each stage can be invoked on its own if its preconditions are met. A stage is a workflow unit, not necessarily a separate agent process; however, the responsible specialist roles must be invoked as real agents when the stage requires their expertise.

### FR-06 — Skills and procedures

Support a portable skill format (prefer Agent Skills conventions) and select compatible skills from a curated catalog. Expose executor-specific discovery through adapters.

### FR-07 — Structural knowledge

Offer provider-neutral interfaces for structural code and cross-artifact knowledge. Use CodeGraph and Graphify as complementary providers, subject to validation of overlapping domains, freshness, provenance, Windows/Linux behavior, and local indexing. Store derived graph outputs with source revision metadata.

### FR-08 — Persistent memory

Integrate ai-memory where compatible, or provide a documented memory adapter. Separate project facts, episodic experiences, handoffs, and harness-learning evidence.

### FR-09 — Model abstraction

Use a provider-neutral model interface. Direct provider use must work. OmniRoute is optional and replaceable.

### FR-10 — Execution policy

Declare allowed tools, file scope, network policy, approval requirements, and resource budgets. Reuse native sandbox/approval features of the executor when available.

### FR-11 — Verification and acceptance

Run project-native build, test, lint, type-check, and other checks discovered during onboarding. Evaluate acceptance criteria and report evidence. Do not claim success when mandatory checks fail or are unavailable without explicit qualification.

### FR-12 — Observability

Instrument pipeline and tool calls using OpenTelemetry-compatible traces where supported. Capture run IDs, stage IDs, model/provider, tool, timing, status, retries, and token/cost metrics where available. Redact secrets and sensitive payloads.

### FR-13 — Evaluation

Combine deterministic checks with optional semantic grading. Preserve task inputs, expected conditions, evaluation results, and evaluator versions.

### FR-14 — Controlled learning

Use Harness²-inspired Propose–Probe–Compose and persistent improvement-harness concepts to test changes to the harness, not to mutate production skills automatically during ordinary feature work.

### FR-15 — Portable installation

Copy managed harness files into the project, record upstream version/revision and schema versions, generate native agent pointers, and keep local runtime state separate from committed project artifacts.

### FR-16 — Safe upgrade

Compare managed files, detect user modifications, preserve overrides, and present a clear upgrade report. No silent overwrite.

### FR-17 — Executor adapters

Each adapter declares supported capabilities, generated files, limitations, and fallback behavior. Never claim feature parity unless verified.

## 6. Non-functional requirements

- Works with relative paths and relocatable directories; no hard-coded machine paths.
- Core workflows remain usable when optional integrations are absent.
- Windows is a first-class target to evaluate, along with Linux and macOS.
- Native Windows support of third-party dependencies must be verified; WSL2 may be a documented compatibility route, not an invisible requirement.
- No secrets in Git, traces, project profile, memory exports, or generated prompts.
- All writes to project files are reviewable and use safe/atomic update patterns where possible.
- Schemas are versioned and migration behavior is documented.
- Every execution has a durable run ID and status.
- Retries and loops are bounded.
- Dependencies and licenses are inventoried.
- Documentation includes setup, recovery, migration, update, uninstall, and troubleshooting.

## 7. Success criteria

The system is successful when a user can introduce it to an unfamiliar repository through an AI agent, receive a useful evidence-backed project profile and native integration, run a complete development pipeline, resume in another session or supported agent, verify the result, and inspect how the run behaved—without manually assembling prompts or invoking a large preflight installer.
