package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopS struct{ sync.Mutex }
type Philo struct {
	id              int
	leftCS, rightCS *ChopS
	eaten           int
}

var wg sync.WaitGroup
var concurrentEaters int = 0
var host = make(chan bool, 2)

func (p Philo) eat() {
	for p.eaten < 3 {
		host <- true

		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("starting to eat <%d> (%d times)\n", p.id, p.eaten+1)
		time.Sleep(1 * time.Second)
		fmt.Printf("finished eating <%d> (%d times)\n", p.id, p.eaten+1)
		p.eaten = p.eaten + 1

		p.rightCS.Unlock()
		p.leftCS.Unlock()

		<-host
	}
	wg.Done()
}

func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{
			i + 1,
			CSticks[i],
			CSticks[(i+1)%5],
			0,
		}
	}

	fmt.Printf("philosophers: %v", philos)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat()
	}
	wg.Wait()
}