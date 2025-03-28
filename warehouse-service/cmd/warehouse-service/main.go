package main

import (
	"log"
	
	"warehouse-service/internal/application"
	"warehouse-service/internal/infrastructure/cache"
	"warehouse-service/internal/infrastructure/db"
	"warehouse-service/internal/infrastructure/grpc"
	"warehouse-service/internal/infrastructure/repository"
	"warehouse-service/internal/infrastructure/search"
)

func main() {
	// DB
	dbConn, err := db.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}
	
	// Redis
	redisCache := cache.NewRedisCache(os.Getenv("REDIS_HOST"))
	
	// Elasticsearch
	es, err := search.NewElasticSearch(os.Getenv("ELASTICSEARCH_HOST"))
	if err != nil {
		log.Fatal(err)
	}
	
	// Repositories
	productRepo := repository.NewProductRepository(dbConn)
	warehouseRepo := repository.NewWarehouseRepository(dbConn)
	
	// App Service
	app := application.NewWarehouseService(warehouseRepo, redisCache, es)
	
	// gRPC Server
	grpc.StartGRPCServer(os.Getenv("GRPC_PORT"), app)
}