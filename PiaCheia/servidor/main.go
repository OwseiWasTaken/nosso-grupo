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
func (p StaticPage) ServeHTTP(w HttpWriter, r HttpReq) {
	http.ServeFile(w, r, p.FileName)
}

func FileSystemAPI(w HttpWriter, r HttpReq) {
	file := r.URL.Path[4:]
	if (file == "") {
		w.Write(Unpack(json.Marshal(Ls("files/"))))
		return
	}
}

func ArticleSystem(w HttpWriter, r HttpReq) {
	http.ServeFile(w, r, "files/articles"+r.URL.Path[10:])
}

func RegisterHandler(w HttpWriter, r HttpReq) {
	if (r.Method == "GET") {
		RegisterPage.ServeHTTP(w, r)
		return
	}
	name := r.FormValue("name")
	password := r.FormValue("password")
	_, taken := NameToAccount[name]
	if (taken) {
		Redirect(w, r, "/register?err=name-taken", "taken")
		return
	}

	NewAccount(name, password)

	SetUid(w, name, Hash(passhash))
	RegisterPage(w, r, "/", "logged in")
}

func LoginHandler(w HttpWriter, r HttpReq) {
	if (r.Method == "GET") {
		LoginPage.ServeHTTP(w, r)
		return
	}
	name := r.FormValue("name")
	password := r.FormValue("password")
	acc, taken := NameToAccount[name]

	if (!taken) {
		Redirect(w, r, "/login?name-inexistent", "taken")
		return
	}

	passhash = Hash(password)
	if (passhash != acc.Passhash) {
		Redirect(w, r, "/login?name-inexistent", "taken")
		return
	}

	SetUid(w, name, passhash)
	RegisterPage(w, r, "/", "logged in")
}

var ( // flags
	ADDR string
	SQLFILE string
)

var (
	LoginPage TemplatedPage
	RegisterPage TemplatedPage
)

func init() {
	flag.StringVar(&ADDR, "addr", "127.0.0.1:80", "endereço de ip do servidor")
	flag.StringVar(&SQLFILE, "sql", "piacheia.db", "Arquivo de sqlite3 para ser usado como base de dados")
	flag.Parse()

	InitSQL(SQLFILE)
	LoginPage = TemplatePage("./pages/login.html", nil, nil)
	RegisterPage = TemplatePage("./pages/register.html", nil, nil)
}

func main() {
	http.Handle("/", TemplatePage("./pages/index.html",
		[]GOTMPlugin{
			{"acc",GOTM_getAccount},
		},
		nil,
	))
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/register", RegisterHandler)

	http.Handle("/convert", StaticPage{"./pages/convert.html"})

	http.Handle("/list-articles", StaticPage{"./pages/articles.html"})
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

