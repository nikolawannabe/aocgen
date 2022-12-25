package year2022

import (
	"math"
)

type Day25 struct{}

func powInt(x, y int64) int64 {
	return int64(math.Pow(float64(x), float64(y)))
}

func snafuToDecimal(snafu string) int64 {
	runes := split(snafu)
	places := int64(len(snafu))
	base := int64(5)

	total := int64(0)
	for i, r := range runes {
		exp := places - int64(i) - 1
		whichPlace := powInt(base, exp)
		switch r {
		case '2':
			total += whichPlace * 2
			break
		case '1':
			total += whichPlace * 1
			break
		case '0':
			total += whichPlace * 0
			break
		case '-':
			total += whichPlace * -1
			break
		case '=':
			total += whichPlace * -2
			break

		}

	}
	return total
}

func decToSnafu(dec int64) string {
	out := make([]string, 0)
	num := dec
	carry := int64(0)
	for num > 0 {
		digit := num%5 + carry
		carry = 0
		if digit > 2 {
			carry = 1
			digit -= 5
		}
		switch digit {
		case 2:
			out = append(out, "2")
			break
		case 1:
			out = append(out, "1")
			break
		case 0:
			out = append(out, "0")
			break
		case -1:
			out = append(out, "-")
			break
		case -2:
			out = append(out, "=")
		}

		num /= 5
	}
	if carry == 1 {
		out = append(out, "1")
	}
	return concatStr(reverse(out))
}

func (p Day25) PartA(lines []string) any {
	total := int64(0)

	for _, snafu := range lines {
		total += snafuToDecimal(snafu)
	}

	return decToSnafu(total)
}

func (p Day25) PartB(lines []string) any {
	return "implement_me"
}
