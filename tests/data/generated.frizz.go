package data

import (
	frizz "frizz.io/frizz"
	packer "frizz.io/tests/packer"
	sub "frizz.io/tests/packer/sub"
)

const Imports imports = 0

type imports int

func (i imports) Path() string {
	return "frizz.io/tests/data"
}
func (i imports) Add(packers map[string]frizz.Packer, types map[string]frizz.Typer) {
	sub.Imports.Add(packers, types)
	packer.Imports.Add(packers, types)
}
