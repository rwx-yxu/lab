package example

import (
	"fmt"
	"log"
	"runtime"
	"sync"

	"github.com/rwxrob/choose"
)

var Choices = []string{
	`Go routine size - Print out the size of an individual empty go routine`,
	`Closure - Invoke a go routine in a for loop which iterates over a
string slice to print out the current item.`,
	`Wait group - Use wait groups to demonstrate the go routines are invoked
according to the runtime scheduler.`,
	`Mutex - Use sync.Mutex to modify a variable concurrently by using locks 
to ensure resource protection when multiple processes want access to modify.`,
}

func Run() int {
	choice, _, err := choose.From(Choices)
	if err != nil {
		log.Println(err)
		return 1
	}
	switch choice {
	case 0:
		GoRoutineSize()
	case 1:
		Closure()
	case 2:
		WG()
	case 3:
		Access()
	}
	return 0
}

func GoRoutineSize() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() { wg.Done(); <-c }

	const numGoroutines = 1e4
	wg.Add(numGoroutines)
	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Printf("%.3fkb\n", float64(after-before)/numGoroutines/1000)

}

func Closure() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}

func WG() {
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v!\n", id)
	}

	const numGreeters = 5
	var wg sync.WaitGroup
	wg.Add(numGreeters)
	for i := 0; i < numGreeters; i++ {
		go hello(&wg, i+1)
	}
	wg.Wait()
}

func Access() {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	// Increment
	var arithmetic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	// Decrement
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	fmt.Println("Arithmetic complete.")
}
