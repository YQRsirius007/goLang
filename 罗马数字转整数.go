package pos

func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}
	sum := 0
	for i := 0; i < len(s)-1; i++ {
		if roman(s[i]) > roman(s[i+1]) {
			sum -= roman(s[i])
		} else {
			sum += roman(s[i])
		}
	}
	sum += roman(s[len(s)-1])
	return sum
}
func roman(c uint8) int {
	switch c {
	case 'M':
		return 1000
	case 'D':
		return 500
	case 'C':
		return 100
	case 'L':
		return 50
	case 'X':
		return 10
	case 'V':
		return 5
	case 'I':
		return 1
	}
	return 0
}
