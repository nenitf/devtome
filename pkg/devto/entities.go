package devto

type Article struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	BodyMarkdown string `json:"body_markdown"`
}
