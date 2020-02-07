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

	sconfig := server.NewConfig()
	if err := sconfig.DecodeJSONConf(configPath); err != nil {
		log.Fatal(err)
	}

	s := server.New(sconfig)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
