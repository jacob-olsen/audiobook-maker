package metahandler

import "time"

type User struct {
	Id       int
	UserName string
	PassWord string
}

type Auther struct {
	Id   int
	Name string
}

type Book struct {
	Id          int
	Name        string
	Description string
	AutherId    int
	PageCount   int
	LastUpdate  time.Time
	Url         []BookUrl
	Tages       []BookTages
}
type BookUrl struct {
	Id        int
	BookId    int
	Path      string
	AddDate   time.Time
	AddByUser []User
}
type BookTages struct {
	Id           int
	Name         string
	ParentTageId int
}
type BookPage struct {
	Id        int
	PageCount int
	AddData   time.Time
	BookId    int
	FileName  string
}
