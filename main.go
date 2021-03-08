package main

import (
	"fmt"
	"log"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/UoYMathSoc/2020-site/structs"
	"github.com/joho/godotenv"
	"github.com/stretchr/graceful"
)

func main() {
	_ = godotenv.Load()

	log.SetFlags(log.Llongfile)

	config := &structs.Config{}
	_, err := toml.DecodeFile("config.toml", config)
	if err != nil {
		log.Fatal(err)
	}

	s, err := NewServer(config)
	if err != nil {
		log.Fatal(err)
	}

	l := fmt.Sprintf("%s:%d", config.Server.Address, config.Server.Port)
	log.Printf("Listening on: %s", l)
	graceful.Run(l, time.Duration(config.Server.Timeout), s)
}
