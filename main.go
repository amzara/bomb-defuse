package main

import (
	"bomb-defuse/internal/services/bomb"
	"fmt"
	"time"
)

func main() {

	c4 := bomb.NewBomb()
	// fmt.Printf(c4.LookAtBomb())
	c4.Plant(5 * time.Second)
	// time.Sleep(2 * time.Second)
	fmt.Printf(c4.LookAtBomb())
	c4.Defuse()

	// app := server.New()
	// app.StartServer()
	// time.Sleep(3 * time.Second)

}

func plantBomb(c chan string) {
	order := <-c
	fmt.Println(order)
	explosion := "boom"
	c <- explosion

}
