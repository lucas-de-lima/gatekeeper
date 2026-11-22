# Build Memory — Portable Engineering Harness

## Canonical context

- Workspace: `engineering-harness-blueprint` (`workspace_kind=blueprint_package`).
- Product state: `under_construction`; executable source repository is `M:\Projetos\portable-engineering-harness`.
- Current request: continue the Harness build, retain durable memory/log/status/run evidence at every checkpoint, and use `.mts` for new executable TypeScript modules.
- Build plan authority: `24_PLANO_MESTRE_DE_CONSTRUCAO_DO_HARNESS.md`.

## Initial discovery checkpoint — 2026-11-19

- Inspected the package instructions, identity contract, master build plan, execution-adapter contract, memory contract, integration plan, and authoritative product decision record.
- Confirmed `H0-01` requires a separate executable Harness source repository and link to this blueprint.
- Identified the repository destination and initial executor/runtime as decisions reserved for Lucas.

## Build bootstrap checkpoint — 2026-11-19

- Lucas approved `M:\Projetos\portable-engineering-harness` as the separate `harness_source` repository and Codex on native Windows as the initial executor candidate.
- Created and initialized the source repository with Git; the initial commit is `199526a` (`chore: bootstrap harness source and build state`).
- Completed `H0-01` and `H0-02` in the source repository: source identity, version and pre-release license state; build manifest; durable run record; and handoff.
- Completed `H0-06`: source storage and recovery policy separates canonical repository artifacts from local sessions, caches, backups, temporary data, and secrets.
- `H0-03` is now prepared: the source contains a capability matrix and controlled synthetic probe for Codex Desktop on native Windows. It remains unverified until actual delegation is observed.
- `H0-03` is complete: a synthetic run observed distinct specialists, result association, controlled failure, interruption, and reconstruction by a fresh coordinator. Codex Desktop/native Windows is `smoke_verified_with_gaps`, not unconditionally supported.
- `H0-04` and `H0-07` are complete: Windows is the primary profile; Linux and WSL2 are explicitly unsupported pending evidence; release 1 begins with one extensible Codex Desktop/native-Windows adapter.
- `H0-05` is complete: Lucas selected local/loopback OmniRoute as the required initial gateway for all model calls. `synthetic`, `public`, `internal` and `confidential` data may be sent only when necessary, context-minimized and matched to an explicitly eligible upstream route; secrets, credentials, authentication material and regulated data remain prohibited by default. No credential, endpoint, model identifier, gateway installation, or external provider call exists yet.
- No provider/model integration or target-system work has been started.

## Current limitation

- The local Python environment lacks a YAML parser module. The source state was structurally checked without installing an unplanned dependency; semantic YAML validation remains a tracked tooling gap.
- This Codex host only permits creating subagents when the user explicitly requests delegation. The required `H0-03` runtime probe is therefore waiting for that explicit authorization; the Harness will not substitute same-context personas.
- The gateway policy changes external data exposure, retention, cost, and provider ownership. Its local installation, upstream route allowlist and smoke evidence remain to be completed without committing secrets.

## Implementation checkpoint — 2026-11-21

- The executable source repository `M:\Projetos\portable-engineering-harness` now contains implemented work through `H3-02`.
- H2-09 completed Master synthesis and layered escalation.
- H3-01 completed durable project storage layout and atomic-persistence behavior.
- H3-02 completed canonical repository memory and a stable artifact index.
- The newly introduced executable modules use `.mts`, not `.mjs`.
- A persistence protocol violation occurred when an earlier session did not fully
  synchronize all memory surfaces before closing an implementation checkpoint.
- This is now an inviolable rule: build state, `MEMORY.md`, `LOG.md`, run,
  handoff, evidence, and commit must move together.
- The next required source work item is `H3-03`: run/event log, evidence index,
  checkpoint, and handoff.
- `H3-03` started on 2026-11-21. At this checkpoint, the source and blueprint
  states agree that H3-03 is active; next checkpoints must also update this
  status file, this memory, and the blueprint build log.
- Intermediate H3-03 checkpoint: event/evidence persistence and recovery
  artifacts are implemented and initial tests pass. Full regression is pending;
  H3-03 remains active in `BLUEPRINT_STATUS.yaml`.
- H3-03 completed after full regression and durable final evidence. H3-04 is now
  active to recover build state and select ready work from canonical plan
  dependencies; `BLUEPRINT_STATUS.yaml` and both repositories' memory/log
  records were advanced at this checkpoint.
- H3-04 intermediate recovery found duplicate top-level active/next keys in the
  source build state. They were normalized, and the recovery parser now rejects
  duplicate fields and mismatched active-run records. The item remains active
  pending full regression.
- H3-04 completed after full regression: build-state recovery resumes the
  matching active run or selects the first dependency-ready plan item. Duplicate
  state keys, unknown dependencies, and conflicting pointers fail explicitly.
  The next required item is H3-05, pipeline graph schema and validator.
- The source CLI verified H3-05 as the actual next item after H3-04 closed;
  H3-05 is now active. Its objective is pre-execution validation of pipeline
  graph dependencies, cycles, inputs/outputs, and roles.
- H3-05 intermediate checkpoint: schema, graph validator, and valid/invalid
  fixtures are implemented; initial tests and 24 contract fixtures passed.
  Full regressions remain pending and the item remains active.
- H3-05 completed after full regression and durable final evidence. H3-06 is
  now active for proportional pipeline composition; the blueprint status,
  memory, and log were advanced at this checkpoint.
- H3-06 completed: pipeline composition selects a deterministic smallest eligible reusable graph or creates a validated temporary graph; undeclared external inputs fail explicitly. Run snapshots are canonical JSON, content-addressed, and hash-verified on recovery. Full H1-H3/H2 regressions and 26 schema-contract fixtures passed.
- H3-07 is now active for per-stage state, gates, pause/resume, retries, and exception routing. Source `harness-build.yaml` and this blueprint's `BLUEPRINT_STATUS.yaml`, `MEMORY.md`, and `LOG.md` were advanced together.
- H3-07 completed: snapshot-linked stage/run state, durable SHA-256 integrity, append-only transitions, dependency/output/evidence gates, pause/resume, bounded retries, and explicit exhaustion routes are implemented. Durable pipeline-definition revision proposals link source and proposed snapshots and verify review evidence; they remain `review_required` and are never applied by H3-07.
- H3-07 final validation passed all H2/H3 regression scripts, 30 schema-contract fixtures, 8 transition invariants, 5 authority/event/compatibility invariants, `.mjs` syntax checks, and whitespace validation. Evidence is indexed in the source H3-07 run.
- H3-08 is complete: grouped review requests, action-specific actors from a SHA-256-pinned H1-05 policy, append-only decisions, evidence-backed escalation, idempotency, deadlines, and explicit timeout states are implemented. Missing/tampered/outside-root evidence, unauthorized actors, late replies, and production requests fail closed. All H2/H3 regressions passed with 32 contract fixtures, 8 transition invariants, and 5 authority/event/compatibility invariants. No live notification or human identity provider exists; actor IDs are policy-checked labels, not authenticated identities.
- H3-09 is complete in source run `build/runs/build-20261121-h3-install-reconciliation/`: evidence-based reconciliation creates immutable request-keyed decisions and one append-only run event, verifies hashes/timestamps and fails closed on toggle conflicts, changed evidence, invalid paths, and duplicate YAML manifest keys. It reads the canonical top-level YAML manifest subset and JSON. All 17 H2/H3 regression scripts, 34 contract fixtures, 8 transition invariants, and 5 authority/event/compatibility invariants passed. No target installation or operator authentication is performed; trusted host attestation remains required. H3-10, explicit temporary-pipeline promotion, is the next required item.
- H3-10 is complete: the source now has explicit, review-gated promotion of temporary pipelines to project/upstream targets with immutable provenance, generalization, review evidence and Lucas-only authority. Repeated and terminal-state persistence is controlled and evidence drift fails closed. The next required item is H4-01.
- H4-01 through H4-10 are complete. The source has nine specialist roles and integrated synthetic pillar artifacts for architecture, legacy analysis, engineering compatibility, backlog, TDD development, QA/review, homologation/GMUD, and Lucas-selected next scope. Production deployment remains prohibited. The next required item is H5-01.
- H5-01 through H5-18 are complete as provider manifest and integration contracts. Live external integrations remain unconfigured or contract-only; local memory handoffs, skill discovery, publication index rules, Spec Kit mapping, isolated Harness2 proposal, trace fallback, federated context assembly and per-role capability matrix are implemented and tested. The next required item is H6-01.
- H6-01 through H6-09 and H7-01 through H7-07 are complete as security and synthetic portability artifacts. Executor/model profiles, untrusted content, secrets, trust policy, provider changes, migration/rollback, retention and SBOM policy are implemented. Windows/WSL/install/recovery/graph/performance scenarios are covered; Linux and WSL2 remain unsupported until real probes.
- H8-01 through H8-11 are complete as synthetic integrated E2E validation. The source module and test cover new-system, legacy, four-pillar, reopen, GMUD, dynamic-pipeline, provider collaboration, Harness2, traceability, operations and readiness evidence.
- H8-12 was decided by Lucas as `release_candidate` version `0.2.0`; `ready_for_target_use=false`. H9 and H10 were a prompt-entry mistake and do not exist in the master plan.
- Live native-Windows evidence marked CodeGraph 1.6.4 and Graphify 0.9.115 as smoke_verified in temporary fixtures. OmniRoute health was observed, but the exact upstream model route remains unconfigured.
- Lucas authorized `ready_for_target_use=true` for version 0.2.0. Provider limitations remain explicit: OmniRoute has no model-call smoke, ai-memory is experimental on native Windows, and external benchmarks/skills must be independently provisioned.

- The blueprint includes the stack-agnostic
  `25_VALIDACAO_AUTONOMA_DE_PROJETOS.md`. It defines target-local startup,
  isolation, tool discovery, deterministic runners, persistent evidence,
  independent review and terminal states.
