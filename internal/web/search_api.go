package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sog01/productdiscovery/internal/search/model"
	"gopkg.in/guregu/null.v4"
)

// Search product from given q
// @Summary      Search product from given q
// @Description  Search product from given q
// @Tags         Search API
// @Accept       json
// @Produce      json
// @Param        q  query  string  false "q"
// @Param        catalog  query  string  false "catalog"
// @Param        size  query  int  false "size"
// @Param        sort_by  query  string  false "sort_by"
// @Param        next_cursor  query  string  false "next_cursor"
// @Success      200  {object}  model.SearchResponse
// @Router       /search [get]
func (api Router) Search(c *gin.Context) {
	r := model.SearchReq{
		Q: c.Query("q"),
	}
	if catalog := c.Query("catalog"); catalog != "" {
		r.Catalog = null.StringFrom(catalog)
	}
	if nextCursor := c.Query("next_cursor"); nextCursor != "" {
		r.NextCursor = null.StringFrom(nextCursor)
	}
	r.Size, _ = strconv.Atoi(c.Query("size"))
	r.SortBy = model.NewSort(c.Query("sort_by"))

	resp, err := api.searchService.Search(c.Request.Context(), r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// SearchAutocomplete from given q
// @Summary      Search autocomplete from given q
// @Description  Search autocomplete from given q
// @Tags         Search API
// @Accept       json
// @Produce      json
// @Param        q  query  string  false "q"
// @Success      200  {object}  model.AutocompleteResp
// @Router       /search/autocomplete [get]
func (api Router) SearchAutocomplete(c *gin.Context) {
	r := model.AutocompleteReq{
		Q: c.Query("q"),
	}
	resp, err := api.searchService.SearchAutocomplete(c.Request.Context(), r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Search product total from given q
// @Summary      Search product total from given q
// @Description  Search product total from given q
// @Tags         Search API
// @Accept       json
// @Produce      json
// @Param        q  query  string  false "q"
// @Success      200  {object}  model.SearchTotalResp
// @Router       /search/total [get]
func (api Router) SearchTotal(c *gin.Context) {
	r := model.SearchTotalReq{
		Q: c.Query("q"),
	}
	resp, err := api.searchService.SearchTotal(c.Request.Context(), r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Search product total from given q
// @Summary      Search product catalogs from given q
// @Description  Search product catalogs from given q
// @Tags         Search API
// @Accept       json
// @Produce      json
// @Param        q  query  string  false "q"
// @Success      200  {object}  model.SearchCatalogsResp
// @Router       /search/catalogs [get]
func (api Router) SearchCatalogs(c *gin.Context) {
	r := model.SearchCatalogsReq{
		Q: c.Query("q"),
	}
	resp, err := api.searchService.SearchCatalogs(c.Request.Context(), r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// BulkInsert bulk insert products
// @Summary      BulkInsert bulk insert products
// @Description  BulkInsert bulk insert products
// @Tags         Products API
// @Accept       json
// @Produce      json
// @Param        request body  model.BulkInsertReq true "request body"
// @Success      200  {object}  model.BulkInsertResp
// @Router       /products/bulk [post]
func (api Router) BulkInsert(c *gin.Context) {
	r := model.BulkInsertReq{}
	if err := json.NewDecoder(c.Request.Body).Decode(&r); err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}
	resp, err := api.searchService.BulkInsert(c.Request.Context(), r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// BulkUpdate bulk update products
// @Summary      BulkUpdate bulk update products
// @Description  BulkUpdate bulk update products
// @Tags         Products API
// @Accept       json
// @Produce      json
// @Param        request body model.BulkUpdateReq true "request body"
// @Success      200  {object}  model.BulkUpdateResp
// @Router       /products/bulk/update [post]
func (api Router) BulkUpdate(c *gin.Context) {
	r := model.BulkUpdateReq{}
	if err := json.NewDecoder(c.Request.Body).Decode(&r); err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}
	resp, err := api.searchService.BulkUpdate(c.Request.Context(), r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
