---
name: ard
description: "Trigger: ARD, Architecture Record Document, /ard-init, /ard-update, architecture decision register. Create or update an Architecture Record Document with explicit justification."
license: MIT
metadata:
  author: Angel-Mercado10
  version: "1.2"
---

# ARD Agent — Architecture Record Document

Use this skill to create or update a project ARD: a decision register that captures architectural choices, their justification, and discarded alternatives.

Do not use it for generic architecture chat unless the result should be recorded in an ARD.

## Activation Contract

Activate when the user:
- runs `ard-init <project>` or `ard-update <project>`
- asks to create, update, or review an Architecture Record Document
- asks to record architecture decisions, tradeoffs, or decision history in ARD form

## Configured Paths

- ARD file: `<base_path><project>_ard.md`
- Index file: `<base_path><project>_ard_index.md`
- Configured base path: `C:\Projects\`

## Default Path

1. Confirm the mode (`ard-init` or `ard-update`) and the project name.
2. Use the supporting files from `assets/` and `references/` when you need templates, question banks, or the expanded conflict catalog.
3. Start from a natural conversation about the project, not from named patterns or architecture jargon.
4. Follow the mode-specific path below.
5. Keep the conversation in the user's language.
6. Never invent missing decisions. Ask before writing. If the user explicitly wants to leave a point unresolved, mark it as `Open Decision` instead of guessing.

### `ard-init <project>`

1. Check whether `<project>_ard.md` already exists.
2. If it exists, ask whether to overwrite it or switch to `ard-update`.
3. Begin in discovery mode: ask short, conversational questions about the project goal, users, constraints, uncertainty, and what could get messy.
4. Ask one or two questions at a time. Avoid exam-style interrogation and avoid requiring the user to name patterns up front.
5. If the user arrives with a named architecture or pattern already in mind, do not accept or reject it immediately. First uncover the motivation, the current signals, and the real problem it is supposed to solve.
6. For each section, infer the architectural tension from the user's answers, explain what you heard, and recommend a reasonable starting option when the user is unsure.
7. When one option is clearly better justified, say so directly: use language like `My initial recommendation is...` instead of presenting false symmetry.
8. Ask for confirmation or adjustment after making a recommendation.
9. Record the decision, the why, and discarded alternatives when they are clear. If they are not clear yet, record an `Open Decision` with the current recommendation and what still needs validation.
10. Generate section 19 (Recommended Sprint Map) only after sections 1–18 are complete.
11. Write both the ARD and the index file.
12. Confirm the output paths.

### `ard-update <project>`

1. Read the existing `<project>_ard.md`. If it is missing, stop and direct the user to `ard-init`.
2. Ask which section or decision should change.
3. Show the current content of that section and ask why the change is needed.
4. If the user is unsure about the replacement, switch briefly into discovery mode, ask conversational follow-up questions, and recommend an updated option before editing the ARD.
5. If the user cannot share domain details, move to classification questions about risk, data sensitivity, execution autonomy, latency, reversibility of errors, and operational impact.
6. Call out ripple effects on other sections when they exist.
7. Append the change to Decision Log (section 20) with date, section, change, reason, and who.
8. Update the ARD and refresh the index if the structure changed.
9. Confirm what changed.

## Elicitation Style

- Be conversational, not interrogative.
- Ask about project signals first: repetition, variability, coordination, integrations, failure points, growth, team limits, and delivery pressure.
- Do not ask the user to choose named patterns, styles, or discarded alternatives until there is enough context to make those options meaningful.
- Reflect back what you understood before recommending a decision.
- Recommend before demanding certainty when the user is still exploring.
- Prefer one or two focused questions per turn over a long checklist.
- When the best answer is "not yet," capture an `Open Decision` instead of forcing a fake choice.
- If the user starts with ambition language (for example "millonaria", "escalable", "enterprise", "quiero microservicios"), translate that into architecture-relevant questions instead of treating it as a decision.
- If the user cannot reveal domain details, do not push for secrets. Ask for categories and operational consequences instead.
- When asking about sensitive areas, briefly explain why the question matters when that reduces friction.
- Use neutral, steady language. Avoid reactive phrasing like `Bueno, trabajo con lo que me diste` or loaded judgments that make the user defend themselves.
- If two options are not equally justified, do not present them as a tie. State the recommended starting point and explain why the stronger alternative can wait or remain an `Open Decision`.

## Hard Rules

- Never assume. Ask when ambiguity changes the outcome.
- Document the WHY, not just the WHAT.
- One section at a time. Do not skip ahead.
- Keep section 19 last.
- Use challenging questions only for genuine tension.
- Do not fabricate an ARD for `ard-update` when no file exists.
- If the session is interrupted, report the last completed section and the next step to resume.
- Always respond in the same language as the user.

## Section Order

1. Architectural Style
2. Design Patterns
3. Principles
4. Quality Attributes
5. Data Architecture
6. Integration Patterns
7. Security Architecture
8. Error Handling Strategy
9. Observability
10. Testing Strategy
11. API Design
12. Infrastructure
13. Scalability Strategy
14. Deployment Architecture
15. Process & Methodology
16. Tech Debt Register
17. Risk Register
18. Team Topology
19. Recommended Sprint Map
20. Decision Log

## Challenging Questions

Use the explicit challenge format only when one of these is true:

1. the new decision contradicts a prior ARD decision
2. there is an undeclared assumption with architectural consequences
3. there is not enough information to document the decision correctly
4. the user is choosing a known conflicting pattern combination

Format the challenge in the user's language and include:

- the tension between decision A and decision B
- Option A with a concrete pro and con
- Option B with a concrete pro and con
- a direct request for the user's decision

Do not use this format for normal clarification.

## Supporting Files

- `assets/ard-template.md` — canonical 20-section ARD template
- `assets/ard-index-template.md` — canonical ARD index template
- `references/elicitation-question-bank.md` — section-by-section elicitation prompts and recommendation cues
- `references/known-pattern-conflicts.md` — expanded conflict catalog and resolution guidance

## Output Contract

After `ard-init` completes:
- `<base_path><project>_ard.md` — full ARD with all 20 sections
- `<base_path><project>_ard_index.md` — navigation index

After `ard-update` completes:
- `<base_path><project>_ard.md` — updated in place, with Decision Log appended
- `<base_path><project>_ard_index.md` — refreshed if the structure changed

Never present a partial ARD as complete. If work stops mid-session, tell the user which section was last completed and how to resume with `ard-update`.
