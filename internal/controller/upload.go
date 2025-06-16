package controller

import (
	"aitools/internal/models"
	"aitools/internal/nas"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UploadHander 是一个POSTQ请求处理函数，用于处理文件上传。
func UploadHandler(c *gin.Context) {
	// Parse the json body to get the update data.
	var updateData models.UpdateData
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Process the file (e.g., save it, analyze it, etc.)
	// For now, we just return a success message.
	nas.UploadFiles(updateData)
	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}
