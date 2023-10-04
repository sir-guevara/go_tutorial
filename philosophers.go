package main

import (
	"fmt"
	"sync"
	"time"
)

const numPhilosophers = 5
const numMeals = 3

var chopsticks = make([]*sync.Mutex, numPhilosophers)
var wg sync.WaitGroup

func main() {
	for i := 0; i < numPhilosophers; i++ {
		chopsticks[i] = &sync.Mutex{}
	}

	for i := 0; i < numPhilosophers; i++ {
		wg.Add(1)
		go philosopher(i + 1)
	}

	wg.Wait()
}

func philosopher(id int) {
	for meal := 0; meal < numMeals; meal++ {
		dine(id)
	}
	wg.Done()
}


func dine(id int) {

	left := id - 1
	right := (id + 1) % numPhilosophers // Corrected the right chopstick index

	chopsticks[left].Lock()
	chopsticks[right].Lock()

	fmt.Printf("Philosopher %d is starting to eat\n", id)
	time.Sleep(time.Millisecond * time.Duration(id*200))
	fmt.Printf("Philosopher %d is finishing eating\n", id)

	chopsticks[left].Unlock()
	chopsticks[right].Unlock()
}
