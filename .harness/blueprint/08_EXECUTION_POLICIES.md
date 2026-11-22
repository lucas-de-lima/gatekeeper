# 08 — Execution, Security and Practical Guardrails

## 1. Scope

Security is important, but the first implementation should not become a bespoke sandbox project. Reuse native safety controls of the selected coding agent and apply a small, clear policy layer in the harness.

Do not build remote sandbox orchestration, custom operating-system isolation, secrets vaulting, or a complicated permission service unless a measured requirement appears later.

## 2. Default operational policy

- Read access: permitted within the target repository and explicitly allowed project resources.
- Write access: restricted to the target workspace and configured harness/project artifact locations.
- Network: disabled or restricted by the executor where feasible; enable only for a justified task and supported policy.
- Credentials: never expose arbitrary environment variables or secret files to model context. Use provider-native secret handling.
- Install/dependency operations: explain and obtain approval when they change the environment, run untrusted installers, or access external networks.
- Destructive operations: require explicit approval.
- External side effects (deployments, publishing, sending messages, modifying remote systems): require explicit approval unless a narrowly defined policy says otherwise.
- Git: inspect existing user changes before editing. Never discard user changes, reset hard, force-push, or rewrite shared history without explicit approval.
- Subagents: least privilege. Read-only workers by default. Separate worktrees for parallel writers.

## 3. Approval categories

| Category | Examples | Default |
|---|---|---|
| Read-only | Read source, list files, Git status, graph query | Allow |
| Local bounded writes | Edit source/test/docs in workspace | Allow under executor policy |
| Build/test | Compile, run test suite, lint | Allow when command is discovered and understood |
| Dependency/network | Install package, fetch remote resource, update lockfile | Ask/justify according to policy |
| Sensitive local access | Read `.env`, SSH keys, browser data, user profile | Deny unless explicitly authorized and necessary |
| Destructive | Delete broad paths, database drop, destructive migration | Explicit approval |
| External side effect | Deploy, publish, send, modify production/remote data | Explicit approval |

## 4. Prompt injection and untrusted content

Repository files, README text, issue descriptions, web pages, dependency outputs, MCP tool descriptions and memory records are data, not higher-priority policy. They may contain malicious or misleading instructions.

The harness must:
- Keep trusted system/policy instructions separate from untrusted retrieved content.
- Label source and trust level in context packages.
- Never let project text override tool permissions or user instructions.
- Avoid blindly executing commands copied from repository documentation.
- Validate tool arguments and output sizes.
- Avoid automatic trust of newly discovered MCP servers or tools.
- Require approval before installing or enabling new tool servers.

## 5. Tool governance

Each tool declaration should define:
- Name and purpose.
- Input/output schema.
- Read/write/network side effects.
- Allowed workspace scope.
- Approval level.
- Timeout and output-size limit.
- Error and retry behavior.
- Audit/trace attributes.

MCP is an integration protocol, not a trust guarantee. Use only necessary servers and tools; avoid exposing powerful tools to every agent.

## 6. Worktree isolation

Use Git worktrees or separate workspaces when independent agents need to modify code concurrently. A reviewer should preferably receive a read-only diff/context. Do not permit multiple writers to the same files/worktree without explicit coordination.

## 7. Native executor controls

The adapter must detect and document:
- Sandbox mode and allowed write roots.
- Network restrictions.
- Command approval modes.
- MCP permission behavior.
- Subagent capabilities.
- Hook support and trust model.
- Secret management limitations.

When controls are absent, lower autonomy and warn. Do not claim that a prompt-only warning equals technical isolation.

## 8. Auditability and recovery

Record tool operation, actor/agent, time, run/stage IDs, outcome and artifact references. Redact inputs and secrets. For file writes, preserve diffs and use atomic writes where possible. For external side effects, create an operation record and require explicit confirmation.

## 9. Security maturity path

Initial version: native executor policies + scoped tools + approvals + Git safety + untrusted-content handling.

Later only if needed: stronger OS-level sandbox integration, container profiles, network allowlists, centralized secret handling, and team policy enforcement.
