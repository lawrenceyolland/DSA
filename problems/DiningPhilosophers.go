package main

import (
	"fmt"
	"sync"
)

const numPhilosophers int = 5
const concurrentDiners int = 2

type chopstick struct{ sync.Mutex }

type philosopher struct{ id, leftChopstick, rightChopstick int }

type host struct {
	chopsticks         []*chopstick
	hostRequestChannel chan bool
}

func displayMessage(msg string, p *philosopher) {
	fmt.Printf("%s %d\n", msg, p.id)
}

func (dh *host) grabChopsticks(p *philosopher) {
	dh.hostRequestChannel <- true
	for {
		dh.chopsticks[p.leftChopstick].Lock()
		if dh.chopsticks[p.rightChopstick].TryLock() {
			displayMessage("starting to eat", p)
			break
		}
		dh.chopsticks[p.leftChopstick].Unlock()
	}
}

func (dh *host) releaseChopsticks(p *philosopher) {
	displayMessage("finishing eating", p)
	dh.chopsticks[p.leftChopstick].Unlock()
	dh.chopsticks[p.rightChopstick].Unlock()
	<-dh.hostRequestChannel
}

func (p philosopher) eat(dh *host, meal *sync.WaitGroup) {
	defer meal.Done()
	for i := 0; i < 3; i++ {
		dh.grabChopsticks(&p)
		dh.releaseChopsticks(&p)
	}
}

func main() {
	var meal sync.WaitGroup
	chopsticks := make([]*chopstick, numPhilosophers)
	philosophers := make([]philosopher, numPhilosophers)

	for i := 0; i < numPhilosophers; i++ {
		chopsticks[i] = new(chopstick)
	}

	dinnerHost := host{
		chopsticks:         chopsticks,
		hostRequestChannel: make(chan bool, concurrentDiners),
	}

	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = philosopher{
			id:             i + 1,
			leftChopstick:  i,
			rightChopstick: (i + 1) % numPhilosophers,
		}
	}

	for i := range philosophers {
		meal.Add(1)
		go philosophers[i].eat(&dinnerHost, &meal)
	}

	meal.Wait()
	close(dinnerHost.hostRequestChannel)
}
