# Example — AI-First Project Onboarding Prompt

Copy/adapt this prompt when opening an unfamiliar repository in a supported coding agent. Replace `<PATH-TO-HARNESS>` with the location of the harness repository or extracted bundle. The path may be on a USB drive. Do not persist that absolute path in project files.

---

You are operating in the root of this software project. I want you to onboard this repository using the Portable Engineering Harness located at `<PATH-TO-HARNESS>`.

**You are responsible for leading the onboarding from this point. Do not run a large preflight scanner or opaque installer before you inspect the project.** First read the harness bootstrap instructions and the relevant architecture/onboarding requirements. Then progressively investigate this repository using your native file, search, Git and terminal tools.

Your job is to:

1. Establish the repository root, Git state, branch/revision and any existing uncommitted work. Do not overwrite existing files.
2. Orient yourself using top-level files, README, existing agent instructions, build/dependency manifests and CI configuration.
3. Decide which additional files or bounded helper commands are needed. Explain material operations. Do not scan every source file without a reason.
4. Identify languages, frameworks, build/test/lint commands, entrypoints, architecture evidence, data stores, external contracts, CI/CD, conventions and known gaps.
5. Distinguish every important project claim as observed, inferred, confirmed or unknown, with evidence paths or command results.
6. Decide whether structural graph indexing is useful. If useful, use the chosen graph provider as a subordinate tool and inspect its output; do not blindly accept inferred edges.
7. Retrieve relevant project memory only if a configured memory provider is available. Do not expose secrets or ingest unrelated personal data.
8. Create a project-specific `.harness/` instance, manifest, profile, architecture summary, glossary/conventions where justified, pipeline/quality configuration, and the minimal native entrypoint for this agent.
9. Configure only integrations that are actually available and useful. Mark unavailable services as unconfigured; do not fabricate successful integrations or require credentials just to create the profile.
10. Validate generated files, schema/config consistency, ignore rules, adapter entrypoints and the discovered build/test commands where safe.
11. Present an onboarding report: what you learned, evidence, created/modified files, assumptions, unknowns, enabled/unconfigured integrations, limitations, and how we can run the E2E pipeline.

Use deterministic scripts only as small, visible, scoped helpers under your control. They may collect metadata, validate schemas, run existing tests, or invoke graph extraction. They must not make architectural decisions or silently modify the project. AI reasoning and review remain responsible for interpreting evidence and creating the project instance.

Do not impose a preferred architecture or language. Do not refactor the project during onboarding. Do not install every tool in the reference catalog. Do not begin a long clarification interview unless an ambiguity materially changes the project profile or integration. Ask me only for genuinely blocking decisions or permissions.

When onboarding is complete, we should be able to run a task through the harness's E2E pipeline, invoke individual stages when needed, resume from artifacts in a later session, and inspect a trace/report of what happened.

---

## Expected interaction pattern

The agent should not treat this prompt as permission for unlimited actions. It should begin with read-only orientation, respect the active executor's native approval/sandbox controls, and request permission before destructive operations, external side effects, credential access, broad dependency installation, or changes to existing project instructions.
