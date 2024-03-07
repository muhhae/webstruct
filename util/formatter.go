package util

import (
	"fmt"
	"maps"

	"github.com/muhhae/webstruct/webtype"
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
		attribute += fmt.Sprintf(" %s=\"%s\"", key, value)
		count++
	}
	return attribute
}

func ChildrenFormat(c webtype.Children) string {
	children := make([]string, len(c))
	for i, e := range c {
		children[i] = e.Html()
	}
	return ListToString(children)
}

func StyleFormat(styles webtype.Style) string {
	style := ""
	for key, val := range styles {
		style += fmt.Sprintf("%s:%s", key, val)
	}
	return style
}

func FormatElement(tag webtype.HtmlTag, attribute webtype.Attribute, children webtype.Children) string {
	return fmt.Sprintf("<%s%s>%s</%s>", tag, AttrFormat(attribute), ChildrenFormat(children), tag)
}

func Combine[K comparable, V any](M ...map[K]V) map[K]V {
	m := map[K]V{}
	for i := 0; i < len(M); i++ {
		maps.Copy(m, M[i])
	}
	return m
}
