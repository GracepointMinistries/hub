package main

import (
	"log"

	"github.com/GracepointMinistries/hub/actions"
)

//go:generate sqlboiler -c sqlboiler.toml psql --struct-tag-casing camel --no-hooks
//go:generate swagger generate spec -o swagger.json

func main() {
	if err := actions.App().Serve(); err != nil {
		log.Fatal(err)
	}
}
