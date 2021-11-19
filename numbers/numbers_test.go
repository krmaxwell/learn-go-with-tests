package numbers

import (
	"fmt"
	"testing"
)

func TestNumbers(t *testing.T) {
	cases := []struct {
		Arabic int
		Want   string
	}{
		{1, "one"},
		{2, "two"},
		{3, "three"},
		{4, "four"},
		{5, "five"},
		{6, "six"},
		{7, "seven"},
		{8, "eight"},
		{9, "nine"},
		{10, "ten"},
		{11, "eleven"},
		{12, "twelve"},
		{13, "thirteen"},
		{14, "fourteen"},
		{15, "fifteen"},
		{16, "sixteen"},
		{17, "seventeen"},
		{18, "eighteen"},
		{19, "nineteen"},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("%d converts to %q", test.Arabic, test.Want), func(t *testing.T) {
			got := SpellNumber(test.Arabic)

			if got != test.Want {
				t.Errorf("got %q want %q", got, test.Want)
			}

		})
	}
}
