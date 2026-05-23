# Elicitation Question Bank

Use these prompts when the canonical `ard` skill requires 2–3 focused questions for the current section. Adapt wording to the user's language and context.

## 1. Architectural Style
- What is the primary concern driving the architectural style: maintainability, team size, deployment flexibility, or domain complexity?
- Are you building a monolith that needs to scale, or do you already have independent bounded contexts that justify distribution?

## 2. Design Patterns
- Which patterns are you already using or planning to use (Repository, Factory, Observer, CQRS, Saga, Outbox, etc.)?
- For each pattern, what concrete problem does it solve in this system?

## 3. Principles
- Which SOLID principles are non-negotiable for this project and which are applied selectively?
- Give a concrete example from the domain where DRY would be violated if the team is not careful.

## 4. Quality Attributes
- What are the concrete SLA targets (response time p99, uptime percentage, error-rate threshold)?
- Which quality attribute has priority when they conflict, and has that priority been validated with stakeholders?

## 5. Data Architecture
- What database engine are you choosing, and what specifically ruled out the alternatives?
- Do you have read-heavy workloads that justify CQRS or audit/replay requirements that justify Event Sourcing?

## 6. Integration Patterns
- What systems does this project integrate with, and what latency/reliability requirements exist per integration?
- Do you need event-driven flows, or is synchronous REST sufficient for every integration?

## 7. Security Architecture
- What is the auth strategy (JWT, session, OAuth, API keys), and what drove that choice?
- Do you have multi-tenancy or role-based access requirements? If yes, is RBAC or ABAC a better fit?

## 8. Error Handling Strategy
- What is the retry strategy for transient failures: exponential backoff, fixed delay, or circuit breaker?
- Do you need a dead-letter strategy for messages that cannot be processed?

## 9. Observability
- What logging format and log-level policy apply per environment?
- Do you need distributed tracing, and what alerting thresholds matter on critical paths?

## 10. Testing Strategy
- What minimum coverage target exists, and how is it enforced?
- Are you following TDD, BDD, or test-after, and why?

## 11. API Design
- REST or GraphQL: what drove the choice and what versioning conventions apply?
- Contract-first or code-first: what tradeoff did the team accept?

## 12. Infrastructure
- Which cloud provider fits the constraints, and why?
- Which orchestration model fits the team: Kubernetes, ECS, Cloud Run, or bare VMs?

## 13. Scalability Strategy
- What are the peak-load projections and how were they validated?
- Are services stateless, or does any workload require sticky sessions or external cache?

## 14. Deployment Architecture
- Blue-green, canary, or rolling deployments: what rollback expectation exists?
- Are feature flags required? If yes, which system will manage them?

## 15. Process & Methodology
- Scrum, Kanban, or hybrid: what drove that choice and cycle length?
- Which ceremonies are non-negotiable for this team size?

## 16. Tech Debt Register
- What existing technical debt is already known?
- What policy governs new tech debt: budget, cap, or recurring remediation slot?

## 17. Risk Register
- What are the top three technical risks that could derail the project?
- For each risk, what probability, impact, and mitigation plan apply?

## 18. Team Topology
- How many developers are involved, how are they organized, and who owns each component?
- Does the team structure match the architecture, or is there a Conway's Law mismatch to fix?
