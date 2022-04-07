package entities

import (
	"fmt"
	"regexp"
)

type Cpf struct {
	Value string
}

func NewCpf(value string) (*Cpf, error) {
	cpf := Cpf{}
	cleanCpf, err := cpf.cleanCpf(value)
	if err != nil {
		return nil, err
	}

	err = cpf.validate(cleanCpf)
	if err != nil {
		return nil, err
	}
	cpf.Value = cleanCpf
	return &cpf, nil
}

func (cpf *Cpf) validate(value string) error {
	size := len([]rune(value))
	if size != 11 {
		return fmt.Errorf("invalid cpf size")
	}
	err := cpf.checkIfDigitsAreAllEqual(value)
	if err != nil {
		return err
	}
	return nil
}

func (cpf *Cpf) checkIfDigitsAreAllEqual(value string) error {
	stringArray := []rune(value)
	count := 0
	for i := 0; i < len(stringArray)-1; i++ {
		if stringArray[i] == stringArray[i+1] {
			count++
		}
	}
	fmt.Println(count == len(stringArray)-1)
	if count == len(stringArray)-1 {
		return fmt.Errorf("invalid cpf")
	}
	return nil
}

func (cpf *Cpf) cleanCpf(value string) (string, error) {
	regex, err := regexp.Compile(`[^\w]`)
	if err != nil {
		return "", err
	}
	return regex.ReplaceAllString(value, ""), nil
}
