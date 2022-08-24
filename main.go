package main

import (
	"command-line-application/app"
	"log"
	"os"
)

func main() {
	// Create new app
	application := app.GenerateApp()

	// Method that runs the application. Receives as param os.Args, which is used to recognize command lines
	// It can throw an error, that needs to be treated
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
