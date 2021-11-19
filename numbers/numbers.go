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

	if arabic >= 1000 {
		thousand := arabic / 1000
		result = numbersUnderTwenty[thousand] + " thousand,"
		arabic -= (arabic / 1000) * 1000
	}

	result = processHundreds(&arabic, result)

	result = processTens(&arabic, result)

	return processUnderTwenty(arabic, result)
}

func processHundreds(arabic *int, result string) string {
	if *arabic >= 100 {
		prependSpace(&result)
		nearestHundred := *arabic / 100
		result += numbersUnderTwenty[nearestHundred] + " hundred"
		*arabic %= 100
		if *arabic > 0 {
			result += " and"
		}
	}
	return result
}

func processTens(arabic *int, result string) string {
	if *arabic >= 20 {
		prependSpace(&result)
		nearestTen := (*arabic / 10) * 10
		result += multiplesOfTen[nearestTen]
		*arabic %= 10
	}
	return result
}

func processUnderTwenty(arabic int, result string) string {
	if arabic > 0 {
		prependSpace(&result)
		result += numbersUnderTwenty[arabic]
	}
	return result
}

func prependSpace(result *string) {
	if len(*result) > 0 {
		*result += " "
	}
}
