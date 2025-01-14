package asset

import (
	"bytes"
	"embed"
	"fmt"
	"image"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// Embed assets folder inside the executable
//go:embed **/*
var FS embed.FS

// Simple icon loading util
func LoadIcon(path string) (image.Image, error) {
	iconBytes, err := FS.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading icon file: %w", err)
	}

	icon, _, err := image.Decode(bytes.NewReader(iconBytes))
	if err != nil {
		return nil, fmt.Errorf("decoding icon file: %w", err)
	}

	return icon, nil
}

func LoadFont(path string) (font.Face, error) {
	bytes, err := FS.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("loading font file: %s", err)
	}

	fontfile, err := opentype.Parse(bytes)
	if err != nil {
		return nil, fmt.Errorf("parsing font file: %s", err)
	}

	face, err := opentype.NewFace(fontfile, &opentype.FaceOptions{
		Size:    11,
		DPI:     72,
		Hinting: font.HintingNone,
	})
	if err != nil {
		return nil, fmt.Errorf("creating font face: %s", err)
	}

	return face, nil
}
