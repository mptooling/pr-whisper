package business

import (
	"github.com/mptooling/pr-whisper/internal/domain"
	"testing"
)

func TestMakeGenericWhispers_ValidConfig(t *testing.T) {
	config := &domain.WhisperConfig{
		Whispers: []domain.WhisperConfigItem{
			{
				Name: "Resource BC break",
				Triggers: []domain.Trigger{
					{
						Check:    "filepath",
						Contains: "pr-whisper/Http/Resources",
					},
				},
				Severity: "caution",
				Message:  "This change may break the API contract. Please review the API documentation.",
			},
		},
	}

	factory := NewGenericWhispererBuilder()
	whispers := factory.MakeGenericWhispers(config)

	if len(whispers) != 1 {
		t.Fatalf("Expected 1 whisper, got %d", len(whispers))
	}

	if whispers[0].Name != "Resource BC break" {
		t.Fatalf("Expected whisper name to be 'Resource BC break', got %s", whispers[0].Name)
	}

	if whispers[0].Severity != domain.Caution {
		t.Fatalf("Expected whisper severity to be 'Caution', got %d", whispers[0].Severity)
	}

	if whispers[0].Message != "This change may break the API contract. Please review the API documentation." {
		t.Fatalf("Expected whisper message to be 'This change may break the API contract. Please review the API documentation.', got %s", whispers[0].Message)
	}
}
