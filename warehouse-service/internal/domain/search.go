package domain

import (
	"context"
)

type SearchService interface {
	IndexProduct(ctx context.Context, product Product) error
	SearchProducts(ctx context.Context, query string) ([]Product, error)
}
