package things

// Items are tangible concepts.
// They can be examined and potentially be interacted with further.
type Item struct {
	Thing
	desc string
}

func (item *Item) SetDesc(newDesc string) {
	item.desc = newDesc
}

// Examines an item by returning its description
func (item *Item) Examine() string {
	return item.desc
}

func mapToItem(rawMap map[string]interface{}) *Item {
	item := new(Item)

	if _, ok := rawMap["name"]; ok {
		item.name, _ = rawMap["name"].(string)
	}

	if _, ok := rawMap["desc"]; ok {
		item.desc, _ = rawMap["desc"].(string)
	}

	return item
}
