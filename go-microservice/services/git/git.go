package git

import (
	"os/exec"
	"strings"
)

func IsValidRepo(path string) bool {
	cmd := exec.Command("git", "-C", path, "rev-parse", "--is-inside-work-tree")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}
	return strings.TrimSpace(string(output)) == "true"
}

func GetGitData(path string) (string, error) {
	// Você pode customizar esta função para retornar o 'git status' ou 'git diff'
	cmd := exec.Command("git", "-C", path, "diff", "--staged")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
