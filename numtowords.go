package numtowords

import "fmt"

const MaxNum = 99

func Convert(number int) (string, error) {

	if number < 0 || number > MaxNum {
		return "", fmt.Errorf("Number out of range")
	}
	ones := []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"fifteen",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen",
	}

	tens := []string{
		"twenty",
		"thirty",
		"forty",
		"fifty",
		"sixty",
		"seventy",
		"eighty",
		"ninety",
	}

	if number < 20 {
		return ones[number], nil
	}

	if number%10 == 0 {
		return tens[(number/10)-2], nil
	}

	return tens[(number/10)-2] + " " + ones[number%10], nil
}
