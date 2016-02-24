// info:{"Path":"kego.io/demo/demo3","Hash":14527783073471760224}
package demo3

import (
	"reflect"

	"golang.org/x/net/context"
	"kego.io/context/jsonctx"
	"kego.io/demo/demo3/images"
	"kego.io/system"
)

// Automatically created basic rule for page
type PageRule struct {
	*system.Object
	*system.Rule
}
type Page struct {
	*system.Object
	Hero  *images.Photo  `json:"hero"`
	Title *system.String `json:"title"`
}
type PageInterface interface {
	GetPage(ctx context.Context) *Page
}

func (o *Page) GetPage(ctx context.Context) *Page {
	return o
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/demo/demo3", 14527783073471760224)
	pkg.InitType("page", reflect.TypeOf((*Page)(nil)), reflect.TypeOf((*PageRule)(nil)), reflect.TypeOf((*PageInterface)(nil)).Elem())
}
