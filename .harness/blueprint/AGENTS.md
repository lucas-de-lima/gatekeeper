# Gatekeeper Target-local Blueprint Instance

This directory is a target-local copy of the official Harness blueprint.

- Workspace: `managed_target`
- Project: `gatekeeper`
- Work mode: `target_discovery`
- Canonical upstream: `M:\Projetos\portable-engineering-harness-blueprint-v1.0.0\engineering-harness-blueprint`
- No project discovery result is authoritative upstream.
- Preserve local baseline/evidence and wait for an explicit implementation request.

This directory is the specification package for the Portable Engineering Harness. It does not yet contain the Harness implementation.

`BLUEPRINT_STATUS.yaml` is the current package state. It explicitly says the Harness is under construction, implementation has not started, and no target system is active. The first build work item is `H0-01` in `24_PLANO_MESTRE_DE_CONSTRUCAO_DO_HARNESS.md`.

## Start here

- For an implementation task, read `AGENT_START_HERE.md` first.
- Read `23_IDENTIDADE_DO_WORKSPACE_E_MODOS_DE_TRABALHO.md` to classify the workspace and invocation before choosing a task.
- Use `24_PLANO_MESTRE_DE_CONSTRUCAO_DO_HARNESS.md` as the ordered source of truth for the build backlog, phase gates, readiness definition, and next work item.
- Treat `13_LUCAS_PRODUCT_VISION_AND_DECISIONS.md` as the source of confirmed product decisions. When an earlier requirement conflicts with a later clarification, the clarified decision governs.
- Read `21_ARQUITETURA_INTEGRADA_E_PLANO_DE_IMPLEMENTACAO.md` for the overall design and acceptance criteria.
- Read `22_PLANO_DE_AMARRACAO_E_INTEGRACAO_FUNCIONAL.md` before implementing or installing providers, skills, memory, observability, or evaluation tools.
- Consult `15`–`22` for the detailed contracts relevant to the task. Use `01`–`12` as the original requirements and design baseline, while checking for superseding decisions in `13`–`24`.

## Preserve product invariants

- The Master Agent is a management coordinator. Real specialist agents form the architecture, engineering, backlog-management, and development pillars; persona switching is not a substitute for delegated agents.
- Until a release-ready decision is recorded, work in a `harness_source` repository is `harness_build`; do not treat the Harness source as a managed target system or invent target-system feature work.
- A fixture used to validate the Harness is `harness_fixture`/`harness_verification`, not the user's undisclosed product.
- Pipelines may be composed dynamically, but every run must preserve authority, evidence, safety, and production-boundary rules.
- An approved architecture baseline cannot be reopened without the architect's evidence-backed proposal and Lucas's explicit authorization.
- The project artifacts are the canonical memory. External memory and knowledge graphs provide retrieval, not authority over project decisions.
- CodeGraph and Graphify are intended to work together; validate their roles, provenance, freshness, and access from specialist agents.
- Development includes appropriate test-first work, verification, independent review, and QA, and stops before production deployment. GMUD is a post-homologation change-authorization process when required, not a branch strategy.

## Work and reporting

- Keep confirmed decisions, design proposals, and unresolved choices clearly distinguished.
- Do not claim an integration is functional until an agent has invoked it successfully and its output and provenance have been checked.
- Before installing an upstream component, verify its official source, pinned version, license, platform support, data handling, installation side effects, and removal path.
- When changing this specification package, update `PACKAGE_MANIFEST.json` and verify the recorded file sizes and SHA-256 hashes.
- Do not infer that production deployment or an irreversible external action is in scope.

## Inviolable memory persistence rule

- A build work item is not complete unless `harness-build.yaml`, `MEMORY.md`,
  `LOG.md`, and the matching run/handoff are synchronized and committed in the
  executable Harness source repository.
- Update this blueprint's `BLUEPRINT_STATUS.yaml`, `MEMORY.md`, and `LOG.md` at
  every checkpoint, including the start, interruption, or phase transition of
  source work. Keep the active item and next item aligned with the source
  `harness-build.yaml`.
- Before closing or reporting work, the agent must verify that canonical
  memory, the build log, build state, and evidence are mutually consistent.
  A conversation summary or transcript is not durable evidence.
- Changes to this blueprint package must update `PACKAGE_MANIFEST.json` and
  verify its recorded byte sizes and SHA-256 hashes.

This file guides work on the blueprint package. The executable Harness repository must receive its own root `AGENTS.md` and executor-specific entrypoints based on these specifications.
