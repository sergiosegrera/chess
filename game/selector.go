package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/sergiosegrera/chess/board"
	"image"
)

type Selector struct {
	sprites      []*ebiten.Image
	board        *board.Board
	x            int
	y            int
	fromPosition *board.Position
	toPosition   *board.Position
}

func NewSelector(board *board.Board) (*Selector, error) {
	sheet, _, err := ebitenutil.NewImageFromFile("./assets/selector.png", ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	var sprites []*ebiten.Image
	sprites = append(sprites, sheet.SubImage(image.Rect(0, 0, 8, 8)).(*ebiten.Image))
	sprites = append(sprites, sheet.SubImage(image.Rect(8, 0, 16, 8)).(*ebiten.Image))
	return &Selector{
		sprites: sprites,
		board:   board,
	}, err
}

func (s *Selector) Update() error {
	// TODO: Selector movement to screen limits
	// Movement
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		s.x += 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		s.x -= 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		s.y -= 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		s.y += 1
	}
	// Selection
	if inpututil.IsKeyJustPressed(ebiten.KeyJ) {
		if s.fromPosition == nil {
			s.fromPosition = &board.Position{X: s.x, Y: s.y}
		} else {
			s.board.Move(&board.Position{s.fromPosition.Y, s.fromPosition.X}, &board.Position{X: s.y, Y: s.x})
			s.fromPosition = nil
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyK) {
		s.fromPosition = nil
		s.toPosition = nil
	}
	return nil
}

func (s *Selector) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(s.x)*8, float64(s.y)*8)
	screen.DrawImage(s.sprites[0], options)
	if s.fromPosition != nil {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(s.fromPosition.X)*8, float64(s.fromPosition.Y)*8)
		screen.DrawImage(s.sprites[1], options)
	}
}
