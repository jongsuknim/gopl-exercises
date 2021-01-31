package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {

	var itemList = template.Must(template.New("itemList").Parse(`
		<table>
		<tr style='text-align: left'>
		  <th>Item</th>
		  <th>Price</th>
		</tr>
		{{range $key, $value := .}}
		<tr>
		  <td>{{ $key }}</td>
		  <td>{{ $value }}</td>
		</tr>
		{{end}}
		</table>
		`))

	if err := itemList.Execute(w, db); err != nil {
		log.Fatal(err)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
