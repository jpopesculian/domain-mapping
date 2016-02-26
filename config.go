package main

import (
	"flag"
	"os"
	"strconv"
)

type Config struct {
	Port uint64
}

var config Config

func InitConfig() {
	defaultPort, err := strconv.ParseUint(os.Getenv("PORT"), 0, 16)
	if err != nil {
		defaultPort = 5002
	}
	port := flag.Uint64("p", defaultPort, "Port to serve on")

	flag.Parse()

	config = Config{
		*port,
	}
}
