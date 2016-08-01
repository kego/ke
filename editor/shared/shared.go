package shared // import "kego.io/editor/shared"

// ke: {"package": {"jstest": true}}

type Info struct {
	// Package path
	Path string
	// Map of alias:path
	Aliases map[string]string
	// Map of data names and relative file names
	Data map[string]string
	// Package object
	Package         []byte
	PackageFilename string
	// Flattened list of all imports
	Imports map[string]ImportInfo
}

type ImportInfo struct {
	Path    string
	Aliases map[string]string
	Types   map[string]TypeInfo
}

type TypeInfo struct {
	File  string
	Bytes []byte
}

type Method string

const Save Method = "Server.Save"

type SaveRequest struct {
	File  string
	Bytes []byte
}

type SaveResponse struct {
	Error error
}

const Data Method = "Server.Data"

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
