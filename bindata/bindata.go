// Code generated by go-bindata.
// sources:
// template/ccd.file.tmpl
// template/client.ovpn.tmpl
// template/dh4096.pem.tmpl
// template/iptables.tmpl
// template/server.conf.tmpl
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

var _templateCcdFileTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x4c\x4b\xce\xcf\x4b\xcb\x4c\xd7\x2d\x28\x2d\xce\x50\xa8\xae\x56\xd0\xf3\x0c\x50\xa8\xad\x05\xb3\xfc\x52\x4b\x7c\x13\x8b\xb3\x15\x6a\x6b\xb9\x94\x33\x8b\xf2\x4b\x4b\x52\x15\x0c\x2d\x8d\xf4\x0c\xcd\x2c\xf4\x2c\x0d\xf4\x0c\x14\x8c\x4c\x4d\xf5\x60\xd8\x80\x0b\x10\x00\x00\xff\xff\xb7\x2b\x33\x90\x4a\x00\x00\x00")

func templateCcdFileTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateCcdFileTmpl,
		"template/ccd.file.tmpl",
	)
}

func templateCcdFileTmpl() (*asset, error) {
	bytes, err := templateCcdFileTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/ccd.file.tmpl", size: 74, mode: os.FileMode(420), modTime: time.Unix(1502796579, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateClientOvpnTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\x8e\xb1\x6a\x03\x31\x0c\x86\x77\x3d\x85\xa0\x4b\x3b\x5c\x3c\x74\x2b\x47\xa1\x74\x29\x94\xd2\x4c\x5d\x4a\x07\x9f\x4f\x49\x4c\x6c\xcb\xc8\x3a\x83\x1b\xfc\xee\xe5\x2e\xa1\x9b\xfe\x4f\x48\xdf\x7f\x87\x7a\xf2\x05\xb9\xe6\x84\x07\x1f\x08\x7d\x41\xbb\x28\x47\xab\xde\xd9\x10\x1a\x1e\x29\x91\x58\xa5\x19\xa7\x86\xdf\x9f\x5f\xfb\x8f\x9f\xfb\x93\x6a\x2e\x4f\xc6\x1c\xbd\x9e\x96\x69\xe7\x38\x1a\x67\x67\xc3\x35\xc7\x07\x00\x17\x3c\x25\x85\x99\x2a\xea\x92\x20\x0b\x2b\xe3\x32\x67\x10\x8a\xac\x84\x97\x0b\xee\xde\xb8\x68\xb2\x91\xb0\xf7\x2d\xef\x59\x14\x7b\x07\xa1\xc2\xa1\x0e\x42\x2a\x0d\x7d\x3a\xf8\xe4\x95\x6e\x87\x83\x23\xd1\x41\x43\xc1\x42\x52\x49\x20\xf1\xe4\xd3\x0c\x99\xa4\xf8\xa2\xc3\x99\xda\xff\xbc\x8a\x1d\xc7\x3c\x84\x5f\x86\x4a\x32\xe1\x23\xc0\xe8\xec\x33\xac\xb6\xd7\x17\xec\x7d\x34\x6b\x1c\xd7\xa7\x37\x4a\x5b\x87\xd1\x5c\xd1\x78\xa6\x76\x5d\xbc\x53\xdb\xf8\x06\xfe\x02\x00\x00\xff\xff\x96\xf9\x3e\xe7\x32\x01\x00\x00")

func templateClientOvpnTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateClientOvpnTmpl,
		"template/client.ovpn.tmpl",
	)
}

func templateClientOvpnTmpl() (*asset, error) {
	bytes, err := templateClientOvpnTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/client.ovpn.tmpl", size: 306, mode: os.FileMode(420), modTime: time.Unix(1502796579, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateDh4096PemTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xd4\x35\x12\xa4\x08\x00\x40\xd1\x9c\x53\x4c\x4e\x6d\xe1\x16\xe2\xee\xde\x19\xd0\x38\x8d\xfb\xe9\xb7\x76\xe2\xfd\xe1\x3f\xc0\xfb\xe7\xbf\x38\x51\x56\xad\x3f\x82\xf2\xc7\x61\x3d\xd6\x14\x03\xd1\xf3\xff\x7e\xc0\x54\x55\x51\x70\x75\x9e\xad\x45\xb6\x2f\xb3\xc9\xd2\xe7\x37\x56\x43\x1c\xa4\x1b\x2d\xac\xa4\x52\xd1\xb4\xbe\xa5\x86\xf7\xd6\x99\x21\xd2\xde\x29\x40\x6c\xec\xf4\x04\x64\x3e\x11\x13\xee\x7e\x38\xf0\xc5\x28\x74\x5d\xa8\x62\x55\x98\x61\x1a\xea\x16\x8e\x9e\xaf\x9b\xd5\x15\x48\x63\x88\xf6\xd1\xa8\x1f\x03\x39\x69\x95\x7a\x63\x88\xe6\x15\x81\x75\xb8\x4a\xdb\x34\x5f\x4b\x1c\xdf\xc5\x06\x57\x17\xc0\xf7\x75\x64\x4a\xc5\x2d\x99\x2d\x90\x54\x89\x30\x0f\x94\x8f\x86\xce\xa5\xda\x3e\x8b\xaa\x8d\x2f\x7c\x14\x4d\xab\xeb\x97\x43\x29\x3e\xd4\xd9\x9e\x62\xf5\x10\x0d\x6a\xcb\xbf\x7c\x5f\xfa\xb5\x77\x01\x52\x8d\x90\x43\x92\xb3\xfa\x3a\xdd\x57\xdd\x33\x6c\xbb\x15\x03\x5d\xe9\xe4\x5e\xb0\x56\x9e\x92\x31\x26\xb3\x57\x2e\x0a\xf2\x73\x82\xf0\x0c\xd7\x9f\xd8\xef\xa7\xfc\x40\xe5\xf0\x9b\xa9\x9b\x42\x02\xdd\xf1\x83\xa0\xe0\x73\x4e\xc7\xd5\x0b\x62\x43\xc2\xc7\x2e\x11\x97\x73\x4d\x8b\xf9\x94\x4c\xd6\x84\x76\x79\xbf\x20\xb5\xab\xd4\x51\x22\xd0\xb0\x96\x60\x63\x64\x61\xec\xb6\xf5\x2b\x96\x04\x99\x03\x95\x47\xc4\xe1\xc7\x6b\x2e\xab\x66\xde\x60\xac\x67\x29\x30\x92\x96\x55\x31\x46\xa3\x6f\x25\x0f\xdb\x2f\xd6\x9f\x10\xef\x0b\x0e\xbf\x90\x49\x13\x9a\xf1\xc3\x3f\x06\xb7\xf9\xfe\xb7\x40\xd4\x6e\x01\x66\x8f\xa4\x08\x8e\x5b\x8c\xcb\x9f\xf0\xc9\xb6\xa0\x09\x09\x6a\x33\x15\x30\x08\xbb\xfd\xac\x4f\xfc\xd0\x38\x7d\x5b\x48\x07\xe3\x3b\x2a\x63\x6b\xd0\x7a\x03\x72\x20\x5a\xd6\xec\xa7\xcc\x9e\x30\x07\xe6\x25\x22\x07\x29\x4c\xa4\x09\x49\xba\x7e\x88\xdd\xad\x58\xd0\xfa\x3a\xf4\xf3\xa2\x12\xfe\x0b\x3b\x63\x9a\x7d\xe1\xf5\x36\xa2\x1c\xca\x8c\x27\xd5\xdf\x13\x2f\x56\x83\xf0\xbc\x0d\xba\xed\xa2\x02\x6c\x54\xc0\xa0\x09\x1f\xbe\xd8\x18\xbb\x44\xcb\x11\xad\x82\x53\xbc\xcc\x0e\x63\xe6\x41\x70\x27\x10\xb5\x3c\xda\x12\xfd\xab\x2d\xf9\x26\x71\xc7\xe4\xc0\x98\x70\xcb\xfe\xfd\x40\x59\xf6\xdb\x18\x1d\x30\xe1\x62\x4f\x14\xea\x8c\xbe\xee\x59\x80\x2b\x4d\xea\xc2\x0e\xa9\x6e\x32\xcd\xc7\x34\x25\x88\x8e\xf7\xbf\x66\x09\xbb\x3e\xd6\xd6\x6e\x87\xc3\x33\xf2\x66\x74\xeb\xd5\xd4\xb7\xc4\x72\xe8\x58\x18\xb0\x8c\x92\x87\x7a\x38\xe5\xe0\xd5\x8a\x3f\x37\x16\x7a\x35\x96\xf8\xd0\x09\x4f\x9a\x0d\x3a\x3b\x64\x13\x21\xda\x1e\xd3\xac\x7c\x66\xf4\x88\x1b\xc6\x28\x5d\x70\x69\x37\xbe\xae\x55\x56\xdb\x7f\x27\x80\x2b\x14\x4a\xd9\x50\x26\x14\xfb\x83\x14\xb4\x8e\xa1\xb3\x2e\x3f\x3f\xd0\x90\xb9\xfb\x66\xc6\xe3\x30\x83\x6e\xce\x08\x0b\x8a\x2f\xe8\xda\x95\xc2\x8d\xf5\x68\x42\x94\x6f\x3f\xca\x13\x67\x27\x05\xc0\x64\xa4\x4a\x11\x14\x5e\x69\xd2\x71\x89\xaf\x5b\xe2\x39\x8e\x47\x69\x0f\xee\x20\x33\x1b\xf7\xc6\x05\xd0\x8a\x91\x6e\x8b\x94\xd2\xb3\xa6\xb5\xd9\x4c\xad\x51\x77\x55\xad\x21\x7b\xb6\x26\x22\x03\x28\x1a\xd9\x19\xda\x0f\x12\x54\x7d\xba\xaf\x83\x48\x3c\x25\xbc\x90\x54\xa3\x74\xa1\x5b\x2e\x44\xfd\xc8\xa5\x29\xc6\x4d\x78\x6f\xcf\xe9\x3c\xf2\xd9\x3e\x7c\xab\x38\xf9\x6f\x3d\x2d\xed\xd5\x46\x07\xa0\xda\x1a\x19\xcd\x3c\x57\x5a\x88\x1d\x40\x87\xd1\x91\x6a\x5a\x3d\x96\x00\x23\xbb\x22\xaa\x33\x70\xf6\x85\xfc\xdd\x93\x4b\xd0\x44\xc1\xaf\x46\x8a\x89\x8b\x5b\x82\x95\x33\x84\x6f\xc1\x19\x50\x0e\x54\x66\x34\x5e\x0c\x9c\x53\x9b\x1b\xa9\x8c\x28\x26\xf7\x3c\x13\xe8\xa1\xa6\xcc\xbd\x07\xef\x56\x65\x3a\x26\x9f\xd8\xfc\x0b\x4a\x49\x2c\x8b\xe7\x9b\xa1\x90\x7d\xd1\x93\x9c\xbe\x53\x3e\x12\x3e\x0b\xd8\xcc\x2d\xce\x05\xfa\xb3\xc2\xc1\xb6\x02\x7c\x6f\x38\xa7\x81\x30\x2b\x75\xe4\xcf\xa8\xb6\x42\x6a\xd1\xa7\xd2\xe9\xad\x3f\xa6\x45\xea\x6d\x79\x3c\xdf\x67\x20\x7e\x3c\x4b\xec\x55\xb2\x21\xbe\x11\x40\x63\xe2\xb9\x59\xa6\x3d\xc9\xf4\xf7\xd0\xdc\x66\x38\xdc\xbb\xcf\x60\xbb\x91\x6a\x38\x0c\xe4\xa6\x2c\x10\x24\xd9\xe0\x36\xd8\x5b\x9a\x3f\x47\x7b\xd6\x97\xa7\x46\xcf\x64\xb0\x47\x58\x65\x70\x07\x50\xa3\x82\xbd\x64\x78\x9e\x8d\x91\x76\x1e\x6c\xe3\xd6\xd4\x26\x14\x89\x98\xfd\x7d\x9a\x45\x27\x28\xca\x67\xe0\xfd\xbe\xb7\x03\x0c\x19\x22\x5d\x64\x0f\x3c\x30\xb9\x55\xfc\x34\xcc\x8d\x8e\x6a\x04\xb0\x8a\x7c\xbb\x51\xd0\x97\xcb\x4d\x1f\x54\xcf\x28\x35\x8c\xee\x09\xa3\xfc\xc8\x8b\xce\x24\x89\x8a\x79\xc7\xd0\x34\xa3\x04\xcd\x58\x19\xec\x7d\xc6\x79\xb2\x9c\x21\x47\x1a\x0b\x52\x10\xfd\x89\x01\xea\x98\x67\x1d\x43\x35\xa9\xbd\xc8\x99\xd6\xf7\x63\xee\x77\xd4\x7e\xbf\x44\xd3\xc0\xd6\x46\x9e\x40\x01\x43\xf8\xf7\x78\x8e\x78\x71\x15\x4d\xbd\xd7\xdb\x64\x47\xe8\x60\x85\x3b\xcd\x70\x74\x5f\x01\xe9\x6e\xca\xfc\x55\xbd\x3a\x07\xc7\x1b\xe6\xc7\x18\x92\x3d\x07\xd3\xbd\x42\xfd\xa8\x2b\xd8\x6e\xdd\xd2\x23\x19\x65\xfe\x45\x17\xf8\xeb\xb0\x68\x09\xff\xa7\xf3\xbf\x01\x00\x00\xff\xff\x2c\x8e\x5a\x0b\xbc\x05\x00\x00")

func templateDh4096PemTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateDh4096PemTmpl,
		"template/dh4096.pem.tmpl",
	)
}

func templateDh4096PemTmpl() (*asset, error) {
	bytes, err := templateDh4096PemTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/dh4096.pem.tmpl", size: 1468, mode: os.FileMode(420), modTime: time.Unix(1502796579, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateIptablesTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func templateIptablesTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateIptablesTmpl,
		"template/iptables.tmpl",
	)
}

func templateIptablesTmpl() (*asset, error) {
	bytes, err := templateIptablesTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/iptables.tmpl", size: 0, mode: os.FileMode(420), modTime: time.Unix(1502796579, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateServerConfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x7a\x5d\x73\xdb\xb8\x92\xf6\x3d\x7f\x45\x97\x74\x11\xbb\x4a\xa2\x65\x25\x99\x93\x8c\xdf\x77\xb7\x34\xb2\x32\x51\x4d\xa2\x68\x63\x79\x3e\xaa\xe6\x06\x22\x5b\x22\x4a\x24\xc0\x01\x40\x29\x3a\xa7\xce\x7f\xdf\xea\x06\x40\x52\xb6\x33\xbb\x3b\x53\xae\x48\x64\x03\xe8\xef\x7e\xba\xa1\xbb\x5a\x1b\x07\xb7\xb7\xef\xdf\x24\xfc\xe9\x5f\xff\x82\x74\x4d\x1f\xfe\xfd\xef\x24\x19\xc2\x66\xbe\x06\x6d\xe0\xf1\x7e\x0d\x16\xcd\x11\xcd\x7f\x26\x77\xb5\xd1\x4e\x83\xcb\xea\xc4\x7f\x6a\xf2\x9a\x48\x07\x39\x1e\xc1\x35\x6a\x00\x27\x59\x96\x90\x19\x14\x0e\x41\x80\xd1\x8d\xc3\x1c\x96\x6b\x7a\xa9\xb0\x1c\xb5\xb4\xa2\x7e\x42\xab\x00\x5d\x81\x46\xa1\x0b\xb4\x69\x32\x84\x47\x8b\x2d\xfd\x64\x00\x72\x07\x67\xdd\x80\x30\xd8\x11\x6f\x8d\xcc\xf7\x52\xed\x93\x21\x08\x95\x43\x21\x8e\x08\xb5\x41\xbf\x6d\x0e\x82\x97\xc2\x51\x1a\xd7\x88\x12\xa4\x72\x68\x76\x22\xc3\x40\xce\xab\x31\x07\xe9\xe0\x24\x5d\x41\xdb\x9b\x6e\xef\x96\x9a\x78\x59\xfa\xc3\x4f\x42\x39\x70\x1a\x32\xad\x9c\xd1\x25\x88\x2c\x43\x6b\xa1\xd6\xa5\xcc\x24\xda\x64\x08\xfa\x88\x06\x5c\x81\xf0\xeb\x7a\x35\xe2\x35\x55\x63\x5d\x14\x74\x27\x0d\x9e\x44\x59\x26\x43\x30\x4d\x89\x16\x76\xda\x53\xd3\xdf\xe6\x71\x75\xb3\x99\xad\x2f\x0f\xfe\xa2\x40\x69\x35\xfe\x4d\xaa\x5c\x9f\x2c\xd8\xb3\x75\x58\x59\xbf\x73\x26\x14\xec\xe5\xd1\x8b\x03\xf8\xad\x26\x2e\x1c\x34\x4a\x3a\x50\x4d\xb5\x45\x33\x02\xdb\x64\x05\x08\x4b\x6a\x9d\x84\xfd\xc2\x5e\x23\x68\x82\x82\xc7\x4a\xe7\x38\x08\xbc\x48\x1b\xc8\x2a\x6d\x5d\x77\x5e\x10\xc9\x9b\x4d\x69\x07\xbb\x46\x65\x4e\x6a\x95\x0c\xa1\x51\x25\x69\x81\x58\xaa\x85\x71\x52\x94\xe5\x99\x9c\x67\xd7\xd0\x87\x5c\x5a\xb1\x2d\x89\x49\xda\x23\x6a\xa0\x95\xfc\x05\xa9\xef\x82\xd1\x93\xe0\x58\xe4\x64\x51\x01\x0a\x31\xb7\x7e\xe1\x6c\x4d\x6a\x79\x3d\x05\x91\x8b\xda\xa1\x01\x25\x2a\x3a\x66\x67\x74\xc5\x14\x2b\x74\x27\x6d\x0e\x30\xd7\x4a\x21\x33\x6b\xa1\x16\x0a\xcb\xe0\x4b\xc9\xd0\x7b\x4c\xa5\x0d\x59\x40\x28\xd0\x0a\x53\x20\xe1\x7f\x5f\xc3\xc3\x7a\x4a\x42\x14\x72\x5f\xa0\x21\xdf\x65\x63\x8a\x33\xb3\x40\x4e\x60\xb1\xa4\x5d\x8f\xd8\x09\x49\xc7\xf6\x98\x7d\x2e\xec\x6c\x1d\xb9\x25\x2d\xaf\x9e\x5b\x16\x1a\xdb\xb0\xfe\x72\xad\x5e\xb9\x70\x16\x1b\xe5\x2e\x5a\x0a\x3e\x9f\x37\x82\x23\xef\xe1\xe1\xd3\xcd\xe6\xd3\x03\x18\xad\x1d\x64\x68\x9c\xdc\xc9\x8c\xfc\xec\x2a\x13\xd7\xa3\xfe\x93\x64\x08\x57\xf4\xf5\x7a\xc4\x9e\x5f\x1b\x79\x24\xba\x03\x9e\xe1\xea\x80\xe7\xeb\x14\x60\x21\xb2\x02\xb2\x52\xa2\x72\x21\x3e\x88\x63\x1f\xfc\xde\x89\x59\x59\xae\x40\x69\x40\x9f\x14\x6f\x4f\x74\xc9\x90\xf7\xd9\xc9\x92\x94\xb7\xe9\x16\xd1\x1e\x24\xbd\xdf\xd4\xb2\xef\x90\xbb\x58\xef\xef\x56\x54\x08\x99\xf0\x0b\x93\x21\xc9\x83\xfe\xcd\x00\x85\x3d\x8f\x8d\x15\x03\xc8\xa5\xc1\xcc\x69\x73\x66\x1d\x0a\xda\x3a\x04\xda\x0e\x6c\x66\x64\xed\x7c\x10\xed\x51\xa1\x11\x4e\xaa\x3d\x7c\x7d\x98\xf5\x45\xb7\x41\x9a\x9e\xcc\x36\x05\xf8\x8a\x15\x52\x88\x90\x25\x1b\xcb\x31\x44\x91\xf3\x57\x83\x30\xd7\x55\xa5\x15\xac\x88\xbf\x68\x39\x2f\x52\xd8\x0a\x49\x55\x7a\xc7\x2f\xbc\x70\x17\x07\x7a\x61\x66\xea\x0c\xbf\xbf\x9d\xbc\x67\xe5\x54\x42\x89\x3d\x56\x44\xe9\x0d\xcd\xc1\xbb\x45\x3a\x3a\xe7\x80\xab\x51\x51\x7c\xd1\x63\x51\x5a\xe6\x09\x04\xac\x7f\x99\x3f\xc0\xf0\x76\x4a\x7c\x54\xc2\x51\x52\x8b\xba\x26\x93\x5a\x44\x18\xd4\x87\xcc\xde\x4e\xa3\xaa\xe4\x11\x41\x2a\x3a\x10\x6a\xb1\xc7\xeb\x34\xb9\xcb\x04\x44\x85\xde\x90\xf4\x37\x99\x48\x33\xe3\x92\x3b\xb6\xe0\xe5\x2b\x2f\xa7\x7f\x4d\x27\xbd\xf8\x96\x5e\xc0\x10\x36\x85\xb4\xcc\x0a\xd8\x42\x37\x65\x4e\xf2\x1c\xb0\x76\x60\x29\x03\xbb\x24\xc9\x04\x97\x95\xf9\x6c\x8e\xc6\xad\x85\x2b\xa8\xb8\xf0\xa1\xfc\xb8\xf7\x90\x76\xa4\x67\xbf\xe0\x39\x3e\x4a\x86\x70\x2f\x77\x3b\x89\x50\x60\x59\x7a\x79\x8c\xa8\xd0\xa1\xe1\x14\xf5\xb3\xb7\x38\xfa\xb4\x4d\x0e\x49\x49\xfc\xc7\x64\x08\x00\xba\x46\x65\x6d\x09\x79\xc1\x6b\x60\xac\x1b\x07\x79\x71\x3b\x99\xbe\x49\x6b\xac\x80\x3e\x90\xbf\x35\x5b\xeb\xa4\x6b\x1c\xc2\x74\xf2\xe6\x1d\x1b\x9b\x5e\xf5\x6b\x4d\x63\x7d\x7d\x61\x82\xad\x74\xde\x7f\x92\x61\x5e\xf4\x36\x4c\xee\xf2\xe2\x89\xaa\xf2\x82\x56\xf0\xcb\xbc\x60\xd9\xee\x3f\xae\x89\x19\xdb\x13\x30\x66\x28\xa7\x6b\x5d\xea\xfd\x99\x78\x6a\x55\x69\x9b\x2d\xd5\xa1\x2b\x91\xe7\x06\x2d\xb1\x01\x47\x29\x60\xb9\xbe\xee\x72\x6e\xcc\x1e\x31\xc4\x8e\xd3\x74\x92\xbe\x67\x27\x2d\xf5\x09\x4d\x88\x59\x9d\x0c\xfd\x8e\x35\xd5\x7a\xcc\xe1\xca\x15\xa8\x40\xa1\x7b\x3d\x19\x81\x4c\x31\x05\x01\x37\xaf\x27\x50\xa3\x09\x7b\xd1\x21\xf7\xb8\x13\x4d\xe9\x2c\xc5\x08\xd3\xc2\x15\xa5\x7e\x83\x99\xae\x2a\x54\x39\xe6\xd7\x49\x64\x3d\xb0\x4b\x52\xcd\xb5\xda\xc9\x7d\x63\xba\xec\x41\x59\x8b\x78\x22\x06\xca\x33\x08\x2e\x25\x61\xc1\x90\xd5\x1e\xfd\xdf\x69\xc8\x8d\x38\xc5\xb0\x0a\xb2\x53\xa5\x34\xba\x22\xb3\xf7\xf2\x0b\xd7\x22\x27\x0e\x08\xb7\x93\xf4\x5d\x3a\x49\x6f\x79\x2b\xe9\x2c\x96\xbb\x51\x28\x38\x06\xad\xf3\x94\x5b\x84\x4a\x10\x23\x47\x21\x4b\x9f\xad\x75\x54\x1c\xed\xdc\x4b\x80\xed\x82\x48\x66\x38\xe4\x2f\x12\x81\x56\xed\xb1\x29\xa7\x0c\x5a\x47\xa9\x1a\x4a\xa9\x10\xc8\xe3\x3a\x37\x4a\x86\xcf\x41\x4b\xda\xe6\xbb\x18\xac\xcc\x3e\x57\x23\xa9\x76\x3a\x4d\xee\x82\xa0\xe1\x9c\x09\x4c\xdf\xbe\x4d\xe3\xdf\x24\x09\x6f\xc9\xb7\x56\x48\xc0\x8d\x3f\x7e\x16\xf6\x10\xdc\xeb\xb3\x90\xca\x09\xa9\x08\x8c\x61\xa6\x4d\x4e\x49\x2b\x48\xf8\xff\xc6\xff\xd1\x02\xa3\xe5\x3a\x2a\x9a\x12\x9c\xb5\x3a\x93\xc2\x97\x4b\xa9\xbc\x48\x21\xb9\x2f\x77\xad\xa1\xf6\x1a\x2d\xe4\x14\x76\x9a\xb4\x21\x2d\x6b\x5a\x90\x77\x8d\xf8\x34\x5f\x73\xd5\xbe\xf5\xcd\x90\xf1\x84\xb5\x72\xaf\x30\x0f\xf6\xe1\x2a\xf0\x9c\x91\xae\x8a\xd7\x5a\x97\x54\x9d\x1d\x9c\x04\xf1\x57\x1b\x3c\x4a\xdd\x58\xf2\xa3\xb0\x55\x9a\xc8\x5d\xc6\x4e\x37\x26\xea\x71\x8d\xc6\x4a\xeb\x40\xd6\x75\xea\xbe\xfd\x8d\x4f\x92\xbe\x9f\xdb\x25\x19\xc2\x1f\x11\xb8\xed\xa4\xb1\x8e\x93\x31\xa7\x99\x2f\x0f\xaf\x6c\x4b\x08\x99\xa8\xc5\x56\x96\xd2\x51\xe0\x3a\x1d\x00\x65\x5b\xe9\x5b\x48\xe3\xe1\x25\x3d\x8e\x87\x51\xec\x2f\xe7\x3d\xd0\xc3\x85\x53\x75\x80\xb1\x12\xca\xe3\x00\x4b\x98\x98\x51\xc5\x72\x7d\xa3\xd0\x55\x64\x5e\xad\x78\xb7\x70\x5e\xbb\xcb\x08\x0a\x34\x08\x27\xf4\x66\x6c\xaa\x36\x32\xde\xdc\x5c\xb8\x4e\x0a\xf0\x41\x2a\xde\x9f\x89\xf9\x48\x3a\x49\x58\xc9\xa1\x4a\x96\x30\x42\xf1\xe6\xde\x03\xda\x68\xbd\x62\x2b\xff\xff\xb0\xf1\xdb\x09\xa0\xca\xe3\xb7\xdb\xc9\xe4\x9a\x14\x21\xca\x52\x07\xd8\xe1\xd1\xf2\x13\x5f\x48\x01\x3e\xa1\x47\x13\x31\x5e\x7c\x4a\x71\xec\x17\x14\x3b\x3d\x50\xf9\x22\xe4\x6f\xa3\x63\x1c\xb4\x10\x25\xbd\x0c\x12\xe8\xf8\xec\x78\xfc\x3f\x3a\x04\x63\x16\xe2\x5e\xc0\xfd\xc7\xf9\x7a\x5c\x1b\xfd\xed\x3c\x82\x13\x2b\x3b\x7a\xb7\x13\xe5\xc1\x8b\x4b\x96\x89\x51\x12\x58\x64\xad\xd2\xda\x2e\x7d\x70\x52\xc9\x50\xb6\x90\xaa\xe7\xfb\x41\x7d\x1e\x5d\x53\xde\xbc\x5f\x3d\xb4\xa0\x2a\x66\xc4\x14\x5e\xf0\x52\x0f\x53\xbf\xef\xa7\xcf\xbd\x94\x62\xf7\xfb\x7e\xfa\xc4\x4b\x19\xb3\x3a\xfc\xd1\xdb\x8d\x35\xa6\x15\x39\x91\x36\x07\x4b\x5e\x19\xb5\x71\x15\xfa\x8e\x0e\x0c\x5f\x47\x85\x75\xa0\xc9\xab\xa5\x87\x89\x41\xd2\x82\xad\x6e\x14\x03\x6c\xaf\xee\x40\xfc\xd4\xde\x64\xc3\x75\x63\x0b\xdf\x66\xda\xa8\xf8\x90\xdd\x82\x0f\x9e\x40\xba\xa8\x6b\xc6\x6c\x24\x56\x8b\x04\xbd\x4b\x5b\xd8\x62\x21\x55\x9b\x8f\x3c\xc4\xe9\x03\x44\xca\x3d\xae\x40\xd6\xee\xd3\xb5\x5c\x28\x18\xb0\x11\x54\xf7\x67\x1d\x94\x3e\xf1\x99\xc4\xd9\x85\x3b\x74\xf0\x3a\x58\x9a\x93\xdb\x55\xcc\xef\x97\x41\x4a\x15\x78\x2b\xb2\xc3\xcb\x3e\x95\x26\x43\xbf\xff\xed\xfb\x69\x7a\xfb\xc3\xbb\xf4\xfd\xb3\x02\x11\x3d\x7e\xfa\x8c\xf4\xf6\x7b\xa4\x6f\x9f\x91\x4e\xbf\x47\xfa\xfa\x19\xe9\xeb\xef\x91\xbe\x49\x86\x35\x99\x6a\x10\x16\xfc\x83\x16\x50\x20\x3e\x5d\x30\x78\x42\x18\x76\xfe\xe1\x25\x42\xd8\xe8\x50\x00\xc0\xd6\x98\x11\xf8\xee\xc5\x90\xf7\x88\xf8\x22\x19\xb6\x9e\x49\x00\x61\x07\xe2\x79\x52\x82\x42\x58\x10\xd1\xc0\xc9\x30\xe2\x2f\xef\x1d\x20\x9d\x77\x84\x00\x75\xd9\xe4\x0c\xae\xc8\x24\x7e\x0e\x30\xea\x37\x37\xcd\xb6\x6b\x5d\x06\x59\x96\xfb\xfe\x3a\x3a\x7e\x8f\xaf\x90\x87\x38\xdc\xb9\xce\x5a\x8f\xeb\x5f\x06\x05\xd7\x29\x39\xfe\xe2\xf7\xd9\xe7\xf5\xa7\xc5\x8f\xf0\x40\x88\xce\xf6\x83\xca\x37\xb5\x24\x16\x3f\xeb\x35\x85\x99\x6f\x6b\xa8\x4b\x86\xc1\xa6\xc0\x52\x2b\xaa\xa3\xa4\xca\x20\x0d\xc9\x6f\x2b\x6a\xd9\x2e\x65\xa7\x60\xef\xf4\x45\xf5\x42\x64\x85\x54\xd8\x8d\x17\xa2\xa5\xde\x4c\xd2\xdb\xe9\xbb\x0b\x3f\x9e\xbe\x79\x47\x79\xe3\x03\xa5\xa8\x11\x34\x2a\x24\x7a\x4e\xf2\x1c\x55\x9c\xff\xed\x8f\xc9\x30\x28\x27\x54\xf2\x5c\x1a\xc8\xb2\xfc\x7f\x76\x73\x0f\x0a\x55\x37\x78\xe2\x96\x24\xcb\xf2\x9b\x4e\xc8\x98\xd9\x42\xb5\xf1\x5d\x82\xbc\xdc\xd9\xf3\x0e\x4f\x78\x4f\x42\x9f\x13\x02\x9d\x72\x4a\xb7\xed\xab\x27\xf9\xc0\x23\xed\x30\x15\x0a\x23\x13\x2e\xed\xd2\x02\x7e\x13\x55\x5d\xa2\xdf\xa8\xcd\x9a\x94\x7e\xbb\x66\x83\x18\x92\x6a\x3f\xe2\x09\x4b\x4c\xdf\x01\xa0\x77\x48\xd2\x17\xa3\x6e\xec\x46\x15\x62\xe0\xb3\x42\xaf\x0b\xb4\x2f\x7b\x4a\x7f\x8c\x15\x86\x47\x3d\x35\x91\xf2\xbe\x61\xce\x4e\xdd\x2b\x48\x7a\x47\xa1\xfc\x9e\xb1\x6e\xb4\xe5\xdf\x99\xf2\xee\x65\x53\x3e\x7f\xca\x7d\xe0\xfc\x3e\xb6\x44\xc1\x92\x22\xcf\x7b\xc8\x80\x10\xc4\x85\x31\x83\xf5\x5a\xc8\x47\x19\x23\xb2\x17\x3f\x4c\x93\x24\x33\xe5\xf8\x88\x46\xee\x7c\x6b\x39\xff\xfa\xa9\x3b\xa6\x8b\x1b\xe1\x2e\x54\x82\x8a\x21\x7f\x2e\x77\x3b\x34\x3e\x9c\xda\x29\xce\x93\x69\x1f\x07\x66\x4b\x08\x7b\xa3\x9b\x3a\x4c\x26\x3a\x98\xb3\xe1\x9a\x47\xb6\x75\x27\x0d\x15\xba\x42\xe7\x2c\xc0\xd5\xed\x35\x7c\x6d\x14\x54\x4d\xe9\x24\xf9\x45\x4c\xef\xb9\xc0\x4a\x2b\x3b\x02\xad\x02\x1e\x11\x59\xc1\x12\x83\x3f\xc3\x8f\x70\x5a\xae\x5e\x1c\xa2\x05\xfa\xb8\xdc\x2f\xbc\xf1\x5b\x83\xa8\x6b\xa3\x6b\x23\x85\xc3\xf2\x4c\xf6\xbc\x9a\x5e\xc3\xd5\x2c\x3f\x0a\x95\x61\x7e\x0d\xf3\x18\x48\x7e\xc0\xc2\x3d\xd9\x59\x89\x4a\x66\x84\x18\xc3\xd6\x95\xce\x49\xb3\x17\x33\x3d\xa9\xa8\x09\xa8\xb5\xb2\x6c\x35\xaf\xaf\xc8\x0a\xe1\xf9\x4e\x5b\x9d\x86\x1e\x7c\xae\x0b\x64\xcf\x33\x1e\xc1\x8a\x12\x85\x51\xe3\xe8\x8d\x9e\xad\x34\xb9\xbb\x7c\x9c\xde\xf8\x17\x89\x9f\xd7\x7a\x4b\xe6\x23\xef\x49\xdd\x70\xc4\x0f\x9e\x23\xfa\xe3\xe4\xd7\x8d\xa8\x18\x29\x78\xda\x00\xcb\x72\xdf\x03\x27\x43\xea\x80\xb9\x59\xdf\x0b\x87\x27\x41\xa2\x1b\xdd\xec\x8b\x6e\xda\x9b\x89\x38\x2c\xa0\x2d\x97\x6b\x70\x46\xec\xa8\x2c\xc5\x2c\x79\xc2\x2d\x6c\x8d\x3e\x79\x30\xc9\x90\x23\xa2\xbb\x52\xeb\x43\x53\x33\x03\x7b\xfd\x74\x6b\x32\xd1\xe6\x19\x02\x88\x59\xf8\x62\x28\xb9\x9a\x6d\xc8\x05\xcd\x05\xca\x7b\xea\x1d\x11\x51\xf0\x03\x8f\xea\x25\xf5\x71\x39\x9a\x76\x0a\x4c\x34\x2c\x2e\xf9\x0a\x9a\xf2\x7c\x9d\x26\x77\xa1\x38\x07\x0d\x8d\xa3\x26\x72\xdc\xdd\xc2\xf6\x5c\x0b\x6b\xc7\x79\x91\xd5\x83\xe4\x7f\x4d\x48\x48\x1c\x0d\x37\xa9\x01\x28\xb6\xc5\xb1\x55\xb8\x45\x47\x59\x91\x3c\x29\xf4\x90\xb4\xbd\x97\x37\x18\xae\xab\x44\xf7\xab\x07\xaf\x80\xdf\x96\x2f\x83\xe6\xf9\xec\xd7\xc5\x6c\x43\x21\x58\x38\x57\xff\x78\x73\xa3\x6b\x54\xc7\x5a\xa5\x0a\xdd\xcd\x4e\xfc\x95\x16\xae\x2a\x87\xc4\x5d\x26\x8e\x28\x9c\x0d\x43\x87\x0e\x58\x6c\x91\xaa\x80\xc1\x9d\x9f\x1e\x72\x93\xda\x6c\x4b\x2e\xe7\x1d\x52\xb7\xa4\xb9\xa3\xcc\x31\x87\xed\x99\x27\x52\xb9\xb2\x69\xa6\xab\x56\x8f\x74\xc6\x58\xd7\x5c\xf8\x69\xdd\x74\xf2\x2e\xfd\xe1\x1f\xe9\x74\x3a\xa5\xbf\xa8\xc4\xa7\x54\xef\x52\xfe\x9f\x55\xf7\xd8\xa6\xe0\x27\x4e\xde\xe2\xdf\x7e\x22\xeb\x39\x79\x6f\xb4\x31\xb0\x88\x03\xe8\xe0\x31\x25\x84\x9f\xce\xd1\xef\x47\x17\xd3\x5b\x5f\xb7\x6c\x18\x5b\xb4\x40\x94\xa0\xd8\x4e\x9b\x0c\xfb\x47\xbc\x40\x3a\x0a\x53\xf7\x4b\xd0\xcc\xcc\xf6\x53\xd2\x45\x6e\x23\x20\xc6\x8b\x5f\xd9\x97\xee\x0a\x42\x49\x71\x7a\x1c\xe0\xcf\xdf\x69\x45\xee\xba\x74\x1b\x38\x25\x30\x23\xf7\x85\x8b\xf8\xa6\x6b\x81\xfc\x78\xba\xc3\x4f\x37\x07\x3c\x73\x35\x20\x84\x46\x50\xae\x83\x53\x36\xd6\x78\x9e\x82\xb4\xc3\x31\x1e\x12\x95\x7e\x72\xed\xd0\x32\xd6\xac\x1b\x43\x85\x87\x56\x7c\xd0\xd4\x8a\xe8\xbc\xe1\xab\x09\x42\x8e\x04\x20\xb1\x37\x85\x0a\x58\x93\x61\xa6\x24\xec\x1a\xa6\xee\x97\x1c\xd5\x42\x1a\x3f\x72\x5e\x7e\x80\x3f\xbe\x3c\xc2\xc7\xd9\xaf\x0b\x58\x7d\xd9\xc0\xcf\x8b\xd5\xe2\xeb\x6c\xb3\xb8\x87\xe5\xea\x7e\xf9\xeb\xf2\xfe\x71\xf6\x89\x02\x6e\xf1\x75\xb3\xfc\xb0\x9c\xcf\x36\x8b\x9b\x5f\x16\x7f\xc0\x7a\xb6\xfc\xfa\x00\x1f\xbe\x7c\x85\xc5\x6c\xfe\x11\xe6\x9f\x96\x8b\xd5\x86\x78\xe1\xaf\x1f\x67\xbf\x2e\x57\x3f\xc3\x72\xf3\x00\x5f\x7e\x5b\xc1\xe3\x6a\xf9\x5f\x8f\x0b\x18\xcc\xbf\x7c\xfe\xfc\x65\x05\xab\xd9\xe7\xc5\x80\x68\x1f\x57\xf4\x64\xb1\xda\xc0\xe6\xe3\xf2\x01\x3e\x2d\x57\x0b\xf8\xf2\xb8\x49\x93\xbb\xbc\xa9\x4b\xe6\x77\x9c\xf1\x55\x0e\x05\xd3\x01\xb1\x16\x25\xd9\xa4\xb3\x0e\xa5\x4f\xb4\x50\x4b\xb5\x1f\x97\xf2\xc0\x73\x09\xb4\x56\xec\x31\xba\xac\x25\xa5\x70\x4b\xc4\x15\x50\x1b\x57\xf0\x85\x5b\x68\xdd\x4a\xa9\x0e\x60\xb5\x2f\xea\xac\x47\x6e\x30\xa9\x1d\xb3\xd4\x7c\xaa\x40\xe7\xbb\x40\x7e\x47\x98\x77\x4f\x85\x36\xd7\x27\x45\x9e\xbc\x26\x23\xe1\x11\xcd\x19\x6e\x27\x60\x31\xd3\x2a\xb7\xa3\x38\x4d\xe1\x9d\x0d\x56\x9a\x1b\x84\x1a\xb9\x65\xf5\x53\x30\xb9\x03\xa5\x99\xf9\xd8\xd7\xe7\x90\x37\x26\xd4\x03\xb8\x9d\xc6\xdd\xc0\xc9\x0a\xa1\x46\x23\x75\x9e\x26\x9d\x22\x6e\x27\x44\x44\x0a\x22\xc7\xc0\x6f\xce\x08\x5a\xd1\x18\x6a\xdb\xb7\x78\xe6\xa5\x74\x7e\x4c\x2b\xd4\x1e\x9e\xe3\x1d\xd0\xa8\x77\x97\x3a\xf8\xf8\x79\x36\x6f\x83\x68\xe0\x9b\xd2\x02\xcb\x1a\xb6\xa5\xce\x0e\x70\xaf\x1f\x40\x38\x27\xb2\x83\x65\x45\x3e\xde\xaf\x81\xef\x80\x77\xa5\xd6\xb9\x9f\x80\xf5\x07\xec\x97\x73\xf5\x63\xad\x60\x3c\xde\xa3\x3a\xe0\x19\xc6\x63\x3f\xeb\x07\x27\x52\xf6\xc6\xcb\x09\x6d\x7b\x59\x12\x1c\xba\xbd\x46\x62\x9d\x64\xba\x3e\xfb\x6b\x14\x69\xe1\x80\xe7\x6e\xbc\xcb\x7a\x6a\xc7\xfd\xbd\xeb\x85\x57\x93\x57\x7e\xf4\xea\x2e\x0f\x79\x75\xfb\x2a\x3e\x6d\x27\xba\x77\xae\xb4\x63\xd1\x50\x38\x33\x73\x30\xb9\xb8\xb3\x90\xb6\xbd\xa7\x18\xc2\x03\xdf\xe5\x11\x4f\xe6\x5c\x3b\xbd\x37\xa2\x2e\x64\x06\x99\xac\x43\x46\xdc\x84\x1e\x68\x27\xf7\x20\x1d\x56\x5e\x94\x2d\x75\x55\xb5\xe4\x1c\x16\x9c\x2b\xde\x07\x79\x52\x3e\x89\x2b\x7f\x59\xa6\xc9\x9d\xdf\x0f\x7e\xfa\x30\x9e\xff\x34\x87\xf0\xdf\x10\x7e\x2a\xf5\x69\x27\x6d\x01\x57\x21\xeb\x5e\xb7\xa4\xb3\xc5\xc3\xf8\x76\xfa\x2e\xd0\x0f\xe9\x7b\xfb\xee\x7e\xf1\x30\x5e\xdc\x2f\x5e\xfb\x97\x43\xd8\x18\xca\x6c\xe3\xfb\xc5\x03\x23\x7f\x8f\x64\x33\x5d\xd5\x7c\x7d\xa0\x55\xd4\x10\x41\x08\x8a\x95\xde\xc5\x76\x80\xbd\xd2\xf1\x04\xb1\x77\x71\x4d\x59\x9a\xb2\x52\xfb\x5e\xaa\xef\x88\x99\x26\x74\xd4\xb8\xfc\xa7\x8e\x41\x5e\x89\x6f\xb2\x6a\xaa\x70\x17\xcd\xb8\x58\xab\xac\x31\x54\x8e\xca\x73\xcc\xb8\xec\xc8\x6d\x91\xc1\x16\x89\x73\xfd\x4a\x93\xbb\x4a\x7c\x1b\xc7\xd7\x61\x76\xb7\x74\xaf\xa8\x55\xd9\x6b\x9d\x83\xcc\x51\x04\xd8\xd6\x64\x17\xd3\x96\x64\x18\xb0\xf4\x2b\xcb\x6d\x9a\x2c\x91\x52\x89\xd8\xf1\xa4\x49\x49\x27\x45\x29\xff\xc9\x7d\xb7\x77\xf8\x3f\xc2\x9d\x7a\x73\x59\x3f\xa8\xbb\xe1\x49\xdc\x0b\x17\xf1\x69\xd2\x58\x34\xa0\xf4\x56\xe7\xe7\x84\x51\x76\xfc\x12\x94\x10\x67\xd1\xbe\x82\x87\x32\xea\x0c\x4f\xe2\xc4\x51\xcb\xbc\x6d\x18\x79\x14\x11\xb0\x90\x41\xab\x1b\x93\x21\x0f\xd5\xc2\x60\x9d\x1d\x4c\x38\x0f\xf7\x34\x94\x5a\xed\xd1\x70\x25\xf7\xcb\xc9\x40\x5b\xe4\x2c\xea\xbb\x10\x06\x26\x51\x70\xce\x52\x7b\x23\x72\x4c\x93\xc0\xd3\x98\x22\x36\x7e\x0e\xf7\xec\x5f\x1a\x57\x37\x14\x08\xb6\xa0\x8c\x60\x9d\x70\x4d\x77\xc9\x77\xf2\xe9\x2c\xd8\xb0\x9d\x09\x70\xbb\xe2\x4c\xa3\x28\xcb\x47\x54\x6b\xf0\x64\xa4\x73\xa8\x42\x3a\xad\xa4\x6a\x1c\xa6\x49\xd8\x32\xa4\x92\xb1\xff\x9a\x96\x7a\x9f\x3c\x01\x1e\xa5\xde\x77\x05\x80\xd5\x46\xc8\xd8\xe3\x2d\x7b\xb6\xf4\xfa\x4a\x87\x9b\x98\xf6\x37\x0d\x72\x07\xa6\x51\x8a\x01\xb6\xf5\xf7\xc5\x47\x99\x21\xff\x80\xe1\xdc\xdb\x25\x44\xeb\xe0\xcf\xb5\xa1\x68\xaf\xe0\x03\xd5\xf5\x3f\x83\xeb\xfc\x59\xea\x7d\xef\xee\xf9\x3a\xfe\x1c\x85\xce\xd4\x86\xfe\x19\x8b\x9a\x10\x1d\x03\x9d\x23\x1a\x43\xc5\xc4\xa3\x0d\xcf\x3e\xad\x18\xf0\x2e\xc1\xe0\x5e\x39\xbe\x4e\xe9\x90\x18\xb4\xea\x20\x3d\x59\xb8\xa9\xa9\x92\x9e\x0a\x7a\x37\xe8\x0e\x09\x7b\x74\x27\x4a\x97\x02\xf3\xa3\x15\x7a\xb0\xdb\x95\xb5\xab\x6d\xe3\xfc\xf8\x40\xbb\x82\x30\x3b\x9d\x16\xff\x8b\x50\x97\xb4\x7d\xd7\x93\xe2\xe2\x05\xa7\x43\x6e\x7f\xfa\xb0\x0c\x4a\x3c\x62\x49\x6e\x45\x34\x1e\x07\xc1\x11\xcd\x56\x5b\xe9\xce\x3e\x80\x26\x9c\x54\x65\x89\xca\x8d\x00\xbf\x65\x58\x3b\x06\x3f\x3b\xe1\x44\x09\x68\x8c\x36\x04\xb8\xde\x78\x9c\x24\xac\xf6\x59\xa5\xbb\xbf\x2f\xa1\x21\x73\x27\x43\x78\xcb\x3e\xf4\x03\xc7\x23\x97\x2e\x6a\x42\x71\xdb\xec\x7b\x4e\x47\x95\x70\x5b\x62\x45\x7b\xbe\x07\x1e\xb1\x38\x83\x15\x21\x48\xcf\x18\x26\xf4\x2f\xbc\x66\x91\x88\xad\x0c\xc1\x60\x8d\xfe\x77\x02\xd1\xb7\x52\x80\x99\xf3\x3f\x77\x99\x4e\x18\x6e\xfe\xd5\xa0\xa2\xec\xd0\xb9\x5f\x88\x26\xc6\x84\xe1\x21\x77\x23\x0e\xf7\xda\x9c\xdb\xeb\x42\xed\xa3\x27\x38\x69\xa9\xf7\x94\xbf\xfc\xf5\x72\xf2\xdf\x01\x00\x00\xff\xff\x7c\x09\x5e\xca\x71\x25\x00\x00")

func templateServerConfTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateServerConfTmpl,
		"template/server.conf.tmpl",
	)
}

func templateServerConfTmpl() (*asset, error) {
	bytes, err := templateServerConfTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/server.conf.tmpl", size: 9585, mode: os.FileMode(420), modTime: time.Unix(1502796579, 0)}
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
	"template/ccd.file.tmpl": templateCcdFileTmpl,
	"template/client.ovpn.tmpl": templateClientOvpnTmpl,
	"template/dh4096.pem.tmpl": templateDh4096PemTmpl,
	"template/iptables.tmpl": templateIptablesTmpl,
	"template/server.conf.tmpl": templateServerConfTmpl,
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
	"template": &bintree{nil, map[string]*bintree{
		"ccd.file.tmpl": &bintree{templateCcdFileTmpl, map[string]*bintree{}},
		"client.ovpn.tmpl": &bintree{templateClientOvpnTmpl, map[string]*bintree{}},
		"dh4096.pem.tmpl": &bintree{templateDh4096PemTmpl, map[string]*bintree{}},
		"iptables.tmpl": &bintree{templateIptablesTmpl, map[string]*bintree{}},
		"server.conf.tmpl": &bintree{templateServerConfTmpl, map[string]*bintree{}},
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

