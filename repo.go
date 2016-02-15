package main

import (
	"fmt"
	"golang.org/x/net/context"
)

func RepoGetDomainByHash(hash string) (Domain, error) {
	var domain Domain
	key := fmt.Sprintf("/domains/map/hashes/%s", hash)
	meta, err := etcd.Get(context.Background(), key, nil)
	if err != nil {
		return domain, err
	}
	name := meta.Node.Value
	domain = Domain{
		hash,
		name,
	}
	return domain, nil
}

func RepoGetDomainByName(name string) (Domain, error) {
	var domain Domain
	key := fmt.Sprintf("/domains/map/names/%s", name)
	meta, err := etcd.Get(context.Background(), key, nil)
	if err != nil {
		return domain, err
	}
	hash := meta.Node.Value
	domain = Domain{
		hash,
		name,
	}
	return domain, nil
}
