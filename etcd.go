package main

import (
	"github.com/coreos/etcd/client"
	"log"
	"time"
)

var etcd client.KeysAPI

func InitEtcd() {
	cfg := client.Config{
		Endpoints:               []string{"http://127.0.0.1:2379"},
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	etcd = client.NewKeysAPI(c)
}
