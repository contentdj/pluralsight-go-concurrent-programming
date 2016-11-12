package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	numOfRequests := 10

	// Sequantial
	// for i := 1; i <= numOfRequests; i++ {
	// 	resp, _ := http.Get("https://jsonplaceholder.typicode.com/posts/" + strconv.Itoa(i))
	// 	defer resp.Body.Close()
	// 	post := new(Post)

	// 	json.NewDecoder(resp.Body).Decode(post)
	// 	fmt.Println(post.Id)
	// }

	// Concurrent
	numOfRequestsCompleted := 0
	for i := 1; i <= numOfRequests; i++ {
		go func(i int) {
			resp, _ := http.Get("https://jsonplaceholder.typicode.com/posts/" + strconv.Itoa(i))
			defer resp.Body.Close()
			post := new(Post)

			json.NewDecoder(resp.Body).Decode(post)
			numOfRequestsCompleted++
			fmt.Println(post.Id)
		}(i)
	}

	for numOfRequestsCompleted < numOfRequests {
		time.Sleep(10 * time.Millisecond)
	}

	elapsed := time.Since(start)
	fmt.Printf("Execution time is %s\n", elapsed)

}

type Post struct {
	UserId int
	Id     int
	Title  string
	Body   string
}
