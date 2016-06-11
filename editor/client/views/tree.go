package views // import "kego.io/editor/client/views"

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/flux"
)

type TreeView struct {
	vecty.Composite
	ctx    context.Context
	app    *stores.App
	notifs chan flux.NotifPayload

	Root     *models.BranchModel
	Selected *models.BranchModel
}

func NewTreeView(ctx context.Context) *TreeView {
	v := &TreeView{
		ctx: ctx,
		app: stores.FromContext(ctx),
	}
	v.Mount()
	return v
}

func (v *TreeView) Reconcile(old vecty.Component) {
	if old, ok := old.(*TreeView); ok {
		v.Body = old.Body
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *TreeView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *TreeView) Mount() {
	v.notifs = v.app.Watch(v.app.Branches.Root(),
		stores.BranchInitialStateLoaded,
	)
	go func() {
		for notif := range v.notifs {
			v.Root = v.app.Branches.Root()
			v.ReconcileBody()
			close(notif.Done)
		}
	}()
}

func (v *TreeView) Unmount() {
	if v.notifs != nil {
		v.app.Delete(v.notifs)
		v.notifs = nil
	}
	v.Body.Unmount()
}

func (v *TreeView) render() vecty.Component {
	if v.Root == nil {
		return elem.Div()
	}
	return elem.Div(
		prop.Class("content tree"),
		NewBranchView(v.ctx, v.Root),
	)
}
