package web

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"regexp"
	"time"
)

func Startserver(addr *string) {
	log.Println("ser")
	n := net.ListenConfig{}
	lis, err := n.Listen(context.Background(), "tcp", *addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.Serve(lis, rout()))
}

func rout() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", validUrl(handler))
	return mux
}

var validPath = regexp.MustCompile("^/(reg|auth)?(/process)?$")

func validUrl(fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			log.Println(r.UserAgent(), "knocking the wrong way")
			http.NotFound(w, r)
			return
		}
		fn(w, r)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.UserAgent())
	buffer := bytes.Buffer{}
	now := time.Now()
	t := now.Round(5 * time.Minute)
	detect := t.Unix() - now.Unix()
	log.Println(now.Unix())
	log.Println(t.Unix())
	log.Println(detect)
	if detect > 0 {
		req := now.Unix() + detect
		log.Println(req)
		buffer.WriteString(fmt.Sprintf("%d", req))
		w.Write(buffer.Bytes())
		return
	} else {
		req := t.Unix() + 300
		log.Println(req)
		buffer.WriteString(fmt.Sprintf("%d", req))
		w.Write(buffer.Bytes())
		return
	}
}
