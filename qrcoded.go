package main

import "io/ioutil"

func GenerateQRCode(code string) []byte {
	return nil
}

func main() {
	qrcode := GenerateQRCode("555-2368")
	ioutil.WriteFile("qrcode.png", qrcode, 0644)
}
