package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/sergiosegrera/chess/game"
	"image"
	"log"
)

func main() {
	ebiten.SetWindowSize(512, 512)
	ebiten.SetWindowTitle("chess")

	_, icon, _ := ebitenutil.NewImageFromFile("./assets/icon128.png", ebiten.FilterDefault)
	ebiten.SetWindowIcon([]image.Image{icon})

	game, err := game.NewGame()
	if err != nil {
		log.Panic(err)
	}

	err = ebiten.RunGame(game)
	if err != nil {
		log.Panic(err)
	}
}
