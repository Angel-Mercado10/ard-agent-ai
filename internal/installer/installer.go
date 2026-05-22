package installer

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/Angel-Mercado10/ard-agent-ai/internal/assets"
)

const placeholder = `C:\Projects\`
const codexInstructionsTitle = "## ARD Agent AI"
const codexBlockStart = "<!-- ard-agent-ai:start -->"
const codexBlockEnd = "<!-- ard-agent-ai:end -->"

// Installer handles copying agent files to platform directories.
type Installer struct {
	basePath  string
	platforms []Platform
}

// New creates an Installer configured with the user's chosen base path and
// the list of platforms to install into.
func New(basePath string, platforms []Platform) *Installer {
	resolvedBasePath, err := ResolveBasePath(basePath)
	if err != nil {
		resolvedBasePath = basePath
	}

	return &Installer{
		basePath:  resolvedBasePath,
		platforms: platforms,
	}
}

// CopySkills writes the shared skill directory to every selected platform.
// This is always step 0 during installation.
func (inst *Installer) CopySkills() error {
	for _, p := range inst.platforms {
		var skillsDir string
		switch p.ID {
		case PlatformClaudeCode:
			skillsDir = filepath.Join(p.Dir, "skills", "ard")
		case PlatformGitHubCopilot:
			skillsDir = filepath.Join(p.Dir, "skills", "ard")
		case PlatformQwenCode:
			skillsDir = filepath.Join(p.Dir, "skills", "ard")
		case PlatformOpenCode:
			skillsDir = filepath.Join(p.Dir, "skills", "ard")
		case PlatformCodex:
			skillsDir = filepath.Join(p.Dir, "skills", "ard")
		}

		if skillsDir == "" {
			continue
		}
		if err := inst.copySkillDir(skillsDir); err != nil {
			return fmt.Errorf("copying skill dir for %s: %w", p.Name, err)
		}
	}
	return nil
}

func (inst *Installer) copySkillDir(destRoot string) error {
	const sourceRoot = "agents/skills/ard"

	return fs.WalkDir(assets.AgentFS, sourceRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath := strings.TrimPrefix(path, sourceRoot)
		relPath = strings.TrimPrefix(relPath, "/")
		destPath := destRoot
		if relPath != "" {
			destPath = filepath.Join(destRoot, filepath.FromSlash(relPath))
		}

		if d.IsDir() {
			return os.MkdirAll(destPath, 0o755)
		}

		data, err := assets.AgentFS.ReadFile(path)
		if err != nil {
			return fmt.Errorf("reading %s: %w", path, err)
		}

		return writeFile(destPath, inst.patchContent(data))
	})
}

// InstallPlatform installs all agent files for a single platform.
func (inst *Installer) InstallPlatform(p Platform) error {
	switch p.ID {
	case PlatformClaudeCode:
		return inst.installClaude(p)
	case PlatformGitHubCopilot:
		return inst.installCopilot(p)
	case PlatformQwenCode:
		return inst.installQwen(p)
	case PlatformOpenCode:
		return inst.installOpenCode(p)
	case PlatformCodex:
		return inst.installCodex(p)
	default:
		return fmt.Errorf("unknown platform: %s", p.ID)
	}
}

// installClaude installs the ARD agent into Claude Code (~/.claude).
func (inst *Installer) installClaude(p Platform) error {
	// Agent
	if err := inst.copyAsset(
		"agents/agents-claude/ard.md",
		filepath.Join(p.Dir, "agents", "ard.md"),
	); err != nil {
		return err
	}

	// Prompts
	if err := inst.copyAsset(
		"agents/prompts-claude/ard.md",
		filepath.Join(p.Dir, "prompts", "ard", "ard.md"),
	); err != nil {
		return err
	}

	// Skill registry (generated, not copied)
	if err := inst.writeClaudeSkillRegistry(p.Dir); err != nil {
		return err
	}

	return nil
}

// installCopilot installs into GitHub Copilot (~/.copilot).
func (inst *Installer) installCopilot(p Platform) error {
	if err := inst.copyAsset(
		"agents/agents-copilot/ard.agent.md",
		filepath.Join(p.Dir, "agents", "ard.agent.md"),
	); err != nil {
		return err
	}

	if err := inst.copyAsset(
		"agents/prompts-copilot/ard.md",
		filepath.Join(p.Dir, "prompts", "ard", "ard.md"),
	); err != nil {
		return err
	}

	return nil
}

// installQwen installs into Qwen Code (~/.qwen).
func (inst *Installer) installQwen(p Platform) error {
	if err := inst.copyAsset(
		"agents/agents-qwen/ard.md",
		filepath.Join(p.Dir, "agents", "ard.md"),
	); err != nil {
		return err
	}

	if err := inst.copyAsset(
		"agents/prompts-qwen/ard.md",
		filepath.Join(p.Dir, "prompts", "ard", "ard.md"),
	); err != nil {
		return err
	}

	return nil
}

// installOpenCode installs into opencode (~/.config/opencode).
func (inst *Installer) installOpenCode(p Platform) error {
	// Prompts
	if err := inst.copyAsset(
		"agents/prompts/ard.md",
		filepath.Join(p.Dir, "prompts", "ard", "ard.md"),
	); err != nil {
		return err
	}

	// Commands
	for _, cmd := range []string{"ard-init.md", "ard-update.md"} {
		if err := inst.copyAsset(
			"agents/commands/"+cmd,
			filepath.Join(p.Dir, "commands", cmd),
		); err != nil {
			return err
		}
	}

	// Skill registry
	if err := inst.writeOpenCodeSkillRegistry(p.Dir); err != nil {
		return err
	}

	// Patch opencode.json
	if err := inst.patchOpenCodeJSON(p.Dir); err != nil {
		return err
	}

	return nil
}

// installCodex installs the ARD agent into OpenAI Codex (~/.codex).
func (inst *Installer) installCodex(p Platform) error {
	// Agent file
	if err := inst.copyAsset(
		"agents/agents-codex/ard.md",
		filepath.Join(p.Dir, "agents", "ard.md"),
	); err != nil {
		return err
	}

	// Append instructions to ~/.codex/instructions.md
	if err := inst.appendCodexInstructions(p.Dir); err != nil {
		return err
	}

	return nil
}

// appendCodexInstructions writes the ARD instructions block to
// ~/.codex/instructions.md. It is idempotent but updateable: if a previous
// managed block exists, it is replaced so changing the base path actually
// updates Codex instead of keeping the old route forever.
func (inst *Installer) appendCodexInstructions(codexDir string) error {
	instructionsPath := filepath.Join(codexDir, "instructions.md")

	// Read existing content if the file exists
	existing, err := os.ReadFile(instructionsPath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("reading instructions.md: %w", err)
	}

	skillsPath := filepath.Join(codexDir, "skills", "ard", "SKILL.md")
	block := inst.codexInstructionsBlock(skillsPath)
	updated := upsertCodexInstructionsBlock(string(existing), block)

	return writeFile(instructionsPath, []byte(updated))
}

func (inst *Installer) codexInstructionsBlock(skillsPath string) string {
	return fmt.Sprintf(`
---

%s
## ARD Agent AI

You are the ARD runtime adapter. When the user asks to create, update, review, or record
architecture decisions in an Architecture Record Document, read the canonical `+"`"+`ard`+"`"+` skill
and follow it as the source of truth.

Configured ARD base path: %s

Skills location: %s
Use supporting files from the same skill directory when the skill refers to `+"`"+`assets/`+"`"+` or `+"`"+`references/`+"`"+`.
%s
`, codexBlockStart, inst.basePath, skillsPath, codexBlockEnd)
}

func upsertCodexInstructionsBlock(existing, block string) string {
	if existing == "" {
		return strings.TrimLeft(block, "\r\n")
	}

	start := strings.Index(existing, codexBlockStart)
	end := strings.Index(existing, codexBlockEnd)
	if start >= 0 && end >= start {
		end += len(codexBlockEnd)
		return strings.TrimRight(existing[:start], "\r\n") + "\n\n" + strings.TrimSpace(block) + strings.TrimLeft(existing[end:], "\r\n")
	}

	legacyStart := strings.Index(existing, codexInstructionsTitle)
	if legacyStart >= 0 {
		nextSection := strings.Index(existing[legacyStart+len(codexInstructionsTitle):], "\n## ")
		if nextSection >= 0 {
			legacyEnd := legacyStart + len(codexInstructionsTitle) + nextSection
			return strings.TrimRight(existing[:legacyStart], "\r\n") + strings.TrimSpace(block) + "\n" + strings.TrimLeft(existing[legacyEnd:], "\r\n")
		}
		return strings.TrimRight(existing[:legacyStart], "\r\n") + strings.TrimSpace(block) + "\n"
	}

	return strings.TrimRight(existing, "\r\n") + "\n\n" + strings.TrimSpace(block) + "\n"
}

// copyAsset reads a file from the embedded FS, patches the placeholder, and
// writes it to dest (creating intermediate directories as needed).
func (inst *Installer) copyAsset(assetPath, dest string) error {
	data, err := assets.AgentFS.ReadFile(assetPath)
	if err != nil {
		return fmt.Errorf("reading asset %s: %w", assetPath, err)
	}
	data = inst.patchContent(data)

	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return fmt.Errorf("creating directory for %s: %w", dest, err)
	}
	return writeFile(dest, data)
}

// patchContent replaces the placeholder base path with the user's chosen path.
func (inst *Installer) patchContent(data []byte) []byte {
	patched := strings.ReplaceAll(string(data), placeholder, inst.basePath)
	return []byte(patched)
}

// writeClaudeSkillRegistry generates ~/.claude/.atl/skill-registry.md.
func (inst *Installer) writeClaudeSkillRegistry(claudeDir string) error {
	skillsDir := filepath.Join(claudeDir, "skills", "ard")
	atlDir := filepath.Join(claudeDir, ".atl")
	if err := os.MkdirAll(atlDir, 0o755); err != nil {
		return fmt.Errorf("creating .atl dir: %w", err)
	}

	content := generateSkillRegistry(inst.basePath, skillsDir)
	return writeFile(filepath.Join(atlDir, "skill-registry.md"), []byte(content))
}

// writeOpenCodeSkillRegistry generates ~/.config/opencode/.atl/skill-registry.md.
func (inst *Installer) writeOpenCodeSkillRegistry(openCodeDir string) error {
	skillsDir := filepath.Join(openCodeDir, "skills", "ard")
	atlDir := filepath.Join(openCodeDir, ".atl")
	if err := os.MkdirAll(atlDir, 0o755); err != nil {
		return fmt.Errorf("creating .atl dir: %w", err)
	}

	content := generateSkillRegistry(inst.basePath, skillsDir)
	return writeFile(filepath.Join(atlDir, "skill-registry.md"), []byte(content))
}

// generateSkillRegistry produces the skill-registry.md content.
func generateSkillRegistry(basePath, skillsDir string) string {
	return fmt.Sprintf(`# Skill Registry — ard-agent-ai

**Auto-generated by ard-agent-ai installer v1.0.0**
**Base path:** %s

## User Skills

| Trigger | Skill | Path |
|---------|-------|------|
| ARD, Architecture Record Document, architecture decision register, /ard-init, /ard-update | ard | %s/SKILL.md |

## Compact Rules

### ard
- Base path for all ARDs: `+"`"+`%s`+"`"+`
- Commands: `+"`"+`ard-init <project>`+"`"+` (full socratic init) · `+"`"+`ard-update <project>`+"`"+` (incremental evolution)
- Document WHY for every decision — never just WHAT
- Section 19 (Sprint Map) is generated LAST after sections 1–18
- Log every `+"`"+`ard-update`+"`"+` change in Decision Log (section 20) with date and reason
- Use explicit challenge blocks only for contradictions, undeclared assumptions, missing information, or known pattern conflicts
- Supporting templates and extended guidance live next to the installed skill in `+"`"+`assets/`+"`"+` and `+"`"+`references/`+"`"+`
`, basePath, skillsDir, basePath)
}

// patchOpenCodeJSON reads opencode.json, injects the ard agent config, and
// writes it back. If the file has an unexpected structure it is left unchanged.
func (inst *Installer) patchOpenCodeJSON(openCodeDir string) error {
	jsonPath := filepath.Join(openCodeDir, "opencode.json")
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		// Config file missing — nothing to patch
		return nil
	}

	var config map[string]interface{}
	if err := json.Unmarshal(data, &config); err != nil {
		// Can't parse — skip silently to avoid corrupting the file
		return nil
	}

	promptsDir := filepath.Join(openCodeDir, "prompts", "ard")
	ardConfig := map[string]interface{}{
		"description": "Architecture Record Document agent — define and evolve software architecture decisions",
		"mode":        "primary",
		"prompt":      fmt.Sprintf("{file:%s/ard.md}", promptsDir),
		"tools":       map[string]interface{}{"read": true, "write": true, "edit": true},
	}

	// Ensure config.agent exists
	agentSection, ok := config["agent"].(map[string]interface{})
	if !ok {
		agentSection = make(map[string]interface{})
	}
	agentSection["ard"] = ardConfig
	config["agent"] = agentSection

	updated, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("marshalling opencode.json: %w", err)
	}
	return writeFile(jsonPath, updated)
}

// writeFile writes data to path, creating intermediate directories.
func writeFile(path string, data []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}
