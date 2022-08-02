package main

import (
	"bytes"
	"image/png"
	"testing"
)

func TestDetermineVersion(t *testing.T) {
	buffer := new(bytes.Buffer)
	err := GenerateQRCode(buffer, "0797280650", Version(1))
	if err != nil {
		return
	}

	img, _ := png.Decode(buffer)
	if width := img.Bounds().Dx(); width != 21 {
		t.Errorf("Version 1 expected 21 but got %d", width)
	}
}
func TestVersionToDetermineSize(t *testing.T) {
	table := []struct {
		Version, expected int
	}{
		{1, 21},
		{2, 25},
		{6, 41},
		{7, 45},
		{14, 73},
		{40, 177},
	}
	for _, test := range table {
		buffer := new(bytes.Buffer)
		err := GenerateQRCode(buffer, "0797280650", Version(test.Version))
		if err != nil {
			return
		}
		img, _ := png.Decode(buffer)
		if width := img.Bounds().Dx(); width != test.expected {
			t.Errorf("Version %2d, expected %3d but got %3d", test.Version, test.expected, width)
		}
	}
}
