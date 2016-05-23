// Code generated by go-bindata.
// sources:
// templates/call.tmpl
// templates/function.tmpl
// templates/header.tmpl
// templates/inline.tmpl
// templates/inputs.tmpl
// templates/message.tmpl
// templates/results.tmpl
// DO NOT EDIT!

package bindata

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesCallTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x8e\x51\xca\xc2\x40\x0c\x84\xaf\x12\x4a\x1f\xfe\x1f\x4a\x0e\x20\x78\x00\x5f\x44\x54\xf4\x79\xd9\xa6\x1a\xa8\xab\xa4\x51\x91\x90\xbb\xdb\x2d\xd5\xf5\x29\x30\xf3\x4d\x66\xcc\x5a\xea\x38\x11\x54\x31\xf4\x7d\xe5\x6e\xf6\x64\x3d\x03\x6e\x29\x12\x3f\x48\xb2\xc2\x1d\xa4\xab\x02\xae\x86\x9d\xca\x3d\xaa\xbb\x2a\x9a\x51\x6a\xb3\xfb\x21\x01\xdd\x8b\x8a\xeb\x70\x21\xf7\x3f\x33\x09\xe9\x44\x50\x73\x03\x35\xf5\xb0\x58\x02\x6e\x82\x8c\xa6\x92\x0c\xf3\xf7\x9a\xdd\x1b\xf8\x66\x4b\xdf\x51\x58\xf3\x86\xdf\xbe\x29\x9d\xcb\x26\x10\xf7\xaf\x1b\x8d\xe4\x21\x08\x87\x96\xe3\xb8\x01\x0b\x3b\x9d\xff\xf9\xbe\x03\x00\x00\xff\xff\x9d\x9f\x57\x19\xec\x00\x00\x00")

func templatesCallTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesCallTmpl,
		"templates/call.tmpl",
	)
}

func templatesCallTmpl() (*asset, error) {
	bytes, err := templatesCallTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/call.tmpl", size: 236, mode: os.FileMode(420), modTime: time.Unix(1464034914, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesFunctionTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x55\x5d\x4f\xdb\x30\x14\x7d\x2e\xbf\xe2\xaa\x42\xa8\x99\x8a\x79\x07\xed\x61\x6c\x6c\xe2\xa5\x45\x14\x8d\x87\x69\x9a\x4c\x72\x03\xd6\x8c\xd3\xd9\x0e\xac\xb2\xf2\xdf\x77\xfd\x11\xa7\xd0\x94\x69\x2f\x7b\x40\x25\x37\xbe\xe7\x9c\x7b\x7c\xec\x38\x57\x61\x2d\x14\xc2\xb4\x6e\x55\x69\x45\xa3\xa6\x5d\x07\xce\x1d\xd6\x70\xfa\x1e\x58\xd7\x1d\x1c\xf8\x17\x54\x61\x37\x68\xec\x82\x3f\x62\xd7\xcd\x2c\xbc\xb3\xf4\x24\xd4\x3d\xbb\x29\xc0\x1d\x4c\x3e\x36\xea\x09\x37\xb3\xe9\x8b\x65\xd3\x39\xd8\x39\xf8\xf6\x59\x58\x34\xf1\x3d\xc6\xe3\x7e\xfb\x6e\xac\x6e\x4b\x0b\x8e\x80\x45\x0d\xaa\xb1\x70\x58\xb3\x2b\x2d\x94\xbd\x54\xeb\xd6\x1a\x22\x9e\x4c\x26\x27\x27\xe0\xe1\xa0\x42\x53\x6a\xb1\xf6\xf2\x7c\x59\x11\x3c\x10\x02\xf1\x53\x3f\xaa\x2a\x48\x7e\x16\xf6\x01\xd8\x35\x96\x28\x9e\x50\x87\x12\x41\xb3\x4b\xb3\x0a\x5c\xb9\xf0\x59\xa0\xac\x06\x82\xbe\x01\xea\x50\x67\x03\xa0\xe6\xea\x1e\x5f\xae\x77\x2e\x3c\x79\x63\x82\x25\x9b\x35\x86\xff\xfa\x16\x94\x06\x77\x90\x63\x63\xe6\xd9\xdb\x9b\x7e\xbd\x46\x3f\xf5\x15\xd7\x34\xa7\x45\x3d\x68\x1d\x4a\xb0\x23\x73\xac\xc5\xb9\x50\xd9\x4b\x49\x54\x8d\x8e\xad\xd7\x68\x5a\x49\xbb\x43\x06\xda\x56\x2b\x73\xa1\x75\xa3\x33\xf1\xc5\xef\x35\x96\x16\x2b\xd0\x69\xd9\x28\x7d\xc2\xe8\xb9\x6f\xb9\xb2\x6f\x51\xff\x85\x0a\x7d\x19\x8e\x8f\x01\x69\x63\xc9\xb7\xd5\x43\xd3\xca\xea\x1c\x17\x42\x7a\xd5\xf1\x71\xd1\xd8\x50\xf1\xcd\xcf\xc4\x47\x58\x31\x71\x94\x24\xd4\x35\x2f\xd1\x75\x73\x60\x8c\x6d\x3d\x17\x29\x3b\x49\x0c\xb5\x76\xae\x4f\xdb\xf2\xd3\xf2\x14\x3e\x54\x15\xf8\xac\x42\xc9\x0d\x1a\xe6\x17\xd0\x5f\x4d\xa4\x39\xad\x14\xab\x05\xff\x89\x7e\x96\x1f\x14\x73\xeb\x53\xdd\x0f\x17\x1d\x89\x61\x0f\xc0\xe9\x74\x58\xcb\x7c\x72\x87\x23\x31\x12\xda\xd7\x99\xf5\xed\xaf\xc3\x13\xa8\x42\x4a\xc8\xd4\xb0\x96\x53\xe7\x51\xa2\x4f\x66\xb3\xaf\x5c\xb6\x64\xb9\x1b\x0f\x32\x61\xb2\x78\x4a\x4f\x49\x3d\xdb\xca\xf5\x1c\x06\x5f\x68\xf0\x91\x80\x26\xb8\xed\xb0\xf5\xb2\x6f\xb5\xb0\xa8\x33\xc5\x90\x3e\xd2\x7c\x74\xb7\x21\x4f\xd8\x79\x5b\xd7\xa8\xdd\x9e\xe4\x73\x55\xc1\x2c\x18\xbc\x54\x72\xb3\x1d\x90\x62\xb7\xbe\x54\x18\xa6\x2c\x20\x33\x5a\x7c\x5c\x4b\x6e\xe9\x3e\x4b\x49\x9d\xd2\xc5\x12\xc0\x87\x37\x25\x97\x32\x97\xdf\xce\xe3\x24\xbe\x78\xad\xa5\xeb\x28\x9c\x71\x1f\xc6\x60\x13\xea\xaa\x99\xd1\x32\x9f\x0e\x96\xa2\x59\xe4\x3d\x1a\x3b\x35\xb0\xc7\xc5\x2f\x8d\x1d\xf6\x3d\x5b\xca\x56\x21\xc4\xb3\xe2\x0c\x88\x28\xaf\x9a\xa7\x83\x71\xf1\xab\xe5\x72\x1e\xf7\xb6\x3f\x88\x45\xba\xa1\x20\xd2\x9c\x73\x23\xca\x78\x30\x87\x59\xe9\x1a\x1e\x71\xd8\x4b\x7e\x21\x63\x18\x5b\x28\x49\xdf\x8f\x38\xf8\x59\x36\xf4\x5f\x25\xfd\x27\x09\xe4\x35\x3e\xde\x49\x1c\x53\x11\x37\x26\x47\xbf\x2b\xe2\xb9\xa7\x5f\xfa\x08\xf6\x2f\xfe\x04\x00\x00\xff\xff\x44\xd0\xf1\x6b\x30\x07\x00\x00")

func templatesFunctionTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesFunctionTmpl,
		"templates/function.tmpl",
	)
}

func templatesFunctionTmpl() (*asset, error) {
	bytes, err := templatesFunctionTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/function.tmpl", size: 1840, mode: os.FileMode(420), modTime: time.Unix(1464036476, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHeaderTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x34\x8b\x41\xaa\x83\x21\x0c\x06\xf7\x9e\x42\x5c\xbd\xb7\xd1\x73\x74\x53\x7a\x85\x54\x53\x95\xa2\xfe\x68\x5a\xf8\x09\xb9\x7b\x43\xa5\xab\x84\xf9\x66\x98\x13\x3e\x6a\x47\xeb\x0a\x42\xc2\xe9\x44\xcc\x01\xf1\x09\x19\x2d\xb3\xbf\xed\x57\xa1\xa9\xed\x18\x93\xec\x9f\x61\x9e\xd0\x75\xf6\x97\x2f\x59\x22\x2a\x5e\xa1\xa9\xb5\x13\x2a\xea\x33\x63\x4f\x22\xde\xba\x5c\xa9\xbc\xee\x3e\x8e\x16\x56\x83\x49\xe7\xa2\x89\x48\x2b\xe4\x11\x47\x7f\xe3\x19\xf6\x71\xe6\xff\x57\x99\x4f\x00\x00\x00\xff\xff\x72\x96\x3e\xa5\x97\x00\x00\x00")

func templatesHeaderTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesHeaderTmpl,
		"templates/header.tmpl",
	)
}

func templatesHeaderTmpl() (*asset, error) {
	bytes, err := templatesHeaderTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/header.tmpl", size: 151, mode: os.FileMode(420), modTime: time.Unix(1464036481, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesInlineTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\xca\xcc\xcb\x01\xd2\x4a\xb5\xb5\x0a\xd5\xd5\x25\xa9\xb9\x05\x39\x89\x25\x40\xd1\xe4\xc4\x9c\x1c\x25\x05\x3d\xb0\x68\x6a\x5e\x4a\x6d\x2d\x20\x00\x00\xff\xff\xaa\xeb\x41\xff\x31\x00\x00\x00")

func templatesInlineTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesInlineTmpl,
		"templates/inline.tmpl",
	)
}

func templatesInlineTmpl() (*asset, error) {
	bytes, err := templatesInlineTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/inline.tmpl", size: 49, mode: os.FileMode(420), modTime: time.Unix(1464034914, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesInputsTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\xcb\x41\x0a\x02\x51\x08\xc6\xf1\xab\xc8\x30\xcb\x78\x07\x08\x3a\x40\xbb\xae\xf0\x60\x3e\x43\x28\x09\xc7\x56\x1f\x73\xf7\x46\x57\xad\x94\x9f\x7f\xc9\x0d\x6a\x0e\x59\xcc\x3f\xdf\xdc\x97\xe3\x20\x57\x95\xeb\x4d\x46\xad\xa6\xb2\xea\x78\x84\x79\xde\x3b\x28\x8c\xe9\x4f\xb4\xcf\x98\x6f\x24\xe2\xe4\xcc\x41\x36\xd4\xe7\x45\x48\xf8\x56\x35\x5e\x3b\xfa\xec\x67\xfb\xef\x35\x7e\x01\x00\x00\xff\xff\x43\x89\x5c\xae\x80\x00\x00\x00")

func templatesInputsTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesInputsTmpl,
		"templates/inputs.tmpl",
	)
}

func templatesInputsTmpl() (*asset, error) {
	bytes, err := templatesInputsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/inputs.tmpl", size: 128, mode: os.FileMode(420), modTime: time.Unix(1464034914, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMessageTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x8e\xd1\xca\x82\x40\x10\x46\x5f\x65\x10\x85\xff\x07\x99\x07\x08\x7a\x80\x6e\x22\x22\xba\x5f\xf2\xd3\x06\x74\xb3\xdd\xd5\x88\x61\xde\x3d\x15\x13\xba\xfa\xe0\x70\x38\x33\xaa\x15\x6a\xf1\xa0\xac\x43\x8c\xae\x41\x66\xa6\x2a\x35\xf9\x47\x22\x3e\x05\xf1\xe9\xe0\xfb\x21\x45\xb3\xe2\xc9\xa4\x0a\x5f\xcd\xc6\x4b\xd2\x9d\xf8\x8c\x1b\x64\x44\x98\x09\x5f\xde\x3d\xf8\xea\xda\x01\x66\xbc\x89\x7c\x74\xdd\x04\xfe\x96\xe8\x6f\x50\x35\x38\xdf\x80\x72\x29\x29\x47\x4b\xbb\xfd\x24\xb8\x30\xf9\x09\x21\xae\x7f\xe4\x62\x56\x7e\xef\x16\xe3\xd6\x5d\xe6\x7f\xdd\x4f\x00\x00\x00\xff\xff\xeb\x6d\x22\x24\xc6\x00\x00\x00")

func templatesMessageTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesMessageTmpl,
		"templates/message.tmpl",
	)
}

func templatesMessageTmpl() (*asset, error) {
	bytes, err := templatesMessageTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/message.tmpl", size: 198, mode: os.FileMode(420), modTime: time.Unix(1464034914, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesResultsTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\xcd\x41\x0a\xc2\x40\x0c\x05\xd0\xab\x7c\x4a\x97\xa5\x07\x10\x5c\x8a\x7b\x6f\x20\x34\x95\x81\x92\x81\x3f\xd3\x55\xc8\xdd\x4d\x6a\x51\x70\x35\x93\xfc\x97\xc4\x6c\x91\xb5\xa8\x60\xa0\xb4\x7d\xeb\x6d\x70\x87\x19\x9f\xfa\x12\x8c\x65\xc2\x28\x1b\x2e\x57\xcc\x8f\x4f\xec\x6e\x56\xd6\x48\xdc\xa7\x70\xa2\x4b\x76\xee\xb5\x63\xce\xcf\x59\x87\x88\x81\xbe\x53\xdb\x8d\xac\x4c\x2c\xe4\x99\xe3\x00\x95\xdf\xa5\xff\x38\x0f\xfe\xec\xf1\xbe\x03\x00\x00\xff\xff\xb0\x4f\xcf\x61\xa8\x00\x00\x00")

func templatesResultsTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesResultsTmpl,
		"templates/results.tmpl",
	)
}

func templatesResultsTmpl() (*asset, error) {
	bytes, err := templatesResultsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/results.tmpl", size: 168, mode: os.FileMode(420), modTime: time.Unix(1464034914, 0)}
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
	"templates/call.tmpl": templatesCallTmpl,
	"templates/function.tmpl": templatesFunctionTmpl,
	"templates/header.tmpl": templatesHeaderTmpl,
	"templates/inline.tmpl": templatesInlineTmpl,
	"templates/inputs.tmpl": templatesInputsTmpl,
	"templates/message.tmpl": templatesMessageTmpl,
	"templates/results.tmpl": templatesResultsTmpl,
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
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"call.tmpl": &bintree{templatesCallTmpl, map[string]*bintree{}},
		"function.tmpl": &bintree{templatesFunctionTmpl, map[string]*bintree{}},
		"header.tmpl": &bintree{templatesHeaderTmpl, map[string]*bintree{}},
		"inline.tmpl": &bintree{templatesInlineTmpl, map[string]*bintree{}},
		"inputs.tmpl": &bintree{templatesInputsTmpl, map[string]*bintree{}},
		"message.tmpl": &bintree{templatesMessageTmpl, map[string]*bintree{}},
		"results.tmpl": &bintree{templatesResultsTmpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
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

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

