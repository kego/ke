package system

import (
	"testing"

	"kego.io/kerr/assert"
)

func TestObjectGetType(t *testing.T) {
	ty := &Type{Object_base: &Object_base{Id: NewReference("a.b/c", "foo")}}

	Register("a.b/c", "t", ty, 0)

	// Clean up for the tests - don't normally need to unregister types
	defer Unregister("a.b/c", "t")

	o := &Object_base{Type: NewReference("a.b/c", "t")}

	gt, ok := o.Type.GetType()
	assert.True(t, ok)
	assert.Equal(t, "a.b/c:foo", gt.Id.Value())

}