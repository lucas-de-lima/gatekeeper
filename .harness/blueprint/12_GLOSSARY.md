# 12 — Glossary

| Term | Meaning in this product |
|---|---|
| Harness | The system around a model/agent: orchestration, tools, prompts, skills, memory, context, policies, evaluation and workflows. |
| Main Agent / Orchestrator | The primary agent responsible for user task, pipeline and convergence. |
| Pipeline | E2E task execution graph of stages, gates, transitions and artifacts. |
| Stage | Independently invocable unit of pipeline work; not necessarily an agent. |
| Agent | LLM-driven executor/persona used where reasoning or delegated context is useful. |
| Subagent / specialist agent | A real delegated agent execution with a scoped role/context/tools. The Harness requires specialist agents; which ones are invoked depends on the task. |
| Master Agent | Delivery manager and coordinator that composes pipelines, delegates, tracks state, and escalates; not a substitute persona for specialist agents. |
| Pillar | One of the four responsibility areas: architecture, engineering, backlog management, development. Each can contain multiple agent specialties. |
| Pipeline instance | A frozen, traceable execution of a selected pipeline definition for a specific request/run. |
| GMUD | Post-homologation audit/change authorization gate when a project's/client's process requires it. It is not a branch strategy. |
| Skill | Reusable procedure/instruction package for a recurring activity. |
| Tool | Concrete operation such as reading a file, querying graph, running tests, or calling an API. |
| Artifact | Durable output from a pipeline stage, such as spec, plan, report, handoff or trace reference. |
| Context Engineering | Selecting, assembling, ranking, budgeting and validating relevant information for a model. |
| Project Profile | Evidence-backed description of a particular repository's stack, architecture, tools and constraints. |
| Project Instance | The harness copy/configuration installed into one target project. |
| Structural Knowledge | Code/document relationship model such as AST graph. |
| Episodic Memory | Record of what happened in a task/session, including failures and progress. |
| Project Memory | Durable decisions, facts and operational knowledge relevant to a repository. |
| Harness Learning | Evidence-based improvement of skills, prompts, policies, routing or pipeline design. |
| Spec-Driven Development | Workflow where intent is formalized into a specification and acceptance criteria before/alongside implementation. |
| Convergence | Final check that implementation and evidence satisfy the original specification. |
| Provenance | Origin of a claim or artifact: file, command, tool, agent, time, revision or confirmation. |
| Observed | Directly supported by source or command evidence. |
| Inferred | Hypothesis derived from evidence, not confirmed. |
| Confirmed | Explicitly validated by a human/project owner. |
| Backpressure | Mechanisms that stop/slow/replan work when budgets, retries, context, concurrency or progress limits are exceeded. |
| Evaluation | Scoring/checking task quality, acceptance, regression, or harness strategy. |
| Observability | Instrumentation that describes execution behavior through traces, events and metrics. |
| Adapter | Mapping between harness-neutral capability contracts and one native executor/provider. |
| Managed copy/vendor model | Physical copy of upstream-managed harness files into the target project, tracked by version/manifest. |
| Upstream | Central generic harness repository maintained by the owner. |
| Project override | Project-owned customization that should not be overwritten by upstream updates. |
