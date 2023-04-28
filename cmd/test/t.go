package main

import (
	"log"
	"time"
)

// fmt.Printf("Ограничение количества горутин: %d\n", runtime.GOMAXPROCS(0))
// n := 40
// fmt.Printf("Установка ограничения количества горутин на %d процессоров\n", n)
// runtime.GOMAXPROCS(n)

func main() {

	now := time.Now()
	t := now.Round(5 * time.Minute)
	detect := t.Unix() - now.Unix()
	log.Println(now.Unix())
	log.Println(t.Unix())
	log.Println(detect)
	if detect > 0 {
		req := now.Unix() + detect
		log.Println(req)
		//reqi := int(req)
		log.Println(time.Unix(req, 0))
		time.Sleep(time.Duration(req) * time.Millisecond)
		log.Println(time.Now().UTC())
	} else {
		req := t.Unix() + 300
		log.Println(req)
		log.Println(time.Unix(req, 0))
		time.Sleep(time.Duration(req) * time.Millisecond)
		log.Println(time.Now().UTC())
	}
}
