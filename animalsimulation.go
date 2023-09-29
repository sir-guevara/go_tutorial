// This program clears a  struct (Animal) and add method (Eat,Move, and Speak) 
// It prompts the user to enter either cow, bird or snake and  the action (eat,move or speak) then returns a print statement for the request.

package main

import (
	"bufio"
    "fmt"
    "os"
	"strings"
)

type Animal struct{
	locomotion string
	sound string
	food string
}

func (a Animal) Eat () string{
return a.food
}

func (a Animal) Move () string{
return a.locomotion
}

func (a Animal) Speak () string{
return a.sound
}

var cow = Animal{
		locomotion:"walk",
		sound:"moo",
		food:"grass",
	}
var bird = Animal {
		locomotion:"fly",
		sound:"peep",
		food:"worms",
	}

var snake = Animal{
		locomotion:"slither",
		sound:"hsss",
		food:"mice",
	}

	var invalid = Animal{
	locomotion:"please enter a valid animal('cow','bird', or 'snake')",
	sound:"please enter a valid animal('cow','bird', or 'snake')",
	food:"please enter a valid animal('cow','bird', or 'snake') ",
}


func exec (animal, action string ){
	var a *Animal;
	var atn string
	switch animal{
	case "cow":
		a = &cow
	case  "bird":
		a = &cow
	case "snake":
		a = &snake
	default:
		a = &invalid
	}
	switch action{
	case "eat":
		atn = a.Eat()
	case  "speak":
		atn = a.Speak()
	case "move":
		atn = a.Move()
	default:
		action = "invalid action use('move','eat', or 'speak')"
	}

	fmt.Println(atn)
}

func main(){
	 reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please type an animal (cow, bird, or snake ) and an action ( eat, move, or speak) e.g cow speak ")
	for{
		fmt.Printf("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		parts := strings.Split(input, " ")
		if len(parts) < 2 {
			fmt.Println("Invalid input. Please provide a string with two parts separated by a space.")
    	}

		animal := parts[0]
    	action := parts[1]
		exec (animal,action)

	}

}