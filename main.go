package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var files, _ = ioutil.ReadDir("views")
var templates = template.Must(template.ParseFiles("views/index.html", "views/buy.html", "views/sell.html", "views/trade.html"))

func main() {

	fmt.Println("SERVER STARTING on port 8080")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/buy", buyHandler)
	http.HandleFunc("/sell", sellHandler)
	http.HandleFunc("/trade", tradeHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

//Page is the data for a server rendered web page
type Page struct {
	Title string
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles("views/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//This is way more efficient, but prevents hot reload.
	// err := templates.ExecuteTemplate(w, tmpl+".html", p)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", &Page{Title: "Home"})
}

func buyHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "buy", &Page{Title: "Buy"})
}

func sellHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "sell", &Page{Title: "Sell"})
}

func tradeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "trade", &Page{Title: "Trade"})
}

func getPostHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)

	switch r.Method {
	case "GET":
		post := Post{
			Type: ForSale,
		}
		var vpost = newVehiclePost(post, 2012, 80000, "Lexus", "GS 350", "myVin")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		enc.Encode(vpost)

	case "POST":
		fmt.Fprintln(w, "No impl")
	case "PUT":
		fmt.Fprintln(w, "No impl")
	case "DELETE":
		fmt.Fprintln(w, "No impl")
	case "PATCH":
		fmt.Fprintln(w, "No impl")
	}

}
