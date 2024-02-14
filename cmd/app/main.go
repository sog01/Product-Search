package main

import (
	"log"

	"github.com/olivere/elastic/v7"
	_ "github.com/sog01/productdiscovery/docs"
	"github.com/sog01/productdiscovery/indices"
	"github.com/sog01/productdiscovery/internal/search/service"
	"github.com/sog01/productdiscovery/internal/web"
)

// @title           Product Search API
// @description     This is a product search API swagger documentation.

// @host      localhost:8080
// @BasePath  /api

func main() {
	es, err := elastic.NewClient(
		elastic.SetHealthcheck(false),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Fatalf("failed create elasticsearch client: %v", err)
	}
	indices.CreateProductSearch(es)
	search := service.NewService(es)
	router := web.NewRouter(search)
	router.Run()
}
