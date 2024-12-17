package domain

type ModelType string

const (
	Online  ModelType = "Online"
	Offline ModelType = "Offline"
)

type Model struct {
	Id   int       `json:"id"`
	Name string    `json:"name"`
	Type ModelType `json:"type"`
}
