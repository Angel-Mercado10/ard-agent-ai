# Known Pattern Conflicts

Use this catalog when the canonical `ard` skill needs to challenge a conflicting architectural combination. Summarize only the entries that matter to the current decision.

| Conflicting combination | Risk | How to resolve |
|------------------------|------|----------------|
| Event Sourcing without CQRS | The event store becomes a read bottleneck and projections grow costly without a separate read model. | Either add CQRS with a dedicated read model or drop Event Sourcing if audit/replay is not a real requirement. |
| Microservices with fewer than 3 developers | Operational overhead exceeds the team's capacity. | Prefer a Modular Monolith until the team and bounded contexts justify distribution. |
| Outbox pattern without a message broker | The outbox depends on a reliable publisher; without one, the system grows a fragile custom dispatcher. | Add a broker or CDC mechanism, or reconsider whether the Outbox pattern is necessary. |
| CQRS without Event Sourcing in simple domains | Two models for CRUD-heavy flows add complexity without meaningful read/write divergence. | Use CQRS only when read and write models genuinely diverge in shape or scale. |
| Hexagonal Architecture with ORM in the domain layer | Infrastructure details leak into the domain and break dependency boundaries. | Keep ORM annotations and persistence mechanisms in infrastructure adapters. |
| Repository pattern plus Active Record | Responsibility boundaries become contradictory and unclear. | Choose Repository for complex, testable domains or Active Record for simple CRUD flows. |
| Saga pattern without idempotent steps | Retries create duplicate side effects such as repeated charges or duplicated records. | Make every step idempotent with keys, natural idempotence, or check-and-set semantics. |
| Synchronous REST between microservices for all communication | Tight temporal coupling turns one service outage into a broader failure cascade. | Reserve synchronous calls for strict query paths and use async messaging for looser workflows. |
