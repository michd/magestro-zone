package things

type Examinable interface {
	Examine() string
}

type Thing struct {
	name string
}

func (thing *Thing) Name() string {
	return thing.name
}

func (thing *Thing) SetName(newName string) {
	thing.name = newName
}
