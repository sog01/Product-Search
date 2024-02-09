package main

import (
	"github.com/olivere/elastic/v7"
	"github.com/sog01/productdiscovery/indices"
	"github.com/sog01/productdiscovery/internal/api"
	"github.com/sog01/productdiscovery/internal/search/service"
)

func main() {
	es, _ := elastic.NewClient()
	indices.CreateProductDiscovery(es)
	search := service.NewService(es)
	router := api.NewRouter(search)
	router.Run()
}
