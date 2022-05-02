//go:build unit
// +build unit

package calculator_test

import (
	"testing"

	"github.com/ajaykakodia/go-testing/calculator"
)

type TestCase struct {
	name     string
	value    int
	expected bool
	actual   bool
}

func TestMultipleCalculateIsArmstrong(t *testing.T) {
	t.Run("test for 3 digit numbers", func(t *testing.T) {
		testCases := []TestCase{
			{name: "test for value: 371", value: 371, expected: true},
			{name: "test for value: 370", value: 370, expected: true},
			{name: "test for value: 153", value: 153, expected: true},
			{name: "test for value: 407", value: 407, expected: true},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
				if testCase.actual != testCase.expected {
					t.Fail()
				}
			})
		}
	})
}

func TestCalculateIsArmstrong(t *testing.T) {

	t.Run("should give true for 371", func(t *testing.T) {
		testCase := TestCase{
			value:    371,
			expected: true,
		}

		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})

	t.Run("should give true for 370", func(t *testing.T) {
		testCase := TestCase{
			value:    370,
			expected: true,
		}

		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})

}

func TestNegativeCalculateArmstrong(t *testing.T) {
	t.Run("should give false for 123", func(t *testing.T) {
		testCase := TestCase{
			value:    123,
			expected: false,
		}
		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})

	t.Run("shoul give false for 256", func(t *testing.T) {
		testCase := TestCase{
			value:    123,
			expected: false,
		}
		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})
}

func benchmarkCalculateArmstrong(input int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		calculator.CalculateIsArmstrong(input)
	}
}

func BenchmarkCalculateIsArmstrong370(b *testing.B) {
	benchmarkCalculateArmstrong(370, b)
}

func BenchmarkCalculateIsArmstrong371(b *testing.B) {
	benchmarkCalculateArmstrong(371, b)
}

func BenchmarkCalculateIsArmstrong153(b *testing.B) {
	benchmarkCalculateArmstrong(153, b)
}

func BenchmarkCalculateIsArmstrong0(b *testing.B) {
	benchmarkCalculateArmstrong(0, b)
}
