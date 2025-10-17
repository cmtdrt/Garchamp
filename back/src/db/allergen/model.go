package allergendb

type Model struct {
	ID   int
	Name string
}

func NewModel(name string) *Model {
	return &Model{
		Name: name,
	}
}
