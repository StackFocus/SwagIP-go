package main

import (
	"log"

	"github.com/StackFocus/swagip/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
