package boards

//Random values board
type Board struct {
	MaxX int
	MaxY int
}

//New random board
func New() Board {
	// rand.Seed(10)
	// return Board{rand.Intn(10), rand.Intn(10)}
	return Board{3, 3}
}
