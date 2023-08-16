package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Deque struct
type Deque struct {
	items []int
	lock  sync.Mutex
}

// adds an element to the front of the queue
func (dq *Deque) PushFront(element int) {
	dq.lock.Lock()
	defer dq.lock.Unlock()
	dq.items = append([]int{element}, dq.items...)
}

// adds an element to the rear end of the queue
func (dq *Deque) PushRear(element int) {
	dq.lock.Lock()
	defer dq.lock.Unlock()
	dq.items = append(dq.items, element)
}

// removes an element from the front of the queue
func (dq *Deque) PopFront() int {
	dq.lock.Lock()
	defer dq.lock.Unlock()
	if len(dq.items) == 0 {
		fmt.Println("Deque is empty")
		return -1
	}

	element := dq.items[0]
	dq.items = dq.items[1:]
	return element
}

// removes an element from the front of the queue
func (dq *Deque) PopRear() int {
	dq.lock.Lock()
	defer dq.lock.Unlock()
	if len(dq.items) == 0 {
		fmt.Println("Deque is empty")
		return -1
	}

	element := dq.items[len(dq.items)-1]
	dq.items = dq.items[:len(dq.items)-1]
	return element
}


func main() {
	dq := &Deque{}
	notify := make(chan int)
	producer1(dq, notify)
	producer2(dq, notify)
	fmt.Println("Press Ctrl + C to force exit program")
	go func() {
		for {
			select {
			case pNo := <-notify:
				fmt.Printf("from producer%d: %d\n", pNo, dq.PopFront())
			}
		}
	}()

	// called here in order to freeze the program
	select {}
}

func producer1(dq *Deque, notify chan int) {
	go func() {
		for {
			dq.PushFront(rand.Intn(100))
			notify <- 1
			time.Sleep(time.Second)
		}
	}()
}

func producer2(dq *Deque, notify chan int) {
	go func() {
		for {
			dq.PushRear(rand.Intn(100))
			notify <- 2
			time.Sleep(time.Second * 2)
		}
	}()
}
