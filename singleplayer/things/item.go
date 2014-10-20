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
