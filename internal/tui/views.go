package tui

import (
	"fmt"
	"strings"

	"github.com/Angel-Mercado10/ard-agent-ai/internal/installer"
)

const checkMark = "✔"
const crossMark = "✗"
const pointer = "▶"
const boxChecked = "[x]"
const boxEmpty = "[ ]"

const dogAscii = `
    /\_____/\
   /  o   o  \
  ( ==  ^  == )
   )         (
  (           )
 ( (  )   (  ) )
(__(__)___(__)__)
  /|  📐 📏  |\
`

func viewWelcome(w *Wizard) string {
	dog := MutedStyle.Render(dogAscii)
	title := TitleStyle.Render(fmt.Sprintf("🏛  ARD-ARCHITECT   v%s", w.version))

	content := strings.Join([]string{
		dog,
		title,
		"",
		SubtitleStyle.Render("Instalador de Architecture Record Documents"),
		MutedStyle.Render("Define y evoluciona decisiones de arquitectura"),
		MutedStyle.Render("de software con asistencia de IA."),
		"",
		KeyStyle.Render("ENTER para continuar · ESC para salir"),
	}, "\n")

	return "\n" + BoxPrimaryStyle.Render(content) + "\n"
}

func viewDetecting(w *Wizard) string {
	var b strings.Builder

	b.WriteString("\n")
	if w.detecting {
		b.WriteString(fmt.Sprintf("  %s  %s\n\n", w.spinner.View(), MutedStyle.Render("Detectando plataformas...")))
	} else {
		b.WriteString(fmt.Sprintf("  %s\n\n", MutedStyle.Render("Detectando plataformas...")))

		for _, p := range w.platforms {
			if p.Detected {
				b.WriteString(fmt.Sprintf("  %s  %-20s %s\n",
					SuccessStyle.Render(checkMark),
					p.Name,
					MutedStyle.Render(p.Dir),
				))
			} else {
				b.WriteString(fmt.Sprintf("  %s  %-20s %s\n",
					ErrorStyle.Render(crossMark),
					MutedStyle.Render(p.Name),
					MutedStyle.Render("no encontrado"),
				))
			}
		}

		b.WriteString("\n")
		b.WriteString(KeyStyle.Render("  ENTER para continuar"))
	}

	return b.String()
}

func viewSelect(w *Wizard) string {
	detected := w.detectedPlatforms()

	var b strings.Builder
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("  %s\n\n", SectionLabelStyle.Render("Selecciona las plataformas a instalar:")))

	for i, p := range detected {
		cursor := "  "
		if w.cursor == i {
			cursor = SelectedStyle.Render(pointer) + " "
		}

		checkbox := boxEmpty
		if w.selected[i] {
			checkbox = CheckboxStyle.Render(boxChecked)
		}

		b.WriteString(fmt.Sprintf("  %s%s %s\n", cursor, checkbox, p.Name))
	}

	b.WriteString("\n")
	b.WriteString(KeyStyle.Render("  SPACE para marcar · ENTER para confirmar · ESC para volver"))
	b.WriteString("\n")

	return b.String()
}

func viewBasePath(w *Wizard) string {
	var b strings.Builder
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("  %s\n\n", SectionLabelStyle.Render("¿Dónde quieres guardar tus documentos ARD?")))
	b.WriteString(fmt.Sprintf("  ❯ %s\n\n", w.textInput.View()))
	b.WriteString(MutedStyle.Render("  Esta es la carpeta raíz donde se crearán todos\n"))
	b.WriteString(MutedStyle.Render("  los archivos ARD.md al usar /ard-init.\n"))
	b.WriteString("\n")
	b.WriteString(KeyStyle.Render("  ENTER para confirmar · ESC para volver"))
	b.WriteString("\n")

	return b.String()
}

func viewConfirm(w *Wizard) string {
	selected := w.selectedPlatforms()

	var platformLines strings.Builder
	for _, p := range selected {
		platformLines.WriteString(fmt.Sprintf("    • %-20s %s\n",
			p.Name,
			MutedStyle.Render(p.Dir),
		))
	}

	content := strings.Join([]string{
		SectionLabelStyle.Render("Listo para instalar"),
		"",
		"  Plataformas:",
		platformLines.String(),
		"  Ruta de documentos ARD:",
		fmt.Sprintf("    %s", SelectedStyle.Render(w.basePath)),
		"",
	}, "\n")

	return "\n" + BoxStyle.Render(content) + "\n\n" +
		KeyStyle.Render("  ENTER para instalar · ESC para volver") + "\n"
}

func viewInstalling(w *Wizard) string {
	var b strings.Builder
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("  %s\n\n", SectionLabelStyle.Render("Instalando...")))

	for i, step := range w.installSteps {
		if w.installDone[i] {
			if w.installErr[i] != nil {
				b.WriteString(fmt.Sprintf("  %s  %s — %s\n",
					ErrorStyle.Render(crossMark),
					step,
					ErrorStyle.Render(w.installErr[i].Error()),
				))
			} else {
				b.WriteString(fmt.Sprintf("  %s  %s\n",
					SuccessStyle.Render(checkMark),
					step,
				))
			}
		} else if i == w.installingIdx {
			b.WriteString(fmt.Sprintf("  %s  %s\n",
				w.spinner.View(),
				MutedStyle.Render(step+"..."),
			))
		} else {
			b.WriteString(fmt.Sprintf("     %s\n", MutedStyle.Render(step)))
		}
	}

	return b.String()
}

func viewDone(w *Wizard) string {
	selected := w.selectedPlatforms()

	var commandsSection strings.Builder
	for _, p := range selected {
		if p.ID == installer.PlatformClaudeCode {
			commandsSection.WriteString(fmt.Sprintf("\n  %s — comandos disponibles:\n", p.Name))
			commandsSection.WriteString(DividerStyle.Render("  ─────────────────────────────────") + "\n")
			commandsSection.WriteString(fmt.Sprintf("  %-30s %s\n",
				SelectedStyle.Render("/ard-init <project>"),
				MutedStyle.Render("Definir arquitectura"),
			))
			commandsSection.WriteString(fmt.Sprintf("  %-30s %s\n",
				SelectedStyle.Render("/ard-update <project>"),
				MutedStyle.Render("Evolucionar arquitectura"),
			))
		}
		if p.ID == installer.PlatformCodex {
			commandsSection.WriteString(fmt.Sprintf("\n  %s — instrucciones instaladas:\n", p.Name))
			commandsSection.WriteString(DividerStyle.Render("  ─────────────────────────────────────────") + "\n")
			commandsSection.WriteString(MutedStyle.Render("  Abre una sesión de Codex y usa /ard-init\n"))
		}
	}

	content := strings.Join([]string{
		"",
		fmt.Sprintf("  %s", SuccessStyle.Render("✔  ¡Instalación completa!")),
		commandsSection.String(),
		MutedStyle.Render("  Reinicia tu editor para que los cambios tomen efecto."),
		"",
	}, "\n")

	return "\n" + BoxPrimaryStyle.Render(content) + "\n\n" +
		KeyStyle.Render("  ENTER o ESC para salir") + "\n"
}
