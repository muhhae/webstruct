package webtype

type HtmlElement interface {
	Html() string
}

type HtmlTag string

const (
	Div  = "div"
	Span = "span"
)

type Style map[string]string
type Attribute map[string]string
