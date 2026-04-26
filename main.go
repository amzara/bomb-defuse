package main

import (
	"fmt"
	"net/http"
	"bomb-defuse/internal/server"
	"bomb-defuse/internal/services/bomb"

	"time"
)

func main() {

	c4 := bomb.NewBomb()
	fmt.Printf(c4.LookAtBomb())
	c4.Plant(5 * time.Second)
	time.Sleep(2 * time.Second)
	fmt.Printf(c4.LookAtBomb())

	generatedPin := bomb.GenerateRandomPin(4)
	fmt.Println(generatedPin)

	app := server.New()
	app.StartServer()
	time.Sleep(3 * time.Second)

	resp, err := http.Get("http://127.0.0.1:1234/healthCheck")
	if err != nil {
		fmt.Printf("Something exploded %v", err)

		return
	}

	defer resp.Body.Close() //

	fmt.Println(resp)

	ch := make(chan string)
	go plantBomb(ch)
	fmt.Println("Planting the bomb now")
	time.Sleep(5 * time.Second)
	pin := "12345"
	fmt.Println("Pressing c4 button now")

	ch <- pin
	x := <-ch
	fmt.Println(x)
	fmt.Println("Terrorist win")

}

// ch <- x // send x to channel
// x := <- ch //s receive value from channel and assign it's value
// current assignment: make a goroutine that waits for an entry, then explode

func plantBomb(c chan string) {
	order := <-c
	fmt.Println(order)
	explosion := "boom"
	c <- explosion

}

//csgo bomb defuse
