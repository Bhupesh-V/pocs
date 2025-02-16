// Copyright 2021 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"math"
	"math/rand"
	"os"
	"time"

	constants "go-game/constants"
	"go-game/spaceship"
	"go-game/stars"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

const radius = 100.0

type Game struct {
	stars                    [constants.StarsCount]stars.Star
	spaceship                *spaceship.Spaceship
	prevCursorX, prevCursorY int
	audioPlayer              *audio.Player
}

func NewGame() *Game {
	g := &Game{}
	// Create an audio context.
	audioContext := audio.NewContext(constants.SampleRate)

	// Open the audio file (replace with your audio file path).
	audioFile, err := os.OpenFile(constants.SoundTrack, os.O_RDONLY, 0)
	if err != nil {
		return nil
	}
	// // Decode the audio file.
	audioDecoder, err := mp3.DecodeWithSampleRate(constants.SampleRate, audioFile)
	if err != nil {
		return nil
	}

	// Create an audio player for the decoded audio.
	g.audioPlayer, err = audioContext.NewPlayer(audioDecoder)
	if err != nil {
		return nil
	}

	for i := 0; i < constants.StarsCount; i++ {
		g.stars[i].Init()
	}

	g.spaceship = spaceship.NewSpaceship()

	// Play background audio
	g.audioPlayer.SetVolume(0.8)

	if err := g.audioPlayer.Rewind(); err != nil {
		log.Fatal(err)
	}
	g.audioPlayer.Play()

	return g
}

func (g *Game) Update() error {
	x, y := ebiten.CursorPosition()
	// Update the spaceship's position
	// Calculate the angle between the spaceship and the cursor
	// Calculate the angle based on the cursor's movement
	dx := float64(x - g.prevCursorX)
	dy := float64(y - g.prevCursorY)
	angle := math.Atan2(dy, dx)

	// Update the spaceship's position
	g.spaceship.X = float64(x) + radius*math.Cos(angle)
	g.spaceship.Y = float64(y) + radius*math.Sin(angle)

	// Store the current cursor position for the next frame
	g.prevCursorX = x
	g.prevCursorY = y

	for i := 0; i < constants.StarsCount; i++ {
		g.stars[i].Update(float32(x*constants.Scale), float32(y*constants.Scale))
	}

	// Check if the audio is playing, and if not, rewind and play it
	if !g.audioPlayer.IsPlaying() {
		if err := g.audioPlayer.Rewind(); err != nil {
			log.Fatal(err)
		}
		g.audioPlayer.Play()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.spaceship.Draw(screen)

	for i := 0; i < constants.StarsCount; i++ {
		g.stars[i].Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return constants.ScreenWidth, constants.ScreenHeight
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(constants.ScreenWidth, constants.ScreenHeight)
	ebiten.SetWindowTitle("Stars (Ebitengine Demo)")

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
