package fizzbuzz

import "fmt"

// Check takes an integer and applies the fizzbuzz logic to it
func Check(value int) (result string) {
	if value%3 == 0 {
		result = "Fizz"
	}
	if value%5 == 0 {
		result += "Buzz"
	}
	if result == "" {
		result = fmt.Sprintf("%d", value)
	}
	return result
}
