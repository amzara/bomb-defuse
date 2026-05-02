package bomb

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type State string

const (
	Idle     State = "idle"
	Planted  State = "planted"
	Defused  State = "defused"
	Exploded State = "exploded"
)

type Bomb struct {
	state    State
	attempts int
	timer    *time.Timer
	ticker   *time.Ticker
	beeping  bool
	pin      []int
	ch       chan string
	cmdCh    chan string
}

func NewBomb(c chan string) *Bomb {
	return &Bomb{
		state:   Idle,
		ch:      c,
		cmdCh:   make(chan string, 1),
		beeping: true,
		pin:     []int{1, 2, 3, 4},
	}
}

func (b *Bomb) Run() {

	go func() {
		limiter := time.NewTicker(time.Second / 10)
		defer limiter.Stop()

		for cmd := range b.ch {
			<-limiter.C
			b.cmdCh <- cmd
		}
		close(b.cmdCh)
	}()

	for {
		select {
		case cmd := <-b.cmdCh:
			b.handle(cmd)

		case <-b.tickChan():
			if b.beeping && b.state == Planted {
				fmt.Println("BEEP")
			}
		}
	}

} /

func (b *Bomb) handle(cmd string) {
	parts := strings.Fields(cmd)
	switch parts[0] {
	case "PLANT":
		b.Plant(10 * time.Second)
	case "STATUS":
		b.Status()
	case "PIN":
		b.Pin()
	case "BEEP":
		b.beeping = !b.beeping
	case "DEFUSE":
		b.Defuse(parts[1:])
	case "TIMEOUT":
		b.explode()
	}
}

func (b *Bomb) GenerateRandomPin() {

	pin := make([]int, 4)
	for i := 0; i < 4; i++ {
		pin[i] = rand.Intn(10)
	}
	b.pin = pin
}

func (b *Bomb) Status() {
	fmt.Println(b.state)

}

func (b *Bomb) Defuse(pin []string) {
	if len(pin) != 4 {
		fmt.Println("Usage: DEFUSE <d> <d> <d> <d>")
		return
	}
	for i := 0; i < 4; i++ {
		digit, err := strconv.Atoi(pin[i])
		if err != nil {
			fmt.Println("Invalid digits")
			return
		}
		if digit != b.pin[i] {
			fmt.Println("Invalid PIN")
			return
		}

	}

	b.stopTicker()
	if b.timer != nil {
		b.timer.Stop()
	}
	b.state = Defused
	fmt.Println("Bomb has been defused")

}

func (b *Bomb) Pin() {
	fmt.Println(b.pin)
}

func (b *Bomb) Plant(d time.Duration) {
	b.state = Planted

	b.GenerateRandomPin()

	b.timer = time.AfterFunc(d, func() {
		b.cmdCh <- "TIMEOUT"
	})

	b.ticker = time.NewTicker(time.Second)

	fmt.Printf("Bomb has been planted! ")
}

func (b *Bomb) stopTicker() {
	if b.ticker != nil {
		b.ticker.Stop()
		b.ticker = nil
	}
}

func (b *Bomb) explode() {
	b.state = Exploded
	b.stopTicker()
	fmt.Println("Boom")
	os.Exit(1)
}

func (b *Bomb) Close() {
	if b.timer != nil {
		b.timer.Stop()
	}
	close(b.ch)
}

func (b *Bomb) tickChan() <-chan time.Time {
	if b.ticker == nil {
		return nil
	}
	return b.ticker.C
}
