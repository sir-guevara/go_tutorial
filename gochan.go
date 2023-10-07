package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"sort"
	"sync"

)

func main() {

	wg := &sync.WaitGroup{}
	fmt.Println("Enter a series of integers separated by spaces:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	integers := parseInput(input)
	partitionSize := len(integers)/4
	ch := make(chan []int, 4)

	for i := 0; i < 4; i++ {
		startIdx := i * partitionSize
		endIdx := startIdx + partitionSize
		if i == 3 {
			endIdx = len(integers) 
		}
		go sortInts(i, integers[startIdx:endIdx], ch, wg)
		wg.Add(1)

	}

	wg.Wait()
	close(ch)
	// Merge the sorted subarrays
	merged := mergeSortedArrays(ch)
	// Print the entire sorted list
	fmt.Println("Sorted List:", merged)

}


func parseInput(input string) []int {
	var integers []int
	strIntegers := strings.Split(input, " ")
	for _, str := range strIntegers {
		num, err := strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		integers = append(integers, num)
	}
	return integers
}

func sortInts(pos int ,ints []int, ch chan []int, wg *sync.WaitGroup){
			fmt.Printf("Goroutine %d sorting: %v\n", pos+1, ints)
			sort.Ints(ints)
			ch <- ints
			wg.Done()
		}

func mergeSortedArrays(ch chan []int) []int {
	merged := make([]int, 0)

	for c := range ch {
		for _, num := range c {
					merged = append(merged, num)
				}
		sort.Ints(merged)
	}

	return merged
}
