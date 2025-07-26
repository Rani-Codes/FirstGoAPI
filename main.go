package main

import (
	"fmt"
	"log"
	"net/http"
)

type Inventory struct {
	id   int
	item string
}

var InvList = []Inventory{
	{id: 1, item: "banana"},
	{id: 2, item: "apple"},
	{id: 3, item: "pear"},
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func (i *Inventory) GetAllItems(w http.ResponseWriter, r *http.Request) {
	for j := 0; j < len(i.item); j++ {
		fmt.Fprintf(w, "The items inside of the inventory: %s", i.item[j])
	}
}

func main() {
	http.HandleFunc("/", HelloWorld)
	http.HandleFunc("/items", GetAllItems)

	fmt.Println("Server running at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
