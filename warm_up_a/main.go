// Warm up problem A:
// Write a recursive function PrintInBinary(int num) that prints the binary representation for a given integer.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello")
}

func binChar(a int) string {
	lastDigit := a % 2
	if a > 1 {
		otherDigits := binChar(a / 2)
		return fmt.Sprintf("%s%d", otherDigits, lastDigit)
	}
	return strconv.Itoa(lastDigit)
}
