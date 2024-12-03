package commit

// CommitTemplates define os padrões de mensagens disponíveis
var CommitTemplates = map[string]string{
	"fix":   "Fix: {description}",
	"feat":  "Feat: {description}",
	"chore": "Chore: {description}",
	"docs":  "Docs: {description}",
}

// GetTemplate retorna o padrão baseado na chave fornecida
func GetTemplate(pattern string) (string, bool) {
	template, exists := CommitTemplates[pattern]
	return template, exists
}
