package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gentleman-programming/ard-agent-ai/internal/installer"
)

// Step constants for the wizard.
const (
	stepWelcome    = 0
	stepDetecting  = 1
	stepSelect     = 2
	stepBasePath   = 3
	stepConfirm    = 4
	stepInstalling = 5
	stepDone       = 6
)

// installStepMsg is sent after each installation sub-step completes.
type installStepMsg struct {
	index int
	err   error
}

// detectedMsg is sent when platform detection finishes.
type detectedMsg struct {
	platforms []installer.Platform
}

// Wizard is the main Bubble Tea model.
type Wizard struct {
	version   string
	step      int
	spinner   spinner.Model
	textInput textinput.Model

	// Detection
	platforms []installer.Platform
	detecting bool

	// Selection (step 2)
	cursor   int
	selected map[int]bool

	// Configuration
	basePath string

	// Installation progress
	installSteps  []string
	installDone   []bool
	installErr    []error
	installingIdx int
	installFinish bool
}

// NewWizard creates and initializes the wizard model.
func NewWizard(version string) *Wizard {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = SelectedStyle

	ti := textinput.New()
	ti.Placeholder = `C:\Projects\`
	ti.CharLimit = 256
	ti.Width = 40

	return &Wizard{
		version:   version,
		step:      stepWelcome,
		spinner:   s,
		textInput: ti,
		selected:  make(map[int]bool),
		basePath:  `C:\Projects\`,
	}
}

func (w *Wizard) Init() tea.Cmd {
	return w.spinner.Tick
}

func (w *Wizard) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		return w.handleKey(msg)

	case spinner.TickMsg:
		var cmd tea.Cmd
		w.spinner, cmd = w.spinner.Update(msg)
		return w, cmd

	case detectedMsg:
		w.platforms = msg.platforms
		w.detecting = false
		// Pre-select all detected platforms using the filtered index
		// (the checkbox list only shows detected platforms, so index 0 is the
		// first detected platform, not necessarily platforms[0]).
		idx := 0
		for _, p := range w.platforms {
			if p.Detected {
				w.selected[idx] = true
				idx++
			}
		}
		return w, nil

	case installStepMsg:
		if msg.err != nil {
			w.installErr[msg.index] = msg.err
		}
		w.installDone[msg.index] = true
		next := msg.index + 1
		if next < len(w.installSteps) {
			w.installingIdx = next
			return w, w.runInstallStep(next)
		}
		// All steps done
		w.installFinish = true
		w.step = stepDone
		return w, nil
	}

	// Delegate textinput updates when on basePath step
	if w.step == stepBasePath {
		var cmd tea.Cmd
		w.textInput, cmd = w.textInput.Update(msg)
		return w, cmd
	}

	return w, nil
}

func (w *Wizard) handleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch w.step {

	case stepWelcome:
		switch msg.String() {
		case "enter":
			w.step = stepDetecting
			w.detecting = true
			return w, tea.Batch(w.spinner.Tick, w.detectPlatforms())
		case "esc", "q", "ctrl+c":
			return w, tea.Quit
		}

	case stepDetecting:
		switch msg.String() {
		case "ctrl+c":
			return w, tea.Quit
		case "enter":
			if !w.detecting {
				w.step = stepSelect
				w.cursor = 0
			}
		}

	case stepSelect:
		switch msg.String() {
		case "ctrl+c":
			return w, tea.Quit
		case "esc":
			w.step = stepDetecting
		case "up", "k":
			if w.cursor > 0 {
				w.cursor--
			}
		case "down", "j":
			detected := w.detectedPlatforms()
			if w.cursor < len(detected)-1 {
				w.cursor++
			}
		case " ":
			w.selected[w.cursor] = !w.selected[w.cursor]
		case "enter":
			w.step = stepBasePath
			w.textInput.SetValue(w.basePath)
			w.textInput.Focus()
			return w, textinput.Blink
		}

	case stepBasePath:
		switch msg.String() {
		case "ctrl+c":
			return w, tea.Quit
		case "esc":
			w.textInput.Blur()
			w.step = stepSelect
		case "enter":
			val := strings.TrimSpace(w.textInput.Value())
			if val == "" {
				val = `C:\Projects\`
			}
			w.basePath = val
			w.textInput.Blur()
			w.step = stepConfirm
		}

	case stepConfirm:
		switch msg.String() {
		case "ctrl+c":
			return w, tea.Quit
		case "esc":
			w.step = stepBasePath
			w.textInput.SetValue(w.basePath)
			w.textInput.Focus()
			return w, textinput.Blink
		case "enter":
			return w, w.startInstallation()
		}

	case stepInstalling:
		switch msg.String() {
		case "ctrl+c":
			return w, tea.Quit
		}

	case stepDone:
		switch msg.String() {
		case "enter", "esc", "q", "ctrl+c":
			return w, tea.Quit
		}
	}

	return w, nil
}

// detectPlatforms runs platform detection in a goroutine and returns a Cmd.
func (w *Wizard) detectPlatforms() tea.Cmd {
	return func() tea.Msg {
		platforms := installer.DetectPlatforms()
		return detectedMsg{platforms: platforms}
	}
}

// detectedPlatforms returns only platforms that were found on the system.
func (w *Wizard) detectedPlatforms() []installer.Platform {
	var result []installer.Platform
	for _, p := range w.platforms {
		if p.Detected {
			result = append(result, p)
		}
	}
	return result
}

// selectedPlatforms returns the platforms the user has checked.
func (w *Wizard) selectedPlatforms() []installer.Platform {
	detected := w.detectedPlatforms()
	var result []installer.Platform
	for i, p := range detected {
		if w.selected[i] {
			result = append(result, p)
		}
	}
	return result
}

// startInstallation sets up the install step list and kicks off the first step.
func (w *Wizard) startInstallation() tea.Cmd {
	selected := w.selectedPlatforms()
	steps := make([]string, 0, len(selected)+1)
	steps = append(steps, "Copying skills")
	for _, p := range selected {
		steps = append(steps, fmt.Sprintf("Installing %s", p.Name))
	}

	w.installSteps = steps
	w.installDone = make([]bool, len(steps))
	w.installErr = make([]error, len(steps))
	w.installingIdx = 0
	w.step = stepInstalling

	return tea.Batch(w.spinner.Tick, w.runInstallStep(0))
}

// runInstallStep performs a single installation step asynchronously.
// selected and basePath are captured by value to avoid data races with the
// Bubble Tea model (Cmds run outside the Update goroutine).
func (w *Wizard) runInstallStep(idx int) tea.Cmd {
	selected := w.selectedPlatforms() // snapshot before goroutine
	basePath := w.basePath
	return func() tea.Msg {
		inst := installer.New(basePath, selected)

		var err error
		if idx == 0 {
			// Step 0 is always "copy skills" (common assets)
			err = inst.CopySkills()
		} else {
			// Steps 1..N correspond to selected[idx-1]
			platformIdx := idx - 1
			if platformIdx < len(selected) {
				err = inst.InstallPlatform(selected[platformIdx])
			}
		}

		return installStepMsg{index: idx, err: err}
	}
}

func (w *Wizard) View() string {
	switch w.step {
	case stepWelcome:
		return viewWelcome(w)
	case stepDetecting:
		return viewDetecting(w)
	case stepSelect:
		return viewSelect(w)
	case stepBasePath:
		return viewBasePath(w)
	case stepConfirm:
		return viewConfirm(w)
	case stepInstalling:
		return viewInstalling(w)
	case stepDone:
		return viewDone(w)
	default:
		return ""
	}
}
