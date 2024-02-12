package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sog01/productdiscovery/internal/search/model"
	"gopkg.in/guregu/null.v4"
)

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
