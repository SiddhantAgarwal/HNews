package main

import (
	"log"
	"os"

	"github.com/SiddhantAgarwal/HNews/internal/app"
)

func main() {
	if err := app.NewHNewsApp().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
