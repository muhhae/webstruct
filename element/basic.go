package element

import (
	"github.com/muhhae/webstruct/util"
	"github.com/muhhae/webstruct/webtype"
)

type CustomElement struct {
	Tag        webtype.HtmlTag
	Attributes webtype.Attribute
	Children   []webtype.HtmlElement
}

func (e CustomElement) Html() string {
	return util.FormatElement(e.Tag, e.Attributes, e.Children)
}

type Div struct {
	Id               string
	Style            webtype.Style
	Class            []string
	CustomAttributes webtype.Attribute
	Children         []webtype.HtmlElement
}

func (d Div) Html() string {
	return CustomElement{
		Tag: webtype.Div,
		Attributes: util.Combine(
			d.CustomAttributes,
			webtype.Attribute{
				"id":    d.Id,
				"style": util.StyleFormat(d.Style),
				"class": util.ListToString(d.Class),
			},
		),
		Children: d.Children,
	}.Html()
}

type Span struct {
	Id               string
	Style            map[string]string
	Class            []string
	CustomAttributes map[string]string
	Children         []webtype.HtmlElement
}

func (s Span) Html() string {
	return CustomElement{
		Tag: webtype.Span,
		Attributes: util.Combine(
			s.CustomAttributes,
			webtype.Attribute{
				"id":    s.Id,
				"style": util.StyleFormat(s.Style),
				"class": util.ListToString(s.Class),
			},
		),
		Children: s.Children,
	}.Html()
}

type RawText string

func (t RawText) Html() string {
	return string(t)
}
