# 07 — Portability, Project Instances and Versioning

## 1. Two lifecycles, two repositories

### Upstream harness repository

The user's toolbox: generic pipelines, stages, skills, policies, schemas, adapters, scripts, documentation and optional integration definitions. It is version-controlled independently and can be cloned or stored on removable media.

### Project instance

A managed copy/configuration inside each target repository. It contains the harness manifest, project profile, project-specific knowledge, selected skills/overrides, pipeline settings, native agent entrypoints and durable project artifacts.

These lifecycles must not be confused. Project-specific facts must not leak into the generic upstream core.

## 2. Installation model

Preferred model: **managed physical copy (vendor-style installation)** of the required harness files into the project. Avoid requiring Git submodules for ordinary use. Git subtree can be evaluated as an optional advanced workflow, but not required.

The install action itself is AI-led; this document does not authorize a blind CLI scan or an opaque installer. The agent decides what is relevant, creates the copy and integration files, and records what it did.

Potential project layout:

```text
project-root/
├── AGENTS.md                     # small native pointer, where applicable
├── .agents/skills/               # compatible portable skill files
├── .harness/
│   ├── manifest.yaml             # upstream identity/version/schema
│   ├── config/
│   ├── pipelines/
│   ├── project/
│   │   ├── PROFILE.md
│   │   ├── ARCHITECTURE.md
│   │   ├── GLOSSARY.md
│   │   ├── CONVENTIONS.md
│   │   └── decisions/
│   ├── overrides/
│   ├── schemas/
│   └── adapters/
├── .harness-local/               # ignored local state, if this location is used
│   ├── runs/
│   ├── graph-cache/
│   ├── traces/
│   ├── temp/
│   └── indexes/
└── ... project files ...
```

The exact location of executor-discovered skill directories and instruction files is adapter-specific. Generate only those required by the active executor. Do not create redundant instructions for every agent automatically.

## 3. Manifest

Example only:

```yaml
schema_version: 1
harness:
  name: portable-engineering-harness
  version: 1.0.0
  source: "<upstream repository URL>"
  revision: "<source commit>"
  install_mode: vendor
project:
  id: "<stable project id>"
  profile_version: 1
  initialized_at: "<ISO-8601>"
execution:
  default_pipeline: feature-development
integrations:
  structural_knowledge:
    provider: graphify
    enabled: true
  memory:
    provider: ai-memory
    enabled: false
  observability:
    provider: opentelemetry
    backend: local-or-phoenix
  inference:
    provider: native
adapters:
  active: [codex]
  capability_profile: "<adapter capability version>"
```

A manifest must distinguish configured, available, connected, and healthy. A provider set to `enabled: true` must not imply the integration actually works; health/capability checks should report reality.

## 4. Versioning rules

Track independently:

- Harness semantic version (major/minor/patch where suitable).
- Source commit/revision.
- Manifest schema version.
- Pipeline schema/version.
- Skill versions/source revisions.
- Adapter version and capabilities.
- Project profile version / last verified revision.
- Graph schema and extractor version.
- Memory server/protocol version where available.

Do not assume a new harness version is automatically compatible with every old project instance. Provide compatibility checks and migrations.

## 5. Upgrade workflow

An AI-led upgrade should:

1. Read installed manifest and inspect current Git status.
2. Identify upstream target version/revision.
3. Compare managed files with installed versions and local modifications.
4. Read release/migration notes and license changes.
5. Compute a proposed update diff.
6. Preserve project-owned profile, memory references, ADRs, overrides and customizations.
7. Detect conflicts and explain options.
8. Apply only after user authorization where changes are material.
9. Validate generated files, schemas and adapter compatibility.
10. Update manifest and generate an upgrade report.

Never silently overwrite modified files. Never automatically upgrade a project just because the upstream core changed.

## 6. Ownership and override model

Every file or config item should have an ownership classification:

- `upstream_managed`: copied from the harness; updateable through controlled upgrade.
- `project_managed`: created for this project and retained by this repository.
- `user_override`: user-owned customization; never overwrite automatically.
- `generated_cache`: reproducible local data; not committed by default.
- `runtime_secret`: external/local secret; never committed.

Keep project-specific overrides outside upstream-managed directories.

## 7. Git policy

Recommended committed content:

- `.harness/manifest.yaml`.
- Project profile and confirmed architecture docs.
- Glossary and ADRs.
- Selected project-specific pipeline config and quality gates.
- Project-specific skill overrides/custom skills.
- Specs and durable handoffs where the team wants them in project history.
- Minimal native agent pointers.

Recommended ignored content:

- `.harness-local/` run state, traces, caches, indexes, temporary files.
- Credentials and tokens.
- Native provider caches and downloaded models.
- Local memory server DB if it is user/machine scoped.
- Worktrees and scratch directories.

The agent must inspect existing `.gitignore` before modifying it and add only the necessary targeted entries.

## 8. USB/removable media

The upstream harness should be relocatable. Do not persist absolute paths to the USB drive in the project instance. The AI bootstrap should locate the source directory at runtime and copy the managed files into the target project.

Once installed, the project instance must continue to work if the USB drive is disconnected, subject only to explicit external dependencies (agent executable, model access, external memory server, graph tool, runtime/build tools). Document offline-capable and network-dependent features.

The harness itself can be copied to removable media; large LLM weights, provider secrets, memory data servers, and graph caches should not be copied implicitly.

## 9. Cross-platform expectations

Use relative paths, platform-aware process invocation, safe path normalization, and no assumptions about Bash being available on Windows. Verify PowerShell and Windows path behavior. Third-party projects have different support levels; for example, ai-memory's README has described WSL2 support and native Windows as experimental—verify the exact current release before relying on native Windows behavior.

## 10. Portability does not mean identical capabilities

Each executor adapter must expose a capability matrix. Some agents may not support the same hooks, subagents, skill loading, MCP, approval or sandbox controls. For optional capabilities, provide a fallback or mark the limitation. For genuine subagent execution, which is a confirmed product requirement, an executor without support is incompatible; do not claim identical behavior or downgrade to personas.
