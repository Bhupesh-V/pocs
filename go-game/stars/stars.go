package stars

import (
	"image/color"
	"math/rand"

	constants "go-game/constants"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Star struct {
	fromx, fromy, tox, toy, brightness float32
}

func (s *Star) Init() {
	s.tox = rand.Float32() * constants.ScreenWidth * constants.Scale
	s.fromx = s.tox
	s.toy = rand.Float32() * constants.ScreenHeight * constants.Scale
	s.fromy = s.toy
	s.brightness = rand.Float32() * 0xff
}

func (s *Star) Update(x, y float32) {
	s.fromx = s.tox
	s.fromy = s.toy
	s.tox += (s.tox - x) / 32
	s.toy += (s.toy - y) / 32

	s.brightness += 3
	if 0xff < s.brightness {
		s.brightness = 0xff
	}
	if s.fromx < 0 || constants.ScreenWidth*constants.Scale < s.fromx || s.fromy < 0 || constants.ScreenHeight*constants.Scale < s.fromy {
		s.Init()
	}
}

func (s *Star) Draw(screen *ebiten.Image) {
	c := color.RGBA{
		R: uint8(0xbb * s.brightness / 0xff),
		G: uint8(0xdd * s.brightness / 0xff),
		B: uint8(0xff * s.brightness / 0xff),
		A: 0xff,
	}
	vector.StrokeLine(screen, s.fromx/constants.Scale, s.fromy/constants.Scale, s.tox/constants.Scale, s.toy/constants.Scale, 1, c, true)
}
