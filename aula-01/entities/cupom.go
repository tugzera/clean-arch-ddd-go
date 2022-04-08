package entities

type Cupom struct {
	DiscountPercentage int
}

func NewCupom(discount int) *Cupom {
	return &Cupom{
		DiscountPercentage: discount,
	}
}
