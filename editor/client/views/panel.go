package views // import "kego.io/editor/client/views"

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/flux"
	"kego.io/system/node"
)

type PanelView struct {
	vecty.Composite
	ctx            context.Context
	app            *stores.App
	notifsBranches chan flux.NotifPayload
	notifsEditors  chan flux.NotifPayload

	branch *models.BranchModel
	node   *node.Node
}

func NewPanelView(ctx context.Context) *PanelView {
	v := &PanelView{
		ctx: ctx,
		app: stores.FromContext(ctx),
	}
	v.Mount()
	return v
}

func (v *PanelView) Reconcile(old vecty.Component) {
	if old, ok := old.(*PanelView); ok {
		v.Body = old.Body
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *PanelView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *PanelView) Mount() {
	v.notifsBranches = v.app.Branches.Watch(nil,
		stores.BranchSelectPostLoad,
	)
	go func() {
		for notif := range v.notifsBranches {
			v.reaction(notif)
		}
	}()

	v.notifsEditors = v.app.Editors.Watch(nil,
		stores.EditorLoaded,
		stores.EditorAdded,
		stores.EditorInitialStateLoaded,
	)
	go func() {
		for notif := range v.notifsEditors {
			v.reaction(notif)
		}
	}()
}

func (v *PanelView) reaction(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.branch = v.app.Branches.Selected()
	v.node = v.app.Nodes.Selected()
	v.ReconcileBody()
	v.Node().Get("parentNode").Set("scrollTop", "0")
}

func (v *PanelView) Unmount() {
	if v.notifsBranches != nil {
		v.app.Branches.Delete(v.notifsBranches)
		v.notifsBranches = nil
	}
	if v.notifsEditors != nil {
		v.app.Editors.Delete(v.notifsEditors)
		v.notifsEditors = nil
	}
	v.Body.Unmount()
}

func (v *PanelView) render() vecty.Component {
	var editor, breadcrumbs vecty.Component
	if v.branch != nil {
		breadcrumbs = NewBreadcrumbsView(v.ctx, v.branch)
	}
	if v.node != nil {
		if v.node.Type.IsNativeMap() {
			editor = NewMapView(v.ctx, v.node)
		} else if v.node.Type.IsNativeArray() {
			editor = NewArrayView(v.ctx, v.node)
		} else {
			editor = NewCompositeView(v.ctx, v.node)
		}
	}

	return elem.Div(
		prop.Class("content panel"),
		breadcrumbs,
		editor,
	)
}
