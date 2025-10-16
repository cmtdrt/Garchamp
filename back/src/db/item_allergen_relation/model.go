package itemallergenrelationdb

type model struct {
	ItemID     int
	AllergenID int
}

func NewModel(itemID, allergenID int) *model {
	return &model{ItemID: itemID, AllergenID: allergenID}
}
