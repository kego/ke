package views

import (
	"code.google.com/p/go.net/context"
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"kego.io/context/envctx"
	"kego.io/editor/client/stores"
)

type Header struct {
	vecty.Composite
	ctx context.Context
	app *stores.App

	Environment *envctx.Env
}

func NewHeader(ctx context.Context, env *envctx.Env) *Header {
	return &Header{
		ctx:         ctx,
		app:         stores.FromContext(ctx),
		Environment: env,
	}
}

func (v *Header) Reconcile(old vecty.Component) {
	if old, ok := old.(*PageView); ok {
		v.Body = old.Body
		v.Environment = old.Environment
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *Header) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *Header) render() vecty.Component {
	return elem.Navigation(
		prop.Class("navbar navbar-inverse navbar-fixed-top"),
		elem.Div(
			prop.Class("container-fluid"),
			elem.Div(
				prop.Class("navbar-header"),
				elem.Button(
					prop.Type("button"),
					prop.Class("navbar-toggle collapsed"),
					vecty.Data("toggle", "collapse"),
					vecty.Data("target", "#navbar"),
					elem.Span(
						prop.Class("sr-only"),
						vecty.Text("Toggle navigation"),
					),
					elem.Span(
						prop.Class("icon-bar"),
					),
					elem.Span(
						prop.Class("icon-bar"),
					),
					elem.Span(
						prop.Class("icon-bar"),
					),
				),
				elem.Anchor(
					prop.Class("navbar-brand"),
					prop.Href("#"),
					vecty.Text(v.Environment.Path),
				),
			),
			elem.Div(
				prop.ID("navbar"),
				prop.Class("navbar-collapse collapse"),
				elem.UnorderedList(
					prop.Class("nav navbar-nav"),
					elem.ListItem(
						prop.Class("active"),
						elem.Anchor(
							prop.Href("#"),
							vecty.Text("Home"),
						),
					),
					elem.ListItem(
						elem.Anchor(
							prop.Href("#"),
							vecty.Text("About"),
						),
					),
					elem.ListItem(
						prop.Class("dropdown"),
						elem.Anchor(
							prop.Href("#"),
							prop.Class("dropdown-toggle"),
							vecty.Data("toggle", "dropdown"),
							vecty.Text("Dropdown"),
							elem.Span(
								prop.Class("caret"),
							),
						),
						elem.UnorderedList(
							prop.Class("dropdown-menu"),
							elem.ListItem(
								elem.Anchor(
									prop.Href("#"),
									vecty.Text("Action"),
								),
							),
							elem.ListItem(
								elem.Anchor(
									prop.Href("#"),
									vecty.Text("Another action"),
								),
							),
							elem.ListItem(
								prop.Class("divider"),
							),
							elem.ListItem(
								prop.Class("dropdown-header"),
								vecty.Text("Nav header"),
							),
							elem.ListItem(
								elem.Anchor(
									prop.Href("#"),
									vecty.Text("Separated"),
								),
							),
						),
					),
				),
				elem.Form(
					prop.Class("navbar-form navbar-right"),
					elem.Input(
						prop.Type("text"),
						prop.Class("form-control"),
						prop.Placeholder("Search..."),
					),
				),
			),
		),
	)
}
