package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/raduq/mars3/boards"
	"github.com/raduq/mars3/walkers"
)

func main() {
	board := boards.New()
	walker := walkers.Walker{0, 0, "N"}
	listen(&board, &walker)
}

func listen(b *boards.Board, w *walkers.Walker) {
	snr := bufio.NewScanner(os.Stdin)
	text := ""
	for fmt.Println(text); snr.Scan(); fmt.Println(text) {
		line := snr.Text()
		if len(line) == 0 {
			break
		}
		walkers.Walk(w, line)

		err := accept(*b, *w)
		if err != nil {
			fmt.Println(err)
			return
		}
		pos := fmt.Sprintf("%d%d", w.AtX, w.AtY)
		fmt.Println(pos)
	}
	if err := snr.Err(); err != nil {
		if err != io.EOF {
			fmt.Fprintln(os.Stderr, err)
		}
	}

}

func accept(b boards.Board, w walkers.Walker) error {
	if w.AtX < 0 || w.AtX > b.MaxX {
		return errors.New("BOOOM > X")
	}
	if w.AtY < 0 || w.AtY > b.MaxY {
		return errors.New("BOOM > Y")
	}
	return nil
}
