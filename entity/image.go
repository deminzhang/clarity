package entity

import (
	"fmt"

	"github.com/leviceccato/clarity/component"
)

func NewImageEntity(path string) (*Entity, error) {
	e := NewEntity()
	e.Position = &component.PositionComponent{X: 0, Y: 0}
	e.Size = &component.SizeComponent{}
	appearance, err := component.NewAppearanceComponent(path, "")
	if err != nil {
		return e, fmt.Errorf("creating appearance component: %s", err)
	}
	e.Appearance = appearance
	return e, nil
}
