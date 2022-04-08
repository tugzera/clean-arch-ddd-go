package entities

type Item struct {
	ID          string
	Name        string
	Description string
	Price       float64
}

func NewItem(id string, name string, price float64, description string) *Item {
	return &Item{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
	}
}
