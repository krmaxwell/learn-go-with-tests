package numbers

var numbersUnderTwenty = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
}

var multiplesOfTen = map[int]string{
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

func SpellNumber(arabic int) string {

	var result string

	if arabic >= 20 {
		ten := (arabic / 10) * 10
		result = multiplesOfTen[ten]
		arabic -= ten
	}
	if arabic > 0 {
		if len(result) > 0 {
			result += " "
		}
		result += numbersUnderTwenty[arabic]
	}
	return result
}
