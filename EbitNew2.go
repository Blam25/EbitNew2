// EbitNew2 project EbitNew2.go
package EbitNew2

import (
	//E "EbitNew"
	//"log"

	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var Images []Image2

type entity1 struct {
	*Position
	*Image
}

var Draws []Draw

func NewEntity1(image *ebiten.Image) *entity1 {
	new := entity1{}
	new.Position = NewPosition(50, 50)
	new.Image = NewImage(image)
	Images = append(Images, &new)
	Draws = append(Draws, &new)
	return &new
}

type Image2 interface {
	Draw()
}

type Image struct {
	Image *ebiten.Image
	Op    ebiten.DrawImageOptions
}

func (s *Image) getImage() *ebiten.Image {
	return s.Image
}

func (s *Image) getOp() *ebiten.DrawImageOptions {
	return &s.Op
}

type Draw interface {
	getPosition() (float64, float64)
	getImage() *ebiten.Image
	getOp() *ebiten.DrawImageOptions
}

func (s *Image) Draw() {
	println("hej")
}

func NewImage(image *ebiten.Image) *Image {
	new := Image{}
	//new.Entity = s
	new.Image = image
	Images = append(Images, &new)
	return &new
}

type Position struct {
	X float64
	Y float64
}

func (s *Position) getPosition() (float64, float64) {
	return s.X, s.Y
}

func NewPosition(x, y float64) *Position {
	new := Position{}
	new.X = x
	new.Y = y
	return &new
}
