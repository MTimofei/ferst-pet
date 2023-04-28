package synchron

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func StartSystem() {

	client := http.Client{}
	res, err := client.Get("http://localhost:7000/")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	ts, err := strconv.Atoi(fmt.Sprintf("%s", body))
	if err != nil {
		log.Fatal(err)
	}
	timestart := int64(ts)
	now := time.Now().Unix()
	timeout := timestart - now
	log.Println(time.Duration(timeout) * time.Second)
	time.Sleep(time.Duration(timeout) * time.Second)
}
