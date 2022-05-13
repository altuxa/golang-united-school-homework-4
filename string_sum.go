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

func StringSum(input string) (output string, err error) {
	a := ""
	sl := []string{}
	num := ""
	zn := []string{}
	// mp := make(map[int]string)
	check := false
	num1 := 0
	num2 := 0
	res := 0
	for _, i := range input {
		// if i != 32 || i != '+' || i != '-' || i < 48 && i > 57 {
		// 	return "", errorEmptyInput
		// }
		if i != 32 {
			a = a + string(i)
		}
	}
	///////////////
	// for _, j := range a {
	// 	// if j == '+' {
	// 	// 	sl = strings.Split(a, "+")
	// 	// 	break
	// 	// } else if j == '-' {
	// 	// 	sl = strings.Split(a, "-")
	// 	// 	break
	// 	// }
	// 	a = a + string(j)
	// 	numb, _ = strconv.Atoi(a)
	// 	if numb != 0 {

	// 	}
	// }
	for i := 0; i < len(a); i++ {
		if a[i] > '0' && a[i] < '9' {
			num = num + string(a[i])
		} else if a[i] < '0' || a[i] > '9' {
			zn = append(zn, string(a[i]))
			// mp[i] = string(a[i])
			if i == 0 {
				check = true
			}
			if num != "" {
				sl = append(sl, num)
				num = ""
			}
		}
	}
	if num != "" {
		sl = append(sl, num)
	}
	if check == true {
		num1, _ = strconv.Atoi(zn[0] + sl[0])
		if len(zn) == 2 {
			num2, _ = strconv.Atoi(sl[1])
			if zn[1] == "+" {
				res = num1 + num2
			} else if zn[1] == "-" {
				res = num1 - num2
			}
		}
	}
	if check == false {
		num1, _ = strconv.Atoi(sl[0])
		if len(zn) == 2 {
			num2, _ = strconv.Atoi(zn[1] + sl[1])
			if zn[0] == "+" {
				res = num1 + num2
			} else if zn[0] == "-" {
				res = num1 - num2
			}
		}
	}

	if len(zn) == 3 {
		num2, _ = strconv.Atoi(zn[2] + sl[1])
		if zn[1] == "+" {
			res = num1 + num2
		} else if zn[1] == "-" {
			res = num1 - num2
		}
	}
	fmt.Println(num1)
	fmt.Println(num2)
	// fmt.Println(mp)
	fmt.Println(zn)
	fmt.Println(sl)
	return strconv.Itoa(res), nil
}
