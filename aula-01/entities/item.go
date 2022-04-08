package entities

type Item struct {
	ID    string
	Name  string
	Price float64
}

func NewItem(id string, name string, price float64) *Item {
	return &Item{
		ID: id,
		Name: name,
		Price: price,
	}
}
