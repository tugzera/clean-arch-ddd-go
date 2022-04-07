package entities

import (
	"fmt"
	"regexp"
	"strconv"
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
	firstDigit, err := cpf.calculateVerifierDigit(value, 10)
	if err != nil {
		return err
	}
	secondDigit, err := cpf.calculateVerifierDigit(value[:9]+strconv.Itoa(firstDigit), 11)
	fmt.Println(firstDigit, secondDigit)
	if err != nil {
		return err
	}
	err = cpf.verifyLastDigits(firstDigit, secondDigit, value)
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

func (cpf *Cpf) calculateVerifierDigit(value string, factor int) (int, error) {
	stringArray := []rune(value)
	sum := 0
	for i := 0; i <= factor-2; i++ {
		numberValue, _ := strconv.Atoi(string(stringArray[i]))
		sum += numberValue * (factor - i)
	}
	rest := sum % 11
	if rest < 2 {
		return 0, nil
	}
	return 11 - rest, nil
}

func (cpf *Cpf) verifyLastDigits(firstDigit int, lastDigit int, value string) error {
	lastDigits := value[len(value)-2:]
	checkString := strconv.Itoa(firstDigit) + strconv.Itoa(lastDigit)
	if lastDigits == checkString {
		return nil
	}
	return fmt.Errorf("invalid cpf")
}
