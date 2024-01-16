package application

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"link_shortener/util"
	"log"
	"net/http"
	"time"
)

type App struct {
	router http.Handler
	rdb    *redis.Client
}

func New() *App {

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Error loading app.env file" + err.Error())
	}

	app := &App{
		rdb: redis.NewClient(&redis.Options{
			Addr:     "redis:" + config.RedisPort,
			Password: "",
			DB:       0,
		}),
	}

	app.loadRoutes()

	return app
}

func (a *App) Start(ctx context.Context) error {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Error loading app.env file" + err.Error())
	}

	server := &http.Server{
		Addr:    ":" + config.AppPort,
		Handler: a.router,
	}

	err = a.rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to connect to redis: %w", err)
	}

	defer func() {
		if err := a.rdb.Close(); err != nil {
			fmt.Println("failed to close redis", err)
		}
	}()

	fmt.Println("Starting server")

	ch := make(chan error, 1)

	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}
}
