# 05 — Component and Source Integration Catalog

This document answers: what does each source contribute, how should it connect, and what should not be duplicated?

## 1. Overview matrix

| Source | Primary responsibility | Integration posture | Avoid |
|---|---|---|---|
| Anthropic agent article | Main-loop orchestration and selective delegation | Adopt coordination/economy principles while preserving Lucas's required specialist organization | Let an external article override product decisions |
| CodeGraph — Colby McHenry | Local code-structure knowledge graph | Use with Graphify as a complementary provider after overlap/freshness validation | Treat its graph as source of truth or duplicate every query without evaluation |
| Google Harness² | Recursive harness refinement and experiments | Adopt learning method; inspect implementation before embedding | Treating it as generic project memory |
| ECC | Broad engineering skills, agents, rules, hooks, security capabilities | Curate selected skills and patterns | Installing entire catalog by default |
| OmniRoute | Multi-provider model gateway and routing | Optional model gateway adapter | Hard dependency or cost-only model selection |
| Graphify | Structural graph connecting code with docs, SQL, configuration, and supported artifacts | Complement CodeGraph with cross-artifact queries; validate overlap, support, and provenance | Assuming inferred relations are facts; full graph in every prompt |
| ai-memory | Persistent cross-session/agent memory and handoffs | Candidate memory provider | Treating every event/transcript as durable project knowledge |
| Matt Pocock skills | Composable engineering disciplines | Curate methodology skills | Duplicating ECC's overlapping TDD/review/plan skills |
| GitHub Spec Kit | Specification/planning/task artifact method | Adopt selected SDD workflow and artifact ideas | Running a second competing pipeline in parallel |
| OpenTelemetry GenAI conventions | Instrumentation vocabulary and trace interoperability | Instrument pipeline/tool/model operations | Treating telemetry as memory or evaluation by itself |
| Arize Phoenix | Trace inspection, datasets, evaluations and experiments UI | Optional observability/analysis backend | Making it a mandatory runtime dependency |
| OWASP MCP security | Practical security guidance | Derive minimum tool/MCP policies | Building a custom sandbox platform in v1 |
| Agent Skills specification | Portable skill packaging/discovery format | Use as format baseline where supported | Assuming every agent has identical discovery behavior |
| SWE-bench / Terminal-Bench | External coding/terminal evaluation references | Optional future benchmark data | Mandatory benchmark platform in initial build |

## 2. Anthropic — when to use subagents

Source: https://claude.com/resources/articles/building-agents-when-to-fold-subagents-into-your-main-agent-loop

**Contribution:** encourages a capable main loop and warns that delegation should have a purpose. Lucas separately requires a Master coordinator and genuine specialist agents across four pillars; the article informs efficiency decisions within that organization, not whether specialists exist.

**Apply:**
- Master Agent owns coordination, state, and convergence; specialist agents retain their domain judgments and authority boundaries.
- Skills/tools are loaded progressively.
- Delegate where tasks are independent, parallelizable, context-heavy, or require isolation.
- Use fresh context for independent code review where beneficial.
- Compare end-to-end outcomes, cost, latency, and quality before adding more agents.

**Boundary:** stages are workflow units and do not map one-to-one to specialist agents. Specialists are real agents; which agents are invoked for a given stage remains proportional to its needs.

## 3. CodeGraph — Colby McHenry

Primary source: https://github.com/colbymchenry/codegraph  
Supplementary article: https://dev.to/arshtechpro/codegraph-stop-your-ai-agent-from-grepping-the-same-files-50-times-3bgm

**Contribution:** current repository describes a local-first graph of code symbols and relationships, exposed through MCP/CLI/library for supported agents. It is paired with Graphify in Lucas's target design.

**Apply:** use the article as a reference when evaluating structural knowledge solutions and graph-query UX.

**Caveat:** repository claims and release support change. Recheck version, license, install chain, telemetry, platform support, index freshness, Windows/WSL2 limitations, and adapter behavior before adoption. Do not rely on secondary articles for compatibility.

## 4. Google Research Harness²

Source: https://github.com/google-research/harness2

**Contribution:** task-level Propose–Probe–Compose and persistent harnessing-level improvement. Its repository describes editable harness components, parallel/sequential recursion, executor integrations, and benchmark runners/judging/reporting.

**Apply:**
- Keep a baseline version of a skill/workflow/policy/model route.
- Generate a candidate change from evidence.
- Run the same task/evaluation cases against baseline and candidate.
- Compose/adopt an improvement only with evidence.
- Preserve lessons and experiment lineage.

**Boundary:** Harness²'s improvement harness is not a general cross-session project memory service. It also does not replace a project's build/test/acceptance checks or runtime tracing. Treat the repository as an experimental reference/integration candidate, not a mandatory dependency until its current interfaces and executor/platform support are verified.

## 5. ECC — Everything Claude Code

Source: https://github.com/affaan-m/ECC

**Contribution:** broad collection of agents, skills, hooks, rules, commands, memory/learning patterns, security checks, and multi-harness adapters.

**Apply:**
- Curate only skills/procedures that fill a real need.
- Study hooks and lifecycle automation where the target executor supports them.
- Study security scanning and adapter patterns.
- Use selected engineering rules per project and language.
- Record source/version and license for adopted content.

**Do not:** install the full catalog or every agent automatically. Do not stack multiple install paths for the same executor. ECC itself warns against duplicate installation layers. Verify current platform feature parity; adapters may have differing capabilities.

## 6. OmniRoute

Source: https://github.com/cheaper-inference/OmniRoute

**Contribution:** a unified gateway for multiple model providers, with routing/fallback/quota and model availability features described by the project.

**Apply:** provide a `ModelProvider` abstraction with an optional OmniRoute adapter. Stages request capabilities/model profiles (e.g. fast classification, code implementation, independent review), while routing policy maps them to providers.

**Do not:** require the gateway to start. A direct model provider must work. Do not optimize on price alone; compare task success, latency, and total cost per successful task. Verify provider terms, privacy, data handling, quota semantics, and current model names.

## 7. Graphify

Source: https://github.com/Graphify-Labs/graphify

**Contribution:** local deterministic AST extraction for code, graph relations and queries, communities, graph reports, and optional semantic processing for documents/media. Its README describes edges distinguished as extracted vs inferred and a graph JSON/report/HTML output.

**Apply:** expose a provider-neutral structural knowledge interface: build/status/query/explain/path/impact. Query scoped subgraphs rather than injecting the entire graph. Store source revision, schema version, extractor version, and freshness status.

**Do not:** treat inferred edges as equivalent to directly extracted facts. Do not commit large generated graph artifacts by default. Do not install it blindly: confirm language coverage, native Windows behavior, update costs, supported integrations, package provenance and license.

## 8. ai-memory

Source: https://github.com/akitaonrails/ai-memory

**Contribution:** long-term, cross-agent and cross-session memory; automatic lifecycle capture via hooks; consolidation to human-readable Git-backed Markdown pages; search/retrieval; handoffs; project and user profile; optional LLM/vector/dream functions; zero-LLM base path described by its README.

**Apply:** use as a `MemoryProvider` for relevant retrieval, experience recording and handoff. Store a separate project memory namespace and personal profile scope. Promote only confirmed durable decisions to project ADR/docs. Keep raw session capture, memory data and project source-control artifacts distinct.

**Platform caveat:** the repository README describes Windows via WSL2 as supported and native Windows as experimental. Verify this against the exact version and target environment. Do not promise native Windows parity without testing. Its persistent memory server/data have a lifecycle separate from a project copy of the harness; do not blindly copy server databases into every repository.

**Do not:** equate memory with Harness² learning. ai-memory remembers project/session experience; Harness² experiments on how the harness should improve.

## 9. Matt Pocock — Skills for Real Engineers

Source: https://github.com/mattpocock/skills

**Contribution:** small, composable engineering skills. The repository emphasizes alignment/grilling, shared language (GLOSSARY), ADRs, TDD, debugging, domain modeling, architecture design, code review, and retrospectives.

**Apply:** curate particularly relevant practices such as `grill-with-docs`, `tdd`, `diagnosing-bugs`, `domain-modeling`, `codebase-design`, `code-review`, `improve-codebase-architecture`, and `retro`, subject to current upstream contents and license.

**Do not:** install a full duplicate skill set alongside ECC. Create a skill registry with one canonical selected procedure for each capability. Where a Matt Pocock skill provides a stronger methodology, use/adapt it as the canonical version; use ECC for complementary hooks, adapters, rules, or utilities.

## 10. GitHub Spec Kit

Source: https://github.com/github/spec-kit

**Contribution:** structured Spec-Driven Development artifacts and workflow concepts (specification, plan, tasks, clarify/checklist/analyze/converge concepts described by the project).

**Apply:** make requirements and acceptance criteria durable before implementation. Link specs to tasks, changes, tests, and convergence results. Allow a short path for small tasks and deeper clarification for ambiguous/high-risk work.

**Do not:** run Spec Kit as a second independent workflow that conflicts with the harness pipeline. Adopt its useful artifacts/concepts behind one canonical Spec interface. Inspect current commands, templates, supported agents, and license before vendoring anything.

## 11. OpenTelemetry

Source: https://opentelemetry.io/docs/specs/semconv/gen-ai/

**Contribution:** common instrumentation concepts for generative AI operations, model calls, tools and agents. Agent-specific conventions may evolve; verify the current stability/development status.

**Apply:** trace a pipeline run as a root span; stages as child spans; model/tool calls as nested spans. Record IDs, provider/model, duration, status, usage metrics where available, retry count and links to artifacts. Redact secrets and sensitive content.

**Do not:** store entire prompts/responses by default. Do not use traces as the canonical memory store.

## 12. Arize Phoenix

Source: https://github.com/Arize-ai/phoenix

**Contribution:** an optional interface and platform for traces, evaluation, datasets, experiments and analysis of LLM/agent applications.

**Apply:** provide an OpenTelemetry-compatible export path and optional Phoenix configuration for local inspection. Use it to diagnose slow, failing, repetitive, or low-quality runs and to compare experiments if useful.

**Do not:** require Phoenix just to execute the pipeline. Keep trace export/provider replaceable and retain a local/no-backend fallback.

## 13. Agent Skills specification

Source: https://agentskills.io/specification

**Contribution:** portable packaging convention for skills, typically a directory with `SKILL.md` and optional scripts/references/assets.

**Apply:** follow the format where possible. Keep skill instructions self-contained and progressive: a concise description/trigger, workflow, relevant references, and optional scripts. Native adapters map skills into each executor's expected discovery location.

**Do not:** assume universal skill discovery or equal tool/hook support across agents. Verify each adapter.

## 14. Security guidance

Source: https://cheatsheetseries.owasp.org/cheatsheets/MCP_Security_Cheat_Sheet.html

**Contribution:** security practices around MCP/tools, untrusted inputs, tool authorization, least privilege, auditability, and prompt-injection boundaries.

**Apply:** minimal tool permissions, explicit network/credential boundaries, approval for sensitive operations, and clear distinction between untrusted repository content and trusted harness policy.

**Do not:** build an independent sandbox product as an initial requirement. Reuse native executor controls and state their limitations.

## 15. SWE-bench Verified and Terminal-Bench

Sources: https://www.swebench.com/verified and https://www.tbench.ai/

**Contribution:** external reference tasks for evaluating real repository issue resolution and terminal-based engineering tasks.

**Apply:** keep them as optional future evaluation references. Harness² already provides benchmark/evaluation mechanisms in its supported framework; a separate benchmark framework is not required for initial delivery. Maintain a small project-specific regression suite and evaluate complete task outcomes.

**Do not:** make these benchmark suites a mandatory dependency or confuse benchmark scores with proof that a particular project task satisfies its acceptance criteria.
