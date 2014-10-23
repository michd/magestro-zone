package things

// Items are tangible concepts.
// They can be examined and potentially be interacted with further.
type Item struct {
	name string
	desc string
}

type Itemlike interface {
	Name() string
	Desc() string
	SetName(newName string)
	SetDesc(newDesc string)
	Examine() string
	Populate(rawMap map[string]interface{})
}

type Populatable interface {
	Populate(rawMap map[string]interface{})
}

type Examinable interface {
	Examine() string
}

func (item *Item) Name() string {
	return item.name
}

func (item *Item) SetName(newName string) {
	item.name = newName
}

func (item *Item) Desc() string {
	return item.desc
}

func (item *Item) SetDesc(newDesc string) {
	item.desc = newDesc
}

// Examines an item by returning its description
func (item *Item) Examine() string {
	return item.desc
}

// Populates the (blank) item with data from a raw map
// Used for loading data into structures from JSON
func (item *Item) Populate(rawMap map[string]interface{}) {
	if _, ok := rawMap["name"]; ok {
		item.name, _ = rawMap["name"].(string)
	}

	if _, ok := rawMap["desc"]; ok {
		item.desc, _ = rawMap["desc"].(string)
	}
}

// Deprecated
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
