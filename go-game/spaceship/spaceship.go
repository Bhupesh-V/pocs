package spaceship

import (
	"go-game/constants"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Define the spaceship struct
type Spaceship struct {
	X, Y  float64
	Image *ebiten.Image
}

// Initialize the spaceship
func NewSpaceship() *Spaceship {
	img, _, err := ebitenutil.NewImageFromFile(constants.Spaceship)
	if err != nil {
		log.Fatal(err)
	}

	return &Spaceship{X: 12, Y: 12, Image: img}
}

// Add a method to update the spaceship's position
func (s *Spaceship) Update(x, y float64) {
	s.X = x
	s.Y = y
}

// Add a method to draw the spaceship
func (s *Spaceship) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(s.X, s.Y)
	screen.DrawImage(s.Image, opts)
}
