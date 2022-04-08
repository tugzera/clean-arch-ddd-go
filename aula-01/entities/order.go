package entities

type Order struct {
	Cpf        *Cpf
	OrderItems []*OrderItem
	Total      float64
	FinalAmount float64
	DiscountValue float64
}

func NewOrder(cpfValue string, orderItems []*OrderItem, cupom *Cupom) (*Order, error) {
	cpf, err := NewCpf(cpfValue)
	if err != nil {
		return nil, err
	}
	var total float64
	for _, value := range orderItems {
		total += value.Total
	}
	discount := (float64(cupom.DiscountPercentage) * total) / 100 
	order := Order{
		Cpf:        cpf,
		OrderItems: orderItems,
		Total:      total,
		DiscountValue: discount,
		FinalAmount: total - discount,
	}
	return &order, nil
}
