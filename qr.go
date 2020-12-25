package main

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"log"
) 

func main() {
	var img string

	fmt.Println("Generador de código QR")
	fmt.Println("inserte el link a convertir")
	fmt.Scanln(&img)

	err := qrcode.WriteFile(img, qrcode.Medium, 256, "qr.png")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Convertido exitosamente")
	}
}
