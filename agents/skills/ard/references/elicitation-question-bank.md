# Elicitation Question Bank

Use this reference when the canonical `ard` skill needs help guiding a section. The goal is not to test the user on architecture vocabulary. The goal is to talk naturally about the project, detect signals, and turn those signals into recommendations.

The example phrasings below are reference prompts, not fixed script. Translate and adapt them to the user's language, tone, and context before asking them.

For each section:
- ask short, conversational questions
- listen for project signals and tensions
- reflect back what you understood
- recommend a reasonable default when the user is unsure
- capture an `Open Decision` when the evidence is still weak

## Special Handling: Named Architecture Up Front
- If the user opens with a solution like microservices, hexagonal, CQRS, or event sourcing, do not force them to defend it immediately.
- First ask what problem they are trying to solve with that choice: speed of growth, team autonomy, domain complexity, resilience, practice, or credibility.
- Then ask what current signals already justify it, versus what is only future ambition.
- If the evidence is weak, recommend the simpler starting point explicitly and leave the heavier option as `Open Decision`.

## Special Handling: Confidential Domain
- If the user cannot reveal domain details, say that is fine and switch to classification questions.
- Ask about categories such as: real money vs analysis, autonomous execution vs human approval, reversible vs irreversible errors, public vs internal users, real-time vs batch, sensitive vs non-sensitive data.
- Briefly explain why the question matters when that helps reduce friction.
- Do not pressure the user for proprietary algorithms, internal strategies, or confidential business logic.

## 1. Architectural Style
- Ask: "Contame qué querés construir y qué parte te preocupa más que se desordene con el tiempo."
- Ask: "¿Esto arranca como algo chico para validar rápido o como una base que ya querés dejar bastante prolija?"
- Ask when the user names a heavy architecture early: "Cuando pensás en crecer, ¿te imaginás partes que de verdad deberían evolucionar o desplegarse por separado, o más bien querés una base ordenada para crecer sin caos?"
- Listen for: delivery speed, domain complexity, team size, future splitting pressure, independent deployability.
- Recommend when: simple early products often start better with layered or modular monolith; strong domain boundaries or heavy integration pressure may justify something more explicit. If the current scope is small and the growth case is still abstract, recommend modular monolith first and leave microservices as `Open Decision`.

## 2. Design Patterns
- Ask: "¿Hay comportamientos o flujos que ya te imaginás que se van a repetir con pequeñas variantes?"
- Ask: "¿Hay partes donde sentís que, si no ordenás desde el inicio, vas a terminar copiando lógica o metiendo demasiados `if`?"
- Listen for: repetition, variability, orchestration, lifecycle complexity, conditional branching.
- Recommend when: repeated creation logic may suggest Factory; behavior that changes by case may suggest Strategy; multi-step cross-service workflows may suggest Saga or Outbox.

## 3. Principles
- Ask: "¿Qué querés cuidar más en el código: claridad, velocidad para cambiar, bajo acoplamiento, simplicidad?"
- Ask: "¿Dónde sentís que el proyecto se puede volver frágil si nadie pone límites?"
- Listen for: maintainability concerns, coupling, duplication, speculative abstraction.
- Recommend when: keep principles concrete; prefer simple guardrails over abstract doctrine.

## 4. Quality Attributes
- Ask: "Si este proyecto sale mal, ¿qué te dolería más: lentitud, caídas, errores de datos, dificultad para mantenerlo, o problemas de seguridad?"
- Ask: "¿Tenés alguna expectativa concreta de volumen, tiempos de respuesta o disponibilidad, aunque sea aproximada?"
- Listen for: dominant quality attribute, tradeoff priority, rough service-level expectations.
- Recommend when: name one or two primary qualities first; avoid pretending all qualities matter equally.

## 5. Data Architecture
- Ask: "¿Qué tipo de información guarda el sistema y qué tan sensible o cambiante es?"
- Ask: "¿Necesitás consultar mucho, auditar cambios, reconstruir historial, o más bien guardar y leer cosas de forma bastante directa?"
- Listen for: relational structure, auditability, read/write asymmetry, historical traceability.
- Recommend when: simple CRUD usually stays simple; audit/replay needs may justify stronger event models later.

## 6. Integration Patterns
- Ask: "¿Con qué sistemas externos o internos va a hablar este proyecto?"
- Ask: "Cuando uno de esos sistemas falle o responda lento, ¿qué pasa en tu flujo principal?"
- Listen for: synchronous dependency pressure, eventual consistency, fragile external systems, latency sensitivity.
- Recommend when: use sync only where the user experience demands it; prefer async for looser workflows.

## 7. Security Architecture
- Ask: "¿Quién usa esto y qué cosas no debería poder ver o tocar cada tipo de usuario?"
- Ask: "¿Hay datos sensibles, multi-tenant, o requisitos de compliance que ya te condicionen?"
- Listen for: identity shape, role complexity, tenancy, sensitive data boundaries.
- Recommend when: simple auth is often enough at first; add stronger authorization models only when the actor model truly demands them.

## 8. Error Handling Strategy
- Ask: "Cuando algo falle, ¿qué error te preocupa más: perder datos, duplicar acciones, dejar al usuario colgado, o romper una integración?"
- Ask: "¿Hay operaciones que no pueden ejecutarse dos veces sin consecuencias?"
- Listen for: idempotency risk, retry safety, partial failure, compensation needs.
- Recommend when: retries need boundaries; idempotency matters before saga-like orchestration.

## 9. Observability
- Ask: "Cuando algo salga mal en producción, ¿cómo te imaginás encontrando la causa?"
- Ask: "¿Qué te gustaría poder ver sí o sí: errores, tiempos, métricas de negocio, trazas entre servicios?"
- Listen for: debugging pain, cross-service visibility, metric maturity.
- Recommend when: structured logs are a strong default; add tracing when the system genuinely becomes distributed.

## 10. Testing Strategy
- Ask: "¿Dónde te da más miedo romper algo cuando empieces a tocar el sistema seguido?"
- Ask: "¿Preferís validar rápido con pocos tests clave o querés una base de pruebas más fuerte desde el arranque?"
- Listen for: regression fear, pace of change, confidence needs, critical flows.
- Recommend when: prioritize high-signal tests around core rules and risky flows before chasing coverage targets.

## 11. API Design
- Ask: "¿Cómo consumen esto otros actores: frontend propio, clientes externos, otros servicios, integraciones de terceros?"
- Ask: "¿Necesitás una interfaz muy estable hacia afuera o algo interno que puedas cambiar seguido?"
- Listen for: public contract stability, consumer diversity, versioning pressure.
- Recommend when: REST is the safer default unless client flexibility clearly dominates.

## 12. Infrastructure
- Ask: "¿Dónde pensás correr esto y qué te condiciona más: costo, simplicidad operativa, experiencia del equipo, requisitos del cliente?"
- Ask: "¿Querés algo fácil de operar al inicio o ya sabés que vas a necesitar más control?"
- Listen for: operational maturity, team capacity, hosting constraints.
- Recommend when: start with the simplest hosting model that satisfies the current constraints.

## 13. Scalability Strategy
- Ask: "Si esto crece, ¿qué creés que va a sufrir primero: base de datos, procesos pesados, tráfico, o coordinación entre módulos?"
- Ask: "¿Tenés una estimación, aunque sea gruesa, de uso simultáneo o carga esperada?"
- Listen for: real growth signals versus hand-wavy future scaling.
- Recommend when: avoid premature distribution; optimize the first likely bottleneck, not every hypothetical one.

## 14. Deployment Architecture
- Ask: "¿Qué tan costoso sería un despliegue fallido para ustedes?"
- Ask: "¿Necesitás poder apagar cambios rápido, hacer rollout gradual, o alcanza con un deploy simple al principio?"
- Listen for: rollback urgency, feature exposure control, operational caution.
- Recommend when: choose the simplest deployment approach that matches the rollback risk.

## 15. Process & Methodology
- Ask: "¿Cómo trabaja hoy el equipo: iteraciones cortas, tickets sobre la marcha, roadmap más formal?"
- Ask: "¿Qué parte del proceso te ayuda de verdad y cuál sentís que es puro ritual?"
- Listen for: team cadence, ceremony tolerance, planning maturity.
- Recommend when: reflect the real team habits instead of idealized process language.

## 16. Tech Debt Register
- Ask: "¿Qué atajos sabés desde ya que probablemente van a aparecer si priorizan velocidad?"
- Ask: "¿Qué tipo de deuda te parece aceptable y cuál no querés dejar pasar?"
- Listen for: deliberate shortcuts, known fragile zones, quality boundaries.
- Recommend when: capture debt as an explicit tradeoff, not as accidental silence.

## 17. Risk Register
- Ask: "¿Qué podría hacer que este proyecto se complique fuerte aunque la idea sea buena?"
- Ask: "¿Qué te genera más incertidumbre hoy: negocio, integración, performance, equipo, plazos?"
- Listen for: real derailers, uncertainty clusters, concentration of risk.
- Recommend when: keep the top risks concrete and actionable; avoid generic fear lists.

## 18. Team Topology
- Ask: "¿Quiénes van a tocar esto y cómo se reparten el trabajo hoy?"
- Ask: "¿Ves partes del sistema que quedarían sin dueño claro o demasiado cruzadas entre varias personas?"
- Listen for: ownership gaps, bottlenecks, Conway misalignment, too much shared surface.
- Recommend when: adjust ownership boundaries before inventing architectural complexity.

## Conversation Moves That Help
- Reflect back: "Lo que estoy entendiendo es..."
- Explain a sensitive question briefly: "Te lo pregunto porque eso cambia bastante el nivel de riesgo, auditoría y tolerancia a fallo."
- Recommend gently but clearly: "Con eso, mi recomendación inicial sería..."
- Mark asymmetry when it exists: "Podemos explorar ambas rutas, pero hoy la mejor justificada me parece esta..."
- Leave room for uncertainty: "Si todavía no querés cerrar esa decisión, lo podemos dejar como `Open Decision`."
- Stay neutral under confidentiality: "Con ese nivel de abstracción ya puedo orientar la decisión; no necesito entrar en lo confidencial."
- Avoid forcing jargon: if the user never names a pattern, the agent can still recommend one from the observed signals.
