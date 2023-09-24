// takes a floating point number from user and prints the truncated number

package main

import (
	"fmt"
)

func main() {
	var floatNum float32

	fmt.Printf("Please Enter a floating point numnber: ")
	fmt.Scan(&floatNum)

	truncateNumber := int(floatNum)
	fmt.Printf("\n Truncated number is :  %d \n", truncateNumber)

}
