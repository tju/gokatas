package gokatas

import (
	"strconv"
	"strings"
)

func calculate(input string) (int, error) {
	split := strings.Split(input, " ")

	var result int
	for _, item := range split {
		parsed, err := strconv.Atoi(item)

		if err != nil {
			return 0, err
		}

		result += parsed
	}

	return result, nil
}
