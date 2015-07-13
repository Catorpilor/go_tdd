package main

import (
	"bytes"
	"image/png"
	"testing"
)

func TestGenerateQRCodeResult(t *testing.T) {
	result := GenerateQRCode("555-2368")

	if result == nil {
		t.Errorf("Generated QRCode nil")
	}

	if len(result) == 0 {
		t.Errorf("Generated QRCode has no data")
	}
}

func TestGenerateQRCodeGeneratesPNG(t *testing.T) {
	result := GenerateQRCode("555-2368")

	bf := bytes.NewBuffer(result)
	_, err := png.Decode(bf)
	if err != nil {
		t.Errorf("Generated QRCode not a png: %s", err)
	}
}
