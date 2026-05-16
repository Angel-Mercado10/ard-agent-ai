package app

import (
	"flag"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gentleman-programming/ard-agent-ai/internal/installer"
	"github.com/gentleman-programming/ard-agent-ai/internal/tui"
)

// Version is set by the build system via ldflags.
var Version = "dev"

// Run is the entry point for the CLI. It routes commands and flags.
func Run(args []string) error {
	if len(args) == 0 || args[0] == "install" {
		if len(args) > 1 {
			return runInstallCLI(args[1:])
		}
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

func runInstallCLI(args []string) error {
	fs := flag.NewFlagSet("install", flag.ContinueOnError)
	basePath := fs.String("path", installer.DefaultBasePath(), "Ruta base donde se guardarán los documentos ARD")
	platformsFlag := fs.String("platform", "", "Plataformas separadas por coma: claude-code,github-copilot,qwen-code,opencode,codex")
	yes := fs.Bool("yes", false, "Instalar sin abrir el asistente interactivo")
	list := fs.Bool("list-platforms", false, "Listar plataformas detectadas")

	if err := fs.Parse(args); err != nil {
		return err
	}

	if *list {
		for _, p := range installer.DetectPlatforms() {
			status := "no detectada"
			if p.Detected {
				status = "detectada"
			}
			fmt.Printf("%-15s %-16s %s\n", p.ID, status, p.Dir)
		}
		return nil
	}

	if !*yes && *platformsFlag == "" {
		return runWizard()
	}

	selected, err := selectPlatforms(*platformsFlag)
	if err != nil {
		return err
	}
	if len(selected) == 0 {
		return fmt.Errorf("no hay plataformas seleccionadas; usa --platform o instala una plataforma compatible")
	}

	resolvedPath, err := installer.ResolveBasePath(*basePath)
	if err != nil {
		return fmt.Errorf("resolving path: %w", err)
	}

	inst := installer.New(resolvedPath, selected)
	if err := inst.CopySkills(); err != nil {
		return err
	}
	for _, p := range selected {
		if err := inst.InstallPlatform(p); err != nil {
			return fmt.Errorf("installing %s: %w", p.Name, err)
		}
	}

	fmt.Printf("Instalación completa. Ruta ARD: %s\n", resolvedPath)
	return nil
}

func selectPlatforms(platformsFlag string) ([]installer.Platform, error) {
	if strings.TrimSpace(platformsFlag) == "" {
		var detected []installer.Platform
		for _, p := range installer.DetectPlatforms() {
			if p.Detected {
				detected = append(detected, p)
			}
		}
		return detected, nil
	}

	var selected []installer.Platform
	for _, rawID := range strings.Split(platformsFlag, ",") {
		id := strings.TrimSpace(rawID)
		if id == "" {
			continue
		}
		p, ok := installer.PlatformByID(id)
		if !ok {
			return nil, fmt.Errorf("plataforma desconocida %q", id)
		}
		selected = append(selected, p)
	}
	return selected, nil
}

func printHelp() {
	fmt.Printf(`ard-agent-ai v%s — Instalador del agente Architecture Record Document

USO:
  ard-agent-ai              Iniciar el asistente de instalación
  ard-agent-ai install      Iniciar el asistente de instalación
  ard-agent-ai install --yes --path "D:\ARD" --platform codex
                            Instalar sin TUI y elegir ruta de salida
  ard-agent-ai install --list-platforms
                            Listar plataformas detectadas
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
