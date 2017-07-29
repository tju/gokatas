package gokatas

import (
	"testing"
)

type testData struct {
	input  string
	output int
}

func TestCalculate(t *testing.T) {
	testTable := []testData{
		{"0", 0},
		{"1", 1},
		{"2", 2},

		{"1 1", 2},
		{"1 3", 4},

		{"1 1 1", 3},
		{"0 0 0", 0}}

	for _, data := range testTable {
		t.Run("test input "+data.input, func(t *testing.T) {
			result, err := calculate(data.input)

			if err != nil {
				t.Error(err)
			}

			if data.output != result {
				t.Errorf("Expected:%v ,Got:%v ", data.output, result)
			}
		})
	}
}
