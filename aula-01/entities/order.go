package entities

type Order struct {
	Cpf *Cpf
	OrderItems []*OrderItem
	Total float64
}

func NewOrder(cpfValue string) (*Order, error) {
	cpf, err := NewCpf(cpfValue)
	if err != nil {
		return nil, err
	}
	order := Order{
		Cpf: cpf,
	}
	return &order, nil
}
