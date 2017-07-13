package boards

import (
	"math/rand"
	"time"
)

//Random values board
type Board struct {
	MaxX int
	MaxY int
}

//New random board
func New() Board {
	rand.Seed(time.Now().UnixNano())
	return Board{rand.Intn(10), rand.Intn(10)}
}
