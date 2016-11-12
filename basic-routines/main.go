package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	fmt.Println("MAX Procs", runtime.GOMAXPROCS(1))
	go func() {
		for index := 0; index < 100; index++ {
			fmt.Println("Hello")
			time.Sleep(10 * time.Microsecond)
		}

	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Go")
			time.Sleep(10 * time.Microsecond)
		}

	}()

	fmt.Println("Erdal was here!")

	time.Sleep(1 * time.Second)
}
