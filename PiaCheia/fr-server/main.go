package main

import (
	"flag"
	"fmt"
	"net/http"
)

var ( // flags
	ADDR string
)

type StaticPage struct{ FileName string }
func (p StaticPage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, p.FileName)
}

func init() {
	flag.StringVar(&ADDR, "addr", "127.0.0.1:80", "endereço de ip do servidor")
	flag.Parse()
}

func main() {
	http.Handle("/", StaticPage{"./files/main.html"})
	http.Handle("/article/", http.StripPrefix("/article/", http.FileServer(http.Dir("pages/"))))
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("files/"))))

	fmt.Println("running")

	panic(http.ListenAndServe(ADDR, nil))
}

