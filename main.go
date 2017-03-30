package main

import (
	"fmt"

	"./roman"
)

func main() {
	var countModeTen int
	var countTen string
	var romanNumeral string
	for i := 1; i <= 100; i++ {
		countModeTen, countTen, romanNumeral = roman.ConvertToRomanNumeral(i, countModeTen, countTen)
		fmt.Printf("%d = %s\n", i, romanNumeral)
	}
}
