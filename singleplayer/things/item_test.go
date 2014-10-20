package things

import (
	"testing"
)

func TestItemExamine(t *testing.T) {
	item := new(Item)
	item.desc = "This item sure looks like a thing."

	if wantedOut, out := item.desc, item.Examine(); out != wantedOut {
		t.Errorf("Examine() want %s, got %s", wantedOut, out)
	}
}
