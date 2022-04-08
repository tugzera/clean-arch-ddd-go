package entities

type OrderItem struct {
	Item *Item
	Count int
	Total float64
}

func NewOrderItem(item *Item, count int) *OrderItem {
	return &OrderItem{
		Item: item,
		Count: count,
		Total: item.Price * float64(count),
	}
}