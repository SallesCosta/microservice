package main

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/sallescosta/fullProject/account"
	"github.com/tinrab/retry"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

func main() {
	var cfg Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var r account.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = account.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()

	port := 8080
	log.Printf("Listening on port %d...", port)
	s := account.NewService(r)

	log.Fatal(account.ListeningGRPC(s, port))
}
