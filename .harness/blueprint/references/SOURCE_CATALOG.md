# Reference Source Catalog

All links below are research starting points and should be rechecked against current official documentation before integration. Repository activity, versions, APIs, platform support, licensing, and installation instructions can change.

## Original eight sources

1. **Anthropic — Building agents: When to fold subagents into your main agent loop**  
   https://claude.com/resources/articles/building-agents-when-to-fold-subagents-into-your-main-agent-loop
2. **Agentic Builders — How to use CodeGraph: turn any folder into a knowledge graph**  
   https://medium.com/agentic-builders/how-to-use-codegraph-turn-any-folder-into-a-knowledge-graph-9958f5a2d077
3. **Google Research — Harness²**  
   https://github.com/google-research/harness2
4. **Affaan Mustafa — ECC / Everything Claude Code**  
   https://github.com/affaan-m/ECC
5. **Cheaper Inference — OmniRoute**  
   https://github.com/cheaper-inference/OmniRoute
6. **Graphify Labs — Graphify**  
   https://github.com/Graphify-Labs/graphify
7. **Akita On Rails — ai-memory**  
   https://github.com/akitaonrails/ai-memory
8. **Matt Pocock — Skills for Real Engineers**  
   https://github.com/mattpocock/skills

## Added references selected during research

9. **GitHub — Spec Kit**  
   https://github.com/github/spec-kit
10. **Arize — Phoenix**  
    https://github.com/Arize-ai/phoenix
11. **OpenTelemetry — GenAI semantic conventions**  
    https://opentelemetry.io/docs/specs/semconv/gen-ai/
12. **Agent Skills — Specification**  
    https://agentskills.io/specification
13. **OWASP — MCP Security Cheat Sheet**  
    https://cheatsheetseries.owasp.org/cheatsheets/MCP_Security_Cheat_Sheet.html
14. **SWE-bench Verified**  
    https://www.swebench.com/verified
15. **Terminal-Bench**  
    https://www.tbench.ai/
16. **OpenAI Codex — AGENTS.md / project instructions**  
    https://developers.openai.com/codex/agents_md
17. **OpenAI Codex — Security and sandboxing**  
    https://developers.openai.com/codex/security
18. **Git documentation — submodules**  
    https://git-scm.com/book/en/v2/Git-Tools-Submodules
19. **Git documentation — subtree**  
    https://github.com/git/git/blob/master/contrib/subtree/git-subtree.txt

## Primary-source checks for clarified product decisions

The following official/current sources were checked on **2026-11-19** for the design documents. Recheck before implementation because current documentation and repository branches can change.

20. **OpenAI — Multi-agent (Responses API)**  
    https://developers.openai.com/api/docs/guides/responses-multi-agent
21. **Anthropic — Claude Code custom subagents**  
    https://code.claude.com/docs/en/sub-agents
22. **CodeGraph — Colby McHenry (official repository)**  
    https://github.com/colbymchenry/codegraph
23. **Graphify — Graphify Labs (official repository)**  
    https://github.com/Graphify-Labs/graphify
24. **CodeGraph — supplemental engineering overview**  
    https://dev.to/arshtechpro/codegraph-stop-your-ai-agent-from-grepping-the-same-files-50-times-3bgm
25. **Google Research — Harness² setup**  
    https://github.com/google-research/harness2/blob/main/docs/SETUP.md
26. **Anthropic — broader public Claude documentation index**  
    https://docs.anthropic.com/en/docs/intro
27. **ai-memory — supported runtime matrix**  
    https://github.com/akitaonrails/ai-memory/blob/main/docs/support-matrix.md
28. **ai-memory — installation instructions**  
    https://github.com/akitaonrails/ai-memory/blob/main/docs/install.md
29. **ai-memory — programmatic MCP and handoff integration**  
    https://github.com/akitaonrails/ai-memory/blob/main/docs/programmatic-memory.md
30. **Everything Claude Code (ECC) — official repository**  
    https://github.com/affaan-m/ECC
31. **Matt Pocock — Skills for Real Engineers (official repository)**  
    https://github.com/mattpocock/skills
32. **GitHub Spec Kit — official repository**  
    https://github.com/github/spec-kit
33. **OmniRoute — official repository**  
    https://github.com/cheaper-inference/OmniRoute
34. **Arize Phoenix — official repository**  
    https://github.com/Arize-ai/phoenix
35. **Graphify MCP service — separate hosted-provider documentation**  
    https://github.com/Graphify-Labs/graphify-mcp

## Research interpretation

The external sources inform methods and integrations but do not overrule product choices recorded by Lucas. For example, recommendations to use fewer agents do not remove the required specialist agents; they inform how many specialists a task should invoke.

The combined sources support the following directions:

- Keep the Master Agent responsible for coordination and use actual specialist delegation; select the invoked specialists to fit each task rather than mapping one agent to every stage.
- Make project context and provenance first-class.
- Treat requirements/specifications, implementation plans and acceptance evidence as durable artifacts.
- Treat memory, structural knowledge, telemetry and harness learning as distinct concerns.
- Use evaluation evidence to adopt improvements rather than relying on intuition.
- Keep tools and permissions explicit; repository content and MCP metadata are untrusted input.
- Make integration portable through adapter boundaries and skills conventions, while documenting feature differences between executors.
- Treat CodeGraph and Graphify as complementary intended providers, while validating overlapping code coverage and querying the underlying sources for critical claims.

## Citation and verification note

This catalog is a human-readable bibliography, not a guarantee of current feature parity. During implementation, the local agent should record:

- URL and official source name.
- Access/verification date.
- Exact release/tag/commit if dependency is adopted.
- License and attribution obligations.
- Required runtime/platforms.
- Security and data-handling implications.
- Adopted interface or file format.
- Features deliberately not integrated.
