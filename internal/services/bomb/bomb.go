package bomb

import (
	"fmt"
	"math/rand/v2"
	"os"
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
}

func NewBomb() *Bomb {
	return &Bomb{
		state: Idle,
	}
}

// https://go.dev/tour/methods/8

func (b *Bomb) Plant(duration time.Duration) {
	b.state = Planted
	fmt.Println("Beep beep beep")
	b.timer = time.AfterFunc(duration, func() {
		fmt.Println("BOOOMMMMMMMMMM")
		os.Exit(1)
	})
}

//

func (b Bomb) LookAtBomb() string {
	var bombStateString string
	switch b.state {
	case Idle:
		return "The bomb is idle"
	case Planted:
		return "The bomb has been planted"
	case Defused:
		return "The bomb has been defused"
	case Exploded:
		return "The bomb has exploded"
	}

	return bombStateString
}

func GenerateRandomPin(l int) []int {

	pin := make([]int, l)
	for i := 0; i < l; i++ {
		pin[i] = rand.IntN(10) // generate random number from 0 to 10

	}

	return pin

}

//learn what is make
// why does slices need make
// slices are dynamic arrays
// relate to channels, they need make too
// https://www.youtube.com/watch?v=FcdTJbIz5p0
