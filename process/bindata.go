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

var _cmd_main_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x50\xcd\x4e\xf3\x30\x10\x3c\xdb\x4f\xb1\x5f\xf4\x1d\x6c\xc9\x4a\xc5\x15\xd4\x03\x07\x8a\x7a\xa0\xea\x1b\x54\x26\xdd\x96\x55\x13\x3b\xd8\x6e\x45\x55\xe5\xdd\x59\x07\xb7\x20\x7e\x0e\xe0\x43\xb2\x9e\x5d\xcf\xce\xcc\xe9\x04\xff\x7b\x9b\x9e\xe0\x7a\x0a\xf5\x32\x17\xc3\x20\x33\x48\x5d\xef\x43\x8a\x23\x3e\x2f\x35\xb7\x7a\xdb\xec\xec\x16\xa1\xb3\xe4\xa4\x7c\x1b\x02\x25\x45\xe5\x63\xc5\xdf\x4d\x97\xf2\x6f\x87\x5b\x5f\x93\x9f\xf4\xc1\x37\x18\xb9\x23\xc5\x0a\x2e\x68\x3c\xc6\x84\x5d\x25\x81\x0f\xaf\xa2\x0d\x38\xcf\x24\xf8\x5c\xa4\x7c\x1e\xd4\x79\xb1\xf8\xca\x30\x49\xc7\x1e\xf3\x5a\x26\x41\xb7\xce\x53\x63\x1d\xac\x63\x85\xc5\xc1\xc2\x76\x68\xce\x97\x65\x71\x7a\x71\x77\x26\x7e\x77\x5c\x32\xa8\x7e\xc0\xbf\xd9\xa9\xa5\xdc\xec\x5d\x33\x46\xa2\x34\x9c\xa4\x58\x53\x30\x90\x30\x26\x03\x01\x9b\x7d\x88\x74\x60\x0d\x07\x0c\x8f\x3e\x72\x91\x4d\x1a\x28\x12\x0c\x60\x08\x59\x53\xc9\xaa\x9e\x3b\x4a\x64\x5b\x8a\xa8\xb4\x14\x1c\x4e\xee\xff\x9b\x82\xa3\x36\x73\x0b\x8e\xb8\x5e\x06\x72\xa9\x75\x8a\x5b\x7a\x8c\x31\x1f\x1f\xeb\xbb\x17\x4a\xea\x8a\x9f\x0d\x97\x97\x1f\x98\xef\xd1\x61\xb0\x09\x67\xd4\x62\x54\x67\x74\xb6\x7a\xb8\x9d\x2f\x0c\xfc\x46\xb4\xbe\xf9\xa3\xaa\xe1\x35\x00\x00\xff\xff\xa7\x64\x28\xf9\x70\x02\x00\x00")

func cmd_main_tmpl_bytes() ([]byte, error) {
	return bindata_read(
		_cmd_main_tmpl,
		"cmd_main.tmpl",
	)
}

func cmd_main_tmpl() (*asset, error) {
	bytes, err := cmd_main_tmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "cmd_main.tmpl", size: 624, mode: os.FileMode(420), modTime: time.Unix(1434920259, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cmd_types_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x91\x51\x4b\xc3\x30\x10\xc7\x9f\x93\x4f\x71\x0e\x1f\x5a\x28\x1d\xbe\x2a\x7b\x50\xd8\xc6\x60\x68\x41\x5f\x7c\x1a\xb1\xbb\xcd\x63\x6d\x52\x93\x6c\x38\x46\xbf\xbb\x97\xda\x95\xea\xf6\x32\xd0\x3c\xb4\x97\xbb\xcb\xff\x7e\xf9\xe7\x70\x80\xeb\x4a\xf9\x77\xb8\x1d\x41\x9a\x85\xa0\xae\x65\x48\x52\x59\x19\xeb\x5d\x93\x9f\xb5\x31\x97\x2a\x95\x6f\xd4\x1a\xa1\x54\xa4\xa5\xfc\x6e\x82\x48\x8a\x81\x71\x03\xfe\xae\x4a\x1f\x7e\x1b\x5c\x9b\x94\xcc\xb0\xb2\x26\x47\xc7\x15\x29\x16\xd0\x65\xdd\xde\x79\x2c\x07\x12\x78\xf1\x28\x5a\x81\x36\x2c\x82\x1f\x2d\xca\xef\xc6\x38\x0c\x16\xa7\x0a\x43\xbf\xaf\x30\x8c\x65\x11\xd4\xcb\xd0\xd5\xc4\x17\x08\x76\xd7\xaf\xeb\x53\x1d\xab\x34\xdf\xb4\x75\xe2\x51\x95\x98\x1c\x37\x59\xeb\x58\xe7\x52\x5f\xaf\xd7\xd2\xa8\x9e\xcd\x9f\x61\x8f\xa5\x5c\x6d\x75\xde\x58\x1b\xc5\x70\x90\x62\x49\x36\x01\x8f\xce\x27\x60\x31\xdf\x5a\x47\x3b\x66\xd8\xa1\x7d\x33\x8e\x83\x00\x9e\x40\x8b\x90\x00\x5a\x1b\x98\x5a\xcf\xd3\x99\x26\x4f\xaa\x20\x87\x51\x2c\x05\x7b\x12\xea\x57\x23\xd0\x54\x04\x6d\xc1\x4f\x95\x66\x96\xb4\x2f\x74\xc4\xa5\xb8\x79\x8e\xb0\x8c\x4b\xc7\x9f\xe4\xa3\x1b\x3e\x56\x77\x27\x7b\xca\x53\xd4\x68\x95\xc7\x09\x15\xe8\xa2\x63\x76\xb2\x78\x79\xcd\xc6\xcf\x09\x5c\x42\x1d\xdf\xfd\x3f\xd6\x74\xfe\xf4\x70\x3f\xff\x7b\x30\x21\x7e\x22\xd5\x5f\x01\x00\x00\xff\xff\xa0\x1b\x81\x1b\x4c\x03\x00\x00")

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

	info := bindata_file_info{name: "cmd_types.tmpl", size: 844, mode: os.FileMode(420), modTime: time.Unix(1434920259, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cmd_validate_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x51\xb1\x6e\xc2\x30\x10\x9d\xed\xaf\xb8\x46\x1d\x6c\x29\x0a\xea\xda\x8a\xb1\x03\x4b\xc5\xd4\x15\xb9\xe1\xa0\x27\x88\x9d\xda\x06\x15\x21\xff\x7b\xcf\xa9\x89\x48\x8b\x04\x03\xb9\xbc\xbb\x7b\xef\xdd\xcb\xf9\x0c\x8f\xbd\x89\x9f\xf0\x3c\x87\x66\x99\x8b\x94\x64\x06\xa9\xeb\x9d\x8f\x61\xc0\x17\xa5\xe6\x56\x6f\xda\x9d\xd9\x22\x74\x86\xac\x94\xbf\x43\xa0\xa4\xa8\x5c\xa8\xf8\x7f\xd3\xc5\xfc\xd8\xe1\xd6\x35\xe4\x66\xbd\x77\x2d\x06\xee\x48\xb1\x82\x11\x0d\xa7\x10\xb1\xab\x6e\x60\xb3\x78\xea\x71\x18\x67\x0b\xb4\x01\xeb\x98\x1c\xbf\x8a\xc5\xbf\x04\x3a\x1b\x12\x99\x65\xbc\x22\xa5\xea\x1f\x72\x21\xcd\x9c\x68\xd7\x79\x69\xa8\xbd\xb1\x7c\x48\x39\xf4\xcd\x74\x58\x5f\x5e\x96\x25\x90\x31\x84\x6b\x9d\xab\x91\x89\xda\x04\xbf\xa1\xa9\xa5\xdc\x1c\x6c\x3b\x24\xa7\x34\x9c\xa5\x58\x93\xaf\x61\x55\x83\xc7\xf6\xe0\x03\x1d\xd9\xc0\x11\xfd\x87\x0b\x5c\x64\xef\x35\x14\xfd\x1a\xd0\xfb\x6c\xa8\xe4\xd9\x2c\x2c\x45\x32\x7b\x0a\xa8\xb4\x14\x1c\x54\xee\x3f\xcc\xc1\xd2\x3e\x13\x0b\xfe\x0c\xcd\xd2\x93\x8d\x7b\xab\xb8\xa5\x25\x94\x9f\x0b\xcd\xeb\x37\x45\xf5\xc4\x6b\x69\xdc\xbc\x62\x7e\x67\xd6\xb5\x89\xa8\x06\x73\x77\x9d\xe9\x97\xbb\xd2\x42\x4c\x45\xd3\x4f\x00\x00\x00\xff\xff\x40\x3d\x2a\x52\x73\x02\x00\x00")

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

	info := bindata_file_info{name: "cmd_validate.tmpl", size: 627, mode: os.FileMode(420), modTime: time.Unix(1434920259, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _global_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x51\xcd\x4a\x03\x31\x10\x3e\x6f\x9e\x62\x28\x1e\x2c\xc8\x7a\x17\x7a\x16\x2f\x4b\xc1\x07\x90\x58\xa7\xdb\xd0\x6e\x26\x26\x51\x90\xb6\xef\x6e\x26\x93\x36\x1b\x0f\xd2\xdb\x64\x66\xbe\xbf\xc9\xf1\x08\x77\x4e\xc7\x1d\x3c\xad\xa0\x5f\x73\x71\x3e\x2b\x6e\x9a\xc9\x91\x8f\x21\xf7\x5f\x4a\x9d\x46\x4e\x6f\xf6\x7a\x44\xb8\xe2\x4e\x60\xf5\x84\x3c\x52\x02\x81\x7b\xd5\xa5\xa9\xd9\x82\xa5\xf4\xc0\xcf\xb2\xb8\xd8\xe3\x48\xbd\xa1\xc7\xf0\x13\x22\x4e\x8b\x25\x63\xba\xee\x6f\x3b\x83\xd1\x7e\x64\x46\xae\xbd\xb6\x49\xaf\xf8\x19\x92\xd6\xc3\xe5\xb1\x2e\xbe\xaf\x5e\x33\x21\x6b\x8b\x91\x19\xa6\x81\x88\x1f\x5e\xae\x52\x4b\xa5\xaa\x96\x23\x63\x23\x7a\xb9\x89\xd4\x42\xfe\xad\xbd\x24\x97\x66\x3f\x48\x74\x58\x35\xdd\x57\xfa\xf2\x1b\x2c\x87\xbc\x44\xa9\xec\x1e\xb7\x29\x42\x90\x25\x96\x78\x3e\xd0\xbb\x3e\xb4\x0a\x69\x49\xd8\x4f\x30\x92\x6d\x64\xc2\xff\xf4\xb7\x1e\x8a\x95\xde\xa0\x7e\x75\xc9\xd2\x0f\x14\x77\xc6\x8e\x73\xf2\xdf\x00\x00\x00\xff\xff\x16\x9e\x56\x8b\x25\x02\x00\x00")

func global_tmpl_bytes() ([]byte, error) {
	return bindata_read(
		_global_tmpl,
		"global.tmpl",
	)
}

func global_tmpl() (*asset, error) {
	bytes, err := global_tmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "global.tmpl", size: 549, mode: os.FileMode(420), modTime: time.Unix(1434920259, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _main_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x54\x4d\x6b\xde\x30\x0c\x3e\xc7\xbf\x42\x84\x31\x92\x51\xb2\xfb\x60\x97\xc1\x36\x7a\x79\x57\xca\xd8\xb5\xb8\x8e\x92\x7a\x4d\xec\xd4\x76\xca\x4a\x92\xff\x3e\x7f\xe5\xc3\x7b\x5f\x68\x77\x53\x24\x3d\xd2\xa3\x47\x56\xa6\x09\xde\x0d\xd4\x3c\xc0\xa7\xcf\x50\xdd\x38\x63\x59\x88\x73\xf2\x7e\x90\xca\x68\xef\xbf\x8e\xb6\x0d\x0d\x94\x3d\xd2\x16\x61\xc3\xcd\x20\x68\x8f\x2e\x44\x02\x04\x0a\x92\xe5\x0a\x9b\x0e\x99\xc9\x89\xb5\x1f\xb1\x95\x15\x97\x1f\x7f\x6b\x29\x9c\xc3\x42\x79\x03\x42\xda\x4c\x7c\x8a\x55\xb6\x24\xfd\xa2\x0d\xf6\x79\xe9\x0a\x66\xd9\xbf\x6e\x0f\x46\x51\xfb\x76\xce\x56\x54\x58\x32\x91\xec\xc9\x12\xb9\x5a\x3f\x6e\xe2\x50\xdb\x20\xbe\xa0\xeb\x1d\x58\x1e\x30\x09\x24\xf0\x71\xc9\x7b\xab\x92\x90\xbd\x97\x1d\xcd\x36\x31\x2f\x03\x7a\x6d\x7e\x5a\x43\xaf\xf9\x76\x2e\x6a\x21\x85\x1b\xce\xa7\x54\xd7\xc2\xa0\x6a\x28\xc3\x32\xf1\xea\x13\x35\xfc\x19\x7f\xd1\x6e\x4c\x23\x5f\xff\xb0\x6e\xac\xb1\xdc\xe8\x86\x4e\x33\xd4\xa8\x99\xe2\x83\xe1\x52\x84\x98\xf7\xaf\x09\xd5\x77\x79\x0a\x6b\x00\x6d\xd4\xc8\x0c\x4c\x36\xe5\x20\x75\xc8\xfa\x42\x35\x67\x01\x9e\x65\x1f\xdc\x4c\xd8\xa0\x42\xc1\xf0\x8e\xd7\x67\x4b\x80\xfc\x9e\x6a\xcc\xa3\x24\xa9\x90\x07\x75\xc2\x47\x54\x07\xfb\x7b\xac\xbd\xee\x61\x1c\xff\x79\x68\x18\x12\x2c\xdd\xdb\xb5\xf3\x6b\xe5\xd3\xfa\x0d\xc7\xae\x8e\x8b\xf6\xf6\xde\xeb\x9b\xfb\x5c\xf1\x5e\xba\x90\x70\x41\xbb\x43\xd8\xcb\x36\x43\x2b\xe3\x33\x76\x92\xb6\xd2\x8b\x1b\xf1\xaf\x8e\x9f\xbc\x95\x03\xf1\x66\x14\x0c\xb8\xe0\xa6\x28\xdd\x3a\xde\xf0\x84\xce\x16\xb6\x3d\x9f\x83\xd0\x9b\x70\xe9\x65\x41\x7e\x8b\x2d\xb7\x8b\x53\xae\xde\x85\xad\x15\x6e\x66\x8b\xb6\x67\x1e\x6e\x78\x86\xa7\x51\x1a\x57\xfa\x0a\xd6\x58\x94\x63\x0f\xc4\x4b\xf6\x24\x7f\x34\xc5\xfb\xf3\x17\x37\x2d\x65\x49\x12\x45\x76\x73\xf1\x97\xf3\xc6\x73\x7f\xa6\x0a\xee\x20\xb8\xaa\x93\x34\x0f\x5c\xb4\x24\x51\xf6\x3f\xcf\x3d\x14\xdc\xff\x66\x91\xf0\x85\xda\xc4\x2f\x23\xfa\xe3\x09\x4d\xcb\xdf\x00\x00\x00\xff\xff\x0d\xf0\x23\xd3\x1d\x05\x00\x00")

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

	info := bindata_file_info{name: "main.tmpl", size: 1309, mode: os.FileMode(420), modTime: time.Unix(1434920438, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _types_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x52\xcb\x6a\xc3\x30\x10\x3c\x47\x5f\xb1\x98\x1e\x62\x08\xce\xbd\xd0\x0f\xe8\xc5\x84\xb6\xf7\x20\xdc\xb5\x23\x82\x25\x47\x92\x0b\xc1\xc9\xbf\x77\xf5\xb2\x55\x97\x40\x7b\x5b\xed\xec\x8c\x66\x1f\xd3\x04\x4f\x03\xb7\x27\x78\x7e\x81\xea\xe0\x82\xfb\x9d\xb9\xa4\xe8\x07\xa5\xad\xf1\xf9\xd7\x18\x47\x48\x2a\xf9\x71\x1d\xd0\x1c\x12\xaf\xce\x13\x54\x34\xf0\xe6\xcc\x3b\x84\x59\xfc\x06\x92\xf7\xe8\x20\x16\x74\x61\xcb\x36\x84\x8a\x16\xa4\xa2\x07\x5e\x62\x61\x71\xc6\x4e\x55\x42\xed\xcd\xd5\x58\xec\x8b\xd2\x71\x36\x9b\x75\xda\x93\x51\x7e\x7a\xc5\xbf\x08\xed\xad\xf3\x97\xe4\x8e\x0f\xe0\x5f\xb2\x9a\x4b\x6a\x23\xce\xa2\xa6\x16\x76\xe9\x91\x7a\x9f\xe7\xe4\x85\x9d\x93\xd0\x5f\xc6\xf9\x41\x09\xee\x52\xb1\xff\xf6\x01\xbe\x58\x29\x19\x6b\x47\xd9\x80\x90\xc2\x6e\x4b\x98\x72\x6b\x83\x12\xd2\xa2\x0e\xeb\x0b\xf1\xe2\x25\xa1\x55\x1d\xa6\xef\xaa\xf2\xf4\xbb\x1a\x75\x83\xab\xdf\x72\x75\x8d\x2d\x75\x6c\x42\x99\xfb\xc2\xaf\x79\xd6\x27\x18\x35\x4a\xc2\xd6\x0b\x82\xe2\x0d\x3b\x41\xa1\x76\x8c\x22\xb6\x95\x0d\x6b\xeb\x7c\x10\x9f\x6e\x2e\xdc\xca\x0d\x2e\xa3\xb2\xce\xcc\x0e\x12\xe6\x6d\xaf\x01\x93\x4c\x97\xb9\x6b\xf2\xfd\xff\x7d\x7d\x71\x0d\x47\x58\xae\x3d\x8e\x89\xee\xd9\x9e\x84\xec\xd8\x2c\xff\x1d\x00\x00\xff\xff\x1c\x4d\x87\x22\x26\x03\x00\x00")

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

	info := bindata_file_info{name: "types.tmpl", size: 806, mode: os.FileMode(420), modTime: time.Unix(1434920259, 0)}
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
	"cmd_main.tmpl":     cmd_main_tmpl,
	"cmd_types.tmpl":    cmd_types_tmpl,
	"cmd_validate.tmpl": cmd_validate_tmpl,
	"global.tmpl":       global_tmpl,
	"main.tmpl":         main_tmpl,
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
	"cmd_main.tmpl":     &_bintree_t{cmd_main_tmpl, map[string]*_bintree_t{}},
	"cmd_types.tmpl":    &_bintree_t{cmd_types_tmpl, map[string]*_bintree_t{}},
	"cmd_validate.tmpl": &_bintree_t{cmd_validate_tmpl, map[string]*_bintree_t{}},
	"global.tmpl":       &_bintree_t{global_tmpl, map[string]*_bintree_t{}},
	"main.tmpl":         &_bintree_t{main_tmpl, map[string]*_bintree_t{}},
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
