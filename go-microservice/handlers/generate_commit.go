package handlers

import (
	"go-microservice/models"
	"go-microservice/services"
	"go-microservice/services/ai"
	"go-microservice/services/git"
	"go-microservice/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// handlers/generate_commit.go

// GenerateCommit
// @Summary      Gera uma mensagem de commit
// @Description  Gera uma mensagem de commit personalizada usando IA
// @Tags         Commit
// @Accept       json
// @Produce      json
// @Param        request  body      models.CommitRequest  true  "Dados para gerar o commit"
// @Success      200      {object}  models.CommitResponse
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /generate-commit [post]
func GenerateCommit(c *gin.Context) {
	var request models.CommitRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validações
	if err := utils.ValidateCommitRequest(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verifica se o repositório Git é válido
	if !git.IsValidRepo(request.RepoPath) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Git repository"})
		return
	}

	// Gera o diff ou status do Git conforme necessário
	gitData, err := git.GetGitData(request.RepoPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Gera a mensagem de commit usando o serviço de IA apropriado
	var commitMessage string
	if ai.IsOnlineModel(request.Model) {
		commitMessage, err = ai.GenerateCommitMessageOnline(request, gitData)
	} else {
		commitMessage, err = ai.GenerateCommitMessageOffline(request, gitData)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Aplica o estilo de commit
	commitMessage = utils.ApplyCommitStyle(commitMessage, request.CommitStyle)

	// Salva a mensagem original antes da tradução
	originalMessage := commitMessage

	// Traduz a mensagem se necessário
	if request.CommitLang != "" && request.CommitLang != "en" {
		commitMessage, err = ai.TranslateMessage(commitMessage, request.CommitLang)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// Registro de logs
	services.LogRequestResponse(c, request, commitMessage)

	// Resposta
	c.JSON(http.StatusOK, models.CommitResponse{
		CommitMessage:   commitMessage,
		OriginalMessage: originalMessage,
	})
}
