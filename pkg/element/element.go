package element

import (
	"fmt"
)

type HtmlElement interface {
	Html() string
}

type Div struct {
	Class      []string
	Attributes map[string]string
	Children   []HtmlElement
}

func (d *Div) Html() string {
	classStrList := ""
	for i, class := range d.Class {
		if i != 0 {
			classStrList += " "
		}
		classStrList += class
	}
	class := ""
	if classStrList != "" {
		class = fmt.Sprintf("class=\"%s\"", classStrList)
	}
	attribute := ""
	if len(d.Attributes) > 0 {
		count := 0
		for key, value := range d.Attributes {
			attribute += fmt.Sprintf("%s=\"%s\"", key, value)
			count++
			if count < len(d.Attributes) {
				attribute += " "
			}
		}
	}
	children := ""
	for _, c := range d.Children {
		children += c.Html()
	}
	return fmt.Sprintf("<div %s %s>%s</div>", class, attribute, children)
}

type RawText struct {
	Text string
}

func (t RawText) Html() string {
	return t.Text
}
