package web

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sog01/productdiscovery/internal/shortener/model"
)

// CreateShortener create url shortener
// @Summary      CreateShortener create url shortener
// @Description  CreateShortener create url shortener
// @Tags         Shortener API
// @Accept       json
// @Produce      json
// @Param        request body  model.CreateShortenerReq true "request body"
// @Success      200  {object}  model.CreateShortenerResp
// @Router       /shortener [post]
func (api Router) CreateShortener(c *gin.Context) {
	r := model.CreateShortenerReq{}
	if err := json.NewDecoder(c.Request.Body).Decode(&r); err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}
	resp, err := api.shortenerService.CreateShortener(c.Request.Context(), r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
