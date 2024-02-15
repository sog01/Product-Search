package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sog01/pipe"
	"github.com/sog01/productdiscovery/internal/search/model"
	shortenermodel "github.com/sog01/productdiscovery/internal/shortener/model"
	"gopkg.in/guregu/null.v4"
)

func (r Router) Index(c *gin.Context) {
	r.templates.ExecuteTemplate(c.Writer, "index", nil)
}

func (r Router) SearchProducts(c *gin.Context) {
	r.renderProductSearchResult(c, "product")
}

func (r Router) SearchProductsResult(c *gin.Context) {
	r.renderProductSearchResult(c, "search_result")
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

	autocompletes := []map[string]any{}
	for _, autocomplete := range resp.Autocompletes {
		autocompletes = append(autocompletes, map[string]any{
			"Highlight": autocomplete.Highlight,
			"Href":      "/product?q=" + autocomplete.Title,
		})
	}

	r.templates.ExecuteTemplate(c.Writer, "autocompletes", map[string]any{
		"Autocompletes": autocompletes,
		"NewPage":       c.Query("newPage"),
	})
}

func (r Router) renderProductSearchResult(c *gin.Context, templateName string) {
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
	r.templates.ExecuteTemplate(c.Writer, templateName, map[string]any{
		"Page":            "Product",
		"Q":               c.Query("q"),
		"SortBy":          c.Query("sort_by"),
		"Catalog":         c.Query("catalog"),
		"AutocompleteURL": "/product/autocomplete",
		"WithSearchInput": true,
		"Products":        productSearchResult.Products,
		"Catalogs":        catalogSearchResult.Catalogs,
		"NextCursor":      productSearchResult.NextCursor,
		"ProductSize":     productSize,
		"TotalProduct":    totalProduct,
		"HasMoreData":     hasMoreData,
	})
}

func (r Router) Catalog(c *gin.Context) {
	var (
		catalogs           model.SearchCatalogsResp
		topProductCatalogs model.SearchTopProductCatalogResp
	)
	exec := pipe.P(
		func(any, pipe.Responses) (response any, err error) {
			resp, err := r.searchService.SearchCatalogs(c.Request.Context(), model.SearchCatalogsReq{})
			catalogs = resp
			return nil, err
		},
		func(any, pipe.Responses) (response any, err error) {
			catalogStrings := []string{}
			for _, catalog := range catalogs.Catalogs {
				catalogStrings = append(catalogStrings, catalog.Catalog)
			}
			resp, err := r.searchService.SearchTopProductCatalogs(c.Request.Context(), model.SearchTopProductCatalogReq{
				Catalogs: catalogStrings,
			})
			topProductCatalogs = resp
			return nil, err
		},
	)

	_, err := exec(nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
		return
	}

	r.templates.ExecuteTemplate(c.Writer, "catalog", map[string]any{
		"Page":     "Catalog",
		"Q":        c.Query("q"),
		"Catalogs": topProductCatalogs.TopProductCatalogs,
	})
}

func (r Router) RedirectShortener(c *gin.Context) {
	realUrl, err := r.shortenerService.GetShortener(c.Request.Context(), shortenermodel.GetShortenerReq{
		Slug: c.Param("slug"),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
		return
	}
	http.Redirect(c.Writer, c.Request, realUrl.RealURL, http.StatusSeeOther)
}
