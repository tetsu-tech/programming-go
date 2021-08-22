package format

import (
	"fmt"
	"testing"
)

func TestAny(t *testing.T) {
	var tests = []struct {
		input interface{}
		want  string
	}{
		{1, "1"},
		{"string", "\"string\""},
		{true, "true"},
	}

	for _, test := range tests {
		if got := Any(test.input); got != test.want {
			fmt.Printf("%s", got)
			t.Errorf("Any(%d) = %v", test.input, test.want)
		}
	}
}
