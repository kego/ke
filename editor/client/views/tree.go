package views // import "kego.io/editor/client/views"

import (
	"code.google.com/p/go.net/context"
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
)

type TreeView struct {
	vecty.Composite
	ctx context.Context
	app *stores.App

	Root     *models.BranchModel
	Selected *models.BranchModel
}

func (v *TreeView) Reconcile(old vecty.Component) {
	if old, ok := old.(*TreeView); ok {
		v.Body = old.Body
		v.Selected = old.Selected
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

func NewTreeView(ctx context.Context) *TreeView {
	p := &TreeView{
		ctx: ctx,
		app: stores.FromContext(ctx),
	}

	go func() {
		for range p.app.Branches.Watch(p.app.Branches.Root()) {
			p.Root = p.app.Branches.Root()
			p.ReconcileBody()
		}
	}()

	return p
}

// Apply implements the vecty.Markup interface.
func (p *TreeView) Apply(element *vecty.Element) {
	element.AddChild(p)
}

func (p *TreeView) render() vecty.Component {
	if p.Root == nil {
		return elem.Div()
	}
	return elem.Div(
		prop.Class("content tree"),
		NewBranchView(p.ctx, p.Root),
	)
}