package element

import (
	"github.com/muhhae/webstruct/pkg/util"
	"github.com/muhhae/webstruct/pkg/webtype"
)

type CustomElement struct {
	Tag        string
	Attributes map[string]string
	Children   []webtype.HtmlElement
}

func (e CustomElement) Html() string {
	return util.FormatElement(e.Tag, util.AttrFormat(e.Attributes), util.ChildrenFormat(e.Children))
}

type Div struct {
	Id               string
	Style            map[string]string
	Class            []string
	CustomAttributes map[string]string
	Children         []webtype.HtmlElement
}

func (d Div) Html() string {
	if d.CustomAttributes == nil {
		d.CustomAttributes = make(map[string]string)
	}
	d.CustomAttributes["class"] = util.ListToString(d.Class)
	d.CustomAttributes["style"] = util.StyleFormat(d.Style)
	return util.FormatElement(webtype.Div, util.AttrFormat(d.CustomAttributes), util.ChildrenFormat(d.Children))
}

type Span struct {
	Id               string
	Style            map[string]string
	Class            []string
	CustomAttributes map[string]string
	Children         []webtype.HtmlElement
}

func (s Span) Html() string {
	if s.CustomAttributes == nil {
		s.CustomAttributes = make(map[string]string)
	}
	s.CustomAttributes["class"] = util.ListToString(s.Class)
	s.CustomAttributes["style"] = util.StyleFormat(s.Style)

	if s.CustomAttributes == nil {
		s.CustomAttributes = map[string]string{}
	}
	return util.FormatElement(webtype.Span, util.AttrFormat(s.CustomAttributes), util.ChildrenFormat(s.Children))
}

type RawText string

func (t RawText) Html() string {
	return string(t)
}
