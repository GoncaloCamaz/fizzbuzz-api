package utils

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("when limit is set to 0", func(t *testing.T) {
		result := FizzBuzz(0, 3, 5, "Fizz", "Buzz")
		if len(result) != 0 {
			t.Errorf("expected length of result to be 0, got %d", len(result))
		}
	})

	t.Run("when using default fizzbuzz exercise values", func(t *testing.T) {
		result := FizzBuzz(15, 3, 5, "Fizz", "Buzz")

		expected := []string{
			"1", "2", "Fizz", "4", "Buzz",
			"Fizz", "7", "8", "Fizz", "Buzz",
			"11", "Fizz", "13", "14", "FizzBuzz",
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("FizzBuzz(15, 3, 5, Fizz, Buzz) = %v, want %v", result, expected)
		}
	})

	t.Run("when replacement strings are equal", func(t *testing.T) {
		result := FizzBuzz(15, 3, 5, "Fizz", "Fizz")

		expected := []string{
			"1", "2", "Fizz", "4", "Fizz",
			"Fizz", "7", "8", "Fizz", "Fizz",
			"11", "Fizz", "13", "14", "FizzFizz",
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("FizzBuzz(15, 3, 5, Fizz, Fizz) = %v, want %v", result, expected)
		}
	})

	t.Run("when multiples are equal", func(t *testing.T) {
		result := FizzBuzz(10, 2, 2, "Foo", "Bar")

		expected := []string{
			"1", "FooBar", "3", "FooBar", "5",
			"FooBar", "7", "FooBar", "9", "FooBar",
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("FizzBuzz(10, 2, 2, Foo, Bar) = %v, want %v", result, expected)
		}
	})
}
