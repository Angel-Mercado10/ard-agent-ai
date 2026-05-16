package assets

import "embed"

// AgentFS holds all embedded agent markdown files.
// The agents/ directory is located alongside this file at
// internal/assets/agents/ — a valid local path for //go:embed.
//
//go:embed all:agents
var AgentFS embed.FS
