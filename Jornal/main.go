package main

include "gutil"

import (
	"net/http"
	"net/url"
//	"html"
)

var (
	cyan = RGB(70, 160, 255)
	nc = RGB(255, 255, 255)
	green = RGB(0, 255, 0)
	red = RGB(255, 0, 0)
)

func HandleMainPage(w http.ResponseWriter, r *http.Request) {
	fprintf(w, "main")
	printf("%sclient%s got the main page\n", green, nc)
}


func handler(w http.ResponseWriter, r *http.Request) {
	printf("\n%sclient%s requested %s with %s\n", cyan, nc, r.URL.Path, r.URL.Query())
	vm, err := url.ParseQuery(r.URL.RawQuery)
	panic(err)
	_=vm

	if len(r.URL.Path) == 1 {
		HandleMainPage(w, r)
		return
	}

}

func main(){
	InitGu()
	http.HandleFunc("/", handler)
	PS("server started")
	http.ListenAndServe(":6969", nil)
	exit(0)
}
