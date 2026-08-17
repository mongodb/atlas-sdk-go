package detectaiagent

import (
	"os"
	"path/filepath"
	"testing"
)

func setAgentDefinitions(t *testing.T, definitions []agentDefinition) {
	t.Helper()

	original := agentDefinitions
	agentDefinitions = definitions

	t.Cleanup(func() {
		agentDefinitions = original
	})
}

func TestDetectEnvironmentMarkers(t *testing.T) {
	for _, tc := range []struct {
		name          string
		markerValue   string
		envValue      string
		expectedAgent DetectedAgent
		expectedOK    bool
	}{
		{
			name:        "matches any non-empty value",
			markerValue: "",
			envValue:    "value",
			expectedAgent: DetectedAgent{
				ID:                  "test_agent",
				UserAgentIdentifier: "ai-agent/test_agent",
			},
			expectedOK: true,
		},
		{
			name:        "rejects empty value",
			markerValue: "",
			envValue:    "",
			expectedOK:  false,
		},
		{
			name:        "matches exact value",
			markerValue: "value",
			envValue:    "value",
			expectedAgent: DetectedAgent{
				ID:                  "test_agent",
				UserAgentIdentifier: "ai-agent/test_agent",
			},
			expectedOK: true,
		},
		{
			name:        "rejects different value",
			markerValue: "value",
			envValue:    "othervalue",
			expectedOK:  false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			setAgentDefinitions(t, []agentDefinition{
				{
					id: "test_agent",
					markers: []marker{
						{envVar: "AGENT_MARKER", value: tc.markerValue},
					},
				},
			})
			t.Setenv("AGENT_MARKER", tc.envValue)

			got, ok := Detect()

			if ok != tc.expectedOK {
				t.Errorf("Detect() ok = %t, expected %t", ok, tc.expectedOK)
			}
			if got != tc.expectedAgent {
				t.Errorf("Detect() agent = %v, expected %v", got, tc.expectedAgent)
			}
		})
	}
}

func TestDetectFileMarker(t *testing.T) {
	tempDir := t.TempDir()
	agentMarkerFile := filepath.Join(tempDir, "agent-marker")
	if err := os.WriteFile(agentMarkerFile, nil, 0o600); err != nil {
		t.Fatalf("failed to create test agent marker file: %v", err)
	}

	for _, tc := range []struct {
		name          string
		filePath      string
		expectedAgent DetectedAgent
		expectedOK    bool
	}{
		{
			name:     "matches existing file",
			filePath: agentMarkerFile,
			expectedAgent: DetectedAgent{
				ID:                  "test_agent",
				UserAgentIdentifier: "ai-agent/test_agent",
			},
			expectedOK: true,
		},
		{
			name:       "rejects missing file",
			filePath:   filepath.Join(tempDir, "missing"),
			expectedOK: false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			setAgentDefinitions(t, []agentDefinition{
				{
					id: "test_agent",
					markers: []marker{
						{filePath: tc.filePath},
					},
				},
			})

			got, ok := Detect()

			if ok != tc.expectedOK {
				t.Errorf("Detect() ok = %t, expected %t", ok, tc.expectedOK)
			}
			if got != tc.expectedAgent {
				t.Errorf("Detect() agent = %v, expected %v", got, tc.expectedAgent)
			}
		})
	}
}

func TestDetectMultipleMarkers(t *testing.T) {
	setAgentDefinitions(t, []agentDefinition{
		{
			id: "first_agent",
			markers: []marker{
				{envVar: "FIRST_AGENT_MARKER"},
				{envVar: "FIRST_AGENT_SECOND_MARKER"},
			},
		},
		{
			id: "second_agent",
			markers: []marker{
				{envVar: "SECOND_AGENT_MARKER"},
			},
		},
	})

	for _, tc := range []struct {
		name          string
		envVars       map[string]string
		expectedAgent DetectedAgent
		expectedOK    bool
	}{
		{
			name: "matches second marker of agent",
			envVars: map[string]string{
				"FIRST_AGENT_MARKER":        "",
				"FIRST_AGENT_SECOND_MARKER": "value",
			},
			expectedAgent: DetectedAgent{
				ID:                  "first_agent",
				UserAgentIdentifier: "ai-agent/first_agent",
			},
			expectedOK: true,
		},
		{
			name: "returns first matching agent",
			envVars: map[string]string{
				"FIRST_AGENT_MARKER":  "value",
				"SECOND_AGENT_MARKER": "value",
			},
			expectedAgent: DetectedAgent{
				ID:                  "first_agent",
				UserAgentIdentifier: "ai-agent/first_agent",
			},
			expectedOK: true,
		},
		{
			name: "skips agent with no matching markers",
			envVars: map[string]string{
				"SECOND_AGENT_MARKER": "value",
			},
			expectedAgent: DetectedAgent{
				ID:                  "second_agent",
				UserAgentIdentifier: "ai-agent/second_agent",
			},
			expectedOK: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for envVar, value := range tc.envVars {
				t.Setenv(envVar, value)
			}

			got, ok := Detect()

			if ok != tc.expectedOK {
				t.Errorf("Detect() ok = %t, expected %t", ok, tc.expectedOK)
			}
			if got != tc.expectedAgent {
				t.Errorf("Detect() agent = %v, expected %v", got, tc.expectedAgent)
			}
		})
	}
}
