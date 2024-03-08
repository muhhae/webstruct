package element

import (
	"github.com/muhhae/webstruct/util"
	"github.com/muhhae/webstruct/webtype"
)

type CustomElement struct {
	Tag        webtype.HtmlTag
	Attributes webtype.Attribute
	Children   webtype.Children
}

func (c CustomElement) Html() string {
	return util.FormatElement(c.Tag, c.Attributes, c.Children)
}

type Div struct {
	Id               string
	Style            webtype.Style
	Class            webtype.Class
	CustomAttributes webtype.Attribute
	Children         webtype.Children
}

func (d Div) ToCustomElement() CustomElement {
	return CustomElement{
		Tag: "div",
		Attributes: util.Combine(
			d.CustomAttributes,
			webtype.Attribute{
				"id":    d.Id,
				"style": util.StyleFormat(d.Style),
				"class": util.ListToString(d.Class),
			},
		),
		Children: d.Children,
	}
}

func (d Div) Html() string {
	return d.ToCustomElement().Html()
}

type Span struct {
	Div
}

func (s Span) Html() string {
	e := s.ToCustomElement()
	e.Tag = "span"
	return e.Html()
}

type P struct {
	Div
}

func (p P) Html() string {
	e := p.ToCustomElement()
	e.Tag = "p"
	return e.Html()
}

type RawText string

func (t RawText) Html() string {
	return string(t)
}

type A struct {
	Div
	Href string
}

func (a A) Html() string {
	e := a.ToCustomElement()
	e.Tag = "a"
	e.Attributes["href"] = a.Href
	return a.Html()
}

type H1 struct {
	Div
}

func (h1 H1) Html() string {
	e := h1.ToCustomElement()
	e.Tag = "h1"
	return e.Html()
}

type H2 struct {
	Div
}

func (h2 H2) Html() string {
	e := h2.ToCustomElement()
	e.Tag = "h2"
	return e.Html()
}

type H3 struct {
	Div
}

func (h3 H3) Html() string {
	e := h3.ToCustomElement()
	e.Tag = "h3"
	return e.Html()
}
