package main

import (
	"context"
	"github.com/safonchikk/link_shortener/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
}
