package game

import (
	"errors"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/sergiosegrera/chess/board"
	"image"
	"image/color"
)

var (
	BLACK = color.RGBA{0x00, 0x00, 0x00, 0xff}
	WHITE = color.RGBA{0xff, 0xff, 0xff, 0xff}
)

type Game struct {
	Board      *board.Board
	Sprites    []*ebiten.Image
	Selector   *Selector
	Background *ebiten.Image
}

func NewGame() (*Game, error) {
	board := board.NewBoard()

	sheet, _, err := ebitenutil.NewImageFromFile("./assets/set.png", ebiten.FilterDefault)

	var sprites []*ebiten.Image
	sw, sh := sheet.Size()
	for y := 0; y < sh/8; y++ {
		for x := 0; x < sw/8; x++ {
			sprites = append(sprites, sheet.SubImage(image.Rect(x*8, y*8, x*8+8, y*8+8)).(*ebiten.Image))
		}
	}

	selector, err := NewSelector(board)
	if err != nil {
		return nil, err
	}

	background, _, err := ebitenutil.NewImageFromFile("./assets/board.png", ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}

	return &Game{
		Board:      board,
		Sprites:    sprites,
		Selector:   selector,
		Background: background,
	}, nil
}

func (g *Game) Update(image *ebiten.Image) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New("Escape called, game exited")
	}
	g.Selector.Update()
	return nil
}

func (g *Game) DrawPieces(screen *ebiten.Image) {
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			position := &board.Position{X: x, Y: y}
			if !g.Board.IsEmpty(position) {
				piece := g.Board.GetPiece(position)
				sprite := piece.Type()
				if piece.Color() == board.BLACK {
					sprite += 6
				}
				options := &ebiten.DrawImageOptions{}
				options.GeoM.Translate(float64(y)*8, float64(x)*8)
				screen.DrawImage(g.Sprites[sprite], options)
			}
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.Background, nil)
	g.DrawPieces(screen)
	g.Selector.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 64, 64
}
