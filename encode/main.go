package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

//Encode is very usefull func
func Encode(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))

	return ciphertext, nil
}

//EncodeHandler is handle /encode?
func EncodeHandler(w http.ResponseWriter, r *http.Request) {
	text := strings.Trim(r.FormValue("text"), " ")
	key := strings.Trim(r.FormValue("key"), " ")

	if text == "" || key == "" {
		fmt.Fprintln(w, "Incorrect request format")
		return
	}

	ciphertext, err := Encode([]byte(key), []byte(text))

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Fprintf(w, "%0x", ciphertext)
}

func main() {
	log.Println("Start working")
	defer log.Println("Finish working")

	router := mux.NewRouter()
	router.HandleFunc("/encode", EncodeHandler)
	http.Handle("/", router)

	http.ListenAndServe(":8001", nil)
}
