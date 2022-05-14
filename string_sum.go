package string_sum

import (
	"errors"
	"fmt"
	"strconv"
)

// use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func DeleteWhitespace(s string) (string, error) {
	a := ""
	for _, i := range s {
		if i != 32 {
			a = a + string(i)
		}
	}
	if len(a) == 0 {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}
	return a, nil
}

func ValidNumberOfOperands(n int) bool {
	return n == 2
}

func SplittingExpression(a string) ([]string, []string, bool) {
	operands := []string{}
	num := ""
	zn := []string{}
	check := false
	for i := 0; i < len(a); i++ {
		if a[i] != '+' && a[i] != '-' {
			num = num + string(a[i])
		} else if a[i] == '+' || a[i] == '-' {
			zn = append(zn, string(a[i]))
			if i == 0 {
				check = true
			}
			if num != "" {
				operands = append(operands, num)
				num = ""
			}
		}
	}
	if num != "" {
		operands = append(operands, num)
	}
	return operands, zn, check
}

func StringSum(input string) (output string, err error) {
	num1 := 0
	num2 := 0
	res := 0
	a, err := DeleteWhitespace(input)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	operands, zn, check := SplittingExpression(a)
	valid := ValidNumberOfOperands(len(operands))
	if !valid {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}
	if len(zn) == 1 {
		zn = append(zn, "")
		zn = append(zn, "")
		zn[1] = zn[0]
		zn[0] = ""
		zn[2] = ""
	}
	if check && len(zn) == 2 {
		zn = append(zn, "")
	} else if !check && len(zn) == 2 {
		zn = append(zn, "")
		zn[2] = zn[1]
		zn[1] = zn[0]
		zn[0] = ""
	}

	num1, err = strconv.Atoi(zn[0] + operands[0])
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	num2, err = strconv.Atoi(zn[2] + operands[1])
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	if zn[1] == "+" {
		res = num1 + num2
	} else if zn[1] == "-" {
		res = num1 - num2
	}
	output = strconv.Itoa(res)
	return output, nil
}
