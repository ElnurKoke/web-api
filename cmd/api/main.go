package main

import (
	"log"

	"github.com/ElnurKoke/web-api.git/internal/apiserver"
)

func main() {
	config, err := apiserver.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
