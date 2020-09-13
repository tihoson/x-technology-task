package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"./types"
	"github.com/gorilla/mux"
)

//URL is url
const URL = "https://iss.moex.com/iss/history/engines/stock/markets/shares/boards/tqbr/securities.xml"

//CourseCache is course cache
var CourseCache types.Cache

//GetDocumnet is very usefull funct
func GetDocumnet() (document *types.Document, err error) {
	now := time.Now()

	CourseCache.RWMutex.Lock()
	defer CourseCache.RWMutex.Unlock()

	if now.Sub(CourseCache.LastUpdate).Hours()/24 > 1 {
		log.Println("Try to update data")

		doc, err := GetNewDocumnet()

		if err != nil {
			return nil, err
		}

		for _, data := range doc.Data {
			if data.ID == "history" {
				if (types.Row{}) != data.Rows.Rows[0] {
					log.Println("Updated!")
					CourseCache.LastUpdate = now
					CourseCache.Data = *doc
				}
			}
		}
	}

	return &CourseCache.Data, nil
}

//GetNewDocumnet is very usefull function
func GetNewDocumnet() (document *types.Document, err error) {
	resp, err := http.Get(URL)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var doc types.Document

	xml.Unmarshal(body, &doc)

	return &doc, nil
}

//GetCourses handle /course
func GetCourses(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	log.Println("Course request")

	doc, err := GetDocumnet()

	if err != nil {
		log.Println(err)
	}

	for _, data := range doc.Data {
		if data.ID == "history" {
			json, err := json.Marshal(data.Rows.Rows)

			if err != nil {
				log.Println(err)
				return
			}

			w.Write(json)
		}
	}
}

func main() {
	CourseCache = types.Cache{
		LastUpdate: time.Time{},
		Data:       types.Document{},
		RWMutex:    sync.RWMutex{},
	}

	log.Println("Start working")
	defer log.Println("Finish working")

	router := mux.NewRouter()
	router.HandleFunc("/course", GetCourses)
	http.Handle("/", router)

	http.ListenAndServe(":8000", nil)
}
