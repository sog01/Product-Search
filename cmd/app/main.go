package main

import (
	"log"

	"github.com/olivere/elastic/v7"
	"github.com/sog01/productdiscovery/indices"
	"github.com/sog01/productdiscovery/internal/search/service"
	"github.com/sog01/productdiscovery/internal/web"
)

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
