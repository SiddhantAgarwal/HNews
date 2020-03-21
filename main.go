package main

import (
	"HNews/internal/app"
	"log"
	"os"
)

func main() {
	if err := app.NewHNewsApp().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
