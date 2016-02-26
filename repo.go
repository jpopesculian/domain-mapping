package main

import (
	"errors"
	"fmt"
	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
	"strings"
)

func RepoGetDomainByName(name string) (Domain, error) {
	var domain Domain
	key := fmt.Sprintf("/domains/map/%s", name)
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
	key := fmt.Sprintf("/domains/users/ids/%s", userId)
	meta, err := etcd.Get(context.Background(), key, nil)
	if err != nil {
		return domains, err
	}
	nodes := meta.Node.Nodes
	for _, node := range nodes {
		nodePath := strings.Split(node.Key, "/")
		name := nodePath[len(nodePath)-1]
		hash := node.Value
		domain := Domain{
			hash,
			name,
		}
		domains = append(domains, domain)
	}
	return domains, nil
}

func RepoGetUserIdByName(name string) (string, error) {
	key := fmt.Sprintf("/domains/users/names/%s", name)
	meta, err := etcd.Get(context.Background(), key, nil)
	if err != nil {
		return "", err
	}
	return meta.Node.Value, nil
}

func RepoCreateDomainMapping(name, hash, userId string) (Domain, error) {
	var domain Domain
	if ok, domain, _ := RepoDomainExistsByName(name); ok {
		return domain, errors.New("Subdomain is already mapped to a hash!")
	}
	if domains, _ := RepoGetDomainsByUserId(userId); len(domains) > 2 {
		return domain, errors.New("Only 3 Subdomains allowed per user!")
	}
	key := fmt.Sprintf("domains/map/%s", name)
	_, err := etcd.Create(context.Background(), key, hash)
	if err != nil {
		return domain, err
	}
	key = fmt.Sprintf("domains/users/ids/%s/%s", userId, name)
	_, err = etcd.Create(context.Background(), key, hash)
	if err != nil {
		return domain, err
	}
	key = fmt.Sprintf("domains/users/names/%s", name)
	_, err = etcd.Create(context.Background(), key, userId)
	if err != nil {
		return domain, err
	}
	domain = Domain{
		hash,
		name,
	}
	return domain, nil
}

func RepoDeleteDomainByName(name string) Domain {
	domain := Domain{
		"",
		name,
	}
	key := fmt.Sprintf("domains/map/%s", name)
	meta, err := etcd.Delete(context.Background(), key, nil)
	if err == nil {
		domain.Hash = meta.PrevNode.Value
	}
	key = fmt.Sprintf("domains/users/names/%s", name)
	meta, err = etcd.Delete(context.Background(), key, nil)
	if err == nil {
		userId := meta.PrevNode.Value
		key = fmt.Sprintf("domains/users/ids/%s/%s", userId, name)
		etcd.Delete(context.Background(), key, nil)
	}
	return domain
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
