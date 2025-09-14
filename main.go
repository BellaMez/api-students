package main

import (
	"github.com/BellaMez/api-students/api"
	"github.com/rs/zerolog/log"
)

func main() {
	server := api.NewServer()

	server.Configureoutes()
	server.Start()

	if err := server.Start(); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server:")
	}

}
