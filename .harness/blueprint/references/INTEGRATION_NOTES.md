# Integration Notes — Overlap, Boundaries and Practical Choices

## 1. Memory versus harness learning

Do not collapse these concepts:

| Concern | Primary purpose | Candidate |
|---|---|---|
| Structural knowledge | What code/docs contain and how entities relate | Graphify / CodeGraph |
| Episodic/project memory | What happened, was decided, failed, or remains open | ai-memory |
| Harness learning | Which procedure/prompt/policy/model strategy performs better | Harness²-inspired loop |
| Telemetry | What happened during runtime and how long/costly/error-prone it was | OpenTelemetry / optional Phoenix |
| Task evaluation | Did this deliverable satisfy its spec? | Project tests + acceptance evaluator |

Example: an idempotency failure in a consumer is an episode in memory, an edge/relationship in structural knowledge only if represented by code structure, a test result in evaluation, and possible evidence for a Harness² improvement proposal.

## 2. ECC versus Matt Pocock

These are overlapping skill libraries, not mutually exclusive products.

Recommended approach:

- Use ECC as a broad source for adapter patterns, hooks, selected capabilities, security and cross-agent compatibility.
- Use Matt Pocock as a source for composable engineering methodology and selected skills such as grilling, TDD, debugging, domain modeling, codebase design, review and retrospective.
- Curate a single capability catalog. Each capability has one canonical selected implementation and metadata describing its origin.
- If two skills cover TDD, review, planning or debugging, compare them and choose/adapt one as the default. Do not load both by default.
- Keep attribution/license metadata for copied or adapted content. Prefer linking to upstream or writing an original equivalent when copying is unnecessary.

## 3. Graphify and CodeGraph together

Lucas expects both tools to participate because they are complementary in his current workflow. The intended design keeps two narrow adapters rather than forcing both behind one interchangeable-provider abstraction: CodeGraph emphasizes code-symbol relationships, while Graphify describes broader connections among code and artifacts such as docs, SQL, configuration, and PDFs. Both overlap on code, so complementarity is a hypothesis to measure rather than assume.

Comparison criteria:
- Is there a maintained executable/official repository, or only an article?
- Supported languages and parser quality.
- Cross-file resolution and relation types.
- Extracted versus inferred edge provenance.
- Incremental refresh and stale-graph detection.
- Query API/CLI and machine-readable output.
- Windows/Linux/WSL2 behavior and local dependency footprint.
- License, maintenance, performance on monorepos, and false-positive rate.

Route queries by task intent, preserve each provider's provenance/freshness, and use source files as final evidence. During evaluation, compare both on overlapping code questions. Do not index every generated/vendor/build directory. Keep graph build/update under AI direction, with deterministic extraction tool execution as a subordinate operation. Never assume Windows and WSL2 can share the same daemon/index without verified support.

## 4. Spec Kit versus Matt Pocock grilling

Use one workflow, not two competing workflows:

- Grilling/requirements interview helps discover ambiguity, domain vocabulary and decisions.
- Spec Kit-style artifacts formalize specification, plan, task list and acceptance criteria.
- The harness Pipeline owns stage order and status.
- Skills provide the procedure inside stages.

For small unambiguous tasks, use a short path; do not force a long interview and full spec package for every trivial edit.

## 5. Phoenix versus Harness²

There is overlap around experiments/evaluations, but different centers of gravity:

- OpenTelemetry instruments execution.
- Phoenix helps inspect traces and organize eval/dataset/experiment analysis.
- Harness² proposes/probes/composes harness changes and preserves lessons to improve future runs.

Avoid maintaining two independent candidate-promotion systems. The local agent should define one evaluation result contract and allow Phoenix/Harness² to consume or emit compatible references where feasible.

## 6. OmniRoute

Optional gateway. The core should operate with direct model provider. Define model capability profiles and provider selection policy separately from the pipeline. Route based on capability and measured successful-task cost, not merely token price.

## 7. Security complexity

Do not build a custom sandbox product in the first version. Use the selected executor's native sandbox and approval mechanisms. Add only a clear policy layer: workspace write scope, network/credential limits, sensitive action approvals, untrusted-content treatment, and Git safety. A prompt that says “be careful” is not a technical sandbox; report when a selected executor lacks enforcement.

## 8. Coding benchmarks versus Harness²

Harness² includes evaluation/benchmark machinery for its supported executor and benchmark setup. Therefore an extra standalone coding-benchmark platform is not a mandatory part of the initial harness.

Still distinguish:
- Project tests and build checks.
- User-task acceptance criteria.
- Runtime traces.
- Harness strategy comparisons.

External benchmark suites can be optional reference datasets later. A small in-house task regression set is useful for checking skills/policy changes, but do not inflate the initial implementation with a separate benchmark platform.

## 9. Context engineering

All information sources should feed a Context Assembler through bounded retrieval. The assembler should rank, filter, deduplicate, label provenance/freshness, budget context, and link source artifacts. Do not send all memory, all graph nodes, all skills, or entire docs to every model call.

## 10. AI-first onboarding versus deterministic helpers

Allowed: agent-directed listing of tracked files, reading manifests, checking Git, running tests, asking Graphify to parse code, validating schemas, generating hashes, and collecting tool versions.

Not allowed: hidden pre-AI megascript that scans the full repository, makes architectural conclusions, installs every integration, generates configuration, and only then starts an agent.

AI decides what evidence is needed and interprets the results. Helpers collect facts or perform bounded operations.
