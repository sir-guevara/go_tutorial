// prompts the user to enter integers and stores the integers in a sorted slice. 

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// initializing the slice
	intSlice := make([]int, 3)

	// loop exits if input is x
	for input := ""; input != "x"; {

		fmt.Printf("Enter an integer (type 'x' to quit): ")
		_, err := fmt.Scanln(&input) //getting user input

		// converting the input to int
		num, err := strconv.Atoi(input)
		if err == nil {
			intSlice = append(intSlice, int(num))
			sort.Ints(intSlice) // sorting the slice
		}

		fmt.Println(intSlice)

		// converting to lowercase so x | X can work
		input = strings.ToLower(input)
	}

}
