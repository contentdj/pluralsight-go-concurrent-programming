package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	in := make(chan int)
	go generate(in)

	for {
		prime := <-in
		fmt.Println(prime)
		out := make(chan int)
		go filter(in, out, prime)
		in = out
	}
}

func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func filter(in, out chan int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}
