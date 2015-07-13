package main

import (
	"bytes"
	"errors"
	"image/png"
	"testing"
)

func TestGenerateQRCodeGeneratesPNG(t *testing.T) {
	buf := new(bytes.Buffer)
	GenerateQRCode(buf, "555-2368", Version(1))
	if buf.Len() == 0 {
		t.Errorf("Generated QRCode has no data")
	}

	_, err := png.Decode(buf)
	if err != nil {
		t.Errorf("Generated QRCode not a png: %s", err)
	}
}

type ErrorWriter struct{}

func (e *ErrorWriter) Write(b []byte) (int, error) {
	return 0, errors.New("Expected error")
}

func TestGenerateQRCodePropagatesErrors(t *testing.T) {
	w := new(ErrorWriter)
	err := GenerateQRCode(w, "555-2368", Version(1))
	if err == nil || err.Error() != "Expected error" {
		t.Errorf("Expected error ,but got %v", err)
	}
}

func TestVersionDeterminesSize(t *testing.T) {
	table := []struct {
		version  int
		expected int
	}{
		{1, 21},
		{2, 25},
		{6, 41},
		{7, 45},
		{14, 73},
		{40, 177},
	}

	for _, test := range table {
		buf := new(bytes.Buffer)
		GenerateQRCode(buf, "555-2368", Version(test.version))
		img, _ := png.Decode(buf)

		if width := img.Bounds().Dx(); width != test.expected {
			t.Errorf("Version %2d, expected %3d but got %3d", test.version, test.expected, width)
		}
	}
}
