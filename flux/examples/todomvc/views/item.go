package views

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/event"
	"github.com/davelondon/vecty/prop"
	"github.com/davelondon/vecty/style"
	"golang.org/x/net/context"
	"kego.io/flux/examples/todomvc/actions"
	"kego.io/flux/examples/todomvc/model"
	"kego.io/flux/examples/todomvc/stores"
)

type ItemView struct {
	vecty.Composite
	ctx context.Context
	app *stores.App

	Index     int
	Item      *model.Item
	editing   bool
	editTitle string
}

func NewItemView(ctx context.Context, index int, item *model.Item) *ItemView {
	return &ItemView{
		ctx:   ctx,
		app:   stores.FromContext(ctx),
		Index: index,
		Item:  item,
	}
}

func (p *ItemView) Apply(element *vecty.Element) {
	element.AddChild(p)
}

func (p *ItemView) Reconcile(old vecty.Component) {
	if old, ok := old.(*ItemView); ok {
		p.Body = old.Body
		p.Index = old.Index
		p.Item = old.Item
		p.editing = old.editing
		p.editTitle = old.editTitle
	}
	p.RenderFunc = p.render
	p.ReconcileBody()
}

func (p *ItemView) onDestroy(event *vecty.Event) {
	go func() {
		p.app.Dispatcher.Dispatch(&actions.Destroy{Index: p.Index})
	}()
}

func (p *ItemView) onToggleCompleted(event *vecty.Event) {
	go func() {
		p.app.Dispatcher.Dispatch(&actions.SetComplete{Index: p.Index, Complete: event.Target.Get("checked").Bool()})
	}()
}

func (p *ItemView) onStartEdit(event *vecty.Event) {
	p.editing = true
	p.editTitle = p.Item.Title
	p.ReconcileBody()
}

func (p *ItemView) onEditInput(event *vecty.Event) {
	p.editTitle = event.Target.Get("value").String()
	p.ReconcileBody()
}

func (p *ItemView) onStopEdit(event *vecty.Event) {
	p.editing = false
	p.ReconcileBody()
	go func() {
		p.app.Dispatcher.Dispatch(&actions.UpdateText{Index: p.Index, Text: p.editTitle})
	}()
}

func (p *ItemView) render() vecty.Component {
	return elem.ListItem(
		vecty.ClassMap{
			"completed": p.Item.Complete,
			"editing":   p.editing,
		},

		elem.Div(
			prop.Class("view"),

			elem.Input(
				prop.Class("toggle"),
				prop.Type(prop.TypeCheckbox),
				prop.Checked(p.Item.Complete),
				event.Change(p.onToggleCompleted),
			),
			elem.Label(
				vecty.Text(p.Item.Title),
				event.DoubleClick(p.onStartEdit),
			),
			elem.Button(
				prop.Class("destroy"),
				event.Click(p.onDestroy),
			),
		),
		elem.Form(
			style.Margin(style.Px(0)),
			event.Submit(p.onStopEdit).PreventDefault(),
			elem.Input(
				prop.Class("edit"),
				prop.Value(p.editTitle),
				event.Input(p.onEditInput),
			),
		),
	)
}