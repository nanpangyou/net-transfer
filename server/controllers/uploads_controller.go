package controllers

import (
	"net/http"

	config "github.com/nanpangyou/net-transfer/server/config"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadsController(c *gin.Context) {
	if path := c.Param("path"); path != "" {
		target := filepath.Join(config.UploadsDir, path)
		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Content-Disposition", "attachment; filename="+path)
		c.Header("Content-Type", "application/octet-stream")
		c.File(target)
	} else {
		c.Status(http.StatusNotFound)
	}
}
