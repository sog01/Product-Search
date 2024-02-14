package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sog01/pipe"
	"github.com/sog01/productdiscovery/internal/search/model"
	"gopkg.in/guregu/null.v4"
)

func (r Router) Index(c *gin.Context) {
	r.indexTemplates.ExecuteTemplate(c.Writer, "index", nil)
}

func (r Router) SearchProducts(c *gin.Context) {
	r.renderSearchResult(c, "product")
}

func (r Router) SearchProductsResult(c *gin.Context) {
	r.renderSearchResult(c, "search_result")
}

func (r Router) SearchProductsAutocomplete(c *gin.Context) {
	resp, err := r.searchService.SearchAutocomplete(c.Request.Context(), model.AutocompleteReq{
		Q: c.Query("q"),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
		return
	}
	r.productTemplates.ExecuteTemplate(c.Writer, "autocompletes", map[string]any{
		"Autocompletes": resp.Autocompletes,
		"NewPage":       c.Query("newPage"),
	})
}

func (r Router) renderSearchResult(c *gin.Context, templateName string) {
	var (
		productSearchResult model.SearchResponse
		catalogSearchResult model.SearchCatalogsResp
		productSize         int
		totalProduct        int64
		hasMoreData         bool
	)
	sizeReq := 10
	exec := pipe.PipeGo(
		func(any, pipe.Responses) (response any, err error) {
			req := model.SearchReq{
				SortBy: model.NewSort(c.Query("sort_by")),
				Size:   sizeReq,
				Q:      c.Query("q"),
			}
			if catalog := c.Query("catalog"); catalog != "" {
				req.Catalog = null.StringFrom(catalog)
			}
			if nextCursor := c.Query("next_cursor"); nextCursor != "" {
				req.NextCursor = null.StringFrom(nextCursor)
			}
			resp, err := r.searchService.Search(c.Request.Context(), req)
			productSearchResult = resp
			productSize = len(resp.Products)
			return productSearchResult, err
		},
		func(args any, responses pipe.Responses) (response any, err error) {
			resp, err := r.searchService.SearchCatalogs(c.Request.Context(), model.SearchCatalogsReq{})
			catalogSearchResult = resp
			return catalogSearchResult, err
		},
		func(args any, responses pipe.Responses) (response any, err error) {
			req := model.SearchTotalReq{
				Q: c.Query("q"),
			}
			if catalog := c.Query("catalog"); catalog != "" {
				req.Catalog = null.StringFrom(catalog)
			}
			resp, err := r.searchService.SearchTotal(c.Request.Context(), req)
			totalProduct = resp.Total
			return totalProduct, err
		},
	)
	_, err := exec(nil, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
		return
	}

	hasMoreData = productSize == sizeReq
	r.productTemplates.ExecuteTemplate(c.Writer, templateName, map[string]any{
		"Q":            c.Query("q"),
		"SortBy":       c.Query("sort_by"),
		"Catalog":      c.Query("catalog"),
		"Products":     productSearchResult.Products,
		"Catalogs":     catalogSearchResult.Catalogs,
		"NextCursor":   productSearchResult.NextCursor,
		"ProductSize":  productSize,
		"TotalProduct": totalProduct,
		"HasMoreData":  hasMoreData,
	})
}

func (r Router) Catalog(c *gin.Context) {
	r.catalogTemplates.ExecuteTemplate(c.Writer, "catalog", nil)
}
