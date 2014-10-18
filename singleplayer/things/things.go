package things

import (
	"strings"
)

type Examinable interface {
	Examine() string
}

type Thing struct {
	name string
}

// Items are tangible concepts.
// They can be examined and potentially be interacted with further.
type Item struct {
	Thing
	desc string
}

// Examines an item by returning its description
func (item Item) Examine() string {
	return item.desc
}

// Areas are anything like a room, which can have items in it to interact with.
type Area struct {
	Thing
	desc  string
	items map[string]*Item
}

// Examines an area by displaying the description and a list of item names.
func (area Area) Examine() (output string) {
	output = area.desc

	itemsInArea := len(area.items)

	if itemsInArea == 0 {
		return
	}

	// TODO: more efficient concatenating
	// http://stackoverflow.com/a/1766304/1019228
	output += " I can see these things: "
	i := 0
	for name, _ := range area.items {
		output += name

		if itemsInArea > 1 && i == itemsInArea-2 {
			output += " and "
		}

		if itemsInArea > 1 && i < itemsInArea-2 {
			output += ", "
		}

		i++
	}

	output = strings.Trim(output, " ")
	return
}
