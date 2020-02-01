package main

import (
	"flag"
	"log"

	"github.com/MishaNiki/go-login-signup/internal/app/server"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.json", "path to config file")
}

func main() {
	flag.Parse()

	config := server.NewConfig()
	if err := config.DecodeJFile(configPath); err != nil {
		log.Fatal(err)
	}

	s := server.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
