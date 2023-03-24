package help

import (
	"fmt"
	"log"
	"net/http"
)

func ServesError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, fmt.Sprintf("serve err:%v", err), http.StatusMethodNotAllowed)
		log.Println(err)
		return
	}
}
