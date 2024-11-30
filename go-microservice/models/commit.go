package models

type CommitRequest struct {
	RepoPath     string `json:"repo_path" binding:"required"`
	Model        string `json:"model" binding:"required"`
	CommitLang   string `json:"commit_lang"`
	CommitLength string `json:"commit_length" binding:"required,oneof=short long"`
	CommitStyle  string `json:"commit_style" binding:"required,oneof=conventional plain-text emoji-style"`
	OpenAIKey    string `json:"openai_key"`
}

type CommitResponse struct {
	CommitMessage   string `json:"commit_message"`
	OriginalMessage string `json:"original_message,omitempty"`
}
