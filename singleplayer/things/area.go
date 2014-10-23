package things

import "strings"

// Areas are anything like a room, which can have items in it to interact with.
type Area struct {
	name  string
	desc  string
	items map[string]Itemlike
}

func (area *Area) Name() string {
	return area.name
}

// SetName updates or sets the name for this area
func (area *Area) SetName(newName string) {
	area.name = newName
}

func (area *Area) Desc() string {
	return area.desc
}

// SetDesc updates or sets the description for this area
func (area *Area) SetDesc(newDesc string) {
	area.desc = newDesc
}

// AddItem adds an item to this area's list of items
// If no map of items exists yet, creates a new map first
// If an item by this name already exists in this area, aborts
// TODO: allow for multiple of the same item in a room or otherwise
// identify things with the same name
func (area *Area) AddItem(item Itemlike) {
	if area.items == nil {
		area.items = make(map[string]Itemlike)
	}

	if area.Has(item.Name()) {
		// TODO: return error or bool
		return
	}

	area.items[item.Name()] = item
}

// Examine() returns the description of the area and a list of items in it
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

// Has tells you whether an item by a given name exists in this area
func (area *Area) Has(itemName string) bool {
	_, ok := area.items[itemName]
	return ok
}

// Item returns an item by name
func (area *Area) Item(itemName string) Itemlike {
	return area.items[itemName]
}

func (area *Area) Populate(rawMap map[string]interface{}) {
	if _, ok := rawMap["name"]; ok {
		area.name, _ = rawMap["name"].(string)
	}

	if _, ok := rawMap["desc"]; ok {
		area.desc, _ = rawMap["desc"].(string)
	}

	if _, ok := rawMap["items"]; ok {
		if rawItems, found := rawMap["items"].(map[string]interface{}); found {
			for _, rawItem := range rawItems {
				if wrappedItemMap, found2 := rawItem.(map[string]interface{}); found2 {
					area.AddItem(MakeItem(wrappedItemMap))
				}
			}
		}
	}
}
