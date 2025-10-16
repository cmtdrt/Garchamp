package itemallergenrelationdb

type Model struct {
	ItemID     int
	AllergenID int
}

func NewModel(itemID, allergenID int) *Model {
	return &Model{ItemID: itemID, AllergenID: allergenID}
}
