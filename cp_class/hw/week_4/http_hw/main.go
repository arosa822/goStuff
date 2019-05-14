package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/add", db.add)
	http.HandleFunc("/remove", db.remove)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) //404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if _, ok := db[item]; ok {
		f, err := strconv.ParseFloat(price, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest) // 404
			fmt.Fprintf(w, "cannot parse price: %v\n", err)
		}
		if f > 0 {
			db[item] = dollars(f)
		} else {
			w.WriteHeader(http.StatusBadRequest) //404
			fmt.Fprintf(w, "price must be > 0: %f\n", f)
		}
	}
}

func (db database) add(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	f, err := strconv.ParseFloat(price, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "cannot parse price: %v\n", err)
	}
	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusBadRequest) // 404
		fmt.Fprintf(w, "item already exists in database")
	} else if f > 0 {
		db[item] = dollars(f)
	} else {
		w.WriteHeader(http.StatusBadRequest) //404
		fmt.Fprintf(w, "price must be > 0: %f\n", f)

	}
}

func (db database) remove(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		delete(db, item)
	} else {
		w.WriteHeader(http.StatusBadRequest) // 404
		fmt.Fprintf(w, "item was not found in database...")
	}

}
