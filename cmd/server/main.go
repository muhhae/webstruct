package main

import (
	"log"
	"net/http"

	"github.com/muhhae/webstruct/pkg/element"
	"github.com/muhhae/webstruct/pkg/webtype"
)

func main() {
	div := element.Div{
		Class: []string{
			"tes",
			"example",
		},
		Id: "test-id",
		Style: map[string]string{
			"color": "red",
		},
		CustomAttributes: map[string]string{
			"custom-attr": "customm-attributes",
		},
		Children: []webtype.HtmlElement{
			element.Span{
				Children: []webtype.HtmlElement{
					element.CustomElement{
						Tag: "button",
						Children: []webtype.HtmlElement{
							element.RawText("Button"),
						},
					},
					element.RawText("Hello World"),
				},
			},
		},
	}

	// fmt.Println(div.Html())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(div.Html()))
		if err != nil {
			panic(err)
		}
	})
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
