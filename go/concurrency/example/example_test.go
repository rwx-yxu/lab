package example

import (
	"sync"
	"testing"
)

func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	sender := func() {
		defer wg.Done()
		<-begin // (1) Corrected: Remove the "1"
		for i := 0; i < b.N; i++ {
			c <- struct{}{} // (2) Corrected: Send an empty struct
		}
	}
	receiver := func() {
		defer wg.Done()
		<-begin // (1) Corrected: Remove the "1"
		for i := 0; i < b.N; i++ {
			<-c // (3) Corrected: Receive from the channel
		}
	}

	wg.Add(2)
	go sender()
	go receiver()
	b.ResetTimer() // (4) Corrected: Use ResetTimer to reset the timer
	close(begin)   // (5) Corrected: Close the 'begin' channel to start the goroutines
	wg.Wait()
}
