package main

import (
	"fmt"
	"sync"
)

// you can create go routines and wait groups to wait for those go routines and you also have to specify when to end the go routine wg.Done()

func Tut3() {
	fmt.Println("from tut3")

	ch := make(chan string)
	var wg sync.WaitGroup

	wg.Add(1)
	go receiveFromChannel(&wg, ch)

	wg.Add(1)
	go pushToChannel(&wg, ch, "Shashank")

	wg.Wait()

}

// for channel we need to make sure that there should be a  lister before pushing data into the channel
// if channel is closed then value reture is 0 and in that case to differentiate between actual input 0 or closed input 0 we should use comma-ok syntax
// channel are used for communication between thread/go routines in other words

func pushToChannel(wg *sync.WaitGroup, ch chan string, value string) {
	defer wg.Done()
	ch <- value
	ch <- value
}

func receiveFromChannel(wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	fmt.Println("from receive channel", <-ch)
	fmt.Println("from receive channel", <-ch)
	close(ch)
}
