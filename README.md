# ard-agent-ai

TUI installer for the **ARD Agent** — an AI coding assistant agent that helps you define and evolve software architecture decisions through Socratic dialogue and generates Architecture Record Documents (ARDs).

## Supported platforms

| Platform | Config directory |
|----------|-----------------|
| Claude Code | `~/.claude` |
| GitHub Copilot | `~/.copilot` |
| Qwen Code | `~/.qwen` |
| opencode | `~/.config/opencode` |

## Installation

Download the binary for your OS from [Releases](https://github.com/Gentleman-Programming/ard-agent-ai/releases) and run:

```sh
# macOS / Linux
./ard-agent-ai

# Windows
ard-agent-ai.exe
```

The TUI wizard will:
1. Detect which AI platforms are installed on your system
2. Let you choose which ones to install into
3. Ask where ARD documents should be saved (base path)
4. Copy all agent files and patch the base path placeholder

## Commands (after installation)

```
/ard-init <project>     Full socratic elicitation → generates <project>_ard.md
/ard-update <project>   Evolve an existing ARD — records changes in Decision Log
```

## Building from source

```sh
go mod tidy
go build -o ard-agent-ai ./cmd/ard-agent-ai
```

## Releasing

Push a `v*` tag — GitHub Actions runs GoReleaser automatically:

```sh
git tag v1.0.0
git push origin v1.0.0
```

## Project structure

```
cmd/ard-agent-ai/       CLI entry point
internal/app/           Command routing
internal/tui/           Bubble Tea wizard (7 screens)
internal/installer/     Platform detection + file installation logic
internal/assets/        //go:embed wrapper for agent markdown files
internal/assets/agents/ Embedded agent files (copied from source)
agents/                 Source agent files (reference copy)
```

## License

MIT
