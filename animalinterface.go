// This program allows the user to create and query animals,
// providing information about their food, locomotion, and spoken sound.

package main

import (
    "fmt"
)

// Animal interface defines methods for Eat, Move, and Speak.
type Animal interface {
    Eat()
    Move()
    Speak()
}

// Cow type satisfies the Animal interface.
type Cow struct{}

func (c Cow) Eat() {
    fmt.Println("grass")
}

func (c Cow) Move() {
    fmt.Println("walk")
}

func (c Cow) Speak() {
    fmt.Println("moo")
}

// Bird type satisfies the Animal interface.
type Bird struct{}

func (b Bird) Eat() {
    fmt.Println("worms")
}

func (b Bird) Move() {
    fmt.Println("fly")
}

func (b Bird) Speak() {
    fmt.Println("peep")
}

// Snake type satisfies the Animal interface.
type Snake struct{}

func (s Snake) Eat() {
    fmt.Println("mice")
}

func (s Snake) Move() {
    fmt.Println("slither")
}

func (s Snake) Speak() {
    fmt.Println("hsss")
}

func main() {
    animalMap := make(map[string]Animal)

    for {
        fmt.Print("> ")
        var command, name, request string
        fmt.Scan(&command)

        if command == "newanimal" {
            fmt.Scan(&name, &request)
            switch request {
            case "cow":
                animalMap[name] = Cow{}
            case "bird":
                animalMap[name] = Bird{}
            case "snake":
                animalMap[name] = Snake{}
            default:
                fmt.Println("Invalid animal type.")
            }
            fmt.Println("Created it!")
        } else if command == "query" {
            fmt.Scan(&name, &request)
            animal, found := animalMap[name]
            if !found {
                fmt.Println("Animal not found.")
                continue
            }
            switch request {
            case "eat":
                animal.Eat()
            case "move":
                animal.Move()
            case "speak":
                animal.Speak()
            default:
                fmt.Println("Invalid query.")
            }
        } else {
            fmt.Println("Invalid command. Please use 'newanimal' or 'query'.")
        }
    }
}
