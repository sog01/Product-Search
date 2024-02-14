package web

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/sog01/productdiscovery/internal/search/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	searchService service.Service
	t             *template.Template
}

func (r Router) Run() {
	g := gin.Default()
	g.StaticFS("/static/", http.Dir(webP+"/static"))
	r.webRouter(g)
	r.apiRouter(g)

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	g.Run()
}

func (r Router) webRouter(g *gin.Engine) {
	searchAPI := g.Group("/")
	{
		searchAPI.GET("/", r.Index)
		searchAPI.GET("/product", r.SearchProducts)
		searchAPI.GET("/product/result", r.SearchProductsResult)
		searchAPI.GET("/product/cards", r.SearchProductsCards)
		searchAPI.GET("/product/autocomplete", r.SearchProductsAutocomplete)
		searchAPI.GET("/catalog", r.Catalog)
	}
}

func (r Router) apiRouter(g *gin.Engine) {
	api := g.Group("/api")
	searchAPI := api.Group("search")
	{
		searchAPI.GET("/", r.Search)
		searchAPI.GET("/autocomplete", r.SearchAutocomplete)
		searchAPI.GET("/total", r.SearchTotal)
		searchAPI.GET("/catalogs", r.SearchCatalogs)
	}
	productAPI := api.Group("/products/bulk")
	{
		productAPI.POST("/", r.BulkInsert)
		productAPI.POST("/update", r.BulkUpdate)
	}
}

func NewRouter(searchService service.Service) Router {
	tt, err := template.ParseGlob(webP + "/templates/*")
	if err != nil {
		log.Fatalf("failed to parse templates: %v", err)
	}
	return Router{
		searchService: searchService,
		t:             tt,
	}
}

var webP = webPath()

func webPath() string {
	path := "./../../web"
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		path = "./web"
	}
	return path
}
