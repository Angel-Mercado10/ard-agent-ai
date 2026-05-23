package installer

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestUpsertCodexInstructionsBlock(t *testing.T) {
	newPath := filepath.Join(t.TempDir(), "new-ard")
	block := New(newPath, nil).codexInstructionsBlock(filepath.Join(t.TempDir(), "skills", "ard", "SKILL.md"))

	tests := []struct {
		name     string
		existing string
	}{
		{
			name:     "appends when missing",
			existing: "# Existing instructions\n\nKeep this.",
		},
		{
			name: "replaces managed block",
			existing: "# Existing\n\n" + codexBlockStart + `
## ARD Agent AI

Configured ARD base path: C:\Projects\
` + codexBlockEnd + "\n\n## Other\nKeep this.",
		},
		{
			name: "replaces legacy block",
			existing: "# Existing\n\n## ARD Agent AI\n\nConfigured ARD base path: C:\\Projects\\\n\nSkills location: old\n" +
				"\n## Other\nKeep this.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := upsertCodexInstructionsBlock(tt.existing, block)

			if !strings.Contains(got, "Configured ARD base path: "+newPath) {
				t.Fatalf("updated block missing new path:\n%s", got)
			}
			if strings.Contains(got, "Configured ARD base path: C:\\Projects\\") {
				t.Fatalf("old forced path survived:\n%s", got)
			}
			if !strings.Contains(got, codexBlockStart) || !strings.Contains(got, codexBlockEnd) {
				t.Fatalf("managed markers missing:\n%s", got)
			}
		})
	}
}

func TestInstallCodexUpdatesExistingPath(t *testing.T) {
	codexDir := t.TempDir()
	instructionsPath := filepath.Join(codexDir, "instructions.md")
	if err := os.WriteFile(instructionsPath, []byte("## ARD Agent AI\n\nConfigured ARD base path: C:\\Projects\\\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	newPath := filepath.Join(t.TempDir(), "custom-ard")
	inst := New(newPath, []Platform{{
		ID:   PlatformCodex,
		Name: "OpenAI Codex",
		Dir:  codexDir,
	}})

	if err := inst.CopySkills(); err != nil {
		t.Fatal(err)
	}
	if err := inst.InstallPlatform(inst.platforms[0]); err != nil {
		t.Fatal(err)
	}

	got, err := os.ReadFile(instructionsPath)
	if err != nil {
		t.Fatal(err)
	}

	content := string(got)
	if !strings.Contains(content, "Configured ARD base path: "+newPath) {
		t.Fatalf("instructions were not updated with custom path:\n%s", content)
	}
	if strings.Contains(content, "Configured ARD base path: C:\\Projects\\") {
		t.Fatalf("old default path survived:\n%s", content)
	}
}

func TestCopySkillsCopiesSupportingFiles(t *testing.T) {
	openCodeDir := t.TempDir()
	newPath := filepath.Join(t.TempDir(), "custom-ard")
	inst := New(newPath, []Platform{{
		ID:   PlatformOpenCode,
		Name: "OpenCode",
		Dir:  openCodeDir,
	}})

	if err := inst.CopySkills(); err != nil {
		t.Fatal(err)
	}

	expectedFiles := []string{
		filepath.Join(openCodeDir, "skills", "ard", "SKILL.md"),
		filepath.Join(openCodeDir, "skills", "ard", "assets", "ard-template.md"),
		filepath.Join(openCodeDir, "skills", "ard", "assets", "ard-index-template.md"),
		filepath.Join(openCodeDir, "skills", "ard", "references", "elicitation-question-bank.md"),
		filepath.Join(openCodeDir, "skills", "ard", "references", "known-pattern-conflicts.md"),
	}

	for _, path := range expectedFiles {
		if _, err := os.Stat(path); err != nil {
			t.Fatalf("expected copied skill file %s: %v", path, err)
		}
	}

	skillContent, err := os.ReadFile(filepath.Join(openCodeDir, "skills", "ard", "SKILL.md"))
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(skillContent), "Configured base path: `"+newPath+"`") {
		t.Fatalf("skill content was not patched with configured base path:\n%s", skillContent)
	}
	if strings.Contains(string(skillContent), "Configured base path: `C:\\Projects\\`") {
		t.Fatalf("placeholder base path survived in installed skill:\n%s", skillContent)
	}
}
