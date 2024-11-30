package utils

import (
	"strings"
)

func ApplyCommitStyle(message, style string) string {
	switch style {
	case "conventional":
		// Aplica o estilo convencional
		// Exemplo: feat: mensagem
		return "feat: " + strings.TrimSpace(message)
	case "emoji-style":
		// Aplica emojis ao commit
		// Exemplo: ✨ mensagem
		return "✨ " + strings.TrimSpace(message)
	case "plain-text":
		// Retorna a mensagem sem modificações
		return strings.TrimSpace(message)
	default:
		// Estilo desconhecido, retorna a mensagem original
		return strings.TrimSpace(message)
	}
}
