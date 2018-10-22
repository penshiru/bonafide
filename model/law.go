package model

import (
	"time"
)

//Law struct with most methods.
type Law struct {
	ID           int       `json:"id" db:"law_id"`
	Name         string    `json:"name" db:"name" `
	Number       int       `json:"number" db:"number"`
	ApprovalDate time.Time `json:"approvalDate" db:"approval_date"`
	PublishDate  time.Time `json:"publishDate" db:"publish_date"`
	Journal      string    `json:"journal" db:"journal"`
	Intro        string    `json:"intro" db:"intro"`
	Reviewed     bool      `json:"reviewed" db:"reviewed"`
	Revision     int       `json:"rev" db:"revision"`
	Books        []Book    `json:"books"`
	Titles       []Title   `json:"titles"`
	Chapters     []Chapter `json:"chapters"`
	Articles     []Article `json:"articles"`
	Init         string    `json:"init" db:"init"`
}

type LawStore interface {
	GetLaws() ([]Law, error)
	InsertLawDB(law *Law) error
	CreateLaw() (uint, error)
	GetLaw(id string) (Law, error)
	AutoComplete(query string) ([]string, error)
}

//AddTitle adds parsed title data to parsed law object
func (law *Law) AddTitle(title Title) []Title {
	law.Titles = append(law.Titles, title)
	return law.Titles
}

//AddBook adds parsed Bookdata to law object
func (law *Law) AddBook(book Book) []Book {
	law.Books = append(law.Books, book)
	return law.Books
}

//AddChapter adds parsed article data to parsed law object
//when there is no title
func (law *Law) AddChapter(chapter Chapter) []Chapter {
	law.Chapters = append(law.Chapters, chapter)
	return law.Chapters
}

//AddArticle adds parsed article data to parsed law object
//when there is no title and no chapter
func (law *Law) AddArticle(article Article) []Article {
	law.Articles = append(law.Articles, article)
	return law.Articles
}
