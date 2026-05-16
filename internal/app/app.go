package app

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gentleman-programming/ard-agent-ai/internal/tui"
)

// Version is set by the build system via ldflags.
var Version = "dev"

// Run is the entry point for the CLI. It routes commands and flags.
func Run(args []string) error {
	if len(args) == 0 || args[0] == "install" {
		return runWizard()
	}

	switch args[0] {
	case "--version", "-v", "version":
		fmt.Printf("ard-agent-ai v%s\n", Version)
		return nil
	case "--help", "-h", "help":
		printHelp()
		return nil
	default:
		fmt.Printf("Comando desconocido: %s\n\n", args[0])
		printHelp()
		return nil
	}
}

func runWizard() error {
	m := tui.NewWizard(Version)
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		return fmt.Errorf("running TUI: %w", err)
	}
	return nil
}

func printHelp() {
	fmt.Printf(`ard-agent-ai v%s — Instalador del agente Architecture Record Document

USO:
  ard-agent-ai              Iniciar el asistente de instalación
  ard-agent-ai install      Iniciar el asistente de instalación
  ard-agent-ai --version    Mostrar versión
  ard-agent-ai --help       Mostrar esta ayuda

DESCRIPCIÓN:
  Instala el Agente ARD en tus asistentes de código con IA
  (Claude Code, GitHub Copilot, Qwen Code, opencode, OpenAI Codex).

  El agente te ayuda a definir y evolucionar decisiones de arquitectura
  de software mediante diálogo socrático y genera Architecture
  Record Documents (ARDs).
`, Version)
}
