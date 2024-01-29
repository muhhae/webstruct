package main

import (
	"gohtml/pkg/element"
	"log"
	"net/http"
)

func main() {

	div := element.Div{
		Class: []string{"tes", "example"},
		Attributes: map[string]string{
			"id":    "test",
			"style": "color: red",
		},
		Children: []element.HtmlElement{
			element.RawText{
				Text: "Hello World",
			},
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(div.Html()))
		if err != nil {
			panic(err)
		}
	})
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
