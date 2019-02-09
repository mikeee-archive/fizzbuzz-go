package fizzbuzz_test

import (
	"github.com/mikeee/fizzbuzz"
	"testing"
)

func TestCheck(t *testing.T) {
	tests := []struct {
		value  int
		expect string
	}{
		{1, "1"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{10, "Buzz"},
		{11, "11"},
		{15, "FizzBuzz"},
	}
	for _, test := range tests {
		fbresult := fizzbuzz.Check(test.value)
		if fbresult != test.expect {
			t.Errorf("Result incorrect, input: %d, got: %d, wanted:%d", test.value, fbresult, test.expect)
		}
	}
}
