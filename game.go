package EbitNew2

import (
	//E "EbitNew"
	//"log"

	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

func (g *Game) Update() error {

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	for _, s := range Draws {
		s.getOp().GeoM.Reset()
		s.getOp().GeoM.Translate(s.getPosition())
		screen.DrawImage(s.getImage(), s.getOp())
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
