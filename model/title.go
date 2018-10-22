package model

//Title struc is the model for a law Title
type Title struct {
	ID       uint      `json:"id" db:"title_id"`
	Name     string    `json:"name" db:"name"`
	Chapters []Chapter `json:"chapters"`
	LawID    uint      `json:"lawID" db:"law_id"`
	BookID   uint      `json:"bookID" db:"book_id"`
	Reviewed bool      `json:"reviewed" db:"reviewed"`
}

type TitleStore interface {
	CreateTitle() (uint, error)
}

//AddChapter adds parsed chapter data to parsed law object
func (t *Title) AddChapter(chapter Chapter) []Chapter {
	t.Chapters = append(t.Chapters, chapter)
	return t.Chapters
}
