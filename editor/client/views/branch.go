package views

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/flux"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
)

type BranchView struct {
	vecty.Composite
	ctx context.Context
	app *stores.App

	model         *models.BranchModel
	cSelect       chan struct{}
	cOpenPostLoad chan struct{}
	cClose        chan struct{}
	cMain         chan struct{}
	cOpen         chan struct{}
	children      vecty.List
}

func NewBranchView(ctx context.Context, model *models.BranchModel) *BranchView {
	if model == nil {
		return nil
	}
	v := &BranchView{
		ctx:   ctx,
		app:   stores.FromContext(ctx),
		model: model,
	}
	v.Mount()
	return v
}

func (v *BranchView) Reconcile(old vecty.Component) {
	if old, ok := old.(*BranchView); ok {
		v.Body = old.Body
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *BranchView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *BranchView) Mount() {
	v.cOpen = v.app.Branches.WatchSingle(stores.BranchOpen, v.model)
	v.cOpenPostLoad = v.app.Branches.WatchSingle(stores.BranchOpenPostLoad, v.model)
	v.cClose = v.app.Branches.WatchSingle(stores.BranchClose, v.model)
	v.cMain = v.app.Branches.Watch(v.model)
	v.cSelect = v.app.Branches.WatchSingle(stores.BranchSelect, v.model)

	go func() {
		for range v.cOpen {
			loaded := LoadBranch(v.ctx, v.app, v.model)
			v.app.Dispatch(&actions.BranchOpenPostLoad{Branch: v.model, Loaded: loaded})
		}
	}()
	go func() {
		for range flux.WatchMulti(v.cOpenPostLoad, v.cClose, v.cMain) {
			v.ReconcileBody()
		}
	}()
	go func() {
		for range v.cSelect {
			loaded := LoadBranch(v.ctx, v.app, v.model)
			if v.model.LastOp == models.BranchOpClickToggle && loaded {
				v.app.Dispatch(&actions.BranchOpen{Branch: v.model})
			}
			v.app.Dispatch(&actions.BranchSelectPostLoad{Branch: v.model, Loaded: loaded})
		}
	}()
}

func (v *BranchView) Unmount() {
	if v.cOpen != nil {
		v.app.Branches.Delete(v.cOpen)
		v.cOpen = nil
	}
	if v.cOpenPostLoad != nil {
		v.app.Branches.Delete(v.cOpenPostLoad)
		v.cOpenPostLoad = nil
	}
	if v.cClose != nil {
		v.app.Branches.Delete(v.cClose)
		v.cClose = nil
	}
	if v.cMain != nil {
		v.app.Branches.Delete(v.cMain)
		v.cMain = nil
	}
	if v.cSelect != nil {
		v.app.Branches.Delete(v.cSelect)
		v.cSelect = nil
	}
	v.Body.Unmount()
}

func (v *BranchView) render() vecty.Component {
	if v.model == nil {
		return elem.Div()
	}

	v.children = vecty.List{}
	if v.model.Open {
		for _, c := range v.model.Children {
			v.children = append(v.children, NewBranchView(v.ctx, c))
		}
	}

	return elem.Div(
		prop.Class("node"),
		NewBranchControlView(v.ctx, v.model),
		elem.Div(
			prop.Class("children"),
			v.children,
		),
	)
}
