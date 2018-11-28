package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", usersHandler)

	log.Fatalln(http.ListenAndServe(":5000", nil))
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(allUsers)
	if err != nil {
		log.Println("Error marshaling users response:", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

var (
	allUsers = []map[string]interface{}{
		map[string]interface{}{"name": "user0", "type": "type0"},
		map[string]interface{}{"name": "user1", "type": "type0"},
		map[string]interface{}{"name": "user2", "type": "type1"},
		map[string]interface{}{"name": "user3", "type": "type1"},
	}
)
