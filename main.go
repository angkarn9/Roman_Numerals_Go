package main

import (
	"fmt"

	"./roman"
)

func main() {
	for i := 1; i <= 100; i++ {
		romanNumeral := roman.ConvertToRomanNumeral(i)
		fmt.Printf("%d = %s\n", i, romanNumeral)
	}
}
