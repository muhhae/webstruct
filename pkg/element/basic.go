package element

import (
	formatter "github.com/muhhae/webstruct/pkg/util"
	"github.com/muhhae/webstruct/pkg/webtype"
)

type Div struct {
	Id         string
	Style      map[string]string
	Class      []string
	Attributes map[string]string
	Children   []webtype.HtmlElement
}

func (d *Div) Html() string {
	d.Attributes["class"] = formatter.ListToString(d.Class)
	d.Attributes["style"] = formatter.StyleFormat(d.Style)
	return formatter.CreateElement(webtype.Div, formatter.AttrFormat(d.Attributes), formatter.ChildrenFormat(d.Children))
}

type RawText struct {
	Text string
}

func (t RawText) Html() string {
	return t.Text
}
