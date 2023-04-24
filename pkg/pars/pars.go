package pars

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type KeshTempl struct {
	Kesh map[string]*template.Template
}

func New(names ...string) (kesh KeshTempl) {
	kesh.Kesh = make(map[string]*template.Template)
	for _, n := range names {
		key := n
		kesh.Kesh[key] = template.New("")
	}
	return kesh
}

func (kesh *KeshTempl) LoadHash(pathdir string) {
	for key := range kesh.Kesh {
		path := fmt.Sprintf("%s%s.html", pathdir, key)
		templ, err := template.ParseFiles(path)
		if err != nil {
			log.Fatal(err)
		}
		kesh.Kesh[key] = templ
	}
}

func ParsPage(w http.ResponseWriter, namepage string, kesh *KeshTempl, content interface{}) (err error) {
	//w.WriteHeader(http.StatusOK)

	err = kesh.Kesh[namepage].Execute(w, content)
	if err != nil {
		return err
	}
	return nil
}
