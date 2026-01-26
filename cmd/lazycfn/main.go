package main

import (
	"log"

	"github.com/myuron/lazycfn/pkg/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatalln(err)
	}
	defer a.Close()
	if err := a.Run(); err != nil {
		log.Fatalln(err)
	}
}
