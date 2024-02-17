package web

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

// UploadFile upload file to server
// @Summary      UploadFile upload file to server
// @Description  UploadFile upload file to server
// @Tags         Upload File API
// @Param        file  formData  file  false "file"
// @Router       /upload/file [post]
func (api Router) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
		return
	}

	dst, err := os.Create(fmt.Sprintf(webP+"/images/%d%s", time.Now().UnixNano(), filepath.Ext(file.Filename)))
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
		return
	}
	defer f.Close()

	_, err = io.Copy(dst, f)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, map[string]any{
		"message": "success upload image",
	})
}
