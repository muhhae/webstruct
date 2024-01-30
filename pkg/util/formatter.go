package util

import (
	"fmt"

	"github.com/muhhae/webstruct/pkg/webtype"
)

func ListToString(strs []string) string {
	str := ""
	for i, s := range strs {
		if i != 0 {
			str += " "
		}
		str += s
	}
	return str
}

func AttrFormat(a map[string]string) string {
	attribute := ""
	count := 0
	for key, value := range a {
		if value == "" {
			continue
		}
		attribute += fmt.Sprintf("%s=\"%s\"", key, value)
		count++
		if count < len(a) {
			attribute += " "
		}
	}
	return attribute
}

func ChildrenFormat(c []webtype.HtmlElement) string {
	children := make([]string, len(c))
	for i, e := range c {
		children[i] = e.Html()
	}
	return ListToString(children)
}

func StyleFormat(styles map[string]string) string {
	style := ""
	for key, val := range styles {
		style += fmt.Sprintf("%s:%s", key, val)
	}
	return style
}

func FormatElement(tag string, attribute string, children string) string {
	return fmt.Sprintf("<%s %s>%s</%s>", tag, attribute, children, tag)
}
