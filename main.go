package main

import (
	"fmt"

	"github.com/leviceccato/clarity/world"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	windowWidth  = 1280
	windowHeight = 720
)

type game struct {
	state *state
}

func (g *game) Update() error {
	g.state.update()
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	g.state.draw(screen)
}

func (g *game) Layout(_, _ int) (int, int) {
	return windowWidth, windowHeight
}

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Clarity")
	mainState := newState()
	mainState.windowWidth = windowWidth
	mainState.windowHeight = windowHeight
	startWorld, err := world.NewStartWorld(mainState)
	if err != nil {
		fmt.Println(err)
	}
	mainState.loadWorld(startWorld)
	mainState.activateWorlds([]string{"start"})
	err = ebiten.RunGame(&game{
		state: mainState,
	})
	if err != nil {
		fmt.Println(err)
	}
}