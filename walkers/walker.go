package walkers

type Walker struct {
	AtX         int
	AtY         int
	Orientation string
}

func Walk(w *Walker, d string) {
	switch d {
	case "w":
		up(w)
	case "W":
		up(w)
	case "a":
		left(w)
	case "A":
		left(w)
	case "d":
		right(w)
	case "D":
		right(w)
	case "s":
		down(w)
	case "S":
		down(w)
	}
}

func up(w *Walker) {
	w.AtY++
}
func down(w *Walker) {
	w.AtY--
}
func right(w *Walker) {
	w.AtX++
}
func left(w *Walker) {
	w.AtX--
}
