package gokatas

import "strconv"

func fizzBuzz(input int) string {

	if input%3 == 0 && input%5 == 0 {
		return "FizzBuzz"
	}

	if input%3 == 0 {
		return "Fizz"
	}

	if input%5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(input)
}
