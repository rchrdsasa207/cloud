package main

import (
	"io"
	"log"
	"net/http"
)

var db *DB

func handleHttp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		raw, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading client's request: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(db.ServePOST(raw, r.Header.Get("Authorization")))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}

func main() {
	db = OpenDB()
	http.HandleFunc("/", handleHttp)
	log.Println("Server is running at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error in server runtime: ", err)
	}
	db.Close()
}
