package main

import "fmt"

func decode(num int) string {
	var result string
	var i int

	var romanMap = map[int]string{1: "I", 5: "V", 10: "X", 50: "L", 100: "C"}

	switch {
	case num < 4:
		for i = 1; i <= num; i++ {
			result += romanMap[1]
		}
	case num == 4 || num == 9 || num == 14:
		result += romanMap[1] + romanMap[num+1]
	case num < 9:
		result += romanMap[5]
		for i = 1; i <= num-5; i++ {
			result += romanMap[1]
		}
	case num == 40:
		result += romanMap[10] + romanMap[50]
	case num == 90:
		result += romanMap[10] + romanMap[100]
	default:
		result += romanMap[num]
	}
	return result
}

func main() {
	var romanMap = map[int]string{1: "I", 5: "V", 10: "X", 50: "L"}
	var countModeTen = 0
	var countTen = ""
	for i := 1; i <= 100; i++ {
		if i%10 == 0 {
			countModeTen++
			if i < 40 {
				countTen += romanMap[10]
			} else if i > 50 && i < 100 {
				countTen += romanMap[10]
			} else {
				countTen = ""
			}
		}
		if i/10 > 0 && i < 100 {
			if i/10 == 4 {
				fmt.Printf("%d = %s%s\n", i, decode(40), decode(i-(countModeTen*10)))
			} else if i >= 50 && i < 90 {
				fmt.Printf("%d = %s%s%s\n", i, decode(50), countTen, decode(i-(countModeTen*10)))
			} else if i/10 == 9 {
				fmt.Printf("%d = %s%s\n", i, decode(90), decode(i-(countModeTen*10)))
			} else {
				fmt.Printf("%d = %s%s\n", i, countTen, decode(i-(countModeTen*10)))
			}
		} else {
			fmt.Printf("%d = %s\n", i, decode(i))
		}
	}
}
