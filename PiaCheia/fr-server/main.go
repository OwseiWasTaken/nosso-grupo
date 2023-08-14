package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"
	"golang.org/x/net/html"
)

var ( // flags
	ADDR string
)

type StaticPage struct{ FileName string }
func (p StaticPage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, p.FileName)
}


func AssertEReaderDoc(content string) {
	doc, err := html.Parse(strings.NewReader(content))
	if (err != nil) { panic(err) }
	wrongNodes := assertEReaderTag(doc, []*html.Node{})
	fmt.Println(wrongNodes)
	//for _, nds := range wrongNodes {
	//	fmt.Printf("%+v\n", nds)
	//}
}

func assertEReaderTag(node *html.Node, errs []*html.Node) ([]*html.Node) {
	if (IinA(node.Data, IMPORTANT_NODES)) {
		var hasalt bool = false

		for _, attr := range node.Attr {
			if (attr.Key == "alt") {
				hasalt = true
			}
			fmt.Println(attr.Key, attr.Val)
		}
		if (!hasalt) {
			errs = append(errs, node)
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if (IinA(node.Data, IMPORTANT_NODES) || IinA(node.Data, CONTAINER_NODES)) {
			nodeErr := assertEReaderTag(c, errs)
			errs = append(errs, nodeErr...)
		}
	}

	return errs
}

func init() {
	flag.StringVar(&ADDR, "addr", "127.0.0.1:80", "endereço de ip do servidor")
	flag.Parse()
}

/* WTF
<video>
	<source>
	<source>
	alt text
</vide>
*/

//TODO <input> only needs .alt if .src is defined
var IMPORTANT_NODES = []string{ "img", "area", "input" }
// contianer nodes that COULD have images that COULD need alt text
var CONTAINER_NODES = []string{ "", "html", "body", "div", "p", "table", "button" }

func main() {
//	AssertEReaderDoc(`
//<!DOCTYPE html>
//<html>
//	<body>
//		<div></div>
//		<img>
//		<a href="http://example.com">Hello!</a>
//	</body>
//</html>
//	`)
	http.Handle("/", StaticPage{"./files/main.html"})
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("files/"))))

	fmt.Println("running")

	panic(http.ListenAndServe(ADDR, nil))
}

func IinA[T comparable](Item T, Arr []T) bool {
	for _, a := range Arr {
		if (Item == a) {
			return true
		}
	}
	return false
}

