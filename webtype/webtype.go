package webtype

type HtmlElement interface {
	Html() string
}

type HtmlTag string

type Class []string
type Style map[string]string
type Attribute map[string]string
type Children []HtmlElement
