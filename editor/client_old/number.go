package editor

import (
	"strconv"

	"golang.org/x/net/context"
	"honnef.co/go/js/dom"
	"kego.io/editor/client_old/mdl"
)

type NumberEditor struct {
	*Node
	*Editor
	original float64
	textbox  *mdl.TextboxStruct
}

var _ EditorInterface = (*NumberEditor)(nil)

func NewNumberEditor(n *Node) *NumberEditor {
	return &NumberEditor{Node: n, Editor: &Editor{}}
}

func (e *NumberEditor) Layout() Layout {
	return Inline
}

func (e *NumberEditor) Initialize(ctx context.Context, holder BranchInterface, layout Layout, fail chan error) error {

	e.Editor.Initialize(ctx, holder, layout, fail)

	e.original = e.ValueNumber

	e.textbox = mdl.Textbox(strconv.FormatFloat(e.ValueNumber, 'f', -1, 64), e.Node.Key)
	e.AppendChild(e.textbox)
	e.textbox.Input.AddEventListener("input", true, func(ev dom.Event) {
		n, err := strconv.ParseFloat(e.textbox.Input.Value, 64)
		if err != nil {
			// display a validation error
		}
		e.update(n)
		e.Send(e)
	})

	return nil
}

func (e *NumberEditor) update(n float64) {
	e.Missing = false
	e.Null = false
	e.ValueNumber = n
	e.Node.Value = n
}

func (e *NumberEditor) Dirty() bool {
	return e.ValueNumber != e.original
}

func (e *NumberEditor) Focus() {
	e.textbox.Input.Focus()
}

func (e *NumberEditor) Value() interface{} {
	return e.Node.Value
}