package inject

import (
	"github.com/mutezebra/tiktok/pkg/discovery"
	"github.com/mutezebra/tiktok/pkg/log"
)

type Registry struct {
	*discovery.Registry
	Key    string
	Prefix string
	TTL    int64
	Addr   string
}

func NewRegistry(registry *Registry) *Registry {
	reg, err := discovery.NewRegistry(registry.Addr, registry.Key, registry.TTL, registry.Prefix)
	if err != nil {
		log.LogrusObj.Panic(err)
		return nil
	}
	return &Registry{Registry: reg}
}

type Resolver struct {
	*discovery.Resolver
}

func NewResolver(endpoint string) (*Resolver, error) {
	re, err := discovery.NewResolver(endpoint)
	return &Resolver{re}, err
}
