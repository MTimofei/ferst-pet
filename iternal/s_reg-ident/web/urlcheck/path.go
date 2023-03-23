package urlcheck

import (
	"log"
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile("^/(reg|identify)/(process)$")

func CheckURL(fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	log.Println("not valid path")
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r)
	}
}
