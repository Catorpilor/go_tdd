package main

import (
	"bytes"
	"image/png"
	"testing"
)

func TestGenerateQRCodeGeneratesPNG(t *testing.T) {
	buf := new(bytes.Buffer)
	GenerateQRCode(buf, "555-2368")
	if buf.Len() == 0 {
		t.Errorf("Generated QRCode has no data")
	}

	_, err := png.Decode(buf)
	if err != nil {
		t.Errorf("Generated QRCode not a png: %s", err)
	}
}
