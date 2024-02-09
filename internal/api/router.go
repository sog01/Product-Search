package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sog01/productdiscovery/internal/search/service"
)

type Router struct {
	searchService service.Service
}

func (r Router) Run() {
	g := gin.Default()
	search := g.Group("/search")
	{
		search.GET("/", r.Search)
		search.GET("/autocomplete", r.SearchAutocomplete)
		search.GET("/total", r.SearchTotal)
		search.GET("/catalogs", r.SearchCatalogs)
	}

	product := g.Group("/products/bulk")
	{
		product.POST("/", r.BulkInsert)
		product.POST("/update", r.BulkUpdate)
	}
	g.Run()
}

func NewRouter(searchService service.Service) Router {
	return Router{
		searchService: searchService,
	}
}
