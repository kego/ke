package shared // import "kego.io/editor/shared"

// ke: {"package": {"jstest": true}}

type Info struct {
	// Package path
	Path string
	// Map of path:alias
	Aliases map[string]string
	// Map of data names and relative file names
	Data map[string]string
	// Package object
	Package []byte
	// Flattened list of all imports
	Imports map[string]*ImportInfo
}

type ImportInfo struct {
	Path    string
	Aliases map[string]string
	Types   map[string][]byte
}

type DataRequest struct {
	File    string
	Name    string
	Package string
}
type DataResponse struct {
	Data    []byte
	Found   bool
	Name    string
	Package string
}

var MESSAGE_TYPE = M_BINARY

type MessageType string

const (
	M_STRING MessageType = "string"
	M_BINARY MessageType = "binary"
)
