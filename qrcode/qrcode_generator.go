package qrcode

import (
	"os"

	"github.com/skip2/go-qrcode"
)

func GenerateQRCode(url string) []byte {
	png, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		// Handle error
		return nil

	}

	os.WriteFile("./qrcode.png", png, os.FileMode(0644))

	return png
}
