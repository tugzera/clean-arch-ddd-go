package entities

type Order struct {
	Cpf        *Cpf
	OrderItems []*OrderItem
	Total      float64
}

func NewOrder(cpfValue string, orderItems []*OrderItem) (*Order, error) {
	cpf, err := NewCpf(cpfValue)
	if err != nil {
		return nil, err
	}
	var total float64
	for _, value := range orderItems {
		total += value.Total
	}
	order := Order{
		Cpf:        cpf,
		OrderItems: orderItems,
		Total:      total,
	}
	return &order, nil
}
