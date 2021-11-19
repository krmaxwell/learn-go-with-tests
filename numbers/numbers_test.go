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
		{20, "twenty"},
		{21, "twenty one"},
		{22, "twenty two"},
		{30, "thirty"},
		{37, "thirty seven"},
		{42, "forty two"},
		{53, "fifty three"},
		{69, "sixty nine"},
		{76, "seventy six"},
		{81, "eighty one"},
		{99, "ninety nine"},
		{100, "one hundred"},
		{110, "one hundred and ten"},
		{135, "one hundred and thirty five"},
		{200, "two hundred"},
		{300, "three hundred"},
		{310, "three hundred and ten"},
		{666, "six hundred and sixty six"},
		{1501, "one thousand, five hundred and one"},
		{12609, "twelve thousand, six hundred and nine"},
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
