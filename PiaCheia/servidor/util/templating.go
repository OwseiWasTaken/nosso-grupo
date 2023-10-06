package util

import (
	"html/template"
)

// GOTM plugins should only rely on static info data
type GOTMPlugin struct {
	Name string
	Plug func(w HttpWriter, r HttpReq, info map[string]any)any
}

//func GOTM_example(w HttpWriter, r HttpReq, info map[string]any) any {
//	return 4
//}

type TemplatedPage struct {
	Template *template.Template
	Info map[string]any
	Plugins []GOTMPlugin
}

func TemplatePage(filename string, plugins []GOTMPlugin, info map[string]any) TemplatedPage {
	if info == nil {
		info = make(map[string]any)
	}
	if info == nil {
		info = make(map[string]any)
	}
	tmpl, e := template.ParseFiles(filename)
	if (e != nil) {panic(e)}
	return TemplatedPage{
		tmpl, info, plugins,
	}
}

func (s TemplatedPage) ServeHTTP (w HttpWriter, r HttpReq) {
	for _, plug := range s.Plugins {
		s.Info[plug.Name] = plug.Plug(w, r, s.Info)
	}

	s.Template.Execute(w, s.Info)
}

