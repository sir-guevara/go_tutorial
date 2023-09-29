/*
	implementing bubbleSort algorithm in go
	takes 1o ints  from user and prints it out on one line in sorted order  from least to greatest
*/
package main
import (
	"fmt"
)


func BubbleSort(aslice []int){
	for i:=0; i < len(aslice); i++ {
		for j:=0; j<len(aslice) -i -1; j++{
			if aslice[j] > aslice[j+1]{
				Swap(aslice,j)
			}
		}
	}

	for _, value := range aslice{
		fmt.Printf("%d ",value)
	}
}

func Swap(vslice []int, j int){
	tmp := vslice[j]
	vslice[j]= vslice[j+1]
	vslice[j+1]= tmp
}


func main(){

	var sl []int
	var num int
	fmt.Println("Please enter a sequence of integers (up to 10)")
	for i :=0; i< 10; i++{
		fmt.Printf("Enter integer #%d: ",i+1)
		fmt.Scan(&num)
		sl = append(sl,num)
	}

	BubbleSort(sl)
}