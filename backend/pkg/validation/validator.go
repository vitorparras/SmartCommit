package validation

import "fmt"

// Validate simula a validação de um repositório local
func Validate(path string) (string, error) {
	// Simulação de validação
	if path == "" {
		return "", fmt.Errorf("invalid repository path")
	}
	return fmt.Sprintf("Repository at path '%s' is valid", path), nil
}
