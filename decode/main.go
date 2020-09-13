package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

//Decode is very usefull func
func Decode(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))

	if err != nil {
		return nil, err
	}

	return data, nil
}

//DecodeHandler is handle /decode?
func DecodeHandler(w http.ResponseWriter, r *http.Request) {
	text := strings.Trim(r.FormValue("text"), " ")
	key := strings.Trim(r.FormValue("key"), " ")

	if text == "" || key == "" {
		fmt.Fprintln(w, "Incorrect request format")
		return
	}

	data, err := hex.DecodeString(text)

	if err != nil {
		log.Println(err)
		return
	}

	ciphertext, err := Decode([]byte(key), []byte(data))

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Fprintf(w, "%s", ciphertext)
}
func main() {
	log.Println("Start working")
	defer log.Println("Finish working")

	router := mux.NewRouter()
	router.HandleFunc("/decode", DecodeHandler)
	http.Handle("/", router)

	http.ListenAndServe(":8002", nil)
}
