package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", usersHandler)

	log.Fatalln(http.ListenAndServe(":5001", nil))
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:5000/users")
	if err != nil {
		log.Println("Error during GET to /users:", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var allUsers []map[string]string
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&allUsers)
	if err != nil {
		log.Println("Error decoding /users response:", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	queryValues := r.URL.Query()
	name := queryValues.Get("name")
	t := queryValues.Get("type")
	if name != "" {
		usersByName(w, name, allUsers)
		return
	} else if t != "" {
		usersByType(w, t, allUsers)
		return
	}

	returnResult(w, allUsers)
}

func usersByName(w http.ResponseWriter, name string, users []map[string]string) {
	users = filterUsers(users, "name", name)
	if len(users) == 0 {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	returnResult(w, users[0])
}

func usersByType(w http.ResponseWriter, t string, users []map[string]string) {
	users = filterUsers(users, "type", t)
	returnResult(w, users)
}

func filterUsers(users []map[string]string, key, value string) (out []map[string]string) {
	out = make([]map[string]string, 0, len(users))
	for _, x := range users {
		if x[key] == value {
			out = append(out, x)
		}
	}
	return out
}

func returnResult(w http.ResponseWriter, result interface{}) {
	js, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
