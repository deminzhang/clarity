package main

import (
	"bytes"
	"fmt"
	"image"
	"io/ioutil"
	"os"

	"github.com/tc-hib/winres"
)

// This package uses tc-hib/winres to create a resource set that gets embedded
// into the final executable for Windows. This allows embedding the application
// icon. This file should not be included in the build and instead should be
// run with `go run ./windows`.
func main() {
	rs := winres.ResourceSet{}

	// Add icons
	icon32, err := loadIcon("../asset/icon/icon_32.png")
	if err != nil {
		fmt.Printf("loading icon_32: %s", err)
	}
	icon16, err := loadIcon("../asset/icon/icon_16.png")
	if err != nil {
		fmt.Printf("loading icon_16: %s", err)
	}
	icon, err := winres.NewIconFromImages([]image.Image{icon32, icon16})
	if err != nil {
		fmt.Printf("creating icon from files: %s", err)
	}
	rs.SetIcon(winres.Name("APPICON"), icon)

	// Output syso file for inclusion in executable
	out, err := os.Create("../rsrc_windows_amd64.syso")
	if err != nil {
		fmt.Printf("creating syso file: %s", err)
	}
	defer out.Close()
	err = rs.WriteObject(out, winres.ArchAMD64)
	if err != nil {
		fmt.Printf("writing resource to syso file: %s", err)
	}
}

// Simple icon loading utility
func loadIcon(path string) (image.Image, error) {
	iconBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading icon file: %s", err)
	}
	icon, _, err := image.Decode(bytes.NewReader(iconBytes))
	if err != nil {
		return nil, fmt.Errorf("decoding icon file: %s", err)
	}
	return icon, nil
}
