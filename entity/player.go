package entity

import (
	"fmt"

	"github.com/leviceccato/clarity/component"
)

func NewPlayerEntity() (*Entity, error) {
	e := NewEntity()
	e.Position = &component.PositionComponent{X: 10, Y: 10}
	e.Size = &component.SizeComponent{Width: 10, Height: 10}
	e.Playable = &component.PlayableComponent{}
	appearance, err := component.NewAppearanceComponent("sprite/player.png", "sprite/player.json")
	if err != nil {
		return e, fmt.Errorf("creating appearance component: %s", err)
	}
	e.Appearance = appearance
	return e, nil
}
