package search

import (
	"context"

	"github.com/ramiroschettino/Go-Store-Microservices/warehouse-service/internal/domain"
)

type ElasticSearchService struct {
	// Aquí puedes definir la conexión a Elasticsearch
}

func NewElasticSearchService() *ElasticSearchService {
	return &ElasticSearchService{}
}

func (e *ElasticSearchService) IndexProduct(ctx context.Context, product domain.Product) error {
	// Lógica para indexar en Elasticsearch
	return nil
}

func (e *ElasticSearchService) SearchProducts(ctx context.Context, query string) ([]domain.Product, error) {
	// Lógica para buscar productos en Elasticsearch
	return nil, nil
}
