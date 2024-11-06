package metahandler

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func SetUp() {
	var err error
	db, err = sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("sql is online")
}

func MakeTabels() {
	if db == nil {
		SetUp()
	}
	db.Exec("CREATE TABLE IF NOT EXISTS user(Id INTEGER,userName TEXT, passWord TEXT, PRIMARY KEY(Id));")
	db.Exec("CREATE TABLE IF NOT EXISTS Auther(Id INTEGER,name TEXT, PRIMARY KEY(Id));")
	db.Exec("CREATE TABLE IF NOT EXISTS book(Id INTEGER,autherId INTEGER,pageCount INTEGER,lastUpdate INTEGER, PRIMARY KEY(Id));")
	db.Exec("CREATE TABLE IF NOT EXISTS bookUrl(Id INTEGER, bookId INTEGER, path TEXT, addDate INTEGER, addByUser INTEGER, PRIMARY KEY(Id));")
	db.Exec("CREATE TABLE IF NOT EXISTS bookTages(Id INTEGER,name TEXT,parentTageId INTEGER, PRIMARY KEY(Id));")
	db.Exec("CREATE TABLE IF NOT EXISTS bookPage(Id INTEGER,pageCount INTEGER, addData INTEGER,bookId INTEGER, fileName TEXT, PRIMARY KEY(Id));")
}

func AddUser(userName, passWord string) {
	if db == nil {
		SetUp()
	}
	db.Exec("INSERT INTO \"user\" (UserName, PassWord) VALUES(?, ?);", userName, passWord)
}
func TestUser(userName, passWord string) (userId int) {
	if db == nil {
		SetUp()
	}
	db.QueryRow("SELECT Id FROM \"user\" WHERE UserName = ? AND PassWord = ?;", userName, passWord).Scan(&userId)
	return
}
