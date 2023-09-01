package main

import (
	"strings"
	"flag"
	"fmt"
	"encoding/json"
	"io/fs"
	"net/http"
	"path/filepath"
	. "piacheia/util"
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
	http.ServeFile(w, r, "files/articles"+r.URL.Path[10:])
}

var ( // flags
	ADDR string
	SQLFILE string
)

func init() {
	flag.StringVar(&ADDR, "addr", "127.0.0.1:80", "endereço de ip do servidor")
	flag.StringVar(&SQLFILE, "sql", "paicheia.db", "Arquivo de sqlite3 para ser usado como base de dados")
	flag.Parse()

	InitSQL(SQLFILE)
	//fmt.Println(Articles)
	//fmt.Println(Comments)
	//fmt.Println(Accounts)
}

func main() {
	http.Handle("/", StaticPage{"./files/main.html"})
	http.Handle("/convert", StaticPage{"./files/convert.html"})

	http.Handle("/list-articles", StaticPage{"./files/articles.html"})
	http.HandleFunc("/articles/", ArticleSystem)
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

