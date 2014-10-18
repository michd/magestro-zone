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

func TestAreaExamine(t *testing.T) {
	area1 := new(Area)
	area1.desc = "Looks like a room to me."

	if wantedOut, out := area1.desc, area1.Examine(); out != wantedOut {
		t.Errorf("Examine() want %s, got %s", wantedOut, out)
	}

	area2 := new(Area)
	area2.desc = "Looks like a generic bathroom."

	toilet := Item{Thing{"toilet"}, "I can pee in this."}
	sink := Item{Thing{"sink"}, "I can wash my hands in there after peeing."}
	items := map[string]*Item{
		"toilet": &toilet,
		"sink":   &sink,
	}
	area2.items = items

	wantedOut := area2.desc + " I can see these things: " +
		area2.items["toilet"].name + " and " + area2.items["sink"].name

	if out := area2.Examine(); out != wantedOut {
		t.Errorf("Examine() want \"%s\", got \"%s\"", wantedOut, out)
	}
}
