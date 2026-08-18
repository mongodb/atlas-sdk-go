// Package detectaiagent identifies AI agents that invoke SDK consumers.
// Detection is based on known environment and filesystem markers used by AI agents.
package detectaiagent

import "os"

const userAgentPrefix = "ai-agent/"

// DetectedAgent describes an AI agent found by Detect.
type DetectedAgent struct {
	// ID is the agent identifier.
	ID string
	// UserAgentIdentifier is the agent identifier formatted for a User-Agent header (ai-agent/<ID>).
	UserAgentIdentifier string
}

type agentDefinition struct {
	id      string
	markers []marker
}

type marker struct {
	envVar   string
	value    string
	filePath string
}

var agentDefinitions = []agentDefinition{
	{
		id: "cursor",
		markers: []marker{
			{envVar: "CURSOR_AGENT"},
			{envVar: "CURSOR_EXTENSION_HOST_ROLE", value: "agent-exec"},
		},
	},
	{
		id: "kimi",
		markers: []marker{
			{envVar: "KIMI_PLUGIN_ROOT"},
		},
	},
	{
		id: "gemini_cli",
		markers: []marker{
			{envVar: "GEMINI_CLI"},
		},
	},
	{
		id: "cline",
		markers: []marker{
			{envVar: "CLINE_ACTIVE"},
		},
	},
	{
		id: "codex_cli",
		markers: []marker{
			{envVar: "CODEX_SANDBOX"},
			{envVar: "CODEX_CI"},
			{envVar: "CODEX_THREAD_ID"},
		},
	},
	{
		id: "antigravity",
		markers: []marker{
			{envVar: "ANTIGRAVITY_AGENT"},
			{envVar: "ANTIGRAVITY_CLI_ALIAS"},
		},
	},
	{
		id: "auggie_cli",
		markers: []marker{
			{envVar: "AUGMENT_AGENT"},
		},
	},
	{
		id: "opencode_client",
		markers: []marker{
			{envVar: "OPENCODE_CLIENT"},
			{envVar: "OPENCODE"},
		},
	},
	{
		id: "goose",
		markers: []marker{
			{envVar: "AGENT", value: "goose"},
		},
	},
	{
		id: "amp",
		markers: []marker{
			{envVar: "AGENT", value: "amp"},
		},
	},
	{
		id: "pi",
		markers: []marker{
			{envVar: "AI_AGENT", value: "pi"},
		},
	},
	{
		id: "claude_code",
		markers: []marker{
			{envVar: "CLAUDECODE"},
			{envVar: "CLAUDE_CODE"},
		},
	},
	{
		id: "copilot",
		markers: []marker{
			{envVar: "COPILOT_AGENT_SESSION_ID"},
		},
	},
	{
		id: "devin",
		markers: []marker{
			{filePath: "/opt/.devin"},
		},
	},
	{
		id: "trae_ai",
		markers: []marker{
			{envVar: "TRAE_AI_SHELL_ID"},
		},
	},
}

// Detect returns the first AI agent detected from known markers.
func Detect() (DetectedAgent, bool) {
	for _, agent := range agentDefinitions {
		for _, marker := range agent.markers {
			found := false

			if marker.envVar != "" {
				value := os.Getenv(marker.envVar)
				found = value != "" && (marker.value == "" || value == marker.value)
			} else {
				_, err := os.Stat(marker.filePath)
				found = err == nil
			}

			if found {
				return DetectedAgent{
					ID:                  agent.id,
					UserAgentIdentifier: userAgentPrefix + agent.id,
				}, true
			}
		}
	}

	return DetectedAgent{}, false
}
