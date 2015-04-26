package system

import (
	"fmt"

	"kego.io/json"
)

type Bool struct {
	Value  bool
	Exists bool
}

func (out *Bool) UnmarshalJSON(in []byte, path string, imports map[string]string) error {
	var b *bool
	if err := json.UnmarshalPlain(in, &b, path, imports); err != nil {
		return err
	}
	if b == nil {
		out.Exists = false
	} else {
		out.Exists = true
		out.Value = *b
	}
	return nil
}

func (b *Bool) MarshalJSON() ([]byte, error) {
	if !b.Exists {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("%v", b.Value)), nil
}
