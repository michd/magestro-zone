package things

import "encoding/json"
import "strings"

// Areas are anything like a room, which can have items in it to interact with.
type Area struct {
	Thing
	desc  string
	items map[string]*Item
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
func (area *Area) Item(itemName string) *Item {
	return area.items[itemName]
}

// mapToArea maps untyped values to its correct fields
// Used to load JSON into an actual area object
func mapToArea(rawMap map[string]interface{}) *Area {
	area := new(Area)

	if _, ok := rawMap["name"]; ok {
		area.name, _ = rawMap["name"].(string)
	}

	if _, ok := rawMap["desc"]; ok {
		area.desc, _ = rawMap["desc"].(string)
	}

	if _, ok := rawMap["items"]; ok {
		if rawItems, found := rawMap["items"].(map[string]interface{}); found {
			for _, rawItem := range rawItems {
				if rawItemMap, found2 := rawItem.(map[string]interface{}); found2 {
					area.AddItem(mapToItem(rawItemMap))
					// needs moar nested control structures
				}
			}
		}
	}

	return area
}

// ParseJson parses json into a fully qualified area object
// Actual mapping is done by mapToArea
func ParseJson(jsonBlob []byte) *Area {
	rawMap := map[string]interface{}{}

	err := json.Unmarshal(jsonBlob, &rawMap)

	if err != nil {
		return nil
	}

	return mapToArea(rawMap)
}
