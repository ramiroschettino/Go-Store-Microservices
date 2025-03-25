package search

import (
	"context"
	
	"github.com/elastic/go-elasticsearch/v8"
)

type ElasticSearch struct {
	client *elasticsearch.Client
}

func NewElasticSearch(addr string) (*ElasticSearch, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{addr},
	}
	
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	
	return &ElasticSearch{client: client}, nil
}

func (e *ElasticSearch) IndexProduct(ctx context.Context, product domain.Product) error {
	// Implementar indexación
}