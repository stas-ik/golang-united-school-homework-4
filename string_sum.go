package string_sum

import (
	"errors"
	"strings"
	"fmt"
	"strconv"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
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
	resString := strings.ReplaceAll(input, " ", "")
	if len(resString) == 0 {
		return "", errorEmptyInput
	}

	resIndex := 1
	var isSum string = "-"
	if strings.Contains(resString, "+") {
		isSum = "+"
		resIndex = 0
	}

	sum1, err1 := strconv.Atoi(strings.Split(resString, isSum)[0 + resIndex])
	sum2, err2 := strconv.Atoi(strings.Split(resString, isSum)[1 + resIndex])

	if err1 != nil || err2 != nil  {
		return "", errorNotTwoOperands
	}

	if resIndex > 0 {
		sum1 = sum1*-1
	}

	if isSum == string('-') {
		return string(sum1 - sum2), nil
	}
	return string(sum1 + sum2), nil
}
