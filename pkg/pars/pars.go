package pars

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type HashTempl struct {
	Hash map[string]*template.Template
}

func New(names ...string) (hash HashTempl) {
	hash.Hash = make(map[string]*template.Template)
	for _, n := range names {
		key := n
		hash.Hash[key] = template.New("")
	}
	return hash
}

func (h *HashTempl) LoadHash(pathdir string) {
	for key := range h.Hash {
		path := fmt.Sprintf("%s%s.html", pathdir, key)
		templ, err := template.ParseFiles(path)
		if err != nil {
			log.Fatal(err)
		}
		h.Hash[key] = templ
	}
}

func ParsPage(w http.ResponseWriter, namepage string, hesh *HashTempl, content interface{}) (err error) {
	//w.WriteHeader(http.StatusOK)

	err = hesh.Hash[namepage].Execute(w, content)
	if err != nil {
		return err
	}
	return nil
}
