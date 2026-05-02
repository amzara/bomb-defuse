package main

import (
	"bomb-defuse/internal/services/bomb"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan string)
	c4 := bomb.NewBomb(ch)
	go c4.Run() //goroutine
	defer c4.Close()

	ch <- "STATUS"

	ch <- "PLANT"
	time.Sleep(1 * time.Second)
	hackPin(ch)

	// ch <- "DEFUSE 1 2 3 4"
	time.Sleep(2 * time.Second)
	fmt.Println("After goroutines")
	ch <- "STATUS"

	time.Sleep(100 * time.Second)

	select {} // keep alive
}

func hackPin(c chan string) {

	const numJobs = 10000
	jobs := make(chan int, numJobs)
	// results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, c)
	}

	for i := 0; i < numJobs; i++ {
		jobs <- i
	}

}

func worker(id int, jobs <-chan int, bomb chan string) {
	for range jobs {
		sendDefuseMessage(bomb)
	}
}

func sendDefuseMessage(c chan string) {
	fmt.Println("Trying to hack!")
	pin := make([]int, 4)
	for i := 0; i < 4; i++ {
		pin[i] = rand.Intn(10)
	}

	message := fmt.Sprintf("DEFUSE %d %d %d %d", pin[0], pin[1], pin[2], pin[3])
	fmt.Printf(message)

	c <- message
}

// spawn worker pool to spam the channel with 4 digit combinations
// goal is to rate limit the spam
// can learn how to manually write worker pool pattern and also how to write rate limiter (albeit channel)
