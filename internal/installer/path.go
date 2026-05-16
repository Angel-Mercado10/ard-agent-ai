package installer

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const basePathEnv = "ARD_AGENT_AI_BASE_PATH"

// DefaultBasePath returns a sensible, user-owned default instead of forcing
// C:\Projects\. Users can override it with ARD_AGENT_AI_BASE_PATH.
func DefaultBasePath() string {
	if envPath := strings.TrimSpace(os.Getenv(basePathEnv)); envPath != "" {
		if resolved, err := ResolveBasePath(envPath); err == nil {
			return resolved
		}
		return envPath
	}

	home, err := os.UserHomeDir()
	if err != nil || home == "" {
		return placeholder
	}

	if runtime.GOOS == "windows" {
		return filepath.Join(home, "Documents", "ARD")
	}
	return filepath.Join(home, "Documents", "ARD")
}

// ResolveBasePath normalizes user input so the saved route is explicit and
// repeatable across installer targets.
func ResolveBasePath(input string) (string, error) {
	path := strings.TrimSpace(input)
	if path == "" {
		path = DefaultBasePath()
	}

	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		if path == "~" {
			path = home
		} else if strings.HasPrefix(path, "~/") || strings.HasPrefix(path, `~\`) {
			path = filepath.Join(home, path[2:])
		}
	}

	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return filepath.Clean(abs), nil
}
