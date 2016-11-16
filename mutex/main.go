package main

import (
	"fmt"
)

func main() {
	mutex := make(chan bool, 1)

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			mutex <- true
			go func() {
				fmt.Printf("%d + %d = %d\n", i, j, i+j)
				<-mutex
			}()
		}
	}

	fmt.Scanln()
}
