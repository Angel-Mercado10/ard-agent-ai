# ARD-ARCHITECT

```
    /\_____/\
   /  o   o  \
  ( ==  ^  == )
   )         (
  (           )
 ( (  )   (  ) )
(__(__)___(__)__)
  /|  📐 📏  |\
```

**ARD-ARCHITECT** es un agente de IA que te ayuda a definir y evolucionar las decisiones arquitectónicas de tu software mediante diálogo socrático. Genera un **Architecture Record Document (ARD)** — un documento vivo donde cada decisión técnica tiene su justificación explícita.

No es documentación descriptiva. Es un registro de **por qué** se tomó cada decisión.

---

## ¿Qué cubre un ARD?

| Sección | Qué captura |
|---------|-------------|
| Architectural Style | Hexagonal, Clean, Microservices, Modular Monolith… |
| Design Patterns | Outbox, CQRS, Saga, Repository — con su justificación |
| Principles | Qué SOLID aplica y cómo, DRY, KISS, YAGNI |
| Quality Attributes | Performance, Scalability, Reliability — con targets medibles |
| Data Architecture | DB engine, ORM vs raw, Event Sourcing |
| Integration Patterns | REST, gRPC, eventos, message queues |
| Security Architecture | Auth, RBAC/ABAC, zero trust, encripción |
| Error Handling | Circuit breaker, retry, dead letter queues |
| Observability | Logging, tracing, métricas, alerting |
| Testing Strategy | Pirámide de tests, TDD/BDD, cobertura mínima |
| API Design | Convenciones, versionado, contract-first |
| Infrastructure | Cloud, containers, CI/CD |
| Scalability | Horizontal vs vertical, caching, CDN |
| Deployment | Blue-green, canary, feature flags |
| Process & Methodology | Agile, Scrum, Kanban, Waterfall |
| Tech Debt Register | Deuda conocida y plan de mitigación |
| Risk Register | Riesgos y estrategia de mitigación |
| Team Topology | Conway's Law, ownership de componentes |
| Sprint Map | Sprints recomendados basados en todas las decisiones |
| Decision Log | Historial de cambios con fecha y razón |

---

## Instalación

### Windows — Scoop (recomendado)

> ⚠️ El bucket de Scoop estará disponible en breve. Por ahora usá el binario directo.

```sh
scoop bucket add gentleman https://github.com/Gentleman-Programming/scoop-bucket
scoop install ard-agent-ai
```

### Descarga directa

Descargá el binario para tu sistema desde [Releases](https://github.com/Angel-Mercado10/ard-agent-ai/releases):

```sh
# macOS / Linux
chmod +x ard-agent-ai
./ard-agent-ai

# Windows
ard-agent-ai.exe
```

### Desde el código fuente

```sh
git clone https://github.com/Angel-Mercado10/ard-agent-ai.git
cd ard-agent-ai
go build -o ard-agent-ai ./cmd/ard-agent-ai
./ard-agent-ai
```

---

## Uso

El wizard detecta automáticamente qué editores de IA tenés instalados y copia los archivos del agente en el lugar correcto.

```sh
ard-agent-ai               # Abre el wizard interactivo (recomendado)
```

### Instalación sin TUI (CI/CD o scripting)

```sh
# Instalar en todas las plataformas detectadas
ard-agent-ai install --yes

# Instalar en plataformas específicas con ruta personalizada
ard-agent-ai install --yes --path "D:\MisProyectos\ARD" --platform claude-code,codex

# Ver qué plataformas están instaladas en tu sistema
ard-agent-ai install --list-platforms
```

### Variable de entorno

```sh
# Setear la ruta base globalmente
export ARD_AGENT_AI_BASE_PATH="$HOME/proyectos/ard"
```

---

## Plataformas soportadas

| Plataforma | Directorio de config | ID para `--platform` |
|------------|---------------------|----------------------|
| Claude Code | `~/.claude` | `claude-code` |
| GitHub Copilot | `~/.copilot` | `github-copilot` |
| Qwen Code | `~/.qwen` | `qwen-code` |
| opencode | `~/.config/opencode` | `opencode` |
| OpenAI Codex | `~/.codex` | `codex` |

---

## Comandos disponibles después de instalar

```
/ard-init <proyecto>     Elicitación socrática completa → genera <proyecto>_ard.md
/ard-update <proyecto>   Evoluciona un ARD existente → registra cambios en Decision Log
```

### Ejemplo de sesión

```
> /ard-init mi-saas

¿Cuál es el estilo arquitectónico que estás considerando?

  A) Hexagonal — ✅ Testeable, límites claros  ❌ Más estructura inicial
  B) Layered — ✅ Simple y familiar  ❌ Acoplamiento entre capas
  C) Microservices — ✅ Escala independiente  ❌ Overhead operacional

[El agente continúa sección por sección, cuestiona supuestos,
 expone tradeoffs y documenta el WHY de cada decisión]
```

---

## Estructura del proyecto

```
cmd/ard-agent-ai/         Punto de entrada CLI
internal/app/             Routing de comandos (TUI vs CLI)
internal/tui/             Wizard Bubble Tea (7 pantallas)
internal/installer/       Detección de plataformas + lógica de instalación
internal/assets/          Wrapper //go:embed para archivos del agente
internal/assets/agents/   Archivos del agente embebidos en el binario
agents/                   Copia de referencia de los archivos del agente
```

---

## Contribuir

1. Fork + clone
2. `go mod tidy && go test ./...`
3. Abrí un PR describiendo el cambio

Cualquier mejora al protocolo socrático del agente, nuevas plataformas o mejoras al TUI son bienvenidas.

---

## Licencia

MIT — [Angel Mercado](https://github.com/Angel-Mercado10)
