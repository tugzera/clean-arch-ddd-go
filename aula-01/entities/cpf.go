package entities

type Cpf struct {
	Value string
}

func NewCpf(value string) *Cpf {
	return &Cpf{
		Value: value,
	}
}


// func (cpf *Cpf) validate(value string) {

// }

