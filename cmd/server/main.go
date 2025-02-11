package main

import (
	"kai_security/internal/actor"
	"kai_security/internal/config"
	"kai_security/internal/logger"
)

func main() {
	_ = config.Get()
	_ = logger.Get()

	server := actor.NewHttpServer()

	server.Start()
}
