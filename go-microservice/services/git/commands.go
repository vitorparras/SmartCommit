// services/git/commands.go

package git

import (
	"fmt"
	"os/exec"
	"strings"
)

// GitStatus retorna o status do repositório
func GitStatus(path string) (string, error) {
	cmd := exec.Command("git", "-C", path, "status", "--short")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("git status error: %v", err)
	}
	return string(output), nil
}

// GitAdd adiciona arquivos ao índice
func GitAdd(path string, files []string) error {
	args := append([]string{"-C", path, "add"}, files...)
	cmd := exec.Command("git", args...)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git add error: %v, output: %s", err, string(output))
	}
	return nil
}

// GitAddAll adiciona todos os arquivos ao índice
func GitAddAll(path string) error {
	cmd := exec.Command("git", "-C", path, "add", ".")
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git add . error: %v, output: %s", err, string(output))
	}
	return nil
}

// GitCommit cria um commit com a mensagem fornecida
func GitCommit(path, message string) error {
	cmd := exec.Command("git", "-C", path, "commit", "-m", message)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git commit error: %v, output: %s", err, string(output))
	}
	return nil
}

// GitPush envia os commits para o repositório remoto
func GitPush(path string, remote string, branch string) error {
	if remote == "" {
		remote = "origin"
	}
	if branch == "" {
		branch = "main"
	}
	cmd := exec.Command("git", "-C", path, "push", remote, branch)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git push error: %v, output: %s", err, string(output))
	}
	return nil
}

// GitDiff retorna o diff das alterações
func GitDiff(path string) (string, error) {
	cmd := exec.Command("git", "-C", path, "diff")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("git diff error: %v", err)
	}
	return string(output), nil
}

// GitLog retorna o log dos commits
func GitLog(path string, n int) (string, error) {
	cmd := exec.Command("git", "-C", path, "log", fmt.Sprintf("-%d", n))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("git log error: %v", err)
	}
	return string(output), nil
}

// IsGitInstalled verifica se o Git está instalado
func IsGitInstalled() bool {
	cmd := exec.Command("git", "--version")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

// ParseGitStatus parseia o resultado do git status
func ParseGitStatus(statusOutput string) []string {
	var files []string
	lines := strings.Split(statusOutput, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			files = append(files, parts[1])
		}
	}
	return files
}
