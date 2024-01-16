package main

import (
	"context"
	"fmt"
	"link_shortener/internal/application"
	"link_shortener/util"
	"log"
	"os"
	"os/signal"
)

func main() {
	_, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Error loading app.env file" + err.Error())
	}
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err = app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
