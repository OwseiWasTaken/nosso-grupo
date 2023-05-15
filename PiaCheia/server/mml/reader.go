package mml

import (
	"fmt"
	"strings"
)

type Id = string
type Obj map[string]string

type Tag struct {
	Name string
	Attributes Obj
	Content string
	HasContent bool
	Children []*Tag
	HasChild bool
	Depth int
}

func (T *Tag) MakeTag(name string, attrs Obj) (Tag) {
	var NewNode Tag = Tag{
		Name:name,
		Attributes:attrs,
		Content:"",
		HasContent: false,
		Children:[]*Tag{},
		HasChild:false,
		Depth: T.Depth+1,
	}
	index := T.InsertTag(&NewNode)
	if id, ok := attrs["id"]; ok {
		TagMap[id] = T.Children[index]
	}
	return NewNode
}

func (T Tag) formatAttrs() (text string) {
	for key, value := range T.Attributes {
		text+=fmt.Sprintf(" %s=\"%s\"", key, value)
	}
	return text
}

func (T Tag) innerHTML() (text string) {
	if (!(T.HasContent||T.HasChild)) {return ""}
	if (T.HasContent) {return T.Content+"\n</"+T.Name+">"}

	for _, tag := range T.Children {
		text+=tag.Format()
	}
	text+="\n</"+T.Name+">"
	return text
}

func (T Tag) Format() (string) {
	return fmt.Sprintf("\n%s<%s%s>%s",
		strings.Repeat("\t", T.Depth),
		T.Name,
		T.formatAttrs(),
		T.innerHTML(),
	)
}

func Root() (T Tag) {
	return Tag{
		Name:"html",
		Attributes:map[string]string{},
		Content:"",
		HasContent: false,
		Children:[]*Tag{},
		HasChild:true,
		Depth: 0,
	}
}

func (T *Tag) InsertTag(Child *Tag) (int){
	T.HasChild = true
	T.Children = append(T.Children, Child)
	return len(T.Children)-1
}

var TagMap = map[Id]*Tag{}
var Document = Root()
func MM() {
	Document.MakeTag("body", Obj{"id":"uwud"})
	fmt.Println(Document)
	fmt.Println(Document.Format())
}
