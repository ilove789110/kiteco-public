// Code generated by go-bindata.
// sources:
// templates/graph.html
// templates/graphwithextras.html
// templates/head.html
// DO NOT EDIT!

package rendered

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

var _templatesGraphHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\x4d\x6f\xe3\x36\x13\xbe\xe7\x57\xf0\x65\x0e\x96\x01\x47\x49\x5e\xec\x66\x1b\x25\xca\x21\x1f\x5b\xb4\xd8\xcd\x1e\xdc\x9e\x16\x85\x31\x26\xc7\x12\x6b\x9a\x54\x29\x4a\x89\x91\xf5\x7f\x2f\x48\xc9\x8a\xbe\x9c\x36\xa8\x2e\x92\x66\xe6\x19\xcd\x3c\xf3\x21\x5e\x73\x51\x12\x26\x21\xcf\x63\xca\xb4\xb2\x20\x14\x9a\x93\x95\x2c\x04\xa7\x37\x47\x84\x10\xd2\xb6\x30\xfa\xa9\x96\xf6\x35\x4c\xcb\x96\xc6\x6b\x97\x85\xb5\x5a\x11\xc1\x63\xba\x2a\xa4\x3c\x49\x0c\x64\xe9\x49\x25\xa5\x37\x9f\x0b\x29\xc9\xcf\x4e\x74\x7d\x5a\xc9\x0e\xc3\x97\xc0\xd6\x4b\xad\xb0\x01\xdf\xd6\x82\x21\xf4\xfa\x94\x8b\xb2\x0e\xbc\xf5\xd8\xcb\xc1\x7b\xcd\x8c\x4e\x0c\x6c\x7c\x58\xf4\xa6\xb6\xae\x6f\x47\xd7\x39\x33\x22\xb3\xc4\x6e\x33\x8c\xa9\xc5\x67\x7b\xfa\x27\x94\x50\x49\xeb\x4c\xb9\x66\xc5\x06\x95\x0d\x81\xf3\x87\x12\x95\xfd\x22\x72\x8b\x0a\x4d\x30\xb9\xff\xf6\xf5\x4e\x2b\xeb\x64\x1a\x38\xf2\xc9\x8c\xac\x0a\xc5\xac\xd0\x2a\x98\x92\x97\x26\xdc\xc6\xc5\x5f\x05\x9a\xed\x1c\x25\x32\xab\x4d\x40\x8f\x87\x8c\x4d\x43\xad\x98\x14\x6c\x4d\xe2\xc6\x17\x09\xb0\x6c\xbb\x73\x97\x41\xc5\xd1\x78\x66\x83\x5f\xe7\xdf\x1e\xc3\x0c\x4c\x8e\x01\x7d\x79\x09\xbd\x70\xb7\xa3\xd3\xe9\x55\x03\xd9\x5d\xfd\x73\x30\x7d\xfe\xff\x6d\x28\x25\x18\xe2\x73\x20\x31\x39\x14\x4a\x07\xe0\x8d\x43\xe4\x09\xe6\x24\x6e\xbf\x85\x2b\x21\x2d\x9a\xc0\xbd\x90\xf8\x86\xb8\x7b\x28\x61\x89\x92\xc4\x31\xa1\x90\xdb\x05\x4b\x85\xe4\x94\xfc\xf8\xd1\x57\x2a\x7c\xb6\x0b\xab\xd7\xa8\x68\x2b\xef\x3e\x55\xfe\x63\xd3\x31\x5a\xde\x41\xe8\xae\xbe\x37\xa4\x0c\xbf\xd0\xa2\x48\xac\x48\xf0\xbf\x2a\x49\xa5\x39\xe6\x63\xf4\x35\x33\x49\xe2\xd7\xfa\x24\x68\x1f\x24\xba\xc7\xdb\xed\x2f\x3c\x98\xb4\x3b\x79\xd2\xcb\xb1\x71\x10\x42\x96\xa1\xe2\x77\x8e\xa6\xa0\x71\xc5\x0c\x82\xc5\xdf\xf0\xd9\x3e\x6a\x8e\x01\x55\xba\xa2\xbd\xd3\x23\x15\x0b\xb6\x30\xea\x95\x9f\xe6\x29\xb7\x5b\x89\xdc\xa1\x5f\x6b\xe6\xd3\x09\x37\x90\x05\x4d\xd7\x3b\xd1\x74\x2c\x3f\xa9\x8d\x03\x76\x55\xee\xf2\x55\x15\xca\xa2\x51\x20\x17\x0e\x4f\x23\x42\x8f\xf1\x82\xff\xff\xf2\x13\x9d\x8d\xdb\x5b\xdc\x88\x8e\xf9\xc3\xed\xc5\x4f\x1f\xce\xc7\xcc\x4b\x30\x02\x96\x12\x17\x45\x0e\x09\xb6\x10\xf7\x77\x97\x1f\x47\x11\x39\xd3\x59\x63\x48\xe8\xf1\xe5\x27\xbc\x80\x0f\x3d\xcb\xdd\xd5\x78\x92\x24\xae\x93\xfd\xee\x1c\x78\x8a\x16\x6e\xbb\xfc\xd1\xb5\x77\x4d\xe1\xed\xfa\xdd\xe0\x2e\x8f\xec\x78\xeb\x82\x77\x23\x25\xf3\xa0\xd6\xbc\x4f\xaf\x8e\x7a\xb5\x7b\x18\x99\xb7\x4e\xed\x9c\x68\xbc\x39\xdf\x2e\x5e\x35\x92\x11\xa1\x4b\x09\x6c\x3d\xc6\x68\x6b\x36\x23\x42\x39\x98\x75\x62\x60\x3b\x66\x29\x9d\x43\x89\xcf\x82\x81\x5c\x14\xb9\x2f\xd5\xe7\x82\xa5\xb9\x80\x31\x73\xa6\x37\x59\x61\x91\x2f\x56\x46\x6f\x9c\xad\x14\x1b\x4c\x0c\xa2\x1a\xb3\xe6\x60\x61\xb1\x92\xfa\xc9\x59\x6a\x03\x2a\xc1\x83\xf5\x77\x26\x06\xf9\xbb\xab\xfe\xba\x95\xde\x53\x71\x8f\xda\x7b\x7a\xa1\xfe\x89\x46\x95\xcf\x19\xa1\xa9\x48\x52\x29\x92\xd4\xee\x65\xbb\xb7\x1a\xc2\x3b\x03\x63\xf4\x93\x2b\x1a\xb5\x9a\x8e\xcd\xb8\x37\x3b\xd0\x30\x2e\x2f\x55\x8f\xba\xc2\x27\x52\x8a\x3c\xbc\x07\x0b\x73\xb4\x41\x6b\x11\xb4\x76\x87\x43\xec\x17\xfa\x38\xc2\xb7\x5f\x0f\xf1\x9f\xd6\x9e\x73\xe0\x2a\x3a\x68\x4c\x1f\x79\x54\xdd\x66\x03\x66\xf2\xa8\xba\xcd\xc6\x7e\x02\xa9\x40\x03\x86\xa5\xae\xfd\x06\x7e\x51\xb9\x3d\xc2\x23\x62\x4d\x81\x5d\xc7\x12\x4b\x94\x73\xcc\xc0\x80\x1b\xa5\x88\x9c\x9f\x9d\xcd\x06\x41\xcd\x33\x60\x42\x25\x11\x39\xff\xd8\xd3\x72\x61\x90\x55\xc8\xc9\xef\xf7\x93\xae\x32\xd7\xc6\x7e\x45\x9b\x6a\x1e\x91\x49\x65\xe9\x8e\x1a\x63\xf1\x3b\x4e\x74\xe6\x1c\x0d\xe7\x55\xc2\x56\x17\x36\x1a\xe9\x3e\x03\x8a\xeb\xcd\x1c\x5d\x6e\xe7\xc3\x81\x10\x9b\xcc\xe8\x12\xf9\x97\xda\xc1\x30\xfd\x3e\x75\x51\xe7\xad\x37\x42\xdd\x57\xbf\xf9\xa1\xca\x7d\x18\x59\xaa\x4b\x34\xd5\x17\xdf\x72\x92\xa5\xdb\x5c\xb0\x3c\x22\x2b\x90\x39\x1e\x64\x46\xa1\x7d\xd2\x66\xdd\x6a\xd1\xc7\x4a\x12\x34\x8d\x38\xf3\x2d\x35\xdb\x93\xd8\x6a\xb7\x1a\x1c\x6a\x15\x50\x6e\x20\x49\x84\x4a\xe8\xac\x75\x30\x72\xc5\xdf\x0c\xfe\xee\x6e\xec\x2b\x4d\xfd\xbb\x94\xa8\x12\x9b\xba\x23\xcb\xf9\xd8\x26\xf0\x71\x0a\x4e\x62\xd2\x46\x7d\x3f\xeb\x6d\x93\x76\x44\x1b\x5d\xa2\xff\xad\x2b\xc1\x67\x7b\x58\xa6\x3d\xb3\x21\x03\x55\x42\x1e\x3e\x1f\x52\x6c\xa7\x87\x56\xc9\xfe\x94\xb3\x3b\xba\x3e\xad\x0e\xc4\x37\x47\x7f\x07\x00\x00\xff\xff\xa2\xd4\x44\x2b\x48\x0c\x00\x00")

func templatesGraphHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesGraphHtml,
		"templates/graph.html",
	)
}

func templatesGraphHtml() (*asset, error) {
	bytes, err := templatesGraphHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/graph.html", size: 3144, mode: os.FileMode(420), modTime: time.Unix(1569458799, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesGraphwithextrasHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x58\xeb\x6f\x1b\xb9\x11\xff\x9e\xbf\x62\xba\x01\x4e\x52\x21\xaf\x24\xd7\xb1\x63\xd9\x12\xd0\xbc\xda\x2b\xce\xc9\x01\x4e\x71\x1f\x82\xc0\xa0\x96\x23\x2d\x6b\x2e\xb9\x25\xb9\x7a\x9c\x4f\xff\x7b\x41\x72\x25\xed\x83\x52\x9c\x14\xa7\x0f\xde\x5d\xce\x6f\x9e\x9c\x19\x0e\x7d\x4b\xd9\x12\x12\x4e\xb4\x9e\x44\x89\x14\x86\x30\x81\xea\x6c\xce\x0b\x46\xa3\xe9\x0b\x00\x80\x2a\x42\xc9\x55\xb9\xda\xa4\x24\x92\x9f\x5d\x56\x68\x8e\x9e\x2b\x04\x46\x2d\x91\x62\x34\x7d\x7a\x8a\xdf\x14\xf3\x39\xaa\xed\xf6\x76\x90\x2b\xac\x08\x1a\x50\xb6\xfc\x01\xb9\x44\x9b\x33\x83\x6b\xe3\x64\xff\xfd\xfe\xf3\x29\xc1\xd5\xd7\xd3\x1e\x59\xc9\x29\x12\x93\x91\x3c\x02\x6d\x36\x1c\x27\x9d\x14\xd9\x22\x35\x63\xb8\x18\x0e\xf3\x75\x67\xfa\x43\x72\x0f\x1e\x35\xfd\x99\x15\xc6\x48\xe1\x14\xcf\x0b\xce\xcf\x16\x8a\xe4\xe9\x99\x5f\x8d\xa6\x1f\x0a\xce\xe1\x1f\x76\xe9\x76\xe0\xd7\x8e\xb3\xcf\x48\xf2\x38\x93\x02\xf7\xcc\x6f\xca\x85\x23\xac\x3b\x7f\x35\x67\x14\x55\x34\xbd\x63\x02\x7e\x73\xbe\xd6\x70\x0e\xcb\x44\x5e\x98\x9d\x17\x33\xa2\x22\x30\x9b\x1c\x27\x91\x22\x62\x81\x91\x13\xe3\x5e\x1d\x2e\x82\x8c\x89\x49\x34\x8c\x87\x11\x64\x64\x3d\x89\x46\xf6\x6d\x49\x78\x81\xe5\xaa\x36\x98\xbb\xd7\x51\x04\x52\x24\xa9\x65\x2d\x25\x38\x58\xec\xc1\xee\x6f\x34\x98\xb6\xed\x91\x85\xb1\x06\xed\xf5\x7a\xe4\x74\x78\x3b\xf0\x94\x86\xab\x8d\x1c\x7b\xd6\x0e\x3a\xe1\xb9\x92\x0b\x45\x32\xb7\x29\xd1\x6e\xeb\xcb\xc7\x8b\x5b\x9d\x28\x96\x9b\x32\x14\x36\x19\x07\xff\x21\x4b\xe2\x57\xcb\x7d\xa6\x32\x29\x32\x14\x26\x26\x94\xbe\x5f\xa2\x30\xbf\x30\x6d\x50\xa0\xea\x76\xde\x7d\xba\x7b\x2b\x85\xb1\x6b\x92\x50\xa4\x9d\x3e\xcc\x0b\x91\x18\x26\x45\xb7\x07\x4f\x7b\x73\xf7\x22\xfe\x5b\xa0\xda\xdc\x23\xc7\xc4\x48\xd5\x8d\x5e\xb6\xf3\xa5\x17\x4b\x91\x70\x96\x3c\xc2\x64\x2f\x0b\xba\xb8\xac\x8a\xb3\x3f\x85\x82\xa2\x72\x79\xd5\xfd\xd7\xfd\xa7\x8f\x71\x4e\x94\xc6\x6e\xf4\xf4\x14\xbb\xc5\xed\x36\xea\xf5\x6e\xf6\x2c\xdb\x9b\x6f\x1b\xd3\xcc\xbe\xe7\x9a\xb2\x24\x0a\x9c\x0f\x30\x81\x63\xa6\xdc\xd4\x38\x1c\x3a\x46\xba\x40\x0d\x93\xea\x57\x3c\x67\xdc\xa0\xea\xda\x0f\x98\x4c\xc1\x3e\x63\x4e\x66\xc8\x61\x32\x01\xdb\x31\x1e\x92\x94\x71\x1a\xc1\x1f\x7f\x34\x89\x02\xd7\xe6\xc1\xc8\x47\x14\x4d\x75\xd5\x58\x39\x65\xe1\xc0\x58\x3f\x5c\x2a\xfe\xec\x4a\x65\x72\x88\xd4\x02\xcd\x7b\x8e\xf6\xf5\xcd\xe6\x67\xda\xad\x16\x4a\x45\xd4\x81\x37\xde\x55\xc4\x37\x22\x77\x5a\x81\xaf\x88\x5e\x6c\xd3\xb2\x4c\x33\x98\x54\xb5\x38\xc0\xcd\x9f\xb6\x17\x19\xc9\xbb\xfb\x74\xb6\x4b\xbd\xa7\x56\x19\x0f\x06\xf0\xc1\x6d\x19\xc8\xc2\x80\x49\x11\xbc\xa4\x15\x33\xa9\xef\x17\x1a\x66\xc8\xe5\xca\xd1\x7c\x9f\xf2\xeb\x7d\xc8\xc8\x23\x02\x01\x5a\x64\xd9\xc6\xb1\x81\x96\x60\x52\xe2\xe5\x78\x2f\xb4\x51\x45\x62\x0a\x85\x40\x25\x6a\xd1\x31\xe0\x03\xdb\x32\x84\xcd\xc1\xd9\xe8\x83\x02\x3f\xfd\x04\x77\xc4\xa4\x31\x99\xe9\xca\x72\x0f\x6e\x5b\xf1\x6b\xee\xc9\xee\x57\xcd\x2e\x88\xa2\x9b\xe3\x20\xaf\x71\x02\xc3\x36\x66\xdb\x5a\x51\x68\x0a\x25\x1c\x63\x1d\xbe\xfd\xb1\xa4\xfd\xce\x3e\xa0\x53\xb9\xfa\x27\x12\x73\x47\xf2\x26\xdc\x1f\x1c\xba\xc2\x60\x4d\x72\x2f\xfb\x1c\xae\xb2\xaf\x3c\xbe\x1a\x3f\xbb\x09\x7f\x29\xd7\xe3\xc4\xa7\x6c\xbb\x71\xd9\x00\xbc\x68\x47\xc8\xa6\x2e\x25\x86\xc0\x04\xbe\xd4\x38\xda\xfb\xf3\xfb\x18\x1a\x5a\xfa\x2d\xcc\xfa\x80\x59\xbb\x6d\xd4\x6d\xcc\xe6\x80\xd9\x1c\xc3\xd8\xa3\x61\x0c\x9d\x72\xa4\xe8\xb4\x01\x89\xe4\x52\xe9\x84\x70\x0b\xfb\x55\x2a\xc3\x89\xa0\x0d\xdc\xc1\xcb\xaf\xf5\x7e\xc3\xc9\x46\xba\x5e\x53\x77\xd2\x30\xe3\xc4\xb9\xb2\x5f\x9b\xf2\x50\xd7\x0d\xa9\x6b\xb2\x66\x7a\x1c\x88\x8f\x66\xd4\x72\x1b\x19\xb2\x97\x14\x46\x66\x44\x2d\x98\x18\x83\x51\x05\x36\x2c\xad\x7f\x6e\x8e\xa9\xf8\x2e\x29\x73\x29\x4c\x48\xc8\x9c\x64\x8c\x6f\xc6\xd0\xb9\x93\x82\x24\xb2\x0f\x99\x14\x52\xe7\x24\xc1\x80\xdd\x9a\xfd\x8e\x63\x18\x5d\x9c\xd4\x54\x99\xf4\xea\x04\x6b\xaf\x97\x50\xb7\xb6\xd1\xff\xf7\x43\xf4\x89\xf6\xbf\xcf\x85\x4a\x5d\xfd\xca\xa5\xe1\x9b\x58\xe0\xca\xbe\x75\xf7\x62\xfa\x2e\xa5\xfb\xe5\x3e\xf7\xe1\xc9\x56\xd0\x3d\x0a\xfa\x59\xbe\xe5\xb2\xa0\xde\x9c\x5d\xed\x6f\x1b\xe5\xd6\xee\x01\xcd\x62\xf3\xfd\x5a\x48\x8a\x3a\x74\x2c\x3f\xcb\x9d\xea\x84\xd4\x69\x74\xa1\xbd\x80\x98\xe4\x39\x0a\xfa\xd6\x9e\xbe\xdd\xbd\xa8\x44\x21\x31\xf8\x19\xd7\xe6\xa3\xa4\xd8\x8d\x84\xf4\x7d\xbb\xd6\x73\xe0\x54\xd5\xbb\x09\x9d\x5a\xee\xc3\xf1\xe3\xdc\xa9\x1f\x3f\x76\xa9\x17\xf2\xcf\xd6\x5e\xab\x7c\xec\xcf\x0d\x0b\x4c\x18\x54\x82\xf0\x07\xcb\x1f\x8d\x21\x7a\x89\x97\xf4\xfc\xfa\x2a\x6a\xa7\x97\xc3\x1b\xcc\x58\x0d\xfe\xfe\xcd\xe5\xeb\x8b\x51\x08\xbe\x24\x8a\x91\x19\xc7\x87\x42\x93\x05\x56\x38\xde\xbd\xbd\x7e\x15\xe4\xd0\x89\xcc\xf7\x40\x88\x5e\x5e\x5f\xe1\x25\xb9\x68\x20\xb7\xed\xf3\xdc\x39\x09\x93\xd2\xd9\x2f\x56\x80\x0b\xd1\x83\x6d\x4d\x5f\xeb\x78\x9b\x14\x0e\x17\x3a\xd6\x1c\x67\x4d\xda\xcd\x91\x2e\xb5\x13\xe5\x38\x38\x2e\x91\xc3\x2d\x0c\x43\x32\x29\x72\x34\x08\x07\xe0\x29\x91\xe5\xe1\x67\xc1\x95\xc3\x6c\x77\xca\xc0\x3e\x1d\xde\x3f\x6b\x1a\x09\xe6\xfb\xe9\x7c\xf0\xc3\xe3\x18\xa2\x19\x27\xc9\x63\x68\x93\x2a\x53\xe4\x18\x22\x4a\xd4\xe3\x42\x91\x4d\x08\xc9\xad\x40\x8e\x6b\x96\x10\xfe\x50\x68\xb7\xfb\x1f\x8a\x24\xd5\x8c\x84\xe0\x89\xcc\xf2\xc2\x20\x7d\x98\x2b\x99\x59\x2c\x67\x19\x2e\x14\xa2\x08\xa1\x6d\xd7\x78\x98\x73\xb9\xb2\x48\xe9\xef\x68\xc7\x52\xca\x42\x14\xd2\x10\xdd\x92\x56\x29\x33\xa8\x33\xf9\xd8\x94\xb0\xad\xc4\x7d\x17\xc0\x5c\x6a\x66\xd8\x12\xc3\x31\x1c\xd9\x04\x9f\xcf\xe7\xf3\xd7\x49\x40\xdb\xb9\xa3\x5e\xd3\xe1\xab\x50\x7d\xfd\xcd\x51\xcf\xaf\xf1\x3c\xe4\xc9\x85\xab\xcd\xab\xcb\xd7\xa3\xd7\x01\xea\x2b\x4b\xa5\x57\xa3\xeb\x51\x53\x6f\xc8\x07\x81\x0b\xf2\x7f\xf9\x70\x3d\xc4\xd9\x75\x28\x9e\xce\x87\xe1\x30\x49\x66\x21\xde\x0b\x4f\x25\x97\x49\x28\x01\x9c\x0f\xe7\xc9\xd5\x6c\x76\xf9\x0c\x1f\x1a\x35\x7f\x18\x46\x03\xf5\x5e\x9d\x6f\xdb\xfe\x36\x06\xe3\x69\xb8\x8c\x77\x7a\x67\x45\xf2\x88\x76\x1e\x71\xd3\x73\x82\x8c\x57\x99\x07\x30\x8c\xcf\x7b\xe1\x51\x78\x67\xf1\x3e\x83\xbe\x78\x59\x5f\x03\x53\x31\x20\xd7\xf8\x7d\x56\x04\xa7\xf9\xe7\xd8\xb3\xcf\x86\x13\xf6\x84\xae\x35\x9f\x96\xa8\x14\xa3\xfe\x42\xe3\xaf\x2c\x30\x97\xee\x48\xd5\xee\x1f\x00\x06\x56\x8c\x9a\xb4\xc5\x5b\xbb\x13\x8c\xda\xda\x1c\xdd\x0d\x78\x30\xf1\x1f\x4c\xcc\xe5\x11\x5c\x2a\x97\xa8\x7e\xb3\x7a\x6a\x57\x49\xa7\xb9\x07\x4f\x65\x43\x75\x9f\xf0\x57\x18\xc5\xe7\x37\xdb\x6f\xf5\xf4\xa3\xc7\x83\x53\xb8\x0b\xdb\x53\xe4\xde\xa2\xb1\x0f\x64\x1f\xa2\x94\x2d\x52\x6e\xa7\xab\xea\x9a\x35\x2f\x1a\xbb\xcf\x93\x8a\x9d\x6c\xa2\x94\x5c\xd9\xaa\x8c\x8c\x6c\x5c\xa8\x1c\xdd\x0e\xd1\x4c\x2c\x9c\xf6\x8c\xd9\x06\x3c\xec\x43\x94\x91\x75\x34\x86\xd1\x70\x1b\x9a\x28\x1a\x17\xa9\xda\x59\xe2\x7a\x41\x39\x58\x08\x5c\xc1\x92\xe9\xf8\x1d\x31\xe4\x1e\x4d\xb7\x32\x76\xf4\xea\x13\xe0\xee\x26\x1c\xe6\x70\x27\x53\xef\x07\x66\xc6\x63\x43\x56\xe5\xd6\x53\xdf\x11\x67\xf9\xd8\x3f\xfa\xad\x58\xe9\xb1\x7f\xd4\x06\xd9\x83\xeb\x29\x43\x45\x54\x92\xda\xb3\xa9\x25\x19\x85\x9d\x5b\x68\x68\x6e\x77\x27\xf8\x3d\xe6\x44\x11\x9b\x68\x63\x18\xbd\x6a\xcc\xd1\xd6\x9e\xfb\x9c\x24\x4c\x2c\x02\x54\xca\x14\x26\x9e\xb3\xf3\xef\x77\x8d\x21\x5e\x4b\x65\xee\xd0\xa4\x92\x8e\xa1\xe3\x91\x58\xbd\x28\x35\x46\x71\x99\x5b\x41\xed\x26\xee\x47\xe9\xd0\x75\x42\x11\x41\x65\x76\x8f\xd6\xb7\x51\xbb\x03\xb3\x2c\x57\x72\x89\xf4\x97\x52\x40\xdb\xfd\x66\xe8\xc6\xb5\xaf\x93\x17\x0f\x37\x69\x12\xef\x7b\xdb\x32\x57\x27\x5e\xe3\x29\x21\x79\xba\xd1\x2c\xd1\x63\x98\x13\xae\x8f\x5f\x52\x04\x9a\x95\x54\x8f\x95\x24\xfd\xe8\x57\xda\xf7\x8e\x32\x88\x95\x84\x2b\x99\x63\x29\xba\x11\x55\x64\xb1\x60\x62\x11\xf5\x2b\xbd\xc5\x6e\x7e\xd6\xba\x4d\xd8\xce\xe1\x29\xe5\x78\xce\x51\x2c\x6c\x57\x9a\xc0\x28\xd4\x4c\x9c\x9d\x8c\xda\x03\xa1\xc2\xf5\x65\x18\x68\xbf\x3b\x8b\x32\xb9\x44\x77\x8d\x10\x8c\xf6\x77\x6c\xb9\x74\x91\x8d\x13\x22\x96\x44\xc7\xeb\x63\x84\x4d\xef\x58\xfb\x39\x5c\xae\x6e\x07\xfe\x1f\xbb\xd3\x17\x2f\xfe\x17\x00\x00\xff\xff\xcb\x03\xdb\x4b\x39\x19\x00\x00")

func templatesGraphwithextrasHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesGraphwithextrasHtml,
		"templates/graphwithextras.html",
	)
}

func templatesGraphwithextrasHtml() (*asset, error) {
	bytes, err := templatesGraphwithextrasHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/graphwithextras.html", size: 6457, mode: os.FileMode(420), modTime: time.Unix(1569565519, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHeadHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x94\xcf\x6f\xe2\x38\x1c\xc5\xcf\xf0\x57\x58\x99\x63\x8b\x9d\x4c\x43\x28\x2c\x54\x2a\x74\x29\x2d\x43\xd9\xf2\x63\x0a\xdc\x9c\xc4\x49\x9c\x38\xb1\xb1\x1d\x20\x8c\xfa\xbf\xaf\x02\x62\xa6\x3b\x4b\x47\xea\xee\x61\x72\x89\xf5\x14\xbd\xf7\x79\xfe\xda\x69\x33\x9a\x25\x40\x12\xd6\x31\x94\x2e\x18\x51\x11\x21\xda\x00\x91\x24\x41\xc7\x88\xb4\x16\xaa\x85\x90\xd2\xd8\x4b\x04\xd6\x11\x74\x39\xd7\x4a\x4b\x2c\x3c\x3f\x83\x1e\x4f\xd1\x77\x01\xd9\xf0\x33\xb4\x90\xa7\xd4\x0f\x0d\xa6\x34\x83\x9e\x52\x06\xa0\x99\x26\xa1\xa4\xba\xe8\x18\x2a\xc2\x57\xd7\x76\xed\xfe\x71\xbf\x5a\xf7\xef\xb7\xae\xf5\x3c\x9b\x3d\x39\xdb\xa2\xde\x0c\x82\xbe\xd5\xcd\xef\x1f\xc5\x97\x29\x6e\xde\x25\xc3\x91\x30\xef\x42\x3a\xba\x4b\x6d\xba\x1c\xc5\x0d\x33\x5c\xbd\x0c\x97\xee\x43\xc3\x74\xf4\xcb\xd4\x00\x9e\xe4\x4a\x71\x49\x43\x9a\x75\x0c\x9c\xf1\xac\x48\x79\xae\x8c\x9b\x6a\x5b\x79\x92\x0a\x0d\x94\xf4\x7e\x74\xf0\xb8\x4f\x60\xbc\xce\x89\x2c\x0e\xe0\xc7\x65\xed\x0a\x5e\x41\x0b\x2a\x46\xd3\x03\x6c\x7c\x96\x75\x7d\x4d\xd1\xe2\xa2\xe9\xd4\xef\xf6\x63\x53\xce\x1a\xd8\x1d\xda\xd6\xe3\x54\x3f\x3f\xdc\xae\xbf\x86\x93\xaf\x7b\xe1\xee\x79\x5d\xa5\x8b\xa1\xb0\x97\xc1\x64\x33\xb8\xb8\xc6\xae\x9e\xfd\x69\xfd\x45\x9d\x98\xee\xf9\xfb\xac\x6d\x74\x64\x7d\x0f\xda\xcf\x62\x05\x3d\xc6\x73\x3f\x60\x58\x92\x03\x39\x8e\xf1\x0e\x31\xea\x2a\x24\xb8\x10\x44\xc2\x58\x21\x0b\x5a\x36\x74\x50\x9e\xfa\x27\xf1\xfd\x36\xdb\xc1\x2d\xed\x07\x13\x36\xea\x17\x0e\xad\x4f\x27\x78\xb7\x09\xc6\x3d\x1a\x74\xe7\xcf\x85\xb5\x1b\xf8\x8f\xa8\xe0\xb4\xd1\x9f\x3c\x2d\x46\x93\x6e\x5e\x7f\x19\xf8\xab\x65\x6e\x45\xb7\xce\x6a\xec\xb2\x30\xd7\xff\xb9\xcd\x07\x8e\x51\xfc\xf3\x29\x3a\x5f\xa5\x6b\xce\x43\x56\x4c\x2e\xe2\x27\xa7\x97\x6c\x36\x0f\xbd\x71\xf7\x73\xcc\x71\x50\x7f\xb0\xd9\x55\x98\x36\xef\xe7\xce\xc0\xb3\x78\xe8\x7c\x51\x0d\xea\xcc\x51\x9a\x24\xd8\xcf\x87\xb8\x1b\xb1\xdb\xc5\xa6\x99\xfc\x9f\xc1\x40\xc1\xb8\x86\xac\x40\xe5\x9b\x15\x35\x86\x35\x51\xfa\x04\x7b\xc6\x40\x17\x82\x74\x0c\x4d\x76\x1a\xc5\x78\x83\x8f\xaa\xf1\xa1\x81\x6f\xa8\x2a\x37\xc8\x82\x66\xb9\x3c\x97\x75\xfe\x3a\xbf\x89\x3e\x5c\xc7\x7f\x5e\xef\x8f\x87\x96\x26\x37\x6d\x54\x86\x95\xfd\xca\xa4\x9f\x33\x6e\xaa\x00\x00\xf0\x49\x48\x1e\x4a\x9c\x86\x12\x8b\x08\x7c\x3b\x68\xe5\x13\x11\x1a\x46\xba\x05\x2c\xd3\x34\xc5\xee\x8f\xef\xba\x8b\xbd\x24\x94\x3c\xcf\xfc\x9a\xc7\x19\x97\x2d\xb0\x8d\xa8\x26\x2a\xe5\x09\x39\x7e\xf5\x5a\xad\x7c\x8a\x08\xd6\x29\x16\xe0\x5b\xb5\x52\x39\x39\xd9\x47\xa3\x4a\xe5\xd7\x16\x95\xd7\x23\x57\xf9\x3b\x78\xc3\x13\xf0\x4c\xd7\x02\x9c\x52\x56\xb4\xc0\x88\x64\x8c\x5f\x82\x11\xcf\xb0\xc7\x2f\x41\x8f\x67\x8a\x33\xac\x2e\x81\xd1\xe3\xb9\xa4\x44\x82\x27\xb2\x35\x2e\x41\xca\x33\xae\x04\xf6\xc8\x2f\xf8\xff\x8d\x7f\x88\xc7\x4a\xd7\xca\xad\xfa\x6d\x04\x50\xf2\xed\x9b\xfe\x02\xfb\x3e\xcd\xc2\x16\xa8\x9f\xa6\xf1\x5a\x6d\xa3\xc3\x60\x6f\xaa\x7f\x07\x00\x00\xff\xff\xd2\xb7\xca\xc4\x24\x06\x00\x00")

func templatesHeadHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHeadHtml,
		"templates/head.html",
	)
}

func templatesHeadHtml() (*asset, error) {
	bytes, err := templatesHeadHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/head.html", size: 1572, mode: os.FileMode(420), modTime: time.Unix(1569389750, 0)}
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
	"templates/graph.html": templatesGraphHtml,
	"templates/graphwithextras.html": templatesGraphwithextrasHtml,
	"templates/head.html": templatesHeadHtml,
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
		"graph.html": &bintree{templatesGraphHtml, map[string]*bintree{}},
		"graphwithextras.html": &bintree{templatesGraphwithextrasHtml, map[string]*bintree{}},
		"head.html": &bintree{templatesHeadHtml, map[string]*bintree{}},
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
