package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Implement the dining philosopher’s problem with the following constraints/modifications.

    There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
    Each philosopher should eat only 3 times
    The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
    In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
    The host allows no more than 2 philosophers to eat concurrently.
    Each philosopher is numbered, 1 through 5.
    When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
    When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
*/

type Chopsticks struct{ sync.Mutex }

type Philosopher struct {
	id, count           int
	leftChop, rightChop *Chopsticks
}

func (p Philosopher) eat(c chan *Philosopher, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		if p.count < 3 {
			c <- &p
			p.leftChop.Lock()
			p.rightChop.Lock()

			fmt.Println("starting to eat", p.id)
			p.count += 1
			fmt.Println("finishing eating", p.id)
			p.leftChop.Unlock()
			p.rightChop.Unlock()
			wg.Done()
		}
	}
}

func host(c chan *Philosopher, wg *sync.WaitGroup) {
	for {
		if len(c) == 2 {
			<-c
			<-c
			time.Sleep(20 * time.Millisecond)
		}
	}
}

func main() {
	Chops := make([]*Chopsticks, 5)
	for i := 0; i < 5; i++ {
		Chops[i] = new(Chopsticks)
	}

	philosophers := make([]*Philosopher, 5)

	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i + 1, 0, Chops[i], Chops[(i+1)%5]}
	}
	// fmt.Println(Chops)
	// fmt.Println(philosophers)

	philo_c := make(chan *Philosopher, 2)

	var wg sync.WaitGroup

	wg.Add(15)
	go host(philo_c, &wg)

	for i := 0; i < 5; i++ {
		go philosophers[i].eat(philo_c, &wg)
	}
	wg.Wait()
	close(philo_c)
}
