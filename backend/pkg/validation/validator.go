package validation

import (
	"os"
	"os/exec"
	"strings"
	"time"
)

// ValidationResult representa o resultado da validação de um repositório
type ValidationResult struct {
	IsValid       bool   `json:"isValid"`
	LastModified  string `json:"lastModified"`
	CurrentBranch string `json:"currentBranch,omitempty"`
	Error         string `json:"error,omitempty"`
}

// ValidateRepository verifica se o caminho contém um repositório Git válido
func ValidateRepository(path string) (ValidationResult, error) {
	// Verifica se o caminho existe
	info, err := os.Stat(path)
	if err != nil || !info.IsDir() {
		return ValidationResult{IsValid: false, Error: "Invalid path or directory not found"}, nil
	}

	// Verifica se é um repositório Git
	if _, err := os.Stat(path + "/.git"); os.IsNotExist(err) {
		return ValidationResult{IsValid: false, Error: "Not a Git repository"}, nil
	}

	// Obtém a última modificação
	lastModified := info.ModTime().Format(time.RFC3339)

	// Obtém a branch atual
	cmd := exec.Command("git", "-C", path, "rev-parse", "--abbrev-ref", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return ValidationResult{IsValid: false, LastModified: lastModified, Error: "Failed to get branch"}, nil
	}

	return ValidationResult{
		IsValid:       true,
		LastModified:  lastModified,
		CurrentBranch: strings.TrimSpace(string(output)),
	}, nil
}
