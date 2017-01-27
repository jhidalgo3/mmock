// Code generated by go-bindata.
// sources:
// tmpl/css/style.css
// tmpl/index.html
// tmpl/js/script.js
// DO NOT EDIT!

package console

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

var _tmplCssStyleCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x90\xcf\x6e\xa3\x30\x10\x87\xcf\xeb\xa7\xb0\x94\xeb\x82\x94\x28\xd2\xae\x8c\xb4\x52\xb6\x6d\x5e\xa3\xf2\x9f\x01\x2c\x26\x1e\x34\x38\x4d\xa2\x28\xef\xde\x81\xd0\xa6\x87\x52\x2e\xc0\xef\x1b\x7f\xe3\x19\x63\x8a\x13\xb8\x2e\xe6\x62\xf0\x4c\x88\xce\xb2\xbe\x2a\x2d\xcf\x29\x86\xdc\x1a\xbd\xde\xf4\xe7\x4a\xdd\xd4\x37\x85\x45\x66\xeb\xbb\xb9\xdc\xc9\x67\xc3\x74\x4c\xa1\xf0\x84\xc4\x46\xaf\x9e\x76\x2f\xdb\xfd\xae\xba\x63\xe2\x00\x5c\x20\xd4\x59\x9c\xfd\x59\x0f\x84\x31\xe8\x95\xf7\x7e\xd1\xde\x1e\x0f\x6e\xd9\xbe\xdd\xfc\x75\xde\xfe\x78\xd8\xb4\xf4\x06\xbc\xac\x58\x3f\xff\xd9\xfe\xdf\x8f\x0a\xa5\xca\x01\x10\x7c\x86\xf0\xca\x74\xd2\x57\x5d\x53\xca\xa2\x8d\x4d\x2b\x17\x76\x84\xa1\xba\x49\x4d\xe6\x98\x1a\xa1\xb3\xa1\x61\x80\x54\x69\x21\x49\xda\x8d\xad\x3e\x48\xb0\xdc\x11\xdb\xd4\xc0\x84\x1d\x11\x82\x4d\x0f\xee\xf0\x08\xf3\x41\xc4\x47\x7c\xb0\x0d\xa4\x6c\x27\xd2\xc1\xe5\x01\x18\xc2\x18\xaa\x6c\x1d\x82\x2e\x25\x45\xdb\x0f\x50\x46\x71\xaa\x5f\x21\x0e\x3d\xda\x8b\x99\x68\x21\x03\x4c\x33\x95\xf7\xdf\xfb\x12\xfe\xe9\xec\x28\x5c\xc6\x37\xcf\xb3\x8e\x85\xe6\x93\x86\xdf\xea\x4b\x3e\x25\xcb\x9b\xab\xeb\x5a\x5a\xbc\x07\x00\x00\xff\xff\xc8\xe4\x9f\xe6\x3d\x02\x00\x00")

func tmplCssStyleCssBytes() ([]byte, error) {
	return bindataRead(
		_tmplCssStyleCss,
		"tmpl/css/style.css",
	)
}

func tmplCssStyleCss() (*asset, error) {
	bytes, err := tmplCssStyleCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/css/style.css", size: 573, mode: os.FileMode(420), modTime: time.Unix(1485530015, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x58\xdd\x72\xdb\xba\xf1\xbf\x8e\x9f\x02\x87\xff\xff\x74\x9c\x49\x48\x5a\x56\x1c\x3b\x8e\xc4\x19\xc7\xd6\x69\x8e\xc7\x8a\x9d\x58\xb5\x7b\xce\x4d\x06\x22\x21\x11\x32\x48\xd0\x00\xa8\x8f\x64\x3c\xd3\x27\xe8\x7b\xf4\xa2\x2f\xd1\xbe\x49\x9f\xa4\x0b\x80\x14\x69\x89\x52\xec\x73\x66\xea\x49\x24\x02\xdc\x5d\x2c\xf6\xeb\xb7\xab\xce\x4f\x67\x97\xa7\x83\x5f\xaf\x7a\xe8\xe3\xa0\x7f\x11\xec\x74\x62\x95\xb0\x60\x07\xbe\x09\x8e\x82\x1d\x04\x7f\x9d\x84\x28\x8c\x62\xa5\x32\x97\xdc\xe7\x74\xda\x75\x4e\x79\xaa\x48\xaa\xdc\xc1\x22\x23\x0e\x0a\xed\xaa\xeb\x28\x32\x57\xbe\xe6\x7f\x8f\xc2\x18\x0b\x49\x54\xf7\x2f\x83\x9f\xdd\x23\x07\xf9\x85\x24\x45\x15\x23\x41\xbf\xcf\xc3\x3b\x04\x42\x24\x67\xa4\xe3\xdb\x4d\x4b\x20\x43\x41\x33\x85\xa4\x08\xbb\x8e\x3e\x51\x1e\xfb\x7e\xc8\x23\xe2\x4d\xee\x73\x22\x16\x5e\xc8\x13\xdf\x3e\xba\x2d\xaf\xb5\xef\xed\x79\x09\x4d\xbd\x89\x74\x82\x8e\x6f\x79\x9f\x2f\x28\xa1\x63\x81\x15\x01\x81\xfb\x5e\xeb\x59\xf2\xa2\x74\x22\xbd\x90\xf1\x3c\x1a\x31\x2c\x88\x11\x8a\x27\x78\xee\x33\x3a\x94\x7e\x8c\xd3\x88\x91\x21\x18\x02\x04\xfa\x5a\xd7\xbd\xc7\x7b\x6b\x87\xfc\xe4\xba\xe8\x02\x54\x91\x0a\x8c\x9a\x64\x94\x91\x08\x01\x03\x02\x9d\xe8\x88\xc2\xe2\xf4\xfa\x1a\xb9\x6e\x41\xcd\x68\x7a\x87\x04\x61\x5d\x47\xaa\x05\x23\x32\x26\x44\x39\x28\x16\x64\x54\xa9\x98\xe0\x39\x68\xe9\x0d\x39\x57\x52\x09\x9c\xe9\x85\xd6\x72\xb9\xe1\xb7\xbd\xb6\xf7\xd6\x0f\xa5\xac\xf6\x8c\x0d\x60\xc7\x41\x14\x3c\x3b\x16\x54\x2d\xe0\x8c\x18\xb7\x8f\xde\xb8\xad\xfb\xa3\x64\x70\x7e\x79\x72\x3d\x3f\x9a\xb4\x4e\xf2\x57\xf8\xe0\xf6\xec\x26\xbd\xa2\xfb\xec\xee\xe7\xd1\x6c\xd6\x3b\xc1\x47\xf1\xd9\x59\x34\xf9\x8d\x65\x17\x64\x3c\x8f\x27\x37\xfd\x5e\x6b\x34\x9e\xdc\x5e\xfd\x39\xb9\xfb\x26\x0f\x1d\xa3\xbb\xfe\x0b\x05\x97\x92\x0b\x3a\xa6\x69\xd7\xc1\x29\x4f\x17\x09\xcf\xc1\x26\x95\x29\x2e\x33\x45\x79\x8a\x19\x52\x31\x49\xc8\xff\xe0\xe2\xae\x39\x68\xdb\xf5\x47\x17\xb7\xfb\x9f\xf6\x5a\xac\x7f\x3f\xc1\x77\x1f\xee\xe6\x6d\xe6\xf7\xdf\xf5\x70\x9c\xcf\xb2\xeb\x11\xf9\x34\xbd\x79\xdb\x3e\x3f\x20\xdf\xd2\x76\xfe\xdb\x37\x9c\x0d\xf6\xf2\xc3\xde\xaf\xf2\xaf\xfd\xc9\xe7\x9b\x57\x7b\xbd\xf4\x40\x3c\xe7\xfa\x5b\x23\xe1\x1c\x4f\xf1\xb5\x8d\xc9\xa5\x5d\x9a\x62\xf4\xa9\x76\x98\xac\xfa\x7f\xd2\x78\xff\xbd\xe4\x7a\x78\x7e\xd6\xfb\x48\x31\x1b\x25\xf9\x87\x0f\x9f\xaf\xde\x9e\xbc\xf9\x2c\x32\x71\x7f\x70\x79\x33\xba\x6d\x1f\x5e\x7d\xf9\xd2\x9e\x1c\xf4\x2e\xee\xe7\x52\xb6\x16\x37\xf7\x97\x2a\x25\x59\xfa\xf1\xe6\xea\x1d\x3e\x3f\x9c\x5f\xff\xf8\xfe\xcd\x79\xa7\xa0\xd2\x14\x05\xa6\xba\xb9\x63\xaf\x0a\xaa\x5b\xaa\xa6\x94\x6a\x8e\x95\x9a\x38\xe3\x68\x1b\x3b\x3a\x16\x0c\x95\xf1\x7e\xbd\xf8\xa5\x38\x01\xfa\x29\x25\xb3\x8c\x0b\x55\x2b\x79\x33\x1a\xa9\xb8\x1b\x91\x29\x0d\x89\x6b\x16\xaf\xc1\x6a\x54\x81\x7d\x5c\x19\x62\x46\xba\x2d\xe7\x29\x37\x09\x0a\xbb\xfc\xff\x2e\x8a\x78\x98\x27\x20\x1c\xbd\xf4\x04\xd4\xe0\xc5\xee\x28\x4f\x43\x9d\x08\xbb\x2f\xd1\xf7\xa5\xf9\x10\x9a\x62\x01\x17\x83\x12\x26\xd5\x05\x1f\x8f\x89\x40\x5d\x94\x92\x19\xfa\x52\xdf\xdb\x7d\xf9\xbe\x60\x79\x61\x39\x66\xb2\x20\xbb\x25\xc3\x6b\x28\xc3\x44\xed\x3a\x33\x1d\x28\x0e\x7a\x85\x18\x0f\xb1\x3e\xc9\x8b\x39\x04\xde\x2b\xe4\xf8\x24\x8c\xb9\x53\x97\x31\x93\x1e\x4f\x13\x22\x25\x1e\x13\x90\xb4\xd4\x8d\x4c\x41\xe5\x4a\xc1\x17\xcb\x03\x2b\xda\xf3\xeb\xcb\x4f\x5e\xa6\x71\xc1\x52\x7b\x11\x56\xb8\x2e\x1b\xfe\x3d\xba\x90\xc7\xf8\xb8\x97\x2a\xb1\xd8\x2d\x84\xbc\x7c\xff\xc2\x12\x57\x3c\x0f\x25\xff\x43\x25\xc9\xda\x7b\x19\x07\x1d\xdf\x62\xd9\x4e\x67\xc8\xa3\x45\xe1\x8d\x58\x14\x0f\x11\x9d\xa2\x90\x61\x29\xc1\xff\xe0\x55\x4c\x53\x22\xdc\x11\xcb\x69\xb4\x74\x8a\x26\x6f\xad\xa2\x16\xec\x54\xaf\x6b\x42\x04\x9f\xd5\x18\xcd\x5b\x2c\x69\x44\xaa\x43\x98\x9b\x44\xee\x3e\xca\x70\xe4\x42\xf8\xc7\xca\xdd\x5b\x61\x30\x4c\x39\x2b\x39\x52\x3c\x85\x00\x9c\xba\x50\x05\x98\x34\x4f\x52\x61\xf0\x5c\xd4\xc0\x56\xc4\x7c\xc9\x8a\xc1\x39\x53\x02\x49\x81\x8b\x10\x9f\x40\xc8\x59\xbb\x1c\xbf\x77\x82\x22\x54\x64\xc7\xc7\x90\x37\x8c\x6e\x94\xb7\x41\x00\xd2\x1e\x74\x15\x38\x0b\x02\xdd\x49\x78\x84\x59\xb9\x87\xc5\x18\xe0\xdf\xf9\xbf\x64\xd1\x37\xdb\xc1\xc9\x90\xe7\x6a\xf3\x41\x1d\x3f\x67\x2b\x66\xf3\x8d\xdd\x56\x36\x1f\xb9\xcb\x58\xb2\xd5\x68\x3e\x28\xa0\x34\x1d\xf2\x39\x52\x9c\x33\xc0\xdb\xb2\x48\x6e\x12\xb6\xee\xb6\x4d\x47\xce\xa5\xdb\xda\xdf\x40\xbb\x4a\x9f\xe5\x8c\x59\x27\x6f\x61\x30\x4c\xb8\x64\x19\xaa\x14\xc1\x7f\x37\x22\x23\x9c\x33\x65\x9e\xe7\x50\xa2\x4c\x97\xd4\x75\xbe\x90\x91\x80\x2a\x06\xb5\x39\x32\xb4\xa7\x8c\x60\x01\xe9\xf2\x03\xf9\xfa\xcf\x90\x6e\xd7\x02\xdc\xb3\xf9\x5e\x3e\x5c\x6c\x83\x89\xb6\xbc\x02\x47\xf8\xda\x6c\x0d\xf6\x37\xef\x9b\xac\x8b\x64\x86\x43\x22\x0e\x74\x35\x6f\x94\xbc\x69\x5b\x1f\x06\x9e\x6c\x74\xb6\x7e\xf7\xac\x90\xc8\x70\x4a\x18\x32\x9f\x4b\x6f\x54\xf5\xc9\x7a\xa0\x5a\x5f\x69\xba\x4d\x21\x04\x67\x97\x75\x90\x51\xa9\xb6\x1a\xa3\x26\xd6\xd5\xb9\xe4\x94\xfa\x28\x3c\x64\xc4\x05\xef\x67\x50\x83\x4c\x5a\x6f\xf6\x95\xa1\x7d\x24\xca\xec\x3c\x96\x85\xac\xc4\x98\x4f\x89\x86\x12\x13\x58\xee\xb2\x08\x22\x13\xbc\x11\x9f\xa5\x80\xb3\x1a\x15\x4b\x0d\x18\x5e\x40\x2a\x1f\xa3\x11\x9d\x93\xc8\x16\xdd\x19\x17\x91\x3b\x83\xd6\xe1\x18\x0d\x01\xb4\xee\x5c\xbd\x01\x05\x66\xa7\xa3\x74\xed\x2d\x4f\x8d\x69\x14\x11\x13\xce\xe5\x3c\x20\x6a\x45\x54\xc5\xf5\x40\x90\x89\xdb\x06\xff\x43\x43\xc2\xd3\x71\x70\x06\xbd\x10\x54\x74\xbb\x80\x91\x21\xde\xc6\xd7\xaa\xf8\x4e\xa1\xe5\xff\x3d\x7c\x7d\xa2\x62\x1e\x3d\x99\xf3\xb0\xe2\xbc\xc2\x2a\x6e\xe0\x83\x27\xa1\xa1\x48\x59\x2c\xea\x28\x8b\x45\xb0\xb1\x7c\xd0\xc6\xd5\x0f\x26\xb4\xcb\xaf\x2a\x6a\x6d\x20\x42\xe4\xd4\x63\x74\x46\x18\x43\xfa\xc3\x95\x48\xf7\x14\x65\xb1\xe9\x90\x24\xb8\xc0\x72\x19\xb0\x28\xcf\x20\x96\x48\x74\x0c\x7d\x08\x08\x32\x91\x01\x12\xd4\xd7\x62\xdf\x09\x6e\x31\x74\x2d\xe9\xd8\xf3\x3c\xd0\x1e\x48\x40\x77\x90\xd1\xa0\x8d\xbe\xf2\xbb\x52\x8f\x35\x7d\xed\xd7\x8e\x21\x45\xa6\xf4\x23\x53\xf9\x57\x15\x37\x60\x81\x46\x38\x22\x36\x91\x4a\x9c\x40\x82\xeb\x48\x8b\xa0\x81\x5a\x16\xb5\x35\x3e\xd7\xbe\x46\x76\x21\x13\xa7\x19\x8a\xed\xeb\xa2\x57\x5b\x05\xe5\x35\x3a\xed\x18\x22\x9a\xb0\x64\x98\x2b\xc5\xd3\xa2\x73\xb3\x8b\x65\x22\xc1\xf4\x27\x49\x01\x7a\x11\x95\x09\x5d\x0a\x74\x82\x3f\x29\x0a\x59\xff\xbe\xe3\x5b\x9e\x06\xc9\xf1\x9b\xc7\x3a\x98\x22\xef\x14\xad\x06\x74\x68\xb5\x76\xe3\xcd\x2a\x3a\xae\x55\xc0\xf5\x2b\xe9\xd0\x6a\xba\x50\x16\xf4\x13\x7d\x02\x95\x30\x54\xc0\xe0\x3f\xb8\x42\x7a\x0d\xee\x47\x38\xcb\x18\xb5\x5d\x20\x1a\x71\x81\xf4\x08\x62\xf6\x61\xf8\x18\xe9\x80\xca\x04\x57\x1c\x4c\x01\x9b\x1d\x3f\x6b\x14\x7e\x01\xcd\x70\x1a\x12\x50\x3e\x5b\x98\x78\x44\xff\xfa\xc7\x7f\xfe\xf6\x77\xb4\xbf\xd7\x7a\x8b\x5c\xfd\x75\xf4\x1a\x9d\x43\x85\xa0\xa8\x8f\x85\xfa\xf7\x3f\x53\xb4\xbb\x6c\x30\xf4\xe0\x02\xed\xe8\x44\xbf\xf6\x28\x07\xe0\x2b\x5a\x89\xaf\x43\x86\xd3\x3b\x27\x58\x21\xd0\x98\xf5\x72\x93\x26\x5f\x08\x60\x9e\x84\x71\x29\x4f\xc1\xb7\xa8\xff\xcb\x00\xea\x2f\x28\x27\xc9\x6b\x24\x09\x41\x8f\x4e\xd5\x5d\xf0\x98\xaa\x38\x1f\xda\x5f\x07\x12\xd0\x8d\xa6\x47\xfb\x7e\xa2\x8d\xe3\x0f\x19\x1f\xc2\x38\x25\x15\x11\xfe\xc5\x2f\xa7\xbd\x4f\xd7\x3d\x2f\x89\xd6\xf5\x2b\xde\xd9\x56\x27\x0b\x56\xb5\xd2\x46\x8d\x60\xa6\xa0\x4c\x7a\xda\xbd\x4a\x50\x08\x0f\x72\xfc\x1c\x5d\x9a\x8d\xb2\x9d\x67\x0d\xdb\x9f\x14\x41\x23\x98\x09\x9f\x9f\x14\x2b\x3d\xcc\x86\xf4\x38\xd5\xa9\xd3\x9c\x1c\x2b\xba\xd5\x96\x2b\x75\xc6\x94\x99\x41\xaf\x7f\x75\x71\x32\xe8\xd9\xdf\x46\xca\x41\xab\x0e\x7e\x44\x4f\x11\x8f\xc6\xbe\xb9\x5b\xfd\x14\xe3\x2a\x92\x64\x0c\xca\x20\xdc\xb3\xc4\xa5\x2a\xbf\x69\x78\x67\xc0\xf2\xfb\xf7\x42\xda\x57\xa8\x80\x5c\x3c\x3c\xac\xb4\xba\x4b\xcc\xb6\xce\x75\x2b\xfa\x34\x4f\x80\xba\x8e\x1f\xd1\x3a\xd2\x55\xe4\xba\x20\x3f\x3c\x00\x20\x44\xdb\x58\x34\x58\x0d\x83\xba\x52\x91\xe1\x1a\x06\x4f\xe0\xac\xd8\x12\x03\x74\x3f\x3e\xee\xb0\xce\x94\x01\xc6\xd5\x59\x2c\xb8\x95\x96\x7b\xd4\xc0\x6c\x30\x46\x29\x5d\x42\x82\x86\x80\x5a\xd0\x7d\x97\x4d\x06\x84\x09\x38\x63\x71\x8c\x52\x9e\x12\xd3\x42\x80\x2e\x9c\x69\x4c\xea\x3a\x6f\xea\x90\x50\x30\xc0\x20\x15\x41\x35\x3a\x46\x7b\xd9\x1c\xb5\x0e\xea\x1f\x9a\x3d\x6e\x07\x67\x56\x0b\xa8\xa2\x6d\xbb\x61\x41\xb0\xd0\x81\xe1\x21\xf4\x79\xe6\xd3\xcd\x04\x85\xb4\x59\x2c\x47\xa3\x25\x16\x1a\xce\x2a\x0a\x33\xb1\x1c\xe8\xe0\x11\x46\x7d\x80\x2d\x66\x3a\xac\xa0\xa3\x3d\x01\xc6\x2a\x6f\xfc\xa0\x2d\x65\xf6\xa0\x1c\x88\x1a\xc2\x6f\xd3\x63\x3d\xd8\x40\x23\xd3\xf8\x91\x3f\xa6\x92\x95\xf1\xbb\x74\x2a\xb3\x39\x80\x6e\xf7\x0f\x29\x01\xd0\xbd\x49\x81\xaa\x8f\xd0\xb1\x55\x36\x4d\xd5\x24\x6f\x9b\x25\x3d\xd2\xeb\x9f\xa9\xff\x1b\x00\x00\xff\xff\x8c\x79\x99\xeb\xbd\x16\x00\x00")

func tmplIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_tmplIndexHtml,
		"tmpl/index.html",
	)
}

func tmplIndexHtml() (*asset, error) {
	bytes, err := tmplIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/index.html", size: 5821, mode: os.FileMode(420), modTime: time.Unix(1485530860, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplJsScriptJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x56\x69\x6f\xe3\x36\x13\xfe\xee\x5f\xc1\xd5\x06\x2b\x2a\xb6\x65\xef\x8b\x7c\x79\x95\xab\x6d\x9a\x22\x28\x36\x5b\x74\x93\xa2\x68\x23\x27\xa0\x25\xda\x56\x23\x91\x2e\x49\x75\xd7\x8d\xfd\xdf\x3b\xa3\xc3\x92\x18\x27\x8e\x00\x1f\x1a\x3e\xf3\x70\x38\x17\x67\x96\x8b\xc8\x24\x52\x90\x2f\xfc\xef\x9c\x6b\xf3\x49\xce\xe7\x5c\x51\x8f\x3c\xf5\x08\x3c\xa3\x51\x24\xb3\x65\x92\x72\x62\x16\xf0\xe1\xd9\x32\x65\x86\x17\x4b\xff\x30\x45\xb4\xcc\x55\xc4\xc9\x29\x39\xa0\xce\x7b\x55\x32\x0c\xb9\x30\x6a\xe5\x78\xfe\xc2\x64\x29\xf5\x8e\xb7\xe0\x5a\x1b\xe0\x57\x4c\xc4\x29\x9f\x32\xa5\xfd\x8a\x9f\x96\x54\x00\x2f\xf0\x66\x91\x68\x5f\xe4\x19\x60\xc7\xc7\x8d\xc4\xac\x96\xa8\xee\x64\x2c\x4a\x84\x91\x7a\xe1\xb4\x16\x23\x99\x4a\x85\xab\x8a\xc7\x4e\xc5\xb3\x3d\xde\x9c\x9b\x0b\x5c\xff\x61\x75\x63\x98\xc9\x35\xd5\xc5\xcf\x85\x8c\x79\x7d\x56\x7c\x92\x19\x69\xad\x90\xd3\x53\xf2\xbf\xf1\x98\xac\xd7\xc4\x16\x7e\x6c\x6b\xe1\xa3\xb8\xc9\x95\x20\x8e\xce\xa3\x88\x6b\x5d\x19\x86\xcf\x86\xf0\x54\xf3\x1d\xd4\x47\xe3\xa3\x97\x58\x62\x26\x20\x0c\xcf\x49\x76\xa3\xbf\x32\x25\x12\x31\x6f\xc3\x7b\xe5\x77\xd7\x09\xf9\x32\x06\xff\xdf\x26\x06\xfc\xdd\xde\xf9\x80\xc6\x32\xca\x33\x08\x9c\xe7\x33\x63\x14\x75\x0c\x62\x9c\x01\x71\x3e\x5f\xfe\x4e\xbe\x5c\xfe\xfa\xdb\xe5\xcd\xed\xbb\x77\x8e\xd7\xec\xa0\xb9\xb9\x4d\x32\x2e\x73\x43\xb7\xfc\x48\xfa\x1a\xd9\xf5\xb5\x8c\x1e\xc9\x85\x14\x5a\x82\xc0\x3b\x26\x9b\x01\xfa\x77\x5c\xd1\xda\xe6\xea\x95\x30\xec\xdb\x55\x32\x5f\xa4\xf0\x31\xf4\x2f\x2d\x45\xdb\x6c\x7c\x87\x78\xe3\x8f\xaf\x38\xe4\x56\xc4\xe9\xe8\xc3\x68\x3e\x20\xee\x07\x96\x2d\x8f\x5d\xaf\x11\x9f\x94\xe2\xd4\x74\xa4\x67\xa5\x74\x8e\xd2\xe6\x68\x95\x5b\xbb\xbc\xd4\xa1\x61\x98\xdf\xb1\xe1\xbf\xdf\x0f\xff\x1c\x0f\xff\x3f\x79\x3a\xda\xac\xc3\xf0\xee\x3e\x9f\xac\xef\xee\xc3\xd0\x99\x78\x87\x00\xd1\x87\x81\x77\xbe\x0e\xa7\xd4\xa8\x9c\xaf\x67\x0c\x82\xb6\x16\x79\x9a\x7a\xe1\x74\x3d\x3c\x0f\xe3\x3e\x3d\x0f\x42\x3f\x8c\x0f\xbd\x73\xf8\x77\xc7\x2f\x27\x77\xfd\x70\x38\xc1\x15\xef\xdc\x43\x73\x1a\x6f\x66\xcc\x44\x0b\x3b\x43\xb0\x92\xa2\x54\xc3\xb9\x5d\xa8\x8f\x29\x57\xee\x71\x67\x1d\xd3\x6c\x74\xef\x8c\x7c\x03\xb5\x58\x51\xd8\x1c\x5b\x5c\x70\xb0\x17\x87\x4f\xb5\xdf\x23\x5f\x59\x9b\xbd\x98\x99\x96\xa6\x36\x0a\xf2\x73\x97\x72\x6f\x07\x55\x61\x5a\xe3\xbf\xbd\x26\x56\x9b\x4c\x25\x64\x15\x13\xd6\x2e\x6d\x4e\x0c\xc4\x5b\xd9\x10\x6b\x53\xed\x2a\x3e\xf7\x44\x2f\x99\x00\x2d\xa6\xf5\xa9\xe3\x92\x7e\x41\xd0\x27\xae\x73\x86\x2f\xc5\x3e\xf8\x7a\x32\x42\xdc\x59\x8b\x72\xf3\x42\xda\x63\xab\xca\x95\x82\x02\xc2\x02\xeb\x14\x6a\x11\xfc\x72\x2d\x2e\x3b\xa9\xe0\x5f\xc9\x8f\xf0\x97\xb6\xf2\x17\x51\xb8\x6c\x40\x1d\x20\x2d\x05\x1f\xb8\x4b\x34\x98\xe4\x8c\x1c\xd2\xef\x1c\x89\x5a\xd0\x6b\x29\xcc\xa2\xc0\x7e\xdc\xad\x60\xe1\x7f\x02\x9f\xfd\xc1\x99\x2a\xe9\xc9\x77\x64\x0f\xfe\x0a\x7a\xbe\x2e\xc1\xc1\x1e\xe8\x75\x22\x72\x08\xdc\xdb\xc0\x37\x3c\x92\x22\xd6\xf4\x79\x4d\xd7\x6e\xd9\xed\xf9\xb2\x3f\x7e\x62\xda\x54\xb7\x61\xe5\xab\x76\xa7\x74\xdf\x43\xa4\xcd\x43\x09\x8d\xa1\x99\x18\xfe\xcd\x50\x3b\x66\xaf\x84\x16\x7c\x8a\x1a\x50\xbc\x03\x34\x87\xd9\xf1\x2d\xef\x08\x88\x1b\x2e\x42\xff\xd1\x4b\xe8\x97\xdc\x6f\xae\x8e\x6e\xa0\x33\x6e\x16\x32\x6e\xe0\x85\xdd\x7e\x29\xed\x22\x97\xcc\x2c\x6c\x1c\xca\xba\xa8\x6a\x05\x80\x76\xfb\xfd\xf9\xe6\x97\xcf\x7e\x59\xc9\xc9\x6c\x45\xdb\x3c\x03\x92\x8b\x98\xcf\x12\xc1\xe3\x01\x39\xf2\x3c\x9b\xb2\x3c\xc3\x9b\x39\x4b\xf8\xeb\xa4\xa9\x9c\xbf\x9d\x2f\x4f\xf7\x98\x58\xcf\x0e\x2f\x8c\x09\x5e\xcf\x4e\xa4\xa7\xda\x51\x0f\x10\xc8\x80\x14\xd1\xac\x24\x01\xd9\x7a\xa5\x3e\x4a\x40\x9a\x43\x29\xb0\x3c\x40\xf3\xb7\x0a\x0f\x98\x4a\xc1\xb3\xba\x6f\xd6\x23\x88\x7a\x50\x25\x46\x23\x2d\x63\x1c\x54\x19\xd0\xc8\x31\xa6\x41\x11\xed\x36\x03\x1c\x2a\xa8\x8e\xb9\xe9\x24\x67\x31\x3b\x81\x39\x97\x38\xb6\x81\x0b\x9a\xfb\xc7\xce\xce\x7a\x28\xeb\xf7\x6d\xe7\x15\x29\x5d\xbb\xaf\xcc\xef\x1a\x5c\x25\x79\x57\x03\x47\x43\x80\xd7\x13\x21\xad\x18\x5a\xa8\xf6\x40\x69\xd8\x14\x07\xd0\xa9\x8c\x71\xac\x5c\xc2\x95\xcc\x45\x4c\x91\xa3\xa5\xd0\x99\x6d\x6c\xf1\xb3\x92\xae\xc6\x43\xf0\x04\x78\xa1\x3d\xb1\x28\xce\xe2\x55\x77\xa2\xe9\xd5\xf6\x4c\x8d\xb8\x80\x1b\x46\xc1\x8c\x0c\x76\x44\x69\x12\x3d\xee\x40\xbe\x6e\x3d\x9c\xd8\xac\x6a\x03\x37\xb5\x1d\xd8\x56\x6a\x3c\xba\x0b\xda\x8a\x14\xd4\x2d\xb6\x70\x61\x42\xa9\x38\x94\xdb\x1e\x0f\xec\xc6\x91\xc4\xc5\x20\x8e\x8e\xf7\x7c\x64\x81\xc9\x8b\x29\x08\x89\x63\xf9\x15\xba\x27\x80\xa1\x73\xc1\xac\xdf\x78\xab\x30\xa6\x87\xdf\xff\x05\x00\x00\xff\xff\xb7\xb0\x93\x44\x14\x0c\x00\x00")

func tmplJsScriptJsBytes() ([]byte, error) {
	return bindataRead(
		_tmplJsScriptJs,
		"tmpl/js/script.js",
	)
}

func tmplJsScriptJs() (*asset, error) {
	bytes, err := tmplJsScriptJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/js/script.js", size: 3092, mode: os.FileMode(420), modTime: time.Unix(1485530015, 0)}
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
	"tmpl/css/style.css": tmplCssStyleCss,
	"tmpl/index.html": tmplIndexHtml,
	"tmpl/js/script.js": tmplJsScriptJs,
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
	"tmpl": &bintree{nil, map[string]*bintree{
		"css": &bintree{nil, map[string]*bintree{
			"style.css": &bintree{tmplCssStyleCss, map[string]*bintree{}},
		}},
		"index.html": &bintree{tmplIndexHtml, map[string]*bintree{}},
		"js": &bintree{nil, map[string]*bintree{
			"script.js": &bintree{tmplJsScriptJs, map[string]*bintree{}},
		}},
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

