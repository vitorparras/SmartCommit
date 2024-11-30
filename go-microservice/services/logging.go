package services

import (
	"encoding/json"
	"log"

	"go-microservice/models"

	"github.com/gin-gonic/gin"
)

func LogRequestResponse(c *gin.Context, request models.CommitRequest, response string) {
	requestData, _ := json.Marshal(request)
	log.Printf("Request: %s", string(requestData))
	log.Printf("Response: %s", response)
}
