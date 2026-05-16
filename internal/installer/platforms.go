package installer

import (
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

// Platform IDs.
const (
	PlatformClaudeCode    = "claude-code"
	PlatformGitHubCopilot = "github-copilot"
	PlatformQwenCode      = "qwen-code"
	PlatformOpenCode      = "opencode"
	PlatformCodex         = "codex"
)

// Platform describes an AI coding assistant that the installer supports.
type Platform struct {
	ID       string
	Name     string
	Dir      string // Absolute path to the platform config directory
	Detected bool
}

// DetectPlatforms probes the user's system for supported AI platforms.
func DetectPlatforms() []Platform {
	home, _ := os.UserHomeDir()

	detectors := []func(string) Platform{
		detectClaudeCode,
		detectGitHubCopilot,
		detectQwenCode,
		detectOpenCode,
		detectCodex,
	}

	platforms := make([]Platform, len(detectors))
	var wg sync.WaitGroup
	for i, detect := range detectors {
		wg.Add(1)
		go func(i int, detect func(string) Platform) {
			defer wg.Done()
			platforms[i] = detect(home)
		}(i, detect)
	}
	wg.Wait()

	return platforms
}

// PlatformByID returns a platform with its conventional config directory even
// when it has not been detected yet. This is useful for explicit CLI installs.
func PlatformByID(id string) (Platform, bool) {
	home, _ := os.UserHomeDir()
	for _, p := range DetectPlatforms() {
		if p.ID == id {
			return p, true
		}
	}

	defaults := map[string]Platform{
		PlatformClaudeCode: {
			ID:   PlatformClaudeCode,
			Name: "Claude Code",
			Dir:  filepath.Join(home, ".claude"),
		},
		PlatformGitHubCopilot: {
			ID:   PlatformGitHubCopilot,
			Name: "GitHub Copilot",
			Dir:  filepath.Join(home, ".copilot"),
		},
		PlatformQwenCode: {
			ID:   PlatformQwenCode,
			Name: "Qwen Code",
			Dir:  filepath.Join(home, ".qwen"),
		},
		PlatformOpenCode: {
			ID:   PlatformOpenCode,
			Name: "opencode",
			Dir:  filepath.Join(home, ".config", "opencode"),
		},
		PlatformCodex: {
			ID:   PlatformCodex,
			Name: "OpenAI Codex",
			Dir:  filepath.Join(home, ".codex"),
		},
	}
	p, ok := defaults[id]
	return p, ok
}

func detectClaudeCode(home string) Platform {
	dir := filepath.Join(home, ".claude")
	p := Platform{
		ID:   PlatformClaudeCode,
		Name: "Claude Code",
		Dir:  dir,
	}
	// Detected if ~/.claude exists (claude binary is optional)
	if _, err := os.Stat(dir); err == nil {
		p.Detected = true
		p.Dir = dir
		return p
	}
	// Fallback: check if claude binary is available
	if _, err := exec.LookPath("claude"); err == nil {
		p.Detected = true
		p.Dir = dir
	}
	return p
}

func detectGitHubCopilot(home string) Platform {
	dir := filepath.Join(home, ".copilot")
	p := Platform{
		ID:   PlatformGitHubCopilot,
		Name: "GitHub Copilot",
		Dir:  dir,
	}
	if _, err := os.Stat(dir); err == nil {
		p.Detected = true
		p.Dir = dir
		return p
	}
	// Also check VS Code with copilot extension marker
	if _, err := exec.LookPath("code"); err == nil {
		if _, err2 := os.Stat(dir); err2 == nil {
			p.Detected = true
			p.Dir = dir
		}
	}
	return p
}

func detectQwenCode(home string) Platform {
	dir := filepath.Join(home, ".qwen")
	p := Platform{
		ID:   PlatformQwenCode,
		Name: "Qwen Code",
		Dir:  dir,
	}
	if _, err := os.Stat(dir); err == nil {
		p.Detected = true
		p.Dir = dir
	}
	return p
}

func detectCodex(home string) Platform {
	dir := filepath.Join(home, ".codex")
	p := Platform{
		ID:   PlatformCodex,
		Name: "OpenAI Codex",
		Dir:  dir,
	}
	if _, err := os.Stat(dir); err == nil {
		p.Detected = true
		p.Dir = dir
		return p
	}
	// Fallback: check if codex binary is available
	if _, err := exec.LookPath("codex"); err == nil {
		p.Detected = true
		p.Dir = dir
	}
	return p
}

func detectOpenCode(home string) Platform {
	configFile := filepath.Join(home, ".config", "opencode", "opencode.json")
	dir := filepath.Join(home, ".config", "opencode")
	p := Platform{
		ID:   PlatformOpenCode,
		Name: "opencode",
		Dir:  dir,
	}
	if _, err := os.Stat(configFile); err == nil {
		p.Detected = true
		p.Dir = dir
	}
	return p
}
