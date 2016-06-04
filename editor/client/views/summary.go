package views // import "kego.io/editor/client/views"

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/stores"
	"kego.io/flux"
	"kego.io/json"
	"kego.io/system"
	"kego.io/system/node"
)

type SummaryView struct {
	vecty.Composite
	ctx    context.Context
	app    *stores.App
	notifs chan flux.NotifPayload

	node   *node.Node
	origin *system.Reference
}

func NewSummaryView(ctx context.Context, node *node.Node, origin *system.Reference) *SummaryView {
	v := &SummaryView{
		ctx:    ctx,
		app:    stores.FromContext(ctx),
		node:   node,
		origin: origin,
	}
	v.Mount()
	return v
}

func (v *SummaryView) Reconcile(old vecty.Component) {
	if old, ok := old.(*SummaryView); ok {
		v.Body = old.Body
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *SummaryView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *SummaryView) Mount() {
	v.notifs = v.app.Branches.Watch(nil,
		nil,
	) //stores.BranchSelectPostLoad,

	go func() {
		for notif := range v.notifs {
			v.reaction(notif)
		}
	}()
}

func (v *SummaryView) reaction(notif flux.NotifPayload) {
	defer close(notif.Done)
	//v.branch = v.app.Branches.Selected()
	v.ReconcileBody()
}

func (v *SummaryView) Unmount() {
	if v.notifs != nil {
		v.app.Branches.Delete(v.notifs)
		v.notifs = nil
	}
	v.Body.Unmount()
}

func (v *SummaryView) render() vecty.Component {

	if v.node == nil || v.node.JsonType != json.J_OBJECT || len(v.node.Map) == 0 {
		return elem.Div()
	}

	head := elem.TableRow(
		elem.TableHeader(vecty.Text("name")),
		elem.TableHeader(vecty.Text("holds")),
		elem.TableHeader(vecty.Text("value")),
		elem.TableHeader(vecty.Text("options")),
	)

	rows := vecty.List{}
	for _, c := range v.node.Map {
		if *c.Origin != *v.origin {
			continue
		}
		rows = append(rows, NewSummaryRowView(v.ctx, c))
	}

	return elem.Div(
		elem.Table(
			prop.Class("table"),
			elem.TableHead(
				head,
			),
			elem.TableBody(
				rows,
			),
		),
	)
}
