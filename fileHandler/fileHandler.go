package filehandler

import (
	"os"
)

var dataRootPath string = "library"
var metaInfoPath string = dataRootPath + "/meta.json"
var bookRootPath string = dataRootPath + "/books"

type BookInfo struct {
	BookId         int
	HasIcon        bool
	PartOfColltion bool

	BookName   string
	FolderName string

	Description string
	wordCount   int

	Url []string
}

func Setup() {
	if !fileExist(dataRootPath) {
		os.Mkdir(dataRootPath, 0770)
	}
	if !fileExist(bookRootPath) {
		os.Mkdir(bookRootPath, 0770)
	}
}

func ReadStoryPage(bookId int, pageId int, rollBackCount int) string {
	//dataRaw, err := io.ReadAll()
	return ""
}

func AddBook(bookName string) (id int) {
	var newBook BookInfo
	newBook.BookName = bookName
	newBook.FolderName = bookName //WIP

	bookList := GetBookList()
	//find next id
	for _, book := range bookList {
		if book.BookId > id {
			id = book.BookId
		}
	}
	id++
	//set book id
	newBook.BookId = id

	return
}

func GetBookList() []BookInfo {
	if !fileExist(metaInfoPath) {
		RebuildBookList()
	}
	return nil
}

func RebuildBookList() {

}

func fileExist(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
