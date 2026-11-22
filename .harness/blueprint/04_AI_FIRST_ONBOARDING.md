# 04 — AI-First Onboarding and Project Specialization

## 1. Non-negotiable requirement

The project installation and specialization process must be managed by an AI agent **from moment zero**. Do not require the user to run a complex scanner, setup CLI, or preflight script before the agent can participate.

The agent is the onboarding orchestrator. It inspects the project, chooses which evidence to collect, runs optional helpers when appropriate, interprets findings, creates the project instance, and validates the result.

## 2. User experience

The intended user experience is conversational:

1. User opens the target repository in a supported coding agent.
2. User points the agent to the harness bootstrap entrypoint (or uses a previously installed global/project skill).
3. User asks the agent to onboard the current repository using the Portable Engineering Harness.
4. The agent explains the plan briefly and begins targeted discovery.
5. The agent asks only questions that require human knowledge or confirmation.
6. The agent creates the project-specific files and adapter instructions.
7. The agent validates the installation and produces a concise onboarding report.

The user should not need to manually edit a dozen files or run a suite of opaque commands.

## 3. Bootstrap entrypoint

Provide one concise, executor-neutral bootstrap instruction file/skill that tells the agent:

- What the harness is.
- Where the core repository/files are.
- How to read the specification and skills.
- How to lead progressive discovery.
- How to invoke optional deterministic helpers.
- How to create and validate `.harness/` and native agent instructions.
- Which actions require approval.
- How to report the result.

Generate thin native entrypoints (for example `AGENTS.md`, `CLAUDE.md`, or provider-specific configuration) through the AI-led onboarding. These files should be short pointers to the versioned project instance, not duplicated giant instructions.

If the harness is on a USB drive or another path, paths must be discovered at runtime and converted to portable references where possible. Do not write absolute machine-specific paths into committed project files.

## 4. Progressive discovery protocol

### Pass 0 — establish boundaries

The agent should confirm the repository root, current Git status, branch/revision, whether there are uncommitted user changes, and whether onboarding may create `.harness/` and adapter files. Do not overwrite existing user files.

### Pass 1 — low-cost orientation

Use native agent tools first:

- List top-level files and directories.
- Read README and existing project/agent instructions.
- Inspect Git metadata and status.
- Identify likely manifests (package files, build files, project files, dependency manifests, lockfiles, CI workflows).
- Detect existing `.harness/`, `.agents/`, `.claude/`, `AGENTS.md`, and other relevant configuration.

Do not read every source file.

### Pass 2 — targeted technical discovery

Based on evidence, inspect only the relevant manifests and representative source/test structure. Identify:

- Languages and runtimes.
- Build, test, lint, format, type-check and package commands.
- Frameworks and major dependencies.
- Main entrypoints and architectural modules.
- Data stores, external integrations, message brokers, API contracts, schemas and migrations.
- CI/CD and deployment workflows.
- Existing testing strategy and gaps.
- Security-sensitive areas and operational constraints.
- Existing ADRs, architecture docs and domain glossary.

Run commands only when safe and useful. Do not execute application code or install dependencies without an explicit reason and appropriate policy.

### Pass 3 — structural intelligence

Decide whether a structural knowledge provider is useful. If yes:

- Explain why graph construction will help this project.
- Use the selected provider (Graphify or an alternative) through a narrow adapter.
- Scope ignored/generated/vendor directories.
- Record source revision, graph version, coverage, and known unsupported languages.
- Store generated indexes/caches outside Git unless a specific project policy says otherwise.
- Do not treat inferred edges as facts.

### Pass 4 — domain and workflow discovery

Use AI reasoning and selected skills to identify project vocabulary, design principles, conventions, and decision history. Use a grilling/interview skill only for material gaps. Generate or update:

- `PROFILE.md` — project summary and observed stack.
- `ARCHITECTURE.md` — architecture evidence and known uncertainties.
- `GLOSSARY.md` — project vocabulary.
- `CONVENTIONS.md` — coding/testing practices discovered.
- `decisions/` — only confirmed durable decisions, not speculative notes.

Do not invent ADRs or claim a particular architecture solely from folder names.

### Pass 5 — integration and configuration

Configure only what is useful and supported:

- Native agent adapter and instruction pointer.
- Selected skills and workflow catalog.
- Pipeline defaults and project quality gates.
- Structural knowledge provider.
- Memory provider (ai-memory if compatible/configured).
- Model provider or optional gateway.
- OpenTelemetry trace export or local fallback.

Missing credentials or unavailable services should be reported as unconfigured, not faked.

### Pass 6 — validate and report

The agent must:

- Read back generated files.
- Validate schemas/configuration.
- Confirm ignored runtime paths are actually ignored.
- Confirm native agent entrypoints point to the right instructions.
- Run harmless self-checks or a small representative discovery/verification action.
- Report created/modified files, detected project facts, evidence, assumptions, missing integrations, user decisions required, and how to run the pipeline.

## 5. Use of scripts and deterministic helpers

Scripts are permitted as **subordinate tools** when the agent decides they are useful. Examples:

- List tracked files while respecting `.gitignore`.
- Read package/build manifests into structured metadata.
- Compute file hashes/revision fingerprints.
- Run known build/test commands.
- Invoke Graphify's deterministic AST extraction.
- Validate a YAML/JSON schema.
- Compare managed files during upgrade.

Requirements:

- Small, documented, inspectable, and scoped.
- Explicit input and output contracts.
- No hidden network calls or package installation.
- No destructive changes without approval.
- No automatic upload of source or traces.
- The agent explains material operations and interprets their results.
- A helper must fail safely and report uncertainty; it must not make architectural decisions.

Forbidden onboarding design:

- One massive script that scans the entire repo, calls multiple services, edits files, generates architecture conclusions, and installs everything before AI review.
- A mandatory CLI installer that must complete before an agent is involved.
- Unbounded recursive scans or indiscriminate source-file ingestion.
- Automatic installation of all upstream projects and their dependencies.

## 6. Project profile confidence model

Each significant claim should use one of:

- `OBSERVED`: direct evidence (file, command, test, code reference).
- `INFERRED`: plausible conclusion, with evidence and confidence.
- `CONFIRMED`: user or owner explicitly validated it.
- `UNKNOWN`: not enough evidence.

Example:

```yaml
architecture:
  style:
    value: hexagonal
    status: INFERRED
    confidence: medium
    evidence:
      - internal/domain/
      - internal/adapters/
    notes: "Naming and dependency direction suggest ports-and-adapters; not yet confirmed."
```

## 7. Project onboarding is not a migration

Onboarding must not:

- Reformat the project.
- Change architecture.
- Install new frameworks.
- Add tests merely to make a score look better.
- Rewrite existing documentation wholesale.
- Modify CI/CD without an explicit project need and approval.

It should first document what exists and identify optional opportunities separately.

## 8. Existing and empty repositories

### Existing / legacy project

Discover actual behavior and constraints. Preserve current conventions unless there is a clear, approved reason to change them. Explicitly record gaps such as missing tests or incomplete documentation.

### Empty/new project

Use the project owner's stated goals and stack decisions to initialize the profile. Mark decisions as confirmed only after confirmation. Do not infer a stack from the harness defaults.

## 9. Re-onboarding / refresh

A later AI-led refresh compares current evidence with the stored profile and graph metadata. It produces a proposed diff, distinguishes observed changes from hypotheses, and updates only approved project facts. Refresh must be incremental where possible and must not silently replace human-maintained documentation.
