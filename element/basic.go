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
	Id               string
	Style            webtype.Style
	Class            webtype.Class
	CustomAttributes webtype.Attribute
	Children         webtype.Children
}

func (s Span) ToCustomElement() CustomElement {
	return CustomElement{
		Tag: "span",
		Attributes: util.Combine(
			s.CustomAttributes,
			webtype.Attribute{
				"id":    s.Id,
				"style": util.StyleFormat(s.Style),
				"class": util.ListToString(s.Class),
			},
		),
		Children: s.Children,
	}
}

func (s Span) Html() string {
	return s.ToCustomElement().Html()
}

type P struct {
	Id               string
	Style            webtype.Style
	Class            webtype.Class
	CustomAttributes webtype.Attribute
	Children         webtype.Children
}

func (p P) ToCustomElement() CustomElement {
	return CustomElement{
		Tag: "p",
		Attributes: util.Combine(
			p.CustomAttributes,
			webtype.Attribute{
				"id":    p.Id,
				"style": util.StyleFormat(p.Style),
				"class": util.ListToString(p.Class),
			},
		),
		Children: p.Children,
	}
}

func (p P) Html() string {
	return p.ToCustomElement().Html()
}

type RawText string

func (t RawText) Html() string {
	return string(t)
}

type A struct {
	Id               string
	Style            webtype.Style
	Class            webtype.Class
	CustomAttributes webtype.Attribute
	Children         webtype.Children
	Href             string
}

func (a A) ToCustomElement() CustomElement {
	return CustomElement{
		Tag: "a",
		Attributes: util.Combine(
			a.CustomAttributes,
			webtype.Attribute{
				"id":    a.Id,
				"style": util.StyleFormat(a.Style),
				"class": util.ListToString(a.Class),
				"href":  a.Href,
			},
		),
		Children: a.Children,
	}
}

func (a A) Html() string {
	return a.ToCustomElement().Html()
}

type H1 struct {
	Id               string
	Style            webtype.Style
	Class            webtype.Class
	CustomAttributes webtype.Attribute
	Children         webtype.Children
}

func (h1 H1) ToCustomElement() CustomElement {
	return CustomElement{
		Tag: "h1.v",
		Attributes: util.Combine(
			h1.CustomAttributes,
			webtype.Attribute{
				"id":    h1.Id,
				"style": util.StyleFormat(h1.Style),
				"class": util.ListToString(h1.Class),
			},
		),
		Children: h1.Children,
	}
}
func (h1 H1) Html() string {
	return h1.ToCustomElement().Html()
}

type H2 struct {
	Id               string
	Style            webtype.Style
	Class            webtype.Class
	CustomAttributes webtype.Attribute
	Children         webtype.Children
}

func (h2 H2) ToCustomElement() CustomElement {
	return CustomElement{
		Tag: "h2.v",
		Attributes: util.Combine(
			h2.CustomAttributes,
			webtype.Attribute{
				"id":    h2.Id,
				"style": util.StyleFormat(h2.Style),
				"class": util.ListToString(h2.Class),
			},
		),
		Children: h2.Children,
	}
}
func (h2 H2) Html() string {
	return h2.ToCustomElement().Html()
}

type H3 struct {
	Id               string
	Style            webtype.Style
	Class            webtype.Class
	CustomAttributes webtype.Attribute
	Children         webtype.Children
}

func (h2 H3) ToCustomElement() CustomElement {
	return CustomElement{
		Tag: "h2.v",
		Attributes: util.Combine(
			h2.CustomAttributes,
			webtype.Attribute{
				"id":    h2.Id,
				"style": util.StyleFormat(h2.Style),
				"class": util.ListToString(h2.Class),
			},
		),
		Children: h2.Children,
	}
}
func (h2 H3) Html() string {
	return h2.ToCustomElement().Html()
}
