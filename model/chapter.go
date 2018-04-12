package model

//Chapter is the model for a Law chapter
type Chapter struct {
	ID       int64     `json:"id" db:"chapter_id"`
	Name     string    `json:"name" db:"name" `
	Articles []Article `json:"articles" db:"articles"`
	TitleID  int64     `json:"titleID" db:"title_id"`
	LawID    int64     `json:"lawID" db:"law_id"`
	Reviewed bool      `json:"reviewed" db:"reviewed"`
}

//AddArticle adds parsed article data to parsed law object
func (c *Chapter) AddArticle(article Article) []Article {
	c.Articles = append(c.Articles, article)
	return c.Articles
}
