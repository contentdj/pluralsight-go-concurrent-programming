package main

import "fmt"

func main() {
	mutex := make(chan bool, 1)

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			for z := 0; z < 2; z++ {
				mutex <- true
				if z%2 == 0 {
					go func() {
						fmt.Printf("%d + %d = %d\n", i, j, i+j)
						<-mutex
					}()
				} else {
					<-mutex
				}

			}
		}
	}

	fmt.Scanln()
}
