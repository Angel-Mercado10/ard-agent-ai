package installer

import (
	"os"
	"os/exec"
	"path/filepath"
)

// Platform IDs.
const (
	PlatformClaudeCode    = "claude-code"
	PlatformGitHubCopilot = "github-copilot"
	PlatformQwenCode      = "qwen-code"
	PlatformOpenCode      = "opencode"
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

	platforms := []Platform{
		detectClaudeCode(home),
		detectGitHubCopilot(home),
		detectQwenCode(home),
		detectOpenCode(home),
	}

	return platforms
}

func detectClaudeCode(home string) Platform {
	dir := filepath.Join(home, ".claude")
	p := Platform{
		ID:   PlatformClaudeCode,
		Name: "Claude Code",
		Dir:  "~/.claude",
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
		Dir:  "~/.copilot",
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
		Dir:  "~/.qwen",
	}
	if _, err := os.Stat(dir); err == nil {
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
		Dir:  "~/.config/opencode",
	}
	if _, err := os.Stat(configFile); err == nil {
		p.Detected = true
		p.Dir = dir
	}
	return p
}
