package roman

import "fmt"

func romanDefault() func(int) string {
	var romanMap = map[int]string{1: "I", 5: "V", 10: "X", 50: "L", 100: "C"}

	return func(key int) string {
		return romanMap[key]
	}
}

func decode(num int) string {
	var result string
	switch {
	case num/10 == 9 || num/10 == 4:
		result += romanDefault()(10) + romanDefault()(num+10)
	case num < 4:
		result += concat1to2(num%5, 0)
	case num == 4 || num == 9 || num == 14:
		result += romanDefault()(1) + romanDefault()(num+1)
	case num < 9:
		result += concat1to2(num%5, 5)
	default:
		result += romanDefault()(0) + romanDefault()(num+0)
	}
	return result
}

func concat1to2(num int, prefix int) string {
	var result = romanDefault()(prefix)
	for i := 1; i <= num; i++ {
		result += romanDefault()(1)
	}
	return result
}

func ConvertToRomanNumeral(i int) string {
	var result string
	var countModeTen int
	var countTen string
	for ii := 1; ii <= i; ii++ {
		if ii%10 == 0 {
			countModeTen++
			if ii == 100 {
				countModeTen = 0
			}
			if ii != 50 && ii != 40 && ii != 100 {
				countTen += romanDefault()(10)
			} else {
				countTen = ""
			}
		}
	}

	if i == 100 {
		result = fmt.Sprintf("%s%s", countTen, decode(i-(countModeTen*10)))
	} else if i >= 90 {
		result = fmt.Sprintf("%s%s", decode(90), decode(i-(countModeTen*10)))
	} else if i >= 50 {
		result = fmt.Sprintf("%s%s%s", decode(50), countTen, decode(i-(countModeTen*10)))
	} else if i >= 40 {
		result = fmt.Sprintf("%s%s", decode(40), decode(i-(countModeTen*10)))
	} else {
		result = fmt.Sprintf("%s%s", countTen, decode(i-(countModeTen*10)))
	}
	return result
}
