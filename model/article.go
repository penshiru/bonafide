package model

//Article Holds the article model and his methods
type Article struct {
	ID        int    `json:"id" db:"article_id"`
	Name      string `json:"name" db:"name" `
	Text      string `json:"text" db:"text"`
	ChapterID uint   `json:"chapterID" db:"chapter_id"`
	LawID     uint   `json:"lawID" db:"law_id"`
	Reviewed  bool   `json:"reviewed" db:"reviewed"`
}
