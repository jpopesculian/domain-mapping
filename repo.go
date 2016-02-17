package main

import (
	"errors"
	"fmt"
	"github.com/coreos/etcd/client"
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

func RepoGetDomainsByUserId(userId string) (Domains, error) {
	domains := make(Domains, 0)
	key := fmt.Sprintf("/domains/users/%s", userId)
	meta, err := etcd.Get(context.Background(), key, nil)
	if err != nil {
		return domains, err
	}
	nodes := meta.Node.Nodes
	for _, node := range nodes {
		domain := Domain{
			node.Key,
			node.Value,
		}
		domains = append(domains, domain)
	}
	return domains, nil
}

func RepoCreateDomainMapping(name, hash, userId string) (Domain, error) {
	var domain Domain
	if ok, domain, _ := RepoDomainExistsByHash(hash); ok {
		return domain, errors.New("Hash is already mapped to a domain!")
	}
	if ok, domain, _ := RepoDomainExistsByName(name); ok {
		return domain, errors.New("Name is already mapped to a domain!")
	}
	key := fmt.Sprintf("domains/map/names/%s", name)
	_, err := etcd.Create(context.Background(), key, hash)
	if err != nil {
		return domain, err
	}
	key = fmt.Sprintf("domains/map/hashes/%s", hash)
	_, err = etcd.Create(context.Background(), key, name)
	if err != nil {
		return domain, err
	}
	key = fmt.Sprintf("domains/users/%s/%s", userId, hash)
	_, err = etcd.Create(context.Background(), key, name)
	if err != nil {
		return domain, err
	}
	domain = Domain{
		hash,
		name,
	}
	return domain, nil
}

func RepoDomainExistsByHash(hash string) (bool, Domain, error) {
	domain, err := RepoGetDomainByHash(hash)
	if err != nil {
		if client.IsKeyNotFound(err) != true {
			return false, domain, err
		}
		return false, domain, nil
	}
	return true, domain, nil
}

func RepoDomainExistsByName(name string) (bool, Domain, error) {
	domain, err := RepoGetDomainByName(name)
	if err != nil {
		if client.IsKeyNotFound(err) != true {
			return false, domain, err
		}
		return false, domain, nil
	}
	return true, domain, nil
}
