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
	json.NewEncoder(w).Encode("Hello World")
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

func CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Method Not Allowed, Use POST"})
		return
	}

	var newItem Inventory
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid Request Body"})
		return
	}

	InvList = append(InvList, newItem)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)

}

func DeleteItemById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	IdString := r.URL.Path[len("/deleteItems/"):]
	id, err := strconv.Atoi(IdString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid ID"})
		return
	}

	for i, inv := range InvList {
		if inv.Id == id {
			InvList = append(InvList[:i], InvList[i+1:]...)
			json.NewEncoder(w).Encode(map[string]string{"message": "deleted"})
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
	http.HandleFunc("/createItems", CreateItem)
	http.HandleFunc("/deleteItems/", DeleteItemById)

	fmt.Println("Server running at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
