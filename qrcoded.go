package main

import (
	"image"
	"image/png"
	"io"
	"os"
)

func GenerateQRCode(w io.Writer, code string) {
	img := image.NewNRGBA(image.Rect(0, 0, 21, 21))
	//buf := new(bytes.Buffer)
	_ = png.Encode(w, img)
	//return buf.Bytes()
}

func main() {
	file, _ := os.Create("qrcode.png")
	defer file.Close()

	GenerateQRCode(file, "555-2368")
	//ioutil.WriteFile("qrcode.png", qrcode, 0644)
}
