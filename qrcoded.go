package main

import (
	"image"
	"image/png"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Create("QRCode.png")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	err = GenerateQRCode(file, "0797280650", 2)
	if err != nil {
		log.Fatal(err)
	}
}
func GenerateQRCode(w io.Writer, code string, version Version) error {
	size := version.PatternSize()
	img := image.NewNRGBA(image.Rect(0, 0, size, size))
	return png.Encode(w, img)
}

type Version int8

func (version Version) PatternSize() int {
	return 4*int(version) + 17
}
