package go_leetcode

func romanToInt(s string) int {
	l := len(s)
	i, sum := 0, 0
	for i < l {
		if i == l-1 {
			sum += romanMap[s[i:i+1]]
			break
		}
		if v, ok := romanMap[s[i:i+2]]; ok {
			sum += v
			i += 2
		} else {
			sum += romanMap[s[i:i+1]]
			i++
		}
	}
	return sum
}

var romanMap = map[string]int{
	"I":  1,
	"V":  5,
	"X":  10,
	"L":  50,
	"C":  100,
	"D":  500,
	"M":  1000,
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
}
