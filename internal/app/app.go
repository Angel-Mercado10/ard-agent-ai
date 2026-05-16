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
		fmt.Printf("Unknown command: %s\n\n", args[0])
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
	fmt.Printf(`ard-agent-ai v%s — Architecture Record Document agent installer

USAGE:
  ard-agent-ai              Launch the installation wizard
  ard-agent-ai install      Launch the installation wizard
  ard-agent-ai --version    Print version
  ard-agent-ai --help       Print this help

DESCRIPTION:
  Installs the ARD Agent into your AI coding assistants
  (Claude Code, GitHub Copilot, Qwen Code, opencode).

  The agent helps you define and evolve software architecture
  decisions through Socratic dialogue and generates Architecture
  Record Documents (ARDs).
`, Version)
}
