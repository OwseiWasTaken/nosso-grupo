package main

import (
	"strings"
	"flag"
	"fmt"
	"encoding/json"
	"io/fs"
	"net/http"
	"path/filepath"
	// "database/sql"
	// _ "github.com/mattn/go-sqlite3"
)

//var db *sql.DB

//const SchemaAccounts = `
//CREATE TABLE IF NOT EXISTS accounts (
//	accountId INTEGER NOT NULL PRIMARY KEY AUTO INCREAMENT,
//	accountName TEXT NOT NULL,
//	passhash INT NOT NULL,
//	isAdmin BOOLEAN NOT NULL DEFAULT false
//);
//`

//const SchemaArticle = `
//CREATE TABLE IF NOT EXISTS articles (
//	articleId INTEGER NOT NULL PRIMARY KEY AUTO INCREMENT,
//	articleName TEXT NOT NULL,
//  path TEXT NOT NULL,
//	lastEditor INTERGER NOT NULL,
	/*unix timestamp*/
//	lastEdit DATE NOT NULL,
//	UNIQUE (path),
//	FOREIGN KEY(lastEditor) REFERENCES accounts(accountId),
//);
//`

//const SchemaComment = `
//CREATE TABLE IF NOT EXISTS comments (
//	commentId INTEGER NOT NULL PRIMARY KEY AUTO INCREAMENT,
//	posterId INTEGER NOT NULL,
//  text TEXT NOT NULL,
//	FOREIGN KEY(posterId) REFERENCES accounts(accountId),
//);
//`

var ( // flags
	ADDR string
)

type StaticPage struct{ FileName string }
func (p StaticPage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, p.FileName)
}

func FileSystemAPI(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Path[4:]

	if (file == "") {
		w.Write(Unpack(json.Marshal(Ls("files/"))))
		return
	}

}

func ArticleSystem(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, p.FileName)
}

func init() {
	flag.StringVar(&ADDR, "addr", "127.0.0.1:80", "endereço de ip do servidor")
	flag.Parse()
}

func main() {
	http.Handle("/", StaticPage{"./files/main.html"})
	http.Handle("/convert", StaticPage{"./files/convert.html"})

	http.Handle("/list-articles", StaticPage{"./files/articles.html"})
	http.HandleFunc("/articles/", ArticleSystemAPI)
	//http.Handle("/articles/", http.StripPrefix("/articles/", http.FileServer(http.Dir("files/articles/"))))
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("files/"))))
	http.HandleFunc("/fs/", FileSystemAPI)

	fmt.Println("running")

	panic(http.ListenAndServe(ADDR, nil))
}

func Ls(dir string) []string {
	var files = []string{}
	Panic(filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error { // download
		if (dir == path) {return nil}
		path = strings.TrimPrefix(path, dir)

		if ( info.IsDir() && !strings.HasSuffix(path, "/")) {
			path+="/"
		}
		if err != nil {
			fmt.Printf("Can't access file from path %q; e: %v\n", path, err)
			return err
		}
		files = append(files, path)
		return nil
	}))
	return files
}

func Panic(e error) {
	if (e != nil) {
		panic(e)
	}
}

func Unpack[T any](v T, e error) T {
	if (e != nil) {
		panic(e)
	}
	return v
}
