package main

import "testing"

func TestGenerateQRCodeResult(t *testing.T) {
	result := GenerateQRCode("555-2368")

	if result == nil {
		t.Errorf("Generated QRCode nil")
	}

	if len(result) == 0 {
		t.Errorf("Generated QRCode has no data")
	}
}
