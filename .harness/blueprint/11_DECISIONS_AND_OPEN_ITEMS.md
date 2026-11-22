# 11 — Decisions and Open Items

## 1. Approved decisions (treat as requirements)

| ID | Decision |
|---|---|
| D-001 | Product is a universal, project-agnostic portable engineering toolbox. |
| D-002 | AI leads onboarding from the first interaction. No mandatory pre-AI scanner or opaque install script. |
| D-003 | Deterministic scripts/helpers are allowed only as scoped tools under AI orchestration. |
| D-004 | Pipeline is the named E2E task-execution flow. |
| D-005 | Pipeline stages are independently invocable units, not necessarily individual agents. |
| D-006 | The Master Agent owns coordination; genuine specialist agents form four pillars, with task-dependent invocation and no persona fallback. |
| D-007 | The harness core is version-controlled upstream and can be cloned or carried on USB. |
| D-008 | Project instances are copied/managed inside each project and versioned with that project's Git repository. |
| D-009 | Project-specific knowledge/customization is separate from generic core. |
| D-010 | Prefer managed physical copy/vendor distribution; submodule is not the default. |
| D-011 | Do not specialize the core in a language, framework, architecture, or model provider. |
| D-012 | ECC and Matt Pocock are curated sources; do not install both full collections blindly. |
| D-013 | CodeGraph and Graphify are both intended as complementary providers; measure their overlapping code domain and validate their distinct roles. |
| D-014 | ai-memory is persistent project/session memory; Harness² is harness-level improvement/experimentation. |
| D-015 | Spec-Driven Development is part of the core workflow. |
| D-016 | OpenTelemetry-compatible observability and evaluation are included. Phoenix is optional. |
| D-017 | No separate coding-benchmark framework is mandatory in the first version; Harness² provides an evaluation/benchmark reference. |
| D-018 | Do not build a custom sandbox platform initially; reuse native executor controls and add minimal policies. |
| D-019 | Superceded on 2026-11-20: OmniRoute is the required local/loopback gateway for initial-release model calls; direct provider is deferred. |
| D-020 | No endless user-facing micro-prototype/approval loop. Deliver a coherent integrated implementation and consolidated E2E validation. |
| D-021 | The final documentation package is the handoff to the local agent; it should not require the local agent to reconstruct decisions from conversation history. |

## 2. Open technical decisions for the implementation agent

The following are intentionally open; resolve them using current official documentation and local environment evidence. Do not reopen approved product decisions.

1. **Core implementation language:** choose the simplest cross-platform option for the selected executor integrations. Document why.
2. **Pipeline representation:** YAML/JSON/Markdown-driven or code-driven with schemas. Prefer declarative definitions with validation and versioning.
3. **Graph integration:** validate CodeGraph and Graphify together for accuracy, complementary query domains, language/input coverage, incremental update, license, Windows/Linux/WSL2 behavior, and local indexing. Define degraded behavior when one provider is unavailable; do not silently replace the pair with a single provider.
4. **Memory provider:** confirm ai-memory current protocol, install path, security, native Windows/WSL2 support and operational model; provide an adapter/fallback if unavailable.
5. **Phoenix deployment:** choose optional local mode/export target; no mandatory hosted service.
6. **Harness² integration:** decide whether direct integration is worth the dependency or whether to implement a narrow Propose–Probe–Compose workflow inspired by its published method. Be precise in the product's claims.
7. **Executor adapters:** implement the first executor actually available; document others honestly. The design must permit additional adapters.
8. **Model routing:** install and smoke-test OmniRoute locally with a pinned version, a secret-backed explicit upstream provider/model allowlist and capability profiles; direct provider remains deferred.
9. **Project artifact paths:** choose a clean layout and ensure it does not collide with existing project conventions.
10. **Windows distribution:** verify native Windows support for every dependency and provide WSL2 guidance where necessary.

## 3. Explicitly rejected defaults

- One-to-one mapping of agents to pipeline stages; specialists are required, but stages choose relevant roles.
- Full ECC plus full Matt Pocock skill installation.
- Neither provider may claim complete/authoritative project knowledge without source provenance and freshness evidence.
- ai-memory as the only memory of harness improvements.
- Harness² as the replacement for project tests or general session memory.
- Automatic OmniRoute fallback that changes the upstream provider/model or its approved data class without an explicit allowlist.
- Phoenix as mandatory runtime service.
- Massive pre-AI repository scanner.
- Complex bespoke sandbox service.
- Separate benchmark platform as mandatory first-version dependency.
- Automatically modifying all project architecture/conventions during onboarding.
- Automatic promotion of self-generated skills/policies.
- Silent upstream synchronization into project copies.
