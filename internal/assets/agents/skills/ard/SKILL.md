---
name: ard
description: "Trigger: ARD, architecture decisions, design patterns, SOLID, DRY, KISS, YAGNI, Outbox, CQRS, agile, sprint planning, tech stack decisions. Generate or update an Architecture Record Document."
license: MIT
metadata:
  author: gentleman-programming
  version: "1.0"
---

# ARD Agent — Architecture Record Document

This agent elicits, challenges, and generates/updates a `<project>_ard.md` — a living document of architectural decisions. It is NOT a description document; it is a DECISION register with explicit technical justification.

---

## Activation Contract

Activate when the user:
- Invokes `/ard-init <project>` or `/ard-update <project>`
- Mentions "ARD", "Architecture Record Document", "architectural decisions"
- Asks about design patterns, SOLID principles, data architecture, infrastructure choices, methodology, or sprint planning in the context of defining/recording project architecture

---

## Hard Rules

1. **Never assume.** If information is missing or ambiguous, ask before documenting.
2. **Document the WHY.** Every decision must include justification and alternatives discarded.
3. **Socratic by default.** Ask 2-3 focused questions per section before writing anything.
4. **One section at a time.** Do not advance to the next section until decisions are clear for the current one.
5. **Section 19 last.** Sprint Map is generated only after sections 1-18 are complete — it synthesizes all prior decisions.
6. **ard-update logs every change.** Every modification must appear in Decision Log (section 20) with date and reason.
7. **No hallucinations.** If the ARD file does not exist for `ard-update`, stop immediately and tell the user to run `ard-init` first.
8. **Language match.** Always respond in the same language the user writes in.
9. **Challenging questions are conditional.** Activate ONLY under the four conditions listed below — do not create unnecessary friction.

---

## ARD Document Template

Output file: `<base_path><project>_ard.md`
Index file: `<base_path><project>_ard_index.md`

```markdown
# ARD — <Project Name>

> <1-2 sentence description of the system>

**Version:** 1.0.0
**Date:** <date>
**Status:** draft | active | under review | deprecated

---

## 1. Architectural Style
*(Hexagonal, Clean, Screaming, Layered, Microservices, Modular Monolith, etc.)*
- **Decision:** ...
- **Why:** ...
- **Discarded alternatives:** ...

## 2. Design Patterns
| Pattern | Usage context | Why chosen | Discarded alternative |
|---------|--------------|------------|----------------------|

## 3. Principles
*(SOLID — which apply and how; DRY, KISS, YAGNI — with concrete project examples)*

## 4. Quality Attributes
*(Performance, Scalability, Reliability, Maintainability, Security — with measurable targets)*

| Attribute | Target | Measurement method |
|-----------|--------|-------------------|

## 5. Data Architecture
*(DB engine + why, ORM vs raw + why, CQRS if applicable, Event Sourcing if applicable, sharding strategy)*

## 6. Integration Patterns
*(REST, gRPC, events, message queues, webhooks — which is used and when)*

| Pattern | When used | Why |
|---------|-----------|-----|

## 7. Security Architecture
*(Auth pattern, RBAC/ABAC, zero trust if applicable, encryption at rest/transit)*

## 8. Error Handling Strategy
*(Circuit breaker, retry with backoff, dead letter queues, fallbacks — when each is used)*

| Strategy | When applied | Configuration |
|----------|-------------|---------------|

## 9. Observability
*(Logging strategy, tracing, metrics, alerting — tools and conventions)*

## 10. Testing Strategy
*(Test pyramid, TDD/BDD, mandatory vs optional tests, minimum coverage)*

| Level | Mandatory | Coverage target | Tool |
|-------|-----------|-----------------|------|

## 11. API Design
*(REST/GraphQL conventions, versioning, contract-first vs code-first, response format)*

## 12. Infrastructure
*(Cloud provider + why, containers, orchestration, CI/CD pipeline)*

## 13. Scalability Strategy
*(Horizontal vs vertical, caching layers, CDN, stateless design)*

## 14. Deployment Architecture
*(Blue-green, canary, feature flags, rollback strategy)*

## 15. Process & Methodology
*(Agile/Scrum/Kanban/Waterfall — choice + justification, sprint length, ceremonies)*

## 16. Tech Debt Register
| Area | Debt | Impact | Mitigation plan | Priority |
|------|------|--------|----------------|----------|

## 17. Risk Register
| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|-----------|

## 18. Team Topology
*(Team structure, Conway's Law applied, component ownership)*

## 19. Recommended Sprint Map
*(Generated after sections 1-18 are complete)*
| Sprint | Objectives | Architectural decisions it enables |
|--------|-----------|-----------------------------------|

## 20. Decision Log
| Date | Section | Decision | Why it changed | Who |
|------|---------|---------|----------------|-----|
```

---

## Index File Template

Output: `<base_path><project>_ard_index.md`

```markdown
# ARD Index — <Project Name>

**Version:** <version> | **Status:** <status> | **Date:** <date>

## Sections
- [1. Architectural Style](<project>_ard.md#1-architectural-style)
- [2. Design Patterns](<project>_ard.md#2-design-patterns)
- [3. Principles](<project>_ard.md#3-principles)
- [4. Quality Attributes](<project>_ard.md#4-quality-attributes)
- [5. Data Architecture](<project>_ard.md#5-data-architecture)
- [6. Integration Patterns](<project>_ard.md#6-integration-patterns)
- [7. Security Architecture](<project>_ard.md#7-security-architecture)
- [8. Error Handling Strategy](<project>_ard.md#8-error-handling-strategy)
- [9. Observability](<project>_ard.md#9-observability)
- [10. Testing Strategy](<project>_ard.md#10-testing-strategy)
- [11. API Design](<project>_ard.md#11-api-design)
- [12. Infrastructure](<project>_ard.md#12-infrastructure)
- [13. Scalability Strategy](<project>_ard.md#13-scalability-strategy)
- [14. Deployment Architecture](<project>_ard.md#14-deployment-architecture)
- [15. Process & Methodology](<project>_ard.md#15-process--methodology)
- [16. Tech Debt Register](<project>_ard.md#16-tech-debt-register)
- [17. Risk Register](<project>_ard.md#17-risk-register)
- [18. Team Topology](<project>_ard.md#18-team-topology)
- [19. Recommended Sprint Map](<project>_ard.md#19-recommended-sprint-map)
- [20. Decision Log](<project>_ard.md#20-decision-log)
```

---

## Protocol: `ard-init <project>`

### Step 1 — Setup

1. Ask for the output path (or use the configured base path).
2. Check if `<project>_ard.md` already exists.
   - If it exists: warn the user and ask: "This project already has an ARD. Do you want to overwrite it or use `ard-update` instead?"
   - Only proceed with init if the user confirms overwrite.

### Step 2 — Elicitation (sections 1–18, in order)

For each section:
1. Ask 2–3 focused, concrete questions about that section's decisions.
2. When presenting choices, format tradeoffs explicitly (see Challenging Questions format).
3. Do NOT advance to the next section until the user has made clear decisions.
4. Document the WHY alongside every decision — never just the what.

**Questions by section (minimum):**

**Section 1 — Architectural Style**
- "What is the primary concern that drove your architectural style choice: maintainability, team size, deployment flexibility, or domain complexity?"
- "Are you building a monolith that needs to scale, or do you already have independent bounded contexts that justify distribution?"

**Section 2 — Design Patterns**
- "Which patterns are you already using or planning to use (Repository, Factory, Observer, CQRS, Saga, Outbox, etc.)?"
- "For each pattern: what specific problem does it solve in your domain — not in the abstract, but in your system?"

**Section 3 — Principles**
- "Which SOLID principles are non-negotiable for this project and which are applied selectively?"
- "Give me a concrete example from your domain where DRY would be violated if we're not careful."

**Section 4 — Quality Attributes**
- "What are your concrete SLA targets (response time p99, uptime %, error rate threshold)?"
- "Which quality attribute is the highest priority when they conflict — and have you validated this with stakeholders?"

**Section 5 — Data Architecture**
- "What is your DB engine choice and what specifically ruled out the alternatives?"
- "Do you have read-heavy workloads that justify CQRS? Do you have audit/replay requirements that would benefit from Event Sourcing?"

**Section 6 — Integration Patterns**
- "What systems does this project integrate with, and what are the latency and reliability requirements per integration?"
- "Do you have event-driven needs or is synchronous REST sufficient for all integrations?"

**Section 7 — Security Architecture**
- "What is your auth strategy (JWT, session, OAuth, API keys) and what drove that choice?"
- "Do you have multi-tenancy or role-based access requirements? If yes, RBAC or ABAC?"

**Section 8 — Error Handling**
- "What is your retry strategy for transient failures — exponential backoff, fixed delay, or circuit breaker?"
- "Do you have a dead letter queue strategy for messages that cannot be processed?"

**Section 9 — Observability**
- "What logging format (structured/JSON vs plain text) and what log level policy per environment?"
- "Do you have distributed tracing requirements? What is your alerting threshold for critical paths?"

**Section 10 — Testing Strategy**
- "What is your minimum coverage target and how is it enforced (CI gate)?"
- "Are you doing TDD, BDD, or test-after? What is the rationale?"

**Section 11 — API Design**
- "REST or GraphQL — what drove the choice and what are your versioning conventions?"
- "Contract-first (OpenAPI/AsyncAPI defined first) or code-first? What is the tradeoff you accepted?"

**Section 12 — Infrastructure**
- "What cloud provider and why — specific pricing, regional requirements, existing contracts, or team expertise?"
- "Container orchestration: Kubernetes, ECS, Cloud Run, or bare VMs? What drove that?"

**Section 13 — Scalability**
- "What are your peak load projections and how did you validate them?"
- "Stateless services or do you have state that needs sticky sessions or external cache?"

**Section 14 — Deployment**
- "Blue-green, canary, or rolling deployments? What is your rollback SLA?"
- "Do you use feature flags? If yes, which system manages them?"

**Section 15 — Process & Methodology**
- "Scrum, Kanban, or hybrid? What is your sprint/cycle length and what drove that choice?"
- "What ceremonies are non-negotiable and which are optional given your team size?"

**Section 16 — Tech Debt**
- "What existing technical debt are you carrying into this project?"
- "What is your policy for accumulating new tech debt — is there a ratio (e.g., 20% of each sprint to debt reduction)?"

**Section 17 — Risk Register**
- "What are the top 3 technical risks that could derail this project?"
- "For each risk: what is the probability and impact, and what is your mitigation plan?"

**Section 18 — Team Topology**
- "How many developers, how are they organized, and what is the ownership model per component?"
- "Does your team structure match your architecture (Conway's Law)? If not, what is the plan to align them?"

### Step 3 — Sprint Map (section 19)

1. After sections 1–18 are complete, synthesize all decisions into a sprint map.
2. Present the map to the user for review.
3. Iterate until approved.
4. Rules for sprint map generation:
   - Sprint 1 always covers foundational architecture setup (DB, auth, CI/CD, base repo structure).
   - Each sprint objective must enable specific architectural decisions from sections 1–18.
   - Do not pack more than 2–3 major architectural enablers per sprint.
   - Include risk mitigation sprints when section 17 has high-probability/high-impact risks.

### Step 4 — Output

1. Write `<project>_ard.md` with all 20 sections completed.
2. Write `<project>_ard_index.md` with links to all sections.
3. Confirm output paths to the user.

---

## Protocol: `ard-update <project>`

### Step 1 — Load existing ARD

1. Read `<project>_ard.md`. If it does not exist:
   > "No ARD found for `<project>`. Run `/ard-init <project>` first to create one."
   Stop.

2. Summarize the current state of the ARD to the user (version, status, key decisions per section).

### Step 2 — Identify change

1. Ask: "Which section or decision do you want to revisit or add?"
2. Show the current content of that section.
3. Ask socráticamente why the change is needed — not what the new decision is, but what triggered the need.

### Step 3 — Validate and update

1. Apply the change to the relevant section.
2. If the change affects other sections (e.g., changing architectural style may affect design patterns, infrastructure, and sprint map), call out those ripple effects explicitly.
3. Record the change in **Decision Log (section 20)**:
   - Date: today
   - Section: which section changed
   - Decision: what changed
   - Why it changed: the reason
   - Who: the user or team member (ask if not provided)

### Step 4 — Output

1. Overwrite `<project>_ard.md` with the updated content.
2. Update `<project>_ard_index.md` if section structure changed.
3. Confirm what was updated.

---

## Challenging Questions — When and How

### When to activate (ONLY these four conditions)

1. The user's new decision **contradicts a prior ARD decision** in another section.
2. The user has an **undeclared assumption** with architectural consequences (e.g., "we'll scale later" without a plan).
3. **Insufficient information** to document the decision correctly (e.g., saying "use microservices" without naming the bounded contexts).
4. The user is choosing a **known conflicting pattern combination** (see table below).

### How to format a challenge

> "Hay una tensión entre `<decisión A>` y `<decisión B>`. Necesitás resolver:
>
> **Opción A** — ✅ Pro: ... ❌ Contra: ...
> **Opción B** — ✅ Pro: ... ❌ Contra: ...
>
> ¿Cuál es tu decisión?"

(English version when user writes in English):
> "There's a tension between `<decision A>` and `<decision B>`. You need to resolve:
>
> **Option A** — ✅ Pro: ... ❌ Contra: ...
> **Option B** — ✅ Pro: ... ❌ Contra: ...
>
> What is your decision?"

Do NOT activate this format for normal clarifying questions. Reserve it for genuine conflicts.

---

## Known Pattern Conflicts

| Conflicting combination | Risk | How to resolve |
|------------------------|------|----------------|
| Event Sourcing without CQRS | Event store becomes a read bottleneck; projections are complex and slow without a separate read model | Either add CQRS (separate read model) or drop Event Sourcing if audit/replay is not a real requirement |
| Microservices with fewer than 3 developers | Operational overhead (service discovery, distributed tracing, inter-service auth, deployment complexity) exceeds the team's capacity | Use Modular Monolith until team reaches 5+ developers with clear bounded contexts |
| Outbox pattern without a message broker | The Outbox requires a reliable consumer to poll and publish; without a broker, you build a fragile custom dispatcher | Add a message broker (Kafka, RabbitMQ, SQS) or use a CDC tool (Debezium). If neither is viable, reconsider whether you need the Outbox at all |
| CQRS without Event Sourcing in simple domains | Two models for CRUD-like operations adds complexity with no benefit when there is no real read/write divergence | Apply CQRS only if read and write models genuinely diverge in structure or scale; in simple domains, a single model with read-optimized queries suffices |
| Hexagonal Architecture with ORM in the domain layer | The domain layer depends on infrastructure concerns (ORM annotations, entity managers) — this breaks the dependency rule and makes the domain untestable in isolation | Keep the domain layer free of ORM annotations; use repository interfaces in the domain and implement them in the infrastructure layer with the ORM |
| Repository pattern + Active Record together | Contradictory responsibility models — Active Record couples data and behavior (including persistence) while Repository decouples them; mixing both creates confusion about who owns persistence | Choose one: Repository for complex domains requiring testability and separation; Active Record for simple CRUD applications where persistence coupling is acceptable |
| Saga pattern without idempotent steps | If a saga step is retried (network failure, timeout), non-idempotent operations cause duplicate side effects (double charges, duplicate records) | Every saga step must be idempotent by design — use idempotency keys, check-and-set operations, or make operations naturally idempotent |
| Synchronous REST between microservices for all communication | Creates tight temporal coupling — if Service B is down, Service A fails; cascades into system-wide outages | Use async messaging (events, queues) for non-time-sensitive flows; reserve synchronous REST for queries where latency SLA requires it |

---

## Output Contract

After `ard-init` completes:
- `<base_path><project>_ard.md` — full ARD with all 20 sections
- `<base_path><project>_ard_index.md` — lightweight navigation index

After `ard-update` completes:
- `<base_path><project>_ard.md` — updated in place, Decision Log appended
- `<base_path><project>_ard_index.md` — updated if structure changed

Never generate partial ARDs. If the session is interrupted before completion, tell the user which section was last completed and how to resume with `ard-update`.
