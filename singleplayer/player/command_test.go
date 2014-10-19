package player

import (
	"testing"
)

func TestParseCommand(t *testing.T) {
	rawInput := "take thing"
	command := ParseCommand(rawInput)

	if wantedOut, out := "take", command.verb; out != wantedOut {
		t.Errorf("ParseCommand() want verb %s, got %s", wantedOut, out)
	}

	if wantedOut, out := "thing", command.subject; out != wantedOut {
		t.Errorf("ParseCommand() want subject %s, got %s", wantedOut, out)
	}
}
