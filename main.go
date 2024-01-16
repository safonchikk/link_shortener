package main

import (
	"context"
	"fmt"
	"link_shortener/internal"
	"link_shortener/internal/application"
	"os"
	"os/signal"
)

func main() {
	internal.InitVariables()
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
