package system

import (
	"fmt"
	"sync"

	"strings"

	kegofmt "kego.io/fmt"
	"kego.io/uerr"
)

var types struct {
	sync.RWMutex
	m map[string]*Type
}
var packages struct {
	sync.RWMutex
	m map[string]*packageInfo
}

type packageInfo struct {
	getSystemTypes func() map[string]*Type
}

var typesInit = false

// initTypes initialises the system types that are needed for code generation and validation.
// They aren't needed to simply unmarshal data, so we don't initialise them on init().
func initTypes() {
	if typesInit {
		return
	}
	packages := getPackages()
	for _, p := range packages {
		registerTypes(p.getSystemTypes())
	}
	typesInit = true
}
func getPackages() map[string]*packageInfo {
	out := map[string]*packageInfo{}
	packages.RLock()
	for k, p := range packages.m {
		out[k] = p
	}
	packages.RUnlock()
	return out
}
func RegisterPackage(name string, getter func() map[string]*Type) {
	packages.Lock()
	if packages.m == nil {
		packages.m = make(map[string]*packageInfo)
	}
	packages.m[name] = &packageInfo{getSystemTypes: getter}
	packages.Unlock()
}
func registerTypes(arr map[string]*Type) {
	types.Lock()
	if types.m == nil {
		types.m = make(map[string]*Type)
	}
	for name, typ := range arr {
		types.m[name] = typ
	}
	types.Unlock()
}
func RegisterType(name string, typ *Type) {
	types.Lock()
	if types.m == nil {
		types.m = make(map[string]*Type)
	}
	types.m[name] = typ
	types.Unlock()
}
func UnregisterType(name string) {
	types.Lock()
	if types.m == nil || name == "" {
		// Added the name == "" condition to make it
		// possible for a test to get in here
		return
	}
	delete(types.m, name)
	types.Unlock()
}
func GetType(name string) (*Type, bool) {
	initTypes()
	types.RLock()
	t, ok := types.m[name]
	types.RUnlock()
	if !ok {
		return nil, false
	}
	return t, true
}
func GetAllTypesInPackage(path string) map[string]*Type {
	initTypes()
	out := map[string]*Type{}
	types.RLock()
	for k, t := range types.m {
		if strings.HasPrefix(k, fmt.Sprintf("%s:", path)) {
			out[k] = t
		}
	}
	types.RUnlock()
	return out
}

func (t *Type) HasExtends() bool {
	// Only the actual system:object type should have an empty string here.
	return t.Extends.Value != ""
}

type nativeTypeClasses int

const (
	nativeValue nativeTypeClasses = iota
	nativeCollection
	nativeType
)

func nativeTypeClass(nativeTypeString string) nativeTypeClasses {
	switch nativeTypeString {
	case "number", "string", "bool":
		return nativeValue
	case "map", "array":
		return nativeCollection
	default:
		return nativeType
	}
}

func nativeGoType(jsonNativeType string) (string, error) {
	switch jsonNativeType {
	case "number":
		return "float64", nil
	case "string":
		return "string", nil
	case "bool":
		return "bool", nil
	default:
		return "", uerr.New("TXQIDRBJRH", nil, "nativeGoType", "Native type not found: %v", jsonNativeType)
	}
}

func (t *Type) IsNativeValue() bool {
	return nativeTypeClass(t.Native.Value) == nativeValue
}
func (t *Type) IsNativeCollection() bool {
	return nativeTypeClass(t.Native.Value) == nativeCollection
}
func (t *Type) IsNativeType() bool {
	return nativeTypeClass(t.Native.Value) == nativeType
}
func (t *Type) NativeValueGolangType() (string, error) {
	return nativeGoType(t.Native.Value)
}

func (t *Type) GoName() string {
	return IdToGoName(t.Id)
}

func (t *Type) FullName() string {
	return fmt.Sprintf("%s:%s", t.Context.Package, t.Id)
}

// GoTypeReference outputs a Go source code reference to the name of this type. If we're in
// the local package, it just outputs the name e.g. "Object". If we're in a different package,
// it looks up the alias of the package in the imports and appends that to the start.
// e.g. "system.Object".
func (t *Type) GoTypeReference(path string, imports map[string]string) (string, error) {
	return IdToGoReference(t.Context.Package, t.Id, path, imports)
}

func (t *Type) GoSyntax(path string, imports map[string]string) string {
	return kegofmt.GoSyntax(path, imports, *t)
}

func (t *Type) Defaulter() (name string, property *Property, ok bool) {
	for name, prop := range t.Properties {
		if prop.Defaulter {
			return name, prop, true
		}
	}
	return "", nil, false
}
