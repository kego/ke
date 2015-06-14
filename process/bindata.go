package process

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _cmd_types_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x51\x31\x4f\xf3\x30\x10\x9d\x7d\xbf\xe2\xbe\xe8\x1b\x62\x09\xa5\x62\x05\x75\x6c\x11\x0b\x8a\x04\x0b\x53\x65\x85\x4b\xb1\x9a\xd8\xc1\x36\x12\x55\xe4\xff\xce\xd9\xb8\x51\x45\xbb\xe0\x21\xb9\xbc\x7b\x79\xef\xee\x79\x9e\xf1\xff\xa4\xc2\x3b\xde\xad\xb1\x69\x53\x11\x23\x24\x50\x8f\x93\x75\xc1\x67\xfc\xb1\xd4\xdc\x9a\x54\x77\x50\x7b\xc2\x51\x69\x03\xf0\x43\xc2\x1a\x44\x65\x7d\xc5\xcf\x7e\x0c\xe9\x75\xa0\xbd\x6d\xb4\x5d\x4d\xce\x76\xe4\xb9\x03\x62\x87\x0b\xea\x8f\x3e\xd0\x58\x01\xf2\x61\x2b\xdd\xa3\xb1\x2c\x42\x1f\x65\x94\xdf\x44\x99\x8c\xc5\xa5\xc2\x2a\x1c\x27\x4a\xb6\x2c\x42\xe6\x2d\xb1\x72\xfd\x07\xc1\x65\xfd\x18\x2f\x75\x9c\x32\xbc\x69\x49\xe2\x49\x8d\x74\x73\xfa\x68\x4b\x62\x4b\x4a\xe7\x7a\x67\x94\xac\x7a\x15\xbf\x32\xbb\x04\xe8\x3f\x4d\x97\xa3\xad\x25\xce\x20\x78\x11\x72\x2e\x19\x95\x20\x9b\x07\x32\xe4\x54\xa0\xad\x1e\xc8\xd7\x27\x74\xbb\x7b\x79\x6d\x37\xcf\xf2\x3e\xd3\xff\xad\xd1\xe8\x21\xfd\x2f\xf8\x3a\x9a\xd6\x69\x13\x06\x53\x73\x4b\xe6\xc8\xd3\xb1\xbe\xd9\x7c\xe9\x50\xdf\x4a\x10\x11\xe2\x77\x00\x00\x00\xff\xff\xbe\x35\x6b\xfb\x06\x02\x00\x00")

func cmd_types_tmpl_bytes() ([]byte, error) {
	return bindata_read(
		_cmd_types_tmpl,
		"cmd_types.tmpl",
	)
}

func cmd_types_tmpl() (*asset, error) {
	bytes, err := cmd_types_tmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "cmd_types.tmpl", size: 518, mode: os.FileMode(420), modTime: time.Unix(1434278840, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cmd_validate_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x50\x3d\x4f\xc3\x30\x10\x9d\xe3\x5f\x61\x22\x06\x5b\x42\xae\x58\x41\x1d\x41\x62\x41\x99\x58\x91\x95\x5e\x8a\xd5\xf8\x03\xdb\x48\x54\x91\xff\x3b\x77\xc6\x8d\x5a\xe8\x92\x9c\xdf\xdd\xbd\x77\xef\x2d\x0b\xbf\x0d\x3a\x7f\xf0\x87\x2d\x57\x03\x15\xa5\x30\x02\x8d\x0d\x3e\xe6\x54\xf1\x97\x56\x63\x2b\xe8\xf1\xa0\xf7\xc0\xad\x36\x8e\xb1\xdf\x21\x2e\x58\xd7\xfb\xd4\xe3\x77\xb2\x99\x7e\x07\xd8\x7b\x65\xfc\x26\x44\x3f\x42\xc2\x0e\xeb\xde\xf9\x8a\xa6\x63\xca\x60\xfb\x2b\xd8\x26\x1f\x03\xd4\x71\x3c\xc1\x4c\xdc\x79\x24\x87\xcf\x76\xe2\x5f\x02\x49\x07\x75\xc4\xb2\xba\x28\xa5\xff\x87\x9c\x48\x89\x13\xdc\x8e\x96\x6a\x1d\xb5\x43\x23\xcd\xe8\xab\xb6\x70\x77\x7a\x0c\x2d\x90\x35\x84\x73\x9d\xb3\x91\x0b\xb5\x0b\xfc\x8a\xa6\x64\x6c\xfa\x72\x63\x4d\x4e\x48\xbe\xb0\x0e\x0d\x42\x8c\x24\xd4\x72\x52\x6f\x7a\x36\x3b\x9d\xe1\xd9\xcc\x90\x84\x7c\xac\xfd\x9b\x2d\x77\x66\xa6\x85\x0e\xe3\x55\x43\x34\x2e\xcf\x4e\x60\x4b\x22\xe4\x93\x7a\xfa\x36\x59\xdc\xe3\xa3\xb0\xf2\x13\x00\x00\xff\xff\xad\xab\x81\xf2\xd0\x01\x00\x00")

func cmd_validate_tmpl_bytes() ([]byte, error) {
	return bindata_read(
		_cmd_validate_tmpl,
		"cmd_validate.tmpl",
	)
}

func cmd_validate_tmpl() (*asset, error) {
	bytes, err := cmd_validate_tmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "cmd_validate.tmpl", size: 464, mode: os.FileMode(420), modTime: time.Unix(1434278836, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _main_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x53\xcf\x6f\x9d\x30\x0c\x3e\xc3\x5f\x61\x59\x4f\x1b\x3c\x55\x70\x9f\xb4\xf3\xd4\x4b\x57\x4d\x95\x76\xce\xc0\xa1\xac\x90\xd0\x10\x26\x55\x94\xff\x7d\x71\x08\x01\x76\x58\xa5\xc7\xc9\xf8\xc7\xf7\xd9\x9f\x9d\x79\x86\xcb\x20\xec\x33\x7c\xf9\x0a\xc5\x23\x1b\xcb\x92\xb2\xb3\xed\x07\x6d\xec\xe8\xfd\xf7\xc1\x76\xa1\x41\x54\x2f\xa2\x21\x88\x75\xef\xa0\x44\x4f\x1c\x4a\xd7\x12\xc8\xd2\x04\x0d\xc9\x8e\x2a\x8b\xa9\xb3\x5f\xa8\xd1\x45\xab\xcb\xdf\xa3\x56\xec\x70\xa5\xad\x04\xa5\x5d\x26\xbd\x06\x94\x98\x34\xbe\x8d\x96\x7a\xcc\x19\x30\x49\xfe\x75\xfb\x62\x52\xb5\xa7\x63\xdb\x08\xe5\x9a\x09\xcd\x3e\xb8\x46\xee\xb6\x9f\xc7\x30\x54\x1c\xc4\x03\x32\xf7\xda\xe5\xa1\xe6\x54\xb2\xf6\xc3\xc9\x3b\x55\x9e\xa6\x3b\x97\x7d\x1b\x28\x30\xb1\xe9\x05\x7a\x72\xc6\x18\x9b\xba\x78\x45\x98\x9b\x13\x8a\xfb\x7a\x83\xbb\xc8\xa9\xeb\x62\xe0\x21\xe8\xe6\x23\x8d\xde\x0b\xbe\xe9\x2d\xb4\xc6\xfe\x08\xe3\xf7\xd0\x8b\x01\x90\x33\x30\x50\x23\x13\x61\xe0\x43\x06\xc7\xc0\x81\x8d\x46\x0f\x8a\x3c\x0e\x6e\x2a\x07\x2d\xf0\xa4\x4a\x9a\x94\xe5\xf5\xf6\x2f\x94\x43\x1c\x7c\x59\x60\x77\xdf\x8c\x7a\xbc\x93\xa0\xa3\xb2\x64\xa4\xa8\x68\x5b\x65\x79\x85\x9f\x04\xb5\x56\x9f\x2d\x28\xa2\x1a\xec\x33\x8d\x04\xbf\x5a\x37\x97\xd4\x06\xda\x58\x60\xfd\x7e\xae\xe5\x76\x03\xee\x96\x86\x4e\x58\x27\xda\x68\xcd\x54\xd9\xc2\xba\x7f\x0c\x4a\x9f\x77\x7f\x38\x38\x39\xa9\xca\x81\xb6\x36\xcb\x61\x3e\x9e\xdf\x7f\x4f\x22\xf9\x68\x10\x0f\x44\x92\x0c\x29\xe7\x3a\x3f\x17\xc0\x1f\xd4\xb4\xee\xf6\xcd\xd3\xba\x77\xbf\xc7\xc3\xf6\x32\x96\x3d\x5e\xd3\x3b\xbc\x4e\xda\x32\xee\x1d\x84\x27\xe8\x1b\xf9\x2e\xb3\x4f\x5b\xe2\x7e\x5e\xf3\x92\xe7\x6b\x7b\x61\xc2\x83\xb9\xfc\x0d\x00\x00\xff\xff\x30\x4f\xdd\x51\x19\x04\x00\x00")

func main_tmpl_bytes() ([]byte, error) {
	return bindata_read(
		_main_tmpl,
		"main.tmpl",
	)
}

func main_tmpl() (*asset, error) {
	bytes, err := main_tmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "main.tmpl", size: 1049, mode: os.FileMode(420), modTime: time.Unix(1432913174, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _struct_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x91\x4f\x4b\xc4\x30\x10\xc5\xcf\xed\xa7\x18\x4a\x0f\x2a\x52\xef\x82\x17\x41\xc5\xcb\x1e\x3c\x78\x95\x6c\x32\x8d\xc1\x36\x53\x92\xa9\xb8\xd4\xfd\xee\xe6\xdf\x6e\x0b\xc2\xde\xde\xbc\x1f\xf3\x26\x33\x59\x16\x68\x27\xc1\x9f\x70\xff\x00\x5d\x12\xc7\x63\x1d\x4d\x33\x4e\xe4\xd8\x27\xff\xa4\x33\x32\x3d\x08\xab\xe0\xca\x12\x43\xc7\x87\x09\xbb\x57\xbf\x13\x6c\xbe\xf1\x5d\x0c\x33\x5e\x6f\xc9\xd3\x8f\x1c\x66\x15\xbc\xdc\x9a\x4c\xf8\x05\x85\x5e\x3a\x33\xb1\x21\x1b\x49\x72\x23\xd6\x14\x4a\xf0\xec\x66\xc9\xb0\xd4\x75\x95\xc7\xad\x79\x8f\xc2\x1b\x19\x5b\xaa\xea\x26\x30\x87\x3d\x3a\xb4\x12\x3f\x8c\x82\xe6\x0b\x35\x75\x86\xee\xfc\xc1\x33\x8e\x0d\x34\x7b\xe1\xb1\x29\xfb\xb5\x9b\x25\x62\x2c\x86\x15\x8a\x74\xc2\x6a\x84\x16\xc7\x3d\xaa\xb4\x6f\x7e\x7a\x2a\xcf\xa3\x32\xee\x5e\xe8\xed\x34\xf3\x72\xf0\x36\xb9\x37\x38\xa8\x9d\x18\xf1\xb6\xe8\x75\xca\x73\x2c\xcf\xbd\x85\xfe\x3b\xd0\xca\x62\x4a\xe0\x9a\x6c\x14\xe1\x58\x81\x68\x4a\x07\x2c\xcd\x17\x5f\x95\x7e\x21\xeb\xbf\x00\x00\x00\xff\xff\x63\xbc\x02\x72\xfa\x01\x00\x00")

func struct_tmpl_bytes() ([]byte, error) {
	return bindata_read(
		_struct_tmpl,
		"struct.tmpl",
	)
}

func struct_tmpl() (*asset, error) {
	bytes, err := struct_tmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "struct.tmpl", size: 506, mode: os.FileMode(420), modTime: time.Unix(1434298692, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _types_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x52\xcb\x6a\xc3\x30\x10\x3c\x5b\x5f\xb1\x98\x1e\x6c\x08\xc9\xbd\xd0\x0f\xe8\xa5\x84\xd2\x7b\x11\xee\xda\x15\xc1\x92\x22\xc9\x87\xe0\xea\xdf\xab\xd5\xc3\x51\x52\x02\xbd\xad\x34\x9a\xd9\x99\xd5\xae\x2b\x3c\x69\xee\xbe\xe1\xf9\x05\xf6\x47\x2a\xbc\x67\x74\x29\x66\xad\x8c\xb3\xf1\xfe\x35\xd7\x01\xd2\x7c\x38\xf1\x09\x61\xe3\xfd\x80\xe4\x33\x12\xc4\x12\x05\x3a\xd6\x04\x54\x8c\x20\x55\x38\xe0\x39\x3f\x6c\x4f\x38\xa9\xbd\x50\x07\x7b\xb1\x0e\xe7\xb6\x27\x4e\xd3\xdc\x5f\x47\x32\xca\xaf\xa8\xf8\x1f\xa1\x83\xbb\x68\xb4\x45\xee\xf3\x01\xfc\x47\xd6\x70\x19\x62\xe4\x98\x6f\x21\xc2\xae\x1c\x8e\x79\x1c\xdb\x08\xa2\x30\x39\x49\xf9\x2a\xce\x0d\x25\xb9\x2b\x8f\x63\xdb\x07\xf8\xd5\x4a\xcf\xd8\xb8\xc8\x01\x84\x14\xae\xeb\x61\xad\xad\x69\x25\xa4\x43\x93\x7e\x26\xd5\x57\x2f\x1b\x7a\xab\x57\xf3\xc9\x41\x0e\x66\xd5\x62\x06\x8c\x4a\x1f\xd1\x57\x91\x31\x38\xa2\x41\x19\xb0\xfb\x7f\x80\xf6\x1d\x27\x11\x4a\x43\x8c\x36\xbb\xaf\x66\xd2\x91\x8b\xd2\x23\xac\xc1\x79\x51\x8e\xf6\x60\x17\x77\x23\x77\xf4\xbe\xaf\xed\xf9\xdf\x00\x00\x00\xff\xff\x6c\x61\x70\x40\x6f\x02\x00\x00")

func types_tmpl_bytes() ([]byte, error) {
	return bindata_read(
		_types_tmpl,
		"types.tmpl",
	)
}

func types_tmpl() (*asset, error) {
	bytes, err := types_tmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "types.tmpl", size: 623, mode: os.FileMode(420), modTime: time.Unix(1434192299, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"cmd_types.tmpl":    cmd_types_tmpl,
	"cmd_validate.tmpl": cmd_validate_tmpl,
	"main.tmpl":         main_tmpl,
	"struct.tmpl":       struct_tmpl,
	"types.tmpl":        types_tmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() (*asset, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"cmd_types.tmpl":    &_bintree_t{cmd_types_tmpl, map[string]*_bintree_t{}},
	"cmd_validate.tmpl": &_bintree_t{cmd_validate_tmpl, map[string]*_bintree_t{}},
	"main.tmpl":         &_bintree_t{main_tmpl, map[string]*_bintree_t{}},
	"struct.tmpl":       &_bintree_t{struct_tmpl, map[string]*_bintree_t{}},
	"types.tmpl":        &_bintree_t{types_tmpl, map[string]*_bintree_t{}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	if err != nil { // File
		return RestoreAsset(dir, name)
	} else { // Dir
		for _, child := range children {
			err = RestoreAssets(dir, path.Join(name, child))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
