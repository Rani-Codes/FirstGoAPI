package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Inventory struct {
	Id   int    `json:"id"`
	Item string `json:"item"`
}

type ErrorResponse struct {
	Error string
}

var InvList = []Inventory{
	{Id: 1, Item: "banana"},
	{Id: 2, Item: "apple"},
	{Id: 3, Item: "pear"},
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(InvList)
}

func GetItemById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	IdString := r.URL.Path[len("/items/"):]
	id, err := strconv.Atoi(IdString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid ID"})
		return
	}

	for _, inv := range InvList {
		if inv.Id == id {
			json.NewEncoder(w).Encode(inv)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(ErrorResponse{Error: "Page Not Found"})
}

func main() {
	http.HandleFunc("/", HelloWorld)
	http.HandleFunc("/items", GetAllItems)
	http.HandleFunc("/items/", GetItemById)

	fmt.Println("Server running at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
