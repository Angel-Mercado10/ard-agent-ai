# ARD-ARCHITECT

[![Release](https://img.shields.io/github/v/release/Angel-Mercado10/ard-agent-ai?style=flat-square)](https://github.com/Angel-Mercado10/ard-agent-ai/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square)](https://opensource.org/licenses/MIT)
[![Go](https://img.shields.io/badge/go-1.22+-00ADD8.svg?style=flat-square&logo=go)](https://go.dev/)

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

```sh
scoop bucket add angel https://github.com/Angel-Mercado10/scoop-bucket
scoop install angel/ard-agent-ai
```

Actualizar a una versión nueva:

```sh
scoop update ard-agent-ai
```

### Go install (cualquier sistema con Go 1.22+)

```sh
go install github.com/Angel-Mercado10/ard-agent-ai/cmd/ard-agent-ai@latest
```

### Descarga directa (macOS / Linux / Windows)

Descarga el binario desde [Releases](https://github.com/Angel-Mercado10/ard-agent-ai/releases/latest):

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

El wizard detecta automáticamente qué editores de IA tienes instalados y copia los archivos del agente en el lugar correcto.

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

Una vez que el wizard termina, reinicia tu editor y vas a tener disponibles:

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

## Filosofía

El agente está construido sobre estos principios:

- **Documentar el WHY, no el WHAT.** El código ya muestra lo que hace; el ARD captura por qué se decidió así.
- **Socrático, no asistente pasivo.** El agente cuestiona supuestos, expone tradeoffs y no permite decisiones sin justificación.
- **Documento vivo.** El ARD evoluciona con el proyecto. Cada cambio queda registrado con fecha y razón.
- **Conflictos arquitectónicos conocidos.** El agente detecta combinaciones problemáticas (Event Sourcing sin CQRS, Microservices con equipo de menos de 3, Hexagonal con ORM en el dominio…) y las desafía antes de aceptarlas.

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

## Stack técnico

| Componente | Tecnología |
|------------|-----------|
| CLI | Go 1.22+ |
| TUI | [Bubble Tea](https://github.com/charmbracelet/bubbletea) + [Lipgloss](https://github.com/charmbracelet/lipgloss) |
| Distribución | [GoReleaser](https://goreleaser.com/) + [Scoop](https://scoop.sh/) |
| CI/CD | GitHub Actions |

---

## Contribuir

1. Fork + clone
2. `go mod tidy && go test ./...`
3. Abre un PR describiendo el cambio

Cualquier mejora al protocolo socrático del agente, nuevas plataformas o mejoras al TUI son bienvenidas.

### Releasing

Crear un tag con prefijo `v` dispara GoReleaser automáticamente:

```sh
git tag v1.0.3
git push origin v1.0.3
```

Después, actualizar el manifest en [scoop-bucket](https://github.com/Angel-Mercado10/scoop-bucket) con los nuevos hashes del release.

---

## Licencia

MIT — [Angel Mercado](https://github.com/Angel-Mercado10)
