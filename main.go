package main

import (
	"flag"
	"log"

	"github.com/buitanlan/redis_go/config"
	"github.com/buitanlan/redis_go/server"
)

func setupFlags() {
	flag.StringVar(&config.Host, "host", "0.0.0.0", "host for the server")
	flag.IntVar(&config.Port, "port", 7379, "port")
	flag.Parse()
}

func main() {
	setupFlags()
	log.Print("Starting go-redis")
	server.RunSyncTCPServer()
}
