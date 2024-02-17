package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/olivere/elastic/v7"
	_ "github.com/sog01/productdiscovery/docs"
	"github.com/sog01/productdiscovery/indices"
	"github.com/sog01/productdiscovery/internal/search/service"
	shortenersvc "github.com/sog01/productdiscovery/internal/shortener/service"
	"github.com/sog01/productdiscovery/internal/web"
)

// @title           Product Search API
// @description     This is a product search API swagger documentation.

// @host      localhost:8080
// @BasePath  /api

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed read .env file: %v", err)
	}

	log.Println("elastic baseURL", os.Getenv("ELASTIC.BASEURL"))
	es, err := elastic.NewClient(
		elastic.SetHealthcheck(false),
		elastic.SetSniff(false),
		elastic.SetURL(os.Getenv("ELASTIC.BASEURL")),
	)
	if err != nil {
		log.Fatalf("failed create elasticsearch client: %v", err)
	}

	indices.Create(es)
	search := service.NewService(es)
	shortener := shortenersvc.NewService(es)
	router := web.NewRouter(search, shortener)
	router.Run()
}
