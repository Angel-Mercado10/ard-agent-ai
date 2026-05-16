package tui

import (
	"fmt"
	"strings"

	"github.com/gentleman-programming/ard-agent-ai/internal/installer"
)

const checkMark = "✔"
const crossMark = "✗"
const pointer = "▶"
const boxChecked = "[x]"
const boxEmpty = "[ ]"

func viewWelcome(w *Wizard) string {
	title := TitleStyle.Render(fmt.Sprintf("🏛  ARD Agent AI   v%s", w.version))

	content := strings.Join([]string{
		title,
		"",
		SubtitleStyle.Render("Architecture Record Document installer"),
		MutedStyle.Render("Define and evolve software architecture"),
		MutedStyle.Render("decisions with AI assistance."),
		"",
		KeyStyle.Render("Press ENTER to continue · ESC to quit"),
	}, "\n")

	return "\n" + BoxPrimaryStyle.Render(content) + "\n"
}

func viewDetecting(w *Wizard) string {
	var b strings.Builder

	b.WriteString("\n")
	if w.detecting {
		b.WriteString(fmt.Sprintf("  %s  %s\n\n", w.spinner.View(), MutedStyle.Render("Detecting platforms...")))
	} else {
		b.WriteString(fmt.Sprintf("  %s\n\n", MutedStyle.Render("Detecting platforms...")))

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
					MutedStyle.Render("not found"),
				))
			}
		}

		b.WriteString("\n")
		b.WriteString(KeyStyle.Render("  Press ENTER to continue"))
	}

	return b.String()
}

func viewSelect(w *Wizard) string {
	detected := w.detectedPlatforms()

	var b strings.Builder
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("  %s\n\n", SectionLabelStyle.Render("Select platforms to install:")))

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
	b.WriteString(KeyStyle.Render("  SPACE to toggle · ENTER to confirm · ESC to go back"))
	b.WriteString("\n")

	return b.String()
}

func viewBasePath(w *Wizard) string {
	var b strings.Builder
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("  %s\n\n", SectionLabelStyle.Render("Where should ARD documents be saved?")))
	b.WriteString(fmt.Sprintf("  ❯ %s\n\n", w.textInput.View()))
	b.WriteString(MutedStyle.Render("  This is the root folder where all ARD.md files\n"))
	b.WriteString(MutedStyle.Render("  will be created when you run /ard-init.\n"))
	b.WriteString("\n")
	b.WriteString(KeyStyle.Render("  ENTER to confirm · ESC to go back"))
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
		SectionLabelStyle.Render("Ready to install"),
		"",
		"  Platforms:",
		platformLines.String(),
		"  ARD documents path:",
		fmt.Sprintf("    %s", SelectedStyle.Render(w.basePath)),
		"",
	}, "\n")

	return "\n" + BoxStyle.Render(content) + "\n\n" +
		KeyStyle.Render("  ENTER to install · ESC to go back") + "\n"
}

func viewInstalling(w *Wizard) string {
	var b strings.Builder
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("  %s\n\n", SectionLabelStyle.Render("Installing...")))

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
			commandsSection.WriteString(fmt.Sprintf("\n  %s — available commands:\n", p.Name))
			commandsSection.WriteString(DividerStyle.Render("  ─────────────────────────────────") + "\n")
			commandsSection.WriteString(fmt.Sprintf("  %-30s %s\n",
				SelectedStyle.Render("/ard-init <project>"),
				MutedStyle.Render("Define architecture"),
			))
			commandsSection.WriteString(fmt.Sprintf("  %-30s %s\n",
				SelectedStyle.Render("/ard-update <project>"),
				MutedStyle.Render("Evolve architecture"),
			))
		}
	}

	content := strings.Join([]string{
		"",
		fmt.Sprintf("  %s", SuccessStyle.Render("✔  Installation complete!")),
		commandsSection.String(),
		MutedStyle.Render("  Restart your editor for changes to take effect."),
		"",
	}, "\n")

	return "\n" + BoxPrimaryStyle.Render(content) + "\n\n" +
		KeyStyle.Render("  Press ENTER or ESC to exit") + "\n"
}
