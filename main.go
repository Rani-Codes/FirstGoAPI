package main

import (
	"fmt"
	"log"
	"net/http"
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

func main() {
	http.HandleFunc("/", HelloWorld)
	http.HandleFunc("/items", GetAllItems)

	fmt.Println("Server running at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
