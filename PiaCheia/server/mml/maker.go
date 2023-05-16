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
	Children []*Tag
	Depth int
	Parrent *Tag
}

func (T Tag) Tab(i... int) (string) {
	if len(i) == 0 {
		return strings.Repeat("\t", T.Depth);
	} else {
		return strings.Repeat("\t", T.Depth+i[0]);
	}
}

func (T Tag) HasContent() (bool) {
	return len(T.Content) != 0
}

func (T Tag) HasChild() (bool) {
	return len(T.Children) != 0
}

func (T *Tag) MakeTag(name string, attrs Obj) (*Tag) {
	var NewNode Tag = Tag{
		Name:name,
		Attributes:attrs,
		Content:"",
		Children:[]*Tag{},
		Depth: T.Depth+1,
		Parrent: T,
	}
	index := T.InsertTag(&NewNode)
	if id, ok := attrs["id"]; ok {
		TagMap[id] = T.Children[index]
	}
	return &NewNode
}

func (T Tag) formatAttrs() (text string) {
	for key, value := range T.Attributes {
		text+=fmt.Sprintf(" %s=\"%s\"", key, value)
	}
	return text
}

func (T Tag) innerHTML() (text string) {
	if (T.HasContent()) {return "\n"+T.Tab(1)+T.Content+"\n"+T.Tab()+"</"+T.Name+">"}

	for _, tag := range T.Children {
		text+=tag.Format()
	}
	text+="\n"+T.Tab()+"</"+T.Name+">"
	return text
}

func (T Tag) Format() (string) {
	return fmt.Sprintf("\n%s<%s%s>%s",
		T.Tab(),
		T.Name,
		T.formatAttrs(),
		T.innerHTML(),
	)
}

func (T *Tag) InsertTag(Child *Tag) (int){
	T.Children = append(T.Children, Child)
	return len(T.Children)-1
}

// return last tag
func (T *Tag) TagTree(Objects... Obj) (*Tag) {
	var current = T
	var name string
	for _, obj := range Objects {
		name = obj["name"] //TODO assert tag has name
		delete(obj, "name")
		current = current.MakeTag(name, obj)
	}
	return T
}

var TagMap = map[Id]*Tag{}
var Document = Tag{
		Name:"html",
		Attributes:map[string]string{},
		Content:"",
		Children:[]*Tag{},
		Depth: 0,
		Parrent: nil,
}

func MM() {
	// body#uwu;
	body := Document.MakeTag("body", Obj{"id":"uwu"})
	// #uwu->ul->li->p#texthere
	body.TagTree(Obj{"name":"ul"}, Obj{"name":"li"}, Obj{"name":"p", "id":"texthere"})
	// #texthere { text; class:"awhn" }
	TagMap["texthere"].Content = "text"
	TagMap["texthere"].Attributes["class"] = "awhn"

	fmt.Println(Document.Format())
}
