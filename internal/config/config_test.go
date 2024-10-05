package config

import (
	"os"
	"testing"
)

// Test the LoadConfig function with a valid YAML file
func TestLoadConfig_ValidConfig(t *testing.T) {
	// Create a sample YAML configuration
	yamlContent := `
whispers:
  - name: "Resource BC break"
    triggers:
      - check: "filepath"
        contains: "pr-whisper/Http/Resources"
    severity: "caution"
    message: "This change may break the API contract. Please review the API documentation."

  - name: "Controller BC break"
    triggers:
      - check: "filepath"
        contains: "pr-whisper/Http/Controllers"
    severity: "caution"
    message: "This change may break the API contract. Please review the API documentation."
`

	// Write the YAML to a temporary file
	tmpFile, err := os.CreateTemp("", "config*.yaml")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // clean up

	_, err = tmpFile.Write([]byte(yamlContent))
	if err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tmpFile.Close()

	// Create a new config and load the file
	config := NewConfig(tmpFile.Name())
	whisperConfig, err := config.LoadConfig()

	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Verify the content of the loaded configuration
	if len(whisperConfig.Whispers) != 2 {
		t.Fatalf("Expected 2 whispers, but got %d", len(whisperConfig.Whispers))
	}

	if whisperConfig.Whispers[0].Name != "Resource BC break" {
		t.Fatalf("Expected first whisper name to be 'Resource BC break', got %s", whisperConfig.Whispers[0].Name)
	}

	if whisperConfig.Whispers[1].Severity != "caution" {
		t.Fatalf("Expected second whisper severity to be 'caution', got %s", whisperConfig.Whispers[1].Severity)
	}
}

// Test the LoadConfig function with an invalid YAML file
func TestLoadConfig_InvalidYAML(t *testing.T) {
	// Create an invalid YAML content
	invalidYAMLContent := `
whispers:
  - name: "Resource BC break"
    triggers
      - check: "filepath"
`

	// Write the invalid YAML to a temporary file
	tmpfile, err := os.CreateTemp("", "invalid_config*.yaml")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	_, err = tmpfile.Write([]byte(invalidYAMLContent))
	if err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tmpfile.Close()

	// Attempt to load the invalid config
	config := NewConfig(tmpfile.Name())
	_, err = config.LoadConfig()

	if err == nil {
		t.Fatalf("Expected error when loading invalid YAML, but got none")
	}
}
