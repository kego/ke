package tree

import (
	"honnef.co/go/js/dom"
	"kego.io/system"
)

type Item interface {
	Initialise(*dom.HTMLDivElement)
}

type AsyncItem interface {
	LoadContent() chan bool
	ContentLoaded() bool
}

type HasNode interface {
	Node() *system.Node
}
