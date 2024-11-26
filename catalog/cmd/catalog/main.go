package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/sallescosta/fullProject/catalog"
	"github.com/tinrab/retry"
	"log"
	"time"
)

type Config struct {
	DatabaseURL string `json:"DATABASE_URL"`
}

func main() {
	var cfg Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var r catalog.Repository

	retry.ForeverSleep(2*time.Second, func(_, int) (err error) {
		r, err = catalog.NewElasticRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()

	port := 8080
	log.Printf("Listening on port %d...", port)
	s := catalog.NewService(r)

	log.Fatal(catalog.ListeningGRPC(s, port))
}
