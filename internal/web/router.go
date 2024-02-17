package web

import (
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/sog01/productdiscovery/internal/search/service"
	shortenersvc "github.com/sog01/productdiscovery/internal/shortener/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	searchService    service.Service
	shortenerService shortenersvc.Service
	templates        *template.Template
}

func (r Router) Run() {
	g := gin.Default()
	g.StaticFS("/static/", http.Dir(webP+"/static"))
	g.StaticFS("/images/", http.Dir(webP+"/images"))
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
		searchAPI.GET("/product/autocomplete", r.SearchProductsAutocomplete)
		searchAPI.GET("/share/:slug", r.RedirectShortener)
		searchAPI.GET("/catalog", r.Catalog)
		searchAPI.POST("/catalog/share", r.ShareCatalog)
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
	productAPI := api.Group("/products")
	{
		productAPI.POST("/bulk", r.BulkInsert)
		productAPI.POST("/bulk/update", r.BulkUpdate)
		productAPI.POST("/upload/csv", r.UploadCSV)
	}
	shortenerAPI := api.Group("/shortener")
	{
		shortenerAPI.GET("/", r.GetShortener)
		shortenerAPI.POST("/", r.CreateShortener)
	}
	uploadAPI := api.Group("/upload")
	{
		uploadAPI.POST("/file", r.UploadFile)
	}
}

func NewRouter(searchService service.Service,
	shortenerService shortenersvc.Service) Router {
	templatesFiles := []string{}
	filepath.Walk(webP+"/templates", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		templatesFiles = append(templatesFiles, path)
		return nil
	})
	return Router{
		searchService:    searchService,
		shortenerService: shortenerService,
		templates:        template.Must(template.ParseFiles(templatesFiles...)),
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
