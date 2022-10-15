package pipeline

import (
	"fmt"
	"time"
)

type Data int
type Square int
type Result int

func Run() {
	bufSize := 3

	//Channel to communicate raw data to stage 1
	raw := make(chan int, bufSize)
	//Channel to communicate data between Stage 1 and Stage 2
	data := make(chan Data, bufSize)
	//Channel to communicate squared values between Stage 2 and Stage 3
	square := make(chan Square, bufSize)

	//Channel to communicate between Stage 3 and 2
	stage3Result := make(chan Result, bufSize)

	//Channel to communicate result between stage 2 and 1
	stage2Result := make(chan Result, bufSize)

	//Go routine to generate and send raw data via raw channel
	go func() {
		for i := 1; i <= 10; i++ {
			raw <- i
		}
	}()

	go Stage1(raw, data, stage2Result)
	go Stage2(data, square, stage3Result, stage2Result)
	go Stage3(square, stage3Result)

	time.Sleep(time.Second)
}

//Stage 1 receives raw data on raw channel, converts it to Data nd send
//to stage 2 via data channel.
func Stage1(raw <-chan int, data chan<- Data, stage2Result <-chan Result) {
	fmt.Println("Starting stage 1")
	for {
		select {
		case rawData := <-raw:
			fmt.Printf("Stage 1: received raw from downstream raw: %v\n", rawData)
			data <- Data(rawData)
		case result := <-stage2Result:
			fmt.Printf("Stage 1: received result from upstream stage 2: %v\n", result)
		}
	}
}

//Stage 2 receives Data from stage1, squares it and send Square to
//stage3 via square channel.
func Stage2(data <-chan Data, square chan<- Square, stage3Result <-chan Result, stage2Result chan<- Result) {
	fmt.Println("Starting Stage 2")
	for {
		select {
		case d := <-data:
			fmt.Printf("Stage 2: received data from downstream stage 1: %v\n", d)
			square <- Square(d * d)
		case res := <-stage3Result:
			fmt.Printf("Stage 2: received from upstream stage 3: %v\n", res)
			stage2Result <- res
		}
	}
}

//Stage 3 receives square channel from stage 2, multiplies result with
//3 and emits results to stdout.
func Stage3(square <-chan Square, stage3Result chan<- Result) {
	fmt.Println("Starting stage 3")
	for {
		select {
		case s := <-square:
			fmt.Printf("Stage 3: received data from downstream stage 2: %v\n", s)
			res := s * 3
			stage3Result <- Result(res)
			fmt.Printf("Stage 3: sent result data to downstream stage 2: %v\n", res)
		}
	}
}
