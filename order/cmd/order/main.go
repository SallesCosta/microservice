package main

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/sallescosta/fullProject/order"
	"github.com/tinrab/retry"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
	AccountURL  string `envconfig:"ACCOUNT_SERVICE_URL"`
	CatalogURL  string `envconfig:"CATALOG_SERVICE_URL"`
}

func main() {
	var cfg Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var r order.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = order.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()

	port := 8080

	log.Printf("Listening on port %d...", port)

	s := order.NewService(r)
	log.Fatal(order.ListenGRPC(s, cfg.AccountURL, cfg.CatalogURL, port))
}
