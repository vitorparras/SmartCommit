package handlers

import (
	"encoding/json"
	"net/http"

	"backend/internal/core/domain"
	"backend/internal/core/services"

	"github.com/go-chi/chi/v5"
)

type HTTPHandler struct {
	commitService *services.CommitService
}

func NewHTTPHandler(commitService *services.CommitService) *HTTPHandler {
	return &HTTPHandler{commitService: commitService}
}

func (h *HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := chi.NewRouter()

	router.Post("/generate", h.handleGenerateCommit)
	router.Get("/commits", h.handleGetCommits)
	router.Get("/ai-models", h.handleGetAIModels)

	router.ServeHTTP(w, r)
}

// @Summary Gerar uma mensagem de commit
// @Description Gerar uma mensagem de commit usando IA com configuração específica e arquivo README
// @Accept multipart/form-data
// @Produce json
// @Param description formData string true "Descrição do commit"
// @Param author formData string true "Autor do commit"
// @Param format formData string true "Formato do commit"
// @Param readmeFile formData file true "Arquivo README.md"
// @Param aiModel formData string true "Modelo de IA"
// @Param style formData string true "Estilo do commit"
// @Param language formData string true "Idioma do commit"
// @Success 200 {object} domain.Commit
// @Router /generate [post]
func (h *HTTPHandler) handleGenerateCommit(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB max
	if err != nil {
		http.Error(w, "Erro ao processar o formulário: "+err.Error(), http.StatusBadRequest)
		return
	}

	description := r.FormValue("description")
	author := r.FormValue("author")
	format := r.FormValue("format")
	aiModel := r.FormValue("aiModel")
	style := r.FormValue("style")
	language := r.FormValue("language")

	file, header, err := r.FormFile("readmeFile")
	if err != nil {
		http.Error(w, "Erro ao obter o arquivo README: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	config := domain.CommitConfig{
		Format:     format,
		ReadmeFile: header,
		AIModel:    aiModel,
		Style:      style,
		Language:   language,
	}

	commit, err := h.commitService.GenerateAndSaveCommit(description, author, config)
	if err != nil {
		http.Error(w, "Erro ao gerar commit: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(commit)
}

// @Summary Obter todos os commits
// @Description Obter todos os commits gerados
// @Produce json
// @Success 200 {array} domain.Commit
// @Router /commits [get]
func (h *HTTPHandler) handleGetCommits(w http.ResponseWriter, r *http.Request) {
	commits, err := h.commitService.GetAllCommits()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(commits)
}

// @Summary Obter modelos de IA disponíveis
// @Description Obter uma lista de modelos de IA disponíveis para geração de commit
// @Produce json
// @Success 200 {array} string
// @Router /ai-models [get]
func (h *HTTPHandler) handleGetAIModels(w http.ResponseWriter, r *http.Request) {
	models := h.commitService.GetAvailableAIModels()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models)
}
