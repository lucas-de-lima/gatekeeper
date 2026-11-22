# Build Log — Portable Engineering Harness

## 2026-11-21 — Stack-agnostic validation procedure

- Added `25_VALIDACAO_AUTONOMA_DE_PROJETOS.md`.
- Documented target identity, isolation, tool discovery, deterministic runners,
  persistent evidence, independent review and terminal states.
- The procedure remains stack-agnostic and does not authorize code changes or
  provider calls by itself.

## 2026-11-21 — Ready-for-target-use 0.2.0

- Lucas approved `ready_for_target_use=true`.
- The executable source records version 0.2.0 and the explicit provider
  limitations instead of converting unconfigured integrations to smoke success.

## 2026-11-21 — H8-12 release-candidate decision

- Lucas selected `release_candidate` version `0.2.0`.
- The result remains below `ready_for_target_use`; live provider evidence is
  required before provider states can be changed from contract-only/unconfigured.
- H9 and H10 do not exist in the master plan and were removed as an entry error.

## 2026-11-21 — H8-01 through H8-11 completed in executable source

- Added synthetic H8 integrated validation for H8-01 through H8-11.
- All H1-H8 tests and contract validators pass; H8-12 remains the next item
  pending explicit Lucas decision on `ready_for_target_use` and release version.
- Live provider integrations and Linux/WSL2 probes remain unclaimed.

## 2026-11-19 — Discovery checkpoint

- **Scope:** assessed the blueprint package before beginning implementation.
- **Evidence reviewed:** `AGENTS.md`, `AGENT_START_HERE.md`, `BLUEPRINT_STATUS.yaml`, documents `13`–`15`, `18`, `21`–`24`, and the integration notes.
- **Outcome:** `H0-01` is the first work item and requires a separate executable repository. The product package deliberately leaves the repository destination and initial executor/runtime as material choices for Lucas.
- **Status:** blocked pending those two decisions; no unsupported assumptions or implementation artifacts were introduced.
- **Next:** create the executable repository and record its identity once the destination and initial executor are confirmed.

## 2026-11-19 — Harness source bootstrap completed

- **Decision recorded:** Lucas approved `M:\Projetos\portable-engineering-harness` and Codex on native Windows as the initial executor candidate.
- **H0-01 completed:** separate executable repository created, Git initialized, source identity/version/license state created, and blueprint link recorded.
- **H0-02 completed:** durable `harness-build.yaml`, run record, and handoff created in the source repository.
- **Evidence:** source commit `199526a` (`chore: bootstrap harness source and build state`); working tree clean after commit.
- **Validation:** required state fields and Git whitespace check passed. Python has no YAML module, so no unplanned parser dependency was added.
- **Next:** `H0-03`, the observed capability probe for real Codex specialist delegation on native Windows.

## 2026-11-19 — Storage policy and executor probe preparation

- **H0-06 completed:** source storage policy, local-only recovery/cache paths, Git exclusions, retention baseline, recovery order, and secret-exclusion rules implemented and verified (`52dfb7a`).
- **H0-03 prepared:** current official OpenAI multi-agent documentation recorded; Codex Desktop/Windows capability matrix and synthetic probe contract added (`7b514a6`).
- **Truthful integration state:** Codex remains `selected_for_validation`; documentation and configuration are not treated as a successful adapter.
- **Blocker:** the host requires an explicit user request before creating subagents. The actual delegation/failure/resume probe therefore cannot run until Lucas authorizes delegated agents.

## 2026-11-19 — Executor validation, platform baseline, and provider-decision preparation

- **H0-03 completed:** a real synthetic Codex run exercised two distinct specialists, preserved a controlled command failure, interrupted a running specialist, and proved artifact-only resumption through a fresh coordinator (`f4abaaf`).
- **H0-04 and H0-07 completed:** Windows is primary with declared gaps; Linux/WSL2 are not claimed supported; release 1 begins with one extensible Codex Desktop/native-Windows adapter (`b826356`).
- **H0-05 prepared:** no direct provider or credential is configured. The source now has a decision record and safe no-call policy (`10a8cd5`).
- **Current blocker:** Lucas must select the initial direct provider/model policy and the maximum project-data classification it may receive; this decision controls external data exposure and cannot be guessed.

## 2026-11-20 — Gateway and data policy decided

- Lucas selected OmniRoute as the mandatory local/loopback gateway for all initial Harness model calls; no direct-provider path is part of the initial release path.
- `synthetic`, `public`, `internal` and `confidential` data may be sent only when necessary and context-minimized, through an exact upstream provider/model route locally allowlisted for the request's classification.
- Secrets, credentials, authentication material and regulated data remain prohibited by default. No eligible route means fail closed, with no silent fallback.
- The source-repository ADR and H0-05 state were updated. Gateway installation and a synthetic/public smoke call still need evidence before an integration is declared functional.

## 2026-11-21 — H2-09, H3-01, and H3-02 progress

- `H2-09` completed Master synthesis and layered escalation in the executable source repository.
- `H3-01` completed the canonical/local/external/regenerable storage layout and validated atomic-persistence behavior.
- `H3-02` completed canonical memory domains and a stable artifact index with ownership, provenance, containment, and SHA-256 evidence.
- Newly created executable modules use `.mts`; `.mjs` was not to be introduced for new work.

## 2026-11-21 — Memory-persistence protocol correction

- An earlier agent left part of the canonical memory synchronized incorrectly at a checkpoint.
- `AGENTS.md` in both the blueprint and executable source repositories now declares
  the rule inviolable: build state, `MEMORY.md`, `LOG.md`, run/handoff, evidence,
  and commit must be updated and verified together.
- The blueprint `MEMORY.md` and `LOG.md` are now updated to match the source's
  actual current state. The next source work item is `H3-03`.

## 2026-11-21 — H3-03 checkpoint started

- The source repository began implementing durable run events, evidence index,
  checkpoint and handoff artifacts.
- `BLUEPRINT_STATUS.yaml` now records H3-03 as active; both repositories'
  memory/log files were updated at the checkpoint boundary.

## 2026-11-21 — H3-03 intermediate checkpoint

- Implemented append-only event persistence, a SHA-256 evidence index, and
  machine-readable checkpoint/handoff storage and recovery.
- Initial unit tests and 18 contract fixtures passed. Full regression remains
  pending, and the work item remains active.

## 2026-11-21 — H3-03 completed; H3-04 started

- H3-03 passed H3/H2 regression suites and 22 contract fixtures, with final
  event, evidence-index, checkpoint, handoff, and validation artifacts stored.
- H3-04 is now the active source work item. `BLUEPRINT_STATUS.yaml`,
  `MEMORY.md`, and this log were updated at the checkpoint boundary.

## 2026-11-21 — H3-04 intermediate recovery checkpoint

- The recovery implementation detected duplicate active/next-state fields in
  the source build manifest. The source state was normalized and duplicate
  detection added to recovery.
- The live CLI recovered H3-04 from durable state and its matching run record;
  final regressions remain pending.

## 2026-11-21 — H3-04 completed; H3-05 ready

- H3-04 passed all relevant H2/H3 regressions and H1 contract/state/policy checks.
- Build-state recovery and dependency-driven work selection are implemented;
  final event/evidence/checkpoint/handoff artifacts are persisted.
- `BLUEPRINT_STATUS.yaml` and both repositories' memory/log files now identify
  H3-05 as the next required source work item.

## 2026-11-21 — H3-05 checkpoint started

- The source build-recovery CLI selected H3-05 after H3-04 completion, and the
  result is durably indexed as evidence.
- H3-05 is active for pipeline graph schema and validation.

## 2026-11-21 — H3-05 intermediate checkpoint

- Implemented the pipeline graph schema and validator for roles, dependencies,
  cycles, inputs/outputs and stable execution order.
- Initial H3-05 tests and 24 contract fixtures passed; final integrated
  regression remains pending.

## 2026-11-21 — H3-05 completed; H3-06 started

- H3-05 passed full H2/H3 regression and H1 schema/state/policy checks. Final
  evidence, checkpoint and handoff were persisted.
- H3-06 is active for proportional pipeline composition.

## 2026-11-21 — H3-06 completed; H3-07 started

- H3-06 delivers deterministic proportional pipeline composition, H3-05 graph
  validation, explicit external-input contracts, provenance, and per-run
  canonical snapshots verified by content hash.
- Full H1-H3/H2 regressions, 26 schema-contract fixtures, 8 transition
  invariants, 5 authority/event/compatibility invariants, syntax checks, and
  whitespace validation passed. Durable evidence is indexed in the source run.
- H3-07 is active for per-stage lifecycle state, gates, pause/resume, retries,
  and exception routing. Both repositories' status and memory records agree.

## 2026-11-21 — H3-07 intermediate lifecycle checkpoint

- Implemented per-run pipeline instances linked to immutable H3-06 snapshots,
  with durable hash verification and append-only transition history.
- Dependency, output, evidence, pause/resume, bounded retry, and exhaustion
  exception-routing behavior passed focused tests and all H1-H3/H2 regressions.
- 28 schema-contract fixtures passed. Review-linked pipeline-definition
  proposals remain open; approval authority validation is deferred to H3-08.

## 2026-11-21 — H3-08 completed; H3-09 started

- H3-08 added policy-pinned grouped human reviews, actor/action checks,
  append-only decisions, evidence-backed escalation, bounded deadlines, and
  explicit timeouts. Silence and late responses cannot approve; production
  deployment is denied.
- Full H2/H3 regressions passed with 32 schema-contract fixtures, 8
  transition invariants, 5 authority/event/compatibility invariants, JavaScript
  syntax, and whitespace checks. Final evidence is in the source H3-08 run.
- No live notification or human identity provider is configured. Actor IDs are
  checked against policy but are not authenticated identities; trusted host
  attestation remains required for real approvals.
- H3-09 is active for idempotent installation/reconciliation through the
  AI-first workflow. Source state and this blueprint's status, memory, and log
  now agree on H3-09.

## 2026-11-21 — H3-07 completed; H3-08 started

- H3-07 added immutable pipeline-definition revision proposals linked to both
  source and proposed snapshots. Persistence verifies proposal, snapshot, and
  review-evidence hashes; proposals remain `review_required`.
- Full H2/H3 regressions passed, with 30 schema-contract fixtures, 8
  transition invariants, 5 authority/event/compatibility invariants, JavaScript
  syntax checks, and whitespace validation. Final evidence is indexed in the
  source H3-07 run.
- H3-08 is active for grouped human-review questions, evidence-backed
  escalation, and H1-05 authority validation. Silence, timeout, or missing
  evidence must never become approval. The source run is
  `build/runs/build-20261121-h3-human-review/`.

## 2026-11-21 — H6-01 through H6-09 and H7-01 through H7-07 completed

- Implemented H6 security policies and H7 portability contracts as synthetic,
  evidence-backed artifacts.
- Native Windows remains the observed primary profile; Linux and WSL2 remain
  unsupported pending real host probes and separate caches.
- H6-H7 focused tests and build recovery passed. H8-01 is next.

## 2026-11-21 — H5-01 through H5-18 completed

- Added provider manifest and integration contracts for H5.
- External providers remain opt-in and honest about unavailable/contract-only state.
- Added local trace, canonical handoff, skill discovery, publication mapping, Spec Kit mapping, Harness2 evaluation isolation, context assembly and capability matrix.
- H5 focused tests passed. H6-01 is next.

## 2026-11-21 — H4-01 through H4-10 completed

- Expanded the source specialist registry and implemented the complete H4 four-pillar workflow.
- The synthetic H4 test covers architecture pipeline and approval, legacy analysis-only mode, compatibility escalation, architecture reopening, backlog, TDD, independent QA/review, homologation, GMUD, and next-scope control.
- Reserved authority and the pre-production boundary are preserved. H5-01 is next.

## 2026-11-21 — H3-09 completed; H3-10 next

- Implemented pure, evidence-based installation reconciliation and a local
  adapter that hashes explicitly named project files and re-verifies them on
  read. Decisions are immutable and keyed by run/request ID; retries return
  the same result, while changed input under a reused request ID is rejected.
- Appended one `installation.reconciled` event per decision. Bad toggles,
  timestamp conflicts, changed evidence, and invalid paths fail closed.
- The adapter reads the canonical top-level YAML manifest subset using the
  existing parser, rejects duplicate keys, and also accepts JSON manifests.
- No installer, managed-target write, external service, or human identity
  attestation is performed by this adapter.
- All 17 H2/H3 regression scripts passed with 34 schema-contract fixtures,
  8 transition invariants, and 5 authority/event/compatibility invariants.
  Final evidence is indexed in the source H3-09 run.
- H3-10, explicit promotion of a temporary pipeline to a project/upstream, is
  the next required item; source and blueprint state now agree.

## 2026-11-21 — H3-10 completed

- Added explicit promotion of temporary pipelines to project or upstream targets.
- The promotion record preserves source snapshot provenance and generalization,
  requires project-contained review evidence, and allows authority only for
  Lucas under the `promote_project_pipeline` action.
- Repeated promotion persistence is idempotent; only review-required can
  become promoted or rejected, and terminal states cannot regress.
- The focused H3-10 test passed. H4-01 is the next required item.
