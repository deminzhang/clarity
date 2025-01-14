package system

import (
	"fmt"
	"image/color"

	"github.com/leviceccato/clarity/game"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type drawSystem struct {
	System
}

func NewDrawSystem(s *game.State) *drawSystem {
	sys := &drawSystem{}
	sys.Components = []string{
		"Appearance",
		"Position",
		"Size",
	}
	return sys
}

func (sys *drawSystem) Load(s *game.State) {}

func (sys *drawSystem) Update(s *game.State) {}

func (sys *drawSystem) Draw(s *game.State, screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0x00, 0x40, 0x80, 0xff})

	for _, e := range sys.Entities {

		// Draw image
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(e.Position.X, e.Position.Y)

		screen.DrawImage(
			e.Appearance.Image.SubImage(*e.Appearance.Frames[e.Appearance.Frame]).(*ebiten.Image),
			options,
		)

		// No text? nothing to do
		if e.Text == nil {
			continue
		}

		// Draw lines of text
		for i, line := range e.Text.Lines {
			x := line.X + e.Text.Padding
			y := (e.Size.Height / 2) - float64((e.Text.LineHeight*len(e.Text.Lines))/2) + float64((i+1)*e.Text.LineHeight)

			text.Draw(screen, line.Content, e.Text.Font, int(e.Position.X+x), int(e.Position.Y+y), e.Text.Color)
		}
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"TPS: %0.2f\nFPS: %0.2f",
		ebiten.CurrentTPS(),
		ebiten.CurrentFPS(),
	))
}

func (sys *drawSystem) Enter(s *game.State) {}

func (sys *drawSystem) Exit(s *game.State) {}
