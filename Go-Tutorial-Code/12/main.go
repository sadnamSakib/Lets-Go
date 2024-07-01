package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var info = `
<h1>User ID: {{.UserID}}</h1>
<h2>ID: {{.ID}}</h2>
<h3>Title: {{.Title}}</h3>
<h4>Completed: {{.Completed}}</h4>
`

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	baseUrl := "https://jsonplaceholder.typicode.com/"
	resp, err := http.Get(baseUrl + path)

	if err != nil {
		fmt.Println("Error", path, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var t todo
	if dErr := json.NewDecoder(resp.Body).Decode(&t); dErr != nil {
		fmt.Println("------------------------")
		panic(dErr)
	}

	fmt.Println("todo", t)
	tmpl := template.New("todo")
	tmpl, _ = tmpl.Parse(info)
	tmpl.Execute(w, t)
}
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
