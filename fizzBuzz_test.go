package gokatas

import (
	"testing"
)

func TestNormalNumbers(t *testing.T) {
	table := []TestData{
		{1, "1"},
		{2, "2"},
		{4, "4"},
	}

	testTable(t, table)
}

func TestBuzz(t *testing.T) {
	table := []TestData{
		{5, "Buzz"},
		{10, "Buzz"},
	}

	testTable(t, table)
}

func TestFizzBuzz(t *testing.T) {
	table := []TestData{
		{15, "FizzBuzz"},
		{30, "FizzBuzz"},
	}

	testTable(t, table)
}

func TestFizzNumbers(t *testing.T) {
	table := []TestData{
		{3, "Fizz"},
		{6, "Fizz"},
	}

	testTable(t, table)
}

type TestData struct {
	input    int
	expected string
}

func testTable(t *testing.T, table []TestData) {
	for _, data := range table {
		var output = fizzBuzz(data.input)

		if data.expected != output {
			t.Errorf("For input: %v expected:%v ,got:%v ", data.input, data.expected, output)
		}
	}
}
