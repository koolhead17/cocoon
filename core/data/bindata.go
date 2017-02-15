// Code generated by go-bindata.
// sources:
// bindata.go
// cocoon.job.json
// DO NOT EDIT!

package data

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

var _bindataGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func bindataGoBytes() ([]byte, error) {
	return bindataRead(
		_bindataGo,
		"bindata.go",
	)
}

func bindataGo() (*asset, error) {
	bytes, err := bindataGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata.go", size: 0, mode: os.FileMode(420), modTime: time.Unix(1487159267, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cocoonJobJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x54\x4d\x6f\xe2\x3c\x10\xbe\xf7\x57\x58\xd6\x7b\x2c\x50\xde\x55\x2f\x48\x7b\x60\x93\x6e\xc5\x8a\x2f\x15\x38\xad\x50\x35\x4d\xa6\xa9\x85\x63\x47\xf6\xc0\x96\x8d\xf2\xdf\x57\x76\x42\xc9\x07\xb4\x5c\x98\xcc\xf3\x8c\xe7\x99\xf1\x78\xf2\x1b\xc6\xf8\x2f\xfd\xc2\x47\xcc\x99\x8c\xf1\x27\x4c\x84\x56\x7c\xc4\x78\x22\xf5\x0b\x48\x7e\x5b\xfa\x27\xa1\xf3\x45\x3a\xd2\x5a\xf5\xf2\x7c\x12\x16\xc5\x09\x9a\x43\x8a\x57\xc1\xf5\x31\xf3\xa0\x45\x73\x10\x11\x9e\xdc\x4b\x23\xb4\x11\x74\xe4\x23\x76\x7f\x57\xf9\xc6\x52\x8e\x69\xa1\x22\xc7\x7f\x05\x69\xb1\xf2\x87\x40\x10\xa1\x22\x34\x96\x8f\xd8\x6f\xef\x64\x8c\xc7\xd1\x90\x7b\x7b\x5b\xf1\x02\xad\x2c\x19\x10\x8a\xea\xbc\xbc\xfa\x67\x8c\x4f\xd7\x60\x12\x24\x27\xe7\xbf\x1c\x88\x4c\x7f\x87\x46\xa1\xec\x2b\x48\xf1\xa4\xb8\xec\xc2\x99\x29\x85\xda\xbf\xd7\xb1\x45\x86\x06\x54\xec\xb0\xef\xbc\x72\x17\x0d\x21\x6b\xb0\xbb\x47\xa3\xf7\xd9\x15\x1d\x9f\x35\xac\xaa\x64\xaf\x5c\xf2\x3c\xf7\x56\x51\x34\xb0\x7a\x95\x6a\x2f\x65\x0d\x74\x99\xeb\x49\x9b\x89\x2f\x25\x27\xb0\xbb\xae\x82\xb2\xef\x46\x1c\xd0\x38\x6e\xac\xa3\x1d\x9a\x36\x1e\x68\xf5\x2a\x92\x8f\xd1\xa9\x21\x22\x85\xc4\x27\xc9\xf3\x89\x33\x3b\x67\x33\x97\x3f\x4d\xab\x36\xbe\x80\x7d\xeb\x12\xc0\x24\xbe\x14\xde\x8b\xf8\x2d\x97\x96\x0d\xa4\x8e\x40\xf2\x6d\x83\x58\xb4\x44\x3d\xa8\xc3\x25\x45\xc1\x22\x58\x2c\xe6\xcf\xd7\xa7\xb8\xcb\x0d\x16\xe1\xc3\xf3\xe6\x69\x5a\xd6\x11\xf8\x98\x40\xc7\xb8\x79\x9a\x7e\x15\xb5\x1e\x3f\xb6\xa3\xd6\x90\x7c\x15\x35\x1d\xcf\x3b\x61\x53\x50\x2e\xee\xd3\x92\x57\xe5\xd3\xf2\xcd\xda\xb6\xb0\x19\x12\xb8\x7e\xb4\x63\xa6\x3a\xb9\x7e\x7d\x33\x78\xff\x29\xa4\x3f\x70\x78\xd7\x91\x5c\xa1\x2b\xf1\x17\x67\x3f\x3c\xe5\x53\x75\x6b\x4c\x33\x09\x74\x59\xde\xd8\x90\x78\x85\xa8\x7c\xb0\x6d\x1d\x8c\xf1\x47\x24\x42\xb3\xd2\x7b\xe3\xb7\x02\x7f\x23\xca\xec\x68\x30\xb0\xa4\x0d\x24\xd8\x4f\xb4\x4e\x24\x42\x26\x6c\x3f\xd2\xe9\x60\x67\x74\x02\x6a\x10\xe3\xa1\x17\x63\x26\xf5\xb1\x17\x69\xa5\x30\x22\x6d\xfa\x17\x46\xcc\xaf\x3b\x09\x24\x0e\x18\xa2\xf5\xaf\xbd\x9a\xb1\x66\x45\x6d\xd5\xa1\xb0\x19\x50\xf4\xb6\x84\xa3\xd4\x10\xfb\xfe\xd6\x18\x67\x7b\x5b\x5f\x29\x68\x7d\x19\xb6\xd5\x71\x1e\x2c\x37\xe5\x53\x5f\x6e\x8a\x46\xf3\xf8\x0c\x53\x6d\x8e\xbe\xc9\x79\x7e\xfa\x68\x71\x42\x61\x77\x15\xa3\x34\x5b\xf8\x64\xb1\x5c\xf1\x11\x6b\x5c\x23\x9f\x23\xfd\xd1\xa6\xdc\x14\xe7\xe7\x54\x34\xd5\x12\x18\x5a\x6a\x29\xa2\x63\x5b\xf1\xc4\xed\xe2\x03\x48\x3e\x62\xdf\xee\xce\xbf\x46\x8a\x31\x11\xa6\x19\x75\x67\x88\x87\x28\xc1\x1d\xf9\xff\xfd\xe5\xc8\x99\x8e\xfd\x5d\xc7\x9e\x77\x51\xdd\xc7\x54\x5f\x5c\xbf\x9b\x2c\x06\xc2\x9a\x68\xbe\x22\x48\x12\xbf\xca\x86\x5d\xb9\x6e\xa0\x97\x60\x40\x4a\x74\x15\x0d\x6f\x4e\x27\x16\x37\xc5\xbf\x00\x00\x00\xff\xff\xb8\x08\x01\xc9\x20\x07\x00\x00")

func cocoonJobJsonBytes() ([]byte, error) {
	return bindataRead(
		_cocoonJobJson,
		"cocoon.job.json",
	)
}

func cocoonJobJson() (*asset, error) {
	bytes, err := cocoonJobJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cocoon.job.json", size: 1824, mode: os.FileMode(420), modTime: time.Unix(1487159260, 0)}
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
	"bindata.go": bindataGo,
	"cocoon.job.json": cocoonJobJson,
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
	"bindata.go": &bintree{bindataGo, map[string]*bintree{}},
	"cocoon.job.json": &bintree{cocoonJobJson, map[string]*bintree{}},
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

