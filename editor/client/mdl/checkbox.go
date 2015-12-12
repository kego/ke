package mdl

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

type CheckboxStruct struct {
	*dom.HTMLLabelElement
	Input *dom.HTMLInputElement
	Span  *dom.HTMLSpanElement
}

func Checkbox(value bool, label string) *CheckboxStruct {

	id := randomId()

	c := &CheckboxStruct{}
	c.HTMLLabelElement = get("label").(*dom.HTMLLabelElement)
	c.Class().SetString("mdl-switch mdl-js-switch mdl-js-ripple-effect")
	c.For = id

	c.Input = get("input").(*dom.HTMLInputElement)
	c.Input.SetID(id)
	c.Input.Class().Add("mdl-switch__input")
	c.Input.Type = "checkbox"
	c.Input.Checked = value

	c.AppendChild(c.Input)

	c.Span = get("span").(*dom.HTMLSpanElement)
	c.Span.Class().Add("mdl-switch__label")
	c.Span.SetTextContent(label)
	c.AppendChild(c.Span)

	js.Global.Get("componentHandler").Call("upgradeElement", c)

	return c
}
