package things

import "strings"

// Areas are anything like a room, which can have items in it to interact with.
type Area struct {
	Thing
	desc  string
	items map[string]*Item
}

func (area *Area) SetDesc(newDesc string) {
	area.desc = newDesc
}

func (area *Area) AddItem(item *Item) {
	if area.items == nil {
		area.items = make(map[string]*Item)
	}

	if area.Has(item.name) {
		// TODO: return error or bool
		return
	}

	area.items[item.name] = item
}

// Examines an area by displaying the description and a list of item names.
func (area *Area) Examine() (output string) {
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

// Has an item identified by itemName in it
func (area *Area) Has(itemName string) bool {
	_, ok := area.items[itemName]
	return ok
}

func (area *Area) Item(itemName string) *Item {
	return area.items[itemName]
}
