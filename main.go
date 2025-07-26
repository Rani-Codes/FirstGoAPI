package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Inventory struct {
	Id   int
	Item string
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
	for _, inv := range InvList {
		fmt.Fprintf(w, "ID: %d, Item: %s\n", inv.Id, inv.Item)
	}
}

func GetItemById(w http.ResponseWriter, r *http.Request) {
	IdString := r.URL.Path[len("/items/"):]
	id, err := strconv.Atoi(IdString)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, inv := range InvList {
		if inv.Id == id {
			fmt.Fprintf(w, "Id %d returned: %s\n", id, inv.Item)
			return
		}
	}

	http.NotFound(w, r)
}

func main() {
	http.HandleFunc("/", HelloWorld)
	http.HandleFunc("/items", GetAllItems)
	http.HandleFunc("/items/", GetItemById)

	fmt.Println("Server running at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
