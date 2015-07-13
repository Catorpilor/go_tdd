package main

import (
	"image"
	"image/png"
	"io"
	"log"
	"os"
)

type Version int8

func (v Version) MappedSize() int {
	return 4*int(v) + 17
}

func GenerateQRCode(w io.Writer, code string, version Version) error {
	size := version.MappedSize()
	img := image.NewNRGBA(image.Rect(0, 0, size, size))
	//buf := new(bytes.Buffer)
	return png.Encode(w, img)
	//return buf.Bytes()
}

func main() {
	file, err := os.Create("qrcode.png")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = GenerateQRCode(file, "555-2368", Version(1))
	if err != nil {
		log.Fatal(err)
	}
	//ioutil.WriteFile("qrcode.png", qrcode, 0644)
}
