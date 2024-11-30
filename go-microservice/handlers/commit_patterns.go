package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCommitPatterns(c *gin.Context) {
	patterns := []string{"conventional", "plain-text", "emoji-style"}
	c.JSON(http.StatusOK, gin.H{"patterns": patterns})
}
