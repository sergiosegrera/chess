package game

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/sergiosegrera/chess/board"
)

type Game struct {
	reader *bufio.Reader
	board  *board.Board
}

func New() *Game {
	return &Game{
		reader: bufio.NewReader(os.Stdin),
		board:  board.NewBoard(),
	}
}

func (g *Game) Play() {
	for {
		g.board.Print()
		fmt.Print("> ")
		input, _ := g.reader.ReadString('\n')
		input = strings.Trim(input, "\n")
		valid, err := regexp.Match(`^[a-h][1-8]-[a-h][1-8]$`, []byte(input))
		if err != nil {
			fmt.Println(err.Error())
		}
		if !valid {
			fmt.Println("Invalid move")
		} else {
			from, to := LongAlgebraicToPositions(input)
			valid := g.board.Move(from, to)
			if !valid {
				fmt.Println("Invalid move")
			}
		}
	}
}
