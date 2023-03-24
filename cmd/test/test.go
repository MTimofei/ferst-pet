package main

import (
	"log"
	"net/http"
	"ser_identification/iternal/s_reg-ident/web"
)

func main() {

	/*type solt

	for i := 0; i < 8; i++ {
		rand2.Seed(time.Now().Unix())
		randNum := rand2.Intn(100)

	}
	var chars [94]byte

	// добавляем латинские символы
	for i := 33; i <= 126; i++ {
		chars[i-33] = byte(i)
	}

	// добавляем специальные символы
	specialChars := []byte{'!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '_', '+', '-', '=', '{', '}', '[', ']', ':', ';', '<', '>', ',', '.', '/', '?', '|', '\\', '~', '`'}
	for i := 0; i < len(specialChars); i++ {
		chars[62+i] = specialChars[i]
	}

	// выводим массив
	fmt.Println(chars)

	fmt.Println(randNum)
	b := []byte(string(rand2.Uint32()))
	fmt.Println(b)
	var key string
	for _, i := range b {
		key += string(i)
	}
	fmt.Println(key)*/

	log.Fatal(http.ListenAndServeTLS(":8888", "cF.crt", "kF.key", web.Router()))
}
