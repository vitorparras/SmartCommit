package domain

import "mime/multipart"

type CommitConfig struct {
	Format     string                `json:"format"`
	ReadmeFile *multipart.FileHeader `json:"-"`
	AIModel    string                `json:"aiModel"`
	Style      string                `json:"style"`
	Language   string                `json:"language"`
}

type Commit struct {
	Message string       `json:"message"`
	Author  string       `json:"author"`
	Date    string       `json:"date"`
	Config  CommitConfig `json:"config"`
}
