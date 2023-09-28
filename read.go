// Takes a text file containing names  and prints the names in the file

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	type Name struct {
		fname string
		lname string
	}
	var fileName string

	fmt.Print("Enter the name of the text file without extension: ")
	_, err := fmt.Scan(&fileName)
	if err != nil {
		fmt.Printf("Error reading filename: %v\n", err)
		return
	}
	var names = make([]Name, 0)

	// Open the text file
	file, err := os.Open(fileName + ".txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Iterate through each line in the file
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into first name and last name
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Printf("Invalid line format: %s\n", line)
			continue
		}

		// Create a Name struct and add it to the slice
		name := Name{
			fname: parts[0],
			lname: parts[1],
		}
		names = append(names, name)
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Print the names
	fmt.Println("Names in the file:")
	for _, n := range names {
		fmt.Printf(" %s %s\n", n.fname, n.lname)
	}

}
