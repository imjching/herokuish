package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _include_buildpack_bash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x56\x5f\x4f\xeb\xb8\x13\x7d\x4e\x3e\xc5\xfc\x42\xc5\x05\xa4\xd0\x1f\xf7\xee\x13\x97\xa2\x05\xca\xdd\xad\x16\x28\x0b\xed\x13\x8b\x2a\xd7\x99\x34\x16\x6e\x9c\xb5\x1d\x0a\x62\xf9\xee\x3b\x4e\xd3\xfc\xa1\xf4\xc2\xc3\x4a\xa0\x26\x1e\xcf\xcc\x39\xe3\x33\xe3\xf8\xd3\x5c\xc8\x28\x63\xfc\x21\x2c\x9e\x76\x76\xe1\xc5\xf7\x22\xe4\x92\x69\x84\x08\x0d\xef\x05\xa7\xce\x00\x2c\x05\x96\x65\x52\x70\x66\x85\x4a\x21\x37\x22\x9d\x81\x48\x8d\x65\x52\x62\x04\x55\x1c\x13\xf8\x1e\xa6\x26\xd7\x18\x66\xcc\x26\xc6\xf7\xea\x14\x06\x6d\x9e\xc1\x31\x74\x23\x7c\xec\xa6\xb9\x94\x4d\x23\x3e\x21\xcf\x2d\xc2\x3f\x14\x35\xc2\xd4\xfa\x5e\xa6\x15\x8f\x85\xc4\xd0\x3e\x67\x68\x6a\xc3\xab\xdf\x80\x5d\x42\x78\x0f\xf8\x60\x69\xaa\xb1\x41\xac\xd5\x1c\x7e\x13\x16\xc6\x37\x17\xc4\x28\x02\x95\x39\x36\x4c\x02\x57\xf3\xb9\xb0\x56\x98\x24\xa8\xc3\xe4\x5a\xf6\x82\xce\x41\x50\x5a\xe9\xf9\x6b\x00\x29\x9b\x23\x3d\x7d\x5b\xe3\x29\x62\xb8\xbb\x83\xff\x41\xd0\x21\xbf\x00\xee\xef\xbf\x83\x4d\x30\xf5\x3d\x8f\x19\x62\x1e\x52\xe5\x88\x01\x97\x79\x84\xdd\xba\x5c\xfb\xf6\xc9\x12\xb5\x45\x42\x44\x41\x23\x8b\x5c\xd6\x32\xe1\x77\x88\x14\xb9\x7b\x6b\x6c\x57\x39\x82\xce\x72\x23\x61\xf1\x22\x95\x22\xfd\x68\xaa\xb1\xa6\xa4\xb1\xf0\x3d\xa9\x38\x51\xb3\x4c\xcf\xd0\x4e\x1c\x4a\xc2\x5d\xc5\x2a\x16\xba\x9d\x17\xc7\xe7\x30\xec\xec\x4c\x99\x41\xf7\x0c\x2e\xf4\xee\x6b\xb0\x22\x54\xe5\x68\x32\x9a\x51\x0d\xb9\xa4\x8c\x35\x92\x46\x1a\x07\x87\x47\xeb\x6b\x85\x57\x82\xfc\x41\xe5\x16\xc2\xf0\xef\x5c\xa0\x6d\x71\x20\xa7\xb0\xad\x0f\x94\x06\x5b\xf9\xc2\x30\xc2\x8c\x98\x1c\x6c\xca\xec\x88\xeb\x39\x84\x3a\x6e\x9b\xba\xfb\x33\x97\xa4\x25\x1e\x29\x8c\x7d\x4f\x39\x17\xb4\xfe\x39\x71\x4b\x03\xa1\x83\xd2\x2e\xeb\x9b\x34\x85\xee\x97\x79\xb6\xe0\xb4\x52\x23\x3e\x65\xc8\x6d\xd1\x4f\x14\x89\xde\x94\xb6\x70\x72\x7d\x3d\xe9\x0f\x6e\xe8\xa4\xa8\xdd\x56\x9c\x4a\xdb\xef\xc3\xcb\xf3\x77\x0d\x37\xe7\x7f\x8e\xcf\x6f\x47\x93\x41\xbf\x17\x14\x69\xc3\xce\xcd\xc9\x55\x7f\x78\x59\x6f\xb9\x1d\x9d\x9c\xfd\xd1\x0b\x38\x46\x4c\x87\x07\xbf\x90\x81\x67\x54\x24\xa8\xc3\x75\xf7\x83\x15\x91\x55\x78\x87\xb7\xaf\x55\x96\x51\x11\x32\x2d\x1e\x49\xa3\x33\x24\xac\xb9\x41\x3d\x57\x74\x58\x61\xa2\x9c\x62\x1c\x30\x48\xd5\x54\x45\xcf\xcd\xe3\x83\xaf\xc7\xdb\x07\x94\x28\x51\x8b\x14\xc2\x9b\x72\xc7\x61\xaa\x66\x5a\xd1\x20\xf8\x8b\x0e\xb6\xc1\xa6\x7c\x6f\x00\x28\x57\x38\x23\xd1\xac\xef\x69\x54\xdb\xe1\x1c\x1b\x8c\x73\x09\x54\x6b\x4b\xa3\xc9\x40\x17\x62\x64\xd4\x0b\x58\xd7\xf6\x8c\xda\x7e\x72\x36\xbc\xba\x3a\x3f\x1b\x4d\x46\x83\xcb\xf3\xe1\x78\xd4\x0b\xbe\xfd\x7f\x19\xa0\x38\x18\x63\x31\x83\x29\x45\x5e\x30\x1d\x19\xd7\x87\x94\x41\x4c\x85\x14\xf6\x79\xd5\x11\x61\xdc\xaa\x1a\xa6\x8f\xad\xde\x30\x2a\xd7\x1c\xd7\xb6\x14\xca\x6c\xc9\xa2\x9c\x78\x4b\x61\x94\xa1\x53\xf2\x3b\x1d\x0f\x2e\xfa\xd7\x74\x5c\x13\xc2\xdb\x0a\x6d\x85\xa5\x29\x11\xfc\x40\xcb\x13\x37\x7f\x79\x6e\x2c\x4d\xb4\x2a\xa4\xeb\x21\x87\x00\x25\x09\x0b\xa3\x0d\x3d\xbf\xf4\x72\x7b\xab\x3e\x69\x79\xb8\x72\x78\x83\x1f\xb7\xbd\x2f\x5b\x5f\xde\x8e\x24\x38\x3a\x3a\x5a\x83\xe8\x7f\x66\x44\xad\xc0\x6e\x37\xfb\xbb\x09\xb6\x1c\xac\x3b\x79\x5a\x49\x2d\x82\x36\xb2\xee\x54\xa4\xe4\x6c\x69\x05\x1a\x42\xd9\x0d\xaa\x41\xb1\x05\xb1\x72\xd5\x4f\x50\xab\x87\x3c\xac\x61\xcd\x73\x69\x05\x58\x05\xa5\x7b\x2c\xb4\xeb\xf1\x98\xba\x90\x9a\x9d\x34\xe2\x4e\x80\xba\x79\xad\x5a\xef\x47\x0a\xd6\x74\x5e\x1d\xd2\x47\x8c\x3e\x15\x7f\x23\xd3\xa2\x03\x3c\x6f\x7b\x1b\x3e\x38\xe5\x0d\xb8\xc9\xdb\x4d\xc8\xc6\x6d\xd5\x3e\xfb\xa6\xdc\xca\x0b\xa4\x1e\x7f\xbd\x9d\xb7\x59\xf6\x76\xdd\x3e\xaa\x79\xe3\x96\x15\x4e\xc4\x2f\xb5\xd7\xdd\xaf\xf7\xaf\xc1\xea\x2e\xfb\x74\x75\x3e\xaa\xc0\x4f\x4b\xd0\xdc\x34\x25\x09\x3f\xb8\xb7\xf2\x86\x74\xf4\xdd\xff\xea\x7a\xdb\xcc\xbf\x6c\xb7\x4e\x0b\xb2\xfb\x0c\x2a\x45\x84\x51\x2d\xbc\x72\xef\x38\x65\x53\xfa\x25\x9d\x2d\x9d\x80\xb5\xbb\x93\xd4\x66\xe1\xa0\x00\xe0\x66\xce\x68\xd8\x1f\x1e\x82\xc5\xa5\x14\x6d\x22\x0c\xd0\x5f\x8a\x1c\x8d\x61\xfa\x19\xa8\xae\xb4\xbe\x40\xe0\xee\xfb\x4b\x2e\xd8\xb3\x81\x8c\x3e\x27\xa0\x43\x23\xa5\x00\x0c\x0b\x61\x13\x77\x9f\x0a\x63\x72\xac\x59\xed\xb8\x9b\xe9\xa4\xde\xb7\xdb\x62\xd6\xaa\x79\xf0\x4e\x93\xb9\xc1\x47\xe6\x37\x37\x02\xb4\x87\x71\x50\x45\xaf\xeb\xf0\x9f\x44\x5e\x1e\xd0\x87\xa1\x34\xad\xd0\x57\xcb\xcf\x41\x1e\xb7\xac\xdd\xfd\x95\x13\xd5\xdf\x24\xf4\xfd\x07\xa1\x21\x69\xda\x99\x54\x53\x70\xcd\xec\x1e\xaa\x2f\x88\x7a\x86\xef\xf9\xde\xfc\xb1\x29\xc4\xee\x5e\x6d\xad\x42\xe5\xeb\xa1\x5e\xfd\x7f\x03\x00\x00\xff\xff\x10\x52\xac\x0c\x65\x0b\x00\x00")

func include_buildpack_bash_bytes() ([]byte, error) {
	return bindata_read(
		_include_buildpack_bash,
		"include/buildpack.bash",
	)
}

func include_buildpack_bash() (*asset, error) {
	bytes, err := include_buildpack_bash_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "include/buildpack.bash", size: 2917, mode: os.FileMode(420), modTime: time.Unix(1418844091, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _include_buildpacks_txt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x90\x31\x92\x85\x20\x10\x44\x73\x4f\xe1\x05\x5c\x57\x70\xd5\xda\xdb\x0c\xc3\x28\x2a\x0a\x85\x40\x95\xb7\xff\x89\xf5\x25\x9d\x49\x88\xde\xeb\x6e\x4c\x8c\xfe\xfa\x6f\xdb\x65\x8d\x26\xa9\x1f\x74\x47\x6b\x28\xb8\x3d\x3d\x4f\xa3\xd2\x6a\xb5\x07\xdc\x1b\xb4\x6e\x4b\x81\xea\xe7\x14\x0a\x35\xeb\xa9\x62\x18\x16\x57\x97\xa7\xc4\xd0\x01\x28\x96\x21\x80\xb6\xdf\x0a\xf5\xd8\xcb\x79\x94\xc8\x31\x6c\x90\xa1\xe8\x90\x05\x6b\xc1\x91\x6c\x5c\x5f\x1a\xb5\x26\x94\x3d\xc7\x70\x3a\x4d\xdb\xf5\xe6\x0f\xbf\x1c\xda\x1b\x5f\xfe\x60\xee\x25\x8b\xb6\x70\x17\x34\x12\x69\x9a\x06\x96\xe1\x8e\xc6\x9d\x6f\xfe\x9f\xe0\xd0\x21\xa9\x32\x3f\x77\x62\xe4\xe0\x17\x82\x85\x72\x7c\x57\x7d\x02\x00\x00\xff\xff\x67\xc4\x1a\xde\xbf\x02\x00\x00")

func include_buildpacks_txt_bytes() ([]byte, error) {
	return bindata_read(
		_include_buildpacks_txt,
		"include/buildpacks.txt",
	)
}

func include_buildpacks_txt() (*asset, error) {
	bytes, err := include_buildpacks_txt_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "include/buildpacks.txt", size: 703, mode: os.FileMode(420), modTime: time.Unix(1418861795, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _include_cedarish_txt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x33\xe2\x02\x04\x00\x00\xff\xff\x9f\x70\x98\xa2\x03\x00\x00\x00")

func include_cedarish_txt_bytes() ([]byte, error) {
	return bindata_read(
		_include_cedarish_txt,
		"include/cedarish.txt",
	)
}

func include_cedarish_txt() (*asset, error) {
	bytes, err := include_cedarish_txt_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "include/cedarish.txt", size: 3, mode: os.FileMode(420), modTime: time.Unix(1418874313, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _include_cmd_bash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x54\x41\x6f\x9c\x3c\x10\x3d\xc3\xaf\x98\xcf\x71\xa2\xe4\x80\xf8\x58\xf5\x44\xb4\x55\xa2\xb6\xb7\xb6\x97\x1c\xc3\x46\x72\xc0\x04\x14\xaf\x59\x61\xd8\xa6\x22\xfc\xf7\xce\x18\x9b\x40\xb2\xbd\xf4\x04\x9e\x19\xcf\x7b\xf3\xfc\xec\xb0\x90\xb9\x12\xad\x84\xe8\x16\xbe\xfc\xf8\x7a\x17\x86\xf9\xbe\x88\x54\x6d\xba\xcb\x2b\x18\xc2\xc0\xa7\x0b\x69\xf2\x2d\xfb\x8e\x71\x03\xe2\x28\x6a\x25\x1e\x95\x84\xbc\xd9\xef\x85\x2e\x0c\x7b\x2b\xd4\x66\xcb\x78\x82\x01\xdf\x27\x7a\x96\xbf\x0d\x30\xae\x0d\x83\x57\x30\xb2\x00\x66\x62\x5c\xa5\x71\xcc\xc2\xf1\x0d\xcf\xd6\xad\x41\xe7\x5e\x65\xd3\xc2\x33\xd4\x1a\xdb\x0c\xff\x11\xcd\xfb\x9b\xdd\xc8\xae\xa1\x68\xc2\x20\x90\x79\xd5\x60\xe2\x99\x48\x34\x5a\x22\xc8\x53\x2b\x0f\xc0\x1e\x08\xc4\x62\x36\x6d\xb7\x42\xd2\x0e\xe7\x5f\xda\x46\x47\x48\xdf\xf5\x94\x2f\x07\x5c\x9d\xd2\xeb\x9b\xcd\xa0\x62\x50\xf6\x3a\xef\xea\x46\x83\xa0\x95\xd3\x6d\x21\x5b\xa9\xed\xa8\x98\xc6\xef\xb0\x49\x23\x9e\x8c\x98\x56\x4d\x2e\x94\xd5\xc1\xa9\xa0\x89\x2e\xbf\x5c\x8c\x72\xb5\xe6\x5b\x6a\xb6\x54\x20\x62\x70\xf1\x19\xe2\x42\x1e\x63\xdd\x2b\x05\x17\x17\x93\xaa\xda\x8d\x15\x06\x76\x6e\x3a\x9e\x94\x0f\xc2\xc4\x67\xf8\x17\xc5\x23\xdb\x6d\x6d\xaf\xf5\x8c\xb3\x72\xef\x4e\xc8\x8d\xcb\x37\xd8\x55\x1e\x91\x31\x06\xa9\x10\x90\xd7\x4a\x91\x8c\x71\xfa\x66\x58\x67\x0d\xa2\x0d\xf0\x04\x32\x96\xf1\x9b\x0c\x85\x0f\x83\xd1\x39\x67\xc2\x83\xe9\xf8\x1d\xc5\xc4\x92\x4a\x66\x4e\x9e\xcc\x9b\x48\x98\xbc\x06\x53\xd5\x65\x07\x3e\x8c\x85\xab\xf8\xeb\x2b\x74\x6d\x2f\x7d\xda\x74\xa2\xeb\xcd\xf6\xff\x30\xa8\x4b\xf0\xaa\xce\x6e\xb5\x32\x3e\x70\x8c\x67\x7c\xa5\xe3\x35\x74\x95\xd4\x38\x04\x1f\x16\xfa\x61\x1d\xdb\x8d\xb8\xfb\x86\x74\x50\x06\x51\xa8\xed\xfd\x3d\x86\x28\x07\xbb\xdd\xbc\xd1\x9d\xd7\xcf\x06\x4c\x9f\x57\xde\x11\x29\xd8\x42\xca\x3b\x66\x1b\x3a\x5a\xe5\xbb\x10\xad\x8f\x4d\xf8\x65\xa9\x23\xd2\x75\xaa\xb8\xa2\xfd\x65\xed\x3c\x31\x5b\xe3\xf6\xc3\xb5\x4d\x6d\x21\xba\x0a\x31\xd7\xbe\x72\x8d\x9c\xb5\x82\x43\x5b\xeb\xae\x04\x06\x70\x1e\x6d\x3e\x19\x38\x37\x19\xda\xcc\x0d\xb5\x82\xff\x28\xc7\x44\xc7\xc2\x98\xfe\xf1\x14\x12\xd5\xcd\x50\x0b\xac\xf7\x68\xd3\xf6\xbf\x00\x62\x26\xf5\x15\x33\xe8\xe4\x70\xff\xf1\x6a\xbc\xd4\x1d\xf0\x49\xde\x90\x74\x72\x76\xaa\xa4\x3a\x9c\xba\xc4\x77\x55\xf3\xcb\x00\x65\x91\x39\x8e\xb1\x17\xf6\x22\xd3\x40\xa7\x6e\xb2\x68\x9f\xc8\x89\x64\x01\x7f\x6a\x14\x5a\x9c\x1b\x4d\xe6\x65\xa7\xb1\xad\x21\xe9\x07\x5f\x99\x33\x50\x02\x45\xc1\x1d\x53\xd9\xc2\xdb\x03\xb5\x89\xcf\x69\xd2\x98\x5e\xa9\x29\xa8\x31\x04\xeb\x35\xc4\x11\x5e\xa3\xc5\x76\xfb\xb8\x9c\x38\x19\xc7\x44\x47\x34\x97\x7b\x3b\x92\xd9\xb9\xee\x7e\xda\x97\x87\x44\x5a\x5c\x4b\x2f\x97\x55\x25\xfc\x13\x00\x00\xff\xff\xd3\x23\x50\x2b\x44\x06\x00\x00")

func include_cmd_bash_bytes() ([]byte, error) {
	return bindata_read(
		_include_cmd_bash,
		"include/cmd.bash",
	)
}

func include_cmd_bash() (*asset, error) {
	bytes, err := include_cmd_bash_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "include/cmd.bash", size: 1604, mode: os.FileMode(420), modTime: time.Unix(1416876705, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _include_fn_bash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x51\x41\x6e\xc2\x30\x10\x3c\xc7\xaf\x18\xad\x22\x01\xaa\xa2\x08\xae\x34\x3d\x56\xea\x1b\x28\x07\xcb\xac\x89\xd5\xd4\x89\x6c\x03\xaa\x28\x7f\xef\x1a\xd2\x92\x43\x55\x55\xb9\x64\x77\x66\x67\x76\xd6\xca\xfa\x4a\x87\x7d\x9c\x2f\x70\x56\xc5\x8e\x4d\xa7\x03\x63\xc7\xd1\x34\xf4\xe2\xe3\xc0\x26\x41\xc3\x1e\xbc\x49\xae\xf7\xb3\x08\x21\x1f\xde\xd9\xa7\x48\xaa\xe8\x7a\xa3\xbb\xdc\xe9\x9c\xe7\xa6\x9c\xa7\x8f\x81\x51\x2e\xf1\x89\x7d\xe0\x01\xdf\x6a\x63\x59\x1d\x41\x53\x03\x12\xa0\x65\xbd\x43\xb5\x5c\xa8\x82\x4d\xdb\xa3\x62\x50\x79\x1e\x05\xeb\x1a\x35\xbd\x7a\xba\x64\xa2\x3e\xbd\xa1\x7a\x6e\x30\xab\x9b\xfa\x3c\x04\xe7\x13\xe8\x91\xca\x25\x3d\xd1\x65\x26\x78\x0a\xc8\x5c\xc8\xa7\x2e\x2a\xa7\xca\x16\xff\x4e\x95\xa1\xe0\x86\x5c\x51\x1e\xc8\x44\xf9\xe1\xa3\xe4\xa3\x5f\x82\x45\x33\x59\x9e\xd6\xb8\x6e\x5f\xe6\xfe\xe8\xee\xbc\xed\xff\x70\x8f\x13\x7b\xba\x73\xac\x6f\x72\x26\xc4\xb6\x3f\xc5\xfe\x10\x0c\x4b\xbd\xa2\xf1\x3a\x54\x5a\x8f\x72\x3e\xbe\x18\xa4\x5a\xfc\x40\xb8\x01\xd7\xcd\x26\x80\x2a\x9c\xc5\x66\x23\xa3\x77\x49\xc2\x76\xbb\x46\x6a\xd9\xab\xa2\xb8\x25\x13\x5d\x39\xa1\x76\x1d\x2a\x8f\x87\x95\xf4\x6f\xc3\xd6\x49\x9c\xaf\x00\x00\x00\xff\xff\x5c\xaf\x12\xe0\x23\x02\x00\x00")

func include_fn_bash_bytes() ([]byte, error) {
	return bindata_read(
		_include_fn_bash,
		"include/fn.bash",
	)
}

func include_fn_bash() (*asset, error) {
	bytes, err := include_fn_bash_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "include/fn.bash", size: 547, mode: os.FileMode(420), modTime: time.Unix(1416714978, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _include_herokuish_bash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x55\xef\x6f\xdb\x36\x10\xfd\x2c\xfd\x15\x57\xce\xa9\x1b\x03\x8a\xb1\xac\xd8\x87\x18\x1a\xe6\x64\xee\x52\xac\x4b\x8d\xa6\xed\x30\x24\x85\xc1\x88\x67\x99\x88\x44\x0a\x24\x95\x38\x48\xf2\xbf\xf7\xa8\x5f\x96\xed\xc4\x9f\xc4\x7b\xef\x8e\xc7\xc7\xc7\x73\x28\x97\x70\x75\x05\x6c\x70\x3a\xbd\x3c\x5f\x7c\x9f\x7d\xb9\xfc\x78\xf1\xe1\x33\x83\x28\x73\xc0\xde\x33\xf8\xf1\x63\x02\x6e\x85\x2a\x0c\x30\x59\x69\x60\x6f\xde\xc0\xff\xba\x34\x60\x1f\xac\xc3\x1c\x4e\xb9\x5d\x81\xb4\xa0\x4b\x07\x7a\x09\x82\x3b\x3c\x81\x5e\xad\xcf\x17\xac\x97\x39\xcf\x90\x5b\x84\xb2\x48\x0d\x17\x08\x4e\xd7\xf9\xef\x41\x1b\x48\x0d\x52\xb2\x39\xf2\xfc\xb5\x74\x70\x1c\x2e\x65\x18\x52\x50\x68\x95\x3d\x00\x2f\x8a\x45\xc1\xdd\x2a\x66\x83\xc7\xe9\x7c\xbe\x98\x4f\xbf\x9e\x9f\x44\x63\x0a\x3f\xb3\x0d\x0b\xd5\x5d\xc7\x9a\x5d\x7c\x6f\x59\x2e\x2f\xc6\x04\xf5\x99\x37\xa5\xcc\x44\xc7\x3d\xfd\xf6\xf1\xd3\x5f\x7d\x76\x05\xf7\xf9\x09\x4f\x56\xd8\xf1\xcf\xa6\x67\xe7\xb3\x3e\xbf\x82\xf7\xea\x17\x3c\xb9\xdd\xde\x63\x3e\x3d\xfb\x67\x6f\x1f\x4f\xb3\x94\xdc\xdb\x0d\x05\x37\xd2\xae\x16\x77\x68\xac\xd4\x8a\xf2\xdf\x71\x6b\xd1\x45\x09\x77\x20\x55\x92\x95\x02\xc7\x2d\xeb\xc8\xad\xdd\x21\xa5\x0b\x4c\x32\x6e\x48\x5f\x55\x18\x79\x27\x33\x4c\x51\x2c\x4a\x8b\x26\x66\x4a\xdf\x68\xf1\xc0\x5e\xa6\xa4\x46\x97\x85\xe7\x54\x1f\x54\x08\x95\x2d\x0d\x46\xbe\x73\xfb\xee\x10\x1e\xc3\x20\xbf\x15\xd2\x40\x54\xc0\x75\x18\x04\x6c\xd0\xde\x06\x6b\xd6\xad\xee\xed\x7a\xa3\x6e\x1b\xd9\xe8\xb7\xc5\xe9\x14\x62\xe1\x73\x18\xf6\x36\x6c\x1b\x15\x68\x93\x98\x5d\xae\xf4\xbd\x05\x0f\x03\x89\xe0\xa4\x4a\x2d\xf9\x84\x8e\xa0\xdc\x12\xd8\x41\xf4\xdb\xb1\x85\x5f\xe0\xc0\x5e\xab\xa6\x7a\x6b\x91\xb8\xd7\xaa\x0f\x17\x45\x26\x49\x42\x92\xb4\xae\x26\x4a\xaa\x91\x82\x29\x95\x93\x39\x36\xc9\xad\x73\xe2\xde\xb9\x28\x3c\xf7\x09\xe4\xd9\x25\xc9\x66\x61\x49\x9e\x15\xb8\x94\xca\xe7\xdf\x78\x5b\x13\x59\x1a\xad\x72\x54\xae\x29\xb4\xb1\x55\xbc\x25\x49\xc0\xfe\xd3\xe6\xd6\x27\x92\xaa\x98\x38\x6d\x1e\xda\x4e\x2a\x9a\x6d\xf2\x37\x36\x8b\xb7\x04\xa4\xca\xad\x78\xb5\x31\x21\xd3\xf5\xa9\xfa\x1b\x77\x5e\x8b\x77\xb5\x86\xee\x2c\x52\x59\xc7\xb3\x0c\xc5\xc6\xb0\xb6\xba\x8a\xc6\x78\xaf\x5d\x06\x34\x38\x70\x25\xc0\x96\x45\xa1\x8d\xa3\x22\x6d\x54\xaa\xa5\xee\x1e\xfe\x0a\x8d\xbe\x2d\xc9\xa7\x2d\x4c\x13\xe2\xb1\x99\x0e\x27\x91\x40\xff\x2e\x1b\x6a\xa2\x73\x6a\x50\xde\x64\xd8\x3d\x01\x22\xef\xbe\x86\x97\xe8\x9b\xf6\x4f\x08\xde\x7f\x2b\x1b\xdc\xbf\x16\x78\x22\x1f\x09\x18\xda\xf1\xd1\xa8\x6e\xef\x7a\x3c\x1e\x0f\x29\xbc\xe6\x26\x25\xa7\x35\xce\x02\x38\x88\x8e\x7f\xb7\xb5\xb3\xbc\x2c\x4e\xba\x0c\x6b\x51\xaa\x16\x06\xc3\x6b\xbc\xfa\xf5\xef\xc8\xff\xfe\x18\xc2\x60\xe4\x49\x52\x09\xf2\x40\xcd\xba\x5f\x91\x5b\xc0\x3f\x6c\xc8\xa4\xc2\x09\x08\x4d\x17\xd4\x0e\x5d\x1f\x62\x10\xc7\x10\x45\x23\x1a\xb5\xf0\xf4\xb4\x13\x8f\xe3\x51\x6f\x04\x07\x5b\x9b\x0e\x2b\x1a\x45\x31\xb3\xb8\x0b\x42\xf5\x1b\xb6\xb5\x08\xa6\x71\x1a\xd0\x74\x41\xdf\x61\xff\xf9\xd7\x7d\x92\x5e\xa5\x14\xa9\x14\x94\xb1\x37\x3f\xc8\x30\x83\x3f\xab\xf3\xe7\x5c\xaa\x2e\x01\x22\xd4\x50\xc8\x02\x97\x5c\x66\x93\xba\xf3\xaf\x5f\xa6\x67\x33\xff\xb7\x01\x6f\xdf\x42\xc5\x59\x87\x61\x90\xe4\x22\xc2\xb5\x37\x49\xf5\xec\xec\x56\xa4\xb9\xd5\x2d\x5a\xa4\xec\xe6\x4a\x81\x7d\xa3\xe7\xe5\x9d\xd6\xd8\x75\xcb\xac\xfd\x52\x5d\x3c\xaa\xbe\x5e\xc1\x9a\x2a\xaf\xa0\x99\xb4\x6e\xaf\x17\x9b\x95\x29\xb0\x7f\xb9\xe2\x29\xfa\xff\xa2\x6e\x8a\x78\x60\xa7\x09\x1f\x8a\x64\xee\xbf\xf7\xe3\x29\x2a\x34\xf4\x47\xb7\x8f\xd4\xdf\x7b\x3b\x17\x46\x27\x7e\xe2\xd4\x22\xcc\x9b\x95\xad\xe4\xa0\xb1\xe5\x9b\x01\x7a\x07\x39\xad\x77\xfa\x68\x33\x23\x3a\xed\x4e\x2b\x1d\x84\x6b\x4c\x5e\x46\x0a\x6e\xbc\xab\x08\xf4\xb3\x8d\x0d\x2e\x67\x9f\x3e\x30\xba\x00\x20\x2b\x8d\xab\x8a\x87\x41\xb0\xbd\x45\xe5\x92\xc9\xc4\x13\x7c\xdd\x3e\xee\xd7\x3d\xb8\x52\x9b\xf0\x9d\xfb\xaa\xc0\x11\xc5\xab\x8e\xe8\xf0\xac\x76\x1e\x85\xd1\xf2\x24\x7c\xfe\x19\x00\x00\xff\xff\x4e\xb4\xaa\x09\xae\x08\x00\x00")

func include_herokuish_bash_bytes() ([]byte, error) {
	return bindata_read(
		_include_herokuish_bash,
		"include/herokuish.bash",
	)
}

func include_herokuish_bash() (*asset, error) {
	bytes, err := include_herokuish_bash_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "include/herokuish.bash", size: 2222, mode: os.FileMode(420), modTime: time.Unix(1418865371, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _include_procfile_bash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x95\xdd\x4f\xe3\x38\x10\xc0\x9f\x93\xbf\x62\x14\x45\xd0\xde\x61\x02\xf7\x78\xa8\xe8\x4e\xe2\x74\xbc\x00\x2b\x5e\x59\x54\xb9\xf1\x34\xb1\x9a\xc4\x59\xdb\x29\xb0\x2c\xff\xfb\xfa\x2b\x6d\x4c\xcb\x53\xe3\xf1\x7c\xfc\xe6\xc3\xd3\xb4\x97\xa2\x5c\xf3\x06\x49\x4f\xa5\xc2\xd9\x1c\xde\xd3\x84\x61\xd9\x50\x89\xc0\x50\x95\x8b\xec\x7f\xd4\x50\x8a\xb6\xa5\x1d\x03\xa5\x25\xef\x2a\x58\x0b\x09\x14\xac\x29\x2a\x05\xfa\xad\x47\x58\x4b\xd1\xc2\xb7\xe0\x2c\xdb\xfb\xb0\x97\x8b\x2c\xbf\x34\xa2\x92\x6a\xc8\x72\xda\xf7\xcb\x9e\xea\xba\xd8\x29\xc3\x2f\x78\xa3\x6d\x43\x2a\xb4\xf7\xd6\x20\x4b\x3f\xd2\x3d\x99\xd2\x54\xea\x63\x64\x8f\x43\x17\x43\x8c\x98\x11\x0c\xe8\x5a\x8a\xa1\xaa\x01\x5f\xb1\x3c\x4e\xb6\x0b\x65\x55\x0c\xc3\x2c\xae\xca\x48\x35\x8f\xb9\xac\xf2\x57\x58\x54\x81\x34\x24\xa2\x3d\x83\xa1\xeb\x25\xdf\x1a\x83\x0a\x19\x0c\x0a\x25\xbc\x70\x5d\xc3\x2d\x4a\xb1\x19\x48\xc3\x37\x08\xd8\x6d\x6d\x7d\xd8\xa4\x3c\x53\x2a\xef\x89\xff\x44\x62\xcd\x27\x17\x0a\x35\xa9\x45\x8b\x13\x51\x23\x28\x23\xc6\xdf\x67\x91\x39\xd9\x43\x9a\x44\x38\xc5\x8a\x77\xc5\x8a\xaa\x1a\x88\xcb\x1b\xb7\xb4\x01\x2c\x6b\x01\xf9\x3f\x9f\xb2\xb5\x15\x50\x3e\x5d\xcd\xb5\x29\x6b\x76\xc3\x55\x29\xb6\xe8\x26\x62\xda\x06\x65\xd8\xf9\x1a\x9e\x9e\x80\xac\x8f\x37\xfc\xf9\xf9\xca\x74\x05\xbb\x34\x49\x1a\x51\x9a\x90\xce\xcc\x9c\xdc\xaf\x69\xca\xcc\xce\xca\xa1\xe5\x38\x29\x1b\x7c\x53\xe6\xfb\x95\xca\x4a\x39\x5c\xc3\x9a\x24\x8e\x3b\xdb\xe9\x86\xa6\x04\x26\x20\xd7\x90\xbf\xbb\xcf\xa2\x80\xe2\x0c\x3e\xac\x89\x44\x3d\x48\x83\xb1\xe6\x3b\x64\x35\x45\x3e\x97\xd8\x20\x55\xc7\x90\x19\xae\xe9\xd0\xe8\xe5\x88\x1e\x9d\x0f\x53\x18\x3d\x45\x29\x8c\x36\xa1\x7a\xde\xf6\x30\x31\x83\x95\xe5\x91\x7f\xcb\x03\x27\x27\xf0\xdd\xdc\x86\xbc\x6f\xfc\x7d\xdc\x09\xf7\x52\x73\x65\x42\x97\x1a\xd9\xb2\xa3\x2d\xfa\x42\x44\xde\xbe\x28\x88\xf7\x7b\x2f\x0e\x5c\x0e\x1d\x8b\x67\x63\x1c\x3a\x3f\x1e\xa1\x90\x76\x9a\x8d\xcc\x4f\xf3\xb4\x7e\x96\x09\x81\x77\x90\xcf\x1a\x05\x3b\x9d\xf9\x15\x30\xe1\xf2\x79\xed\x85\xb4\xab\x00\x17\xa1\x8a\xa3\x4a\x91\xa3\x2b\x08\x13\x1d\x3a\xc6\x03\x88\x30\xe6\x1e\x44\xd5\xa2\xd7\xb6\xa3\xdd\xd0\x34\x55\x23\x56\x69\xd2\x6e\x18\x97\x40\xfa\xa8\xc9\xc1\xe8\xdc\x64\xe5\xe0\xdc\xf8\x58\xbe\x23\x2a\xc5\x1f\xe7\xaa\x0e\xa4\x4a\x0c\xb2\xb4\xdb\x61\xdc\x79\x0e\x2b\x44\x1d\x26\x51\x6b\xf7\xc2\x64\xbc\xd5\xc2\xdb\xf5\xa8\x21\xe7\xdb\x87\xbb\xff\x16\xd1\x12\xb0\x2f\xbe\x15\x0c\x88\x53\x9e\x2e\x08\xf3\x3d\x7d\xcb\x4b\xab\x99\xc1\x35\x14\x0c\xb7\x85\x8d\x0d\x7f\x5d\x9f\x5c\x9a\xb5\x52\x8b\x97\x0e\xc8\xe3\x31\xfd\xbf\x63\x51\x65\x16\x65\x9f\x45\x5b\x68\x8a\x1c\xef\x21\x0f\xee\x1f\x83\x3d\x73\x66\xa7\x7e\xf6\xf8\xef\xfd\xcd\xc3\xdd\x9f\x97\x17\x17\x17\x73\xdb\xac\xbd\x82\x9d\xbe\x45\x36\xe4\xef\x5e\xdb\xcc\x5b\x9a\x50\xc6\x5c\x50\x93\xdf\x8f\x81\x9b\x3f\x00\x42\x2a\x6e\x27\xc7\xeb\x64\xe1\xcb\x9a\x66\x4e\xdb\x2d\x50\x3b\xf6\x84\xa8\x1a\x4d\x92\xfb\x15\xe6\xa5\x8c\x2b\xba\x6a\xd0\x8c\x02\x55\xea\x45\x48\x16\xe4\xa6\xb1\x25\x92\x15\x65\xee\x11\x78\x59\x27\x48\x29\x91\x6a\xf4\xc5\xf5\xc2\x21\x8a\xef\x65\xd5\x31\x19\x96\x42\xc1\xe9\x69\x38\x7a\x7c\xff\x7d\xd0\x2a\x2b\x8e\x32\x89\xf7\xb0\xeb\xc5\x22\x52\x38\xec\x4b\x74\xff\xf1\x3b\x00\x00\xff\xff\x86\xee\xe1\xff\xb9\x07\x00\x00")

func include_procfile_bash_bytes() ([]byte, error) {
	return bindata_read(
		_include_procfile_bash,
		"include/procfile.bash",
	)
}

func include_procfile_bash() (*asset, error) {
	bytes, err := include_procfile_bash_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "include/procfile.bash", size: 1977, mode: os.FileMode(420), modTime: time.Unix(1418754982, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _include_slug_bash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x53\xef\x6f\xd3\x30\x10\xfd\xdc\xfc\x15\x87\x55\x69\xed\x07\x37\x8c\xaf\x28\x48\x68\x20\x34\x09\x01\x62\xad\x84\x34\xa6\xca\x24\xd7\xc4\x92\x13\x5b\xb6\x03\x5d\x80\xff\x9d\x73\x7e\x2c\xc9\xda\x4d\xf4\x53\x73\xbe\xf7\xfc\xfc\xee\x5e\x64\x51\x64\xba\x52\xf7\xe0\x54\x9d\xef\x8d\xf0\x45\xc2\x62\x5f\x9a\x38\x7c\x6f\x7c\xde\xb0\x28\x0a\x7f\xb9\x2c\x8d\xb6\x7e\xb5\x86\xdf\xd1\x22\xc3\x54\x09\x8b\x90\xa1\x4b\x13\x76\xdd\x9e\x80\x80\xbc\x91\xc6\x60\xd6\x52\x81\x17\xf6\x87\x50\x0a\x0e\x56\x97\xb0\xfb\xfa\x11\xb4\x85\x9b\xed\xbb\xeb\x4f\xc0\x46\x82\xda\xaa\x84\x2d\x2f\xa9\x82\x95\xab\x2d\xf2\x20\xc0\x45\x0b\x79\x80\xdb\x5b\x60\xcb\x95\x72\xc0\xdf\xc2\x52\x18\xd3\x6a\x5b\x33\xb8\xbb\x7b\x0d\xbe\xc0\x2a\x5a\x2c\x2c\xfa\xda\x56\x70\x49\x68\x35\x20\x88\x71\xd6\x93\x52\x01\x38\xb1\x70\xea\xb6\xf7\xf0\x6a\xe8\xf9\x13\x14\x02\x3f\x36\x57\x54\x19\xf8\x83\x10\xe5\x30\xe0\x84\x7f\xb2\xe5\x20\xa3\xbf\xbd\x2b\x39\x56\x68\x85\xc7\x73\xbe\x7c\xe8\xcf\x9e\x75\x86\x64\x02\x69\xb4\x58\x91\x83\xc6\x9c\x38\xa1\x74\x2a\x14\xa4\xba\x34\x16\x9d\xdb\x6b\xe3\xa5\xae\x12\xc6\x69\x2e\xc1\xa5\x5f\x85\x4c\x0b\x30\x32\x6f\xe0\x0d\xc4\x19\xfe\x8c\xab\x5a\xa9\xf1\xf5\x27\x38\x5e\x3b\xe4\x43\x99\x1b\xab\x73\x2b\xca\x24\x10\x74\x0f\xeb\x2f\x0c\x42\x65\x5e\x69\x8b\x3d\x74\x98\x09\x3f\x4c\xbc\x88\x37\x63\xdf\xcc\xf5\x13\x38\xdd\xfc\x0d\xce\xe3\xba\x6b\x83\xd3\xcb\x47\x72\x61\x79\xc2\x03\xdf\x89\x9d\x73\x3c\xa6\xaa\xce\x30\xb9\xd8\xe4\xd2\x5f\x74\xc5\xd9\x90\xba\x52\x1a\xc4\x3e\x2c\x76\x28\x02\xfd\x16\x9b\xe9\x2b\xf7\x4e\x36\x48\x4b\xb8\xca\x6a\xe0\x37\x05\x8c\xfd\x34\xff\xb4\xf6\xf4\xe2\xcb\x35\x89\xf4\xd2\x2b\x04\x76\x45\x12\xa5\x1a\x46\x19\xb0\x20\x5d\x0f\x0a\x5f\xec\x61\x35\xf0\xf8\x54\x60\xde\xb7\x27\x30\xec\xce\xa3\xb5\xf0\xba\x8d\xcb\xea\xcb\x6e\xbb\xee\x43\xf3\x79\xb7\xfd\xef\xd0\xbc\xe8\x46\x34\x79\xf5\xd9\xc4\x04\xcf\x9f\xcf\xcc\xcb\x36\x36\x7a\xdc\xaa\x49\x86\x68\x98\x24\x0f\xf8\x76\x7e\x53\x47\x35\xcb\xd0\xf4\xbc\x4b\xce\xbf\x00\x00\x00\xff\xff\x5e\x30\xc4\xb4\x73\x04\x00\x00")

func include_slug_bash_bytes() ([]byte, error) {
	return bindata_read(
		_include_slug_bash,
		"include/slug.bash",
	)
}

func include_slug_bash() (*asset, error) {
	bytes, err := include_slug_bash_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "include/slug.bash", size: 1139, mode: os.FileMode(420), modTime: time.Unix(1418873487, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	"include/buildpack.bash": include_buildpack_bash,
	"include/buildpacks.txt": include_buildpacks_txt,
	"include/cedarish.txt": include_cedarish_txt,
	"include/cmd.bash": include_cmd_bash,
	"include/fn.bash": include_fn_bash,
	"include/herokuish.bash": include_herokuish_bash,
	"include/procfile.bash": include_procfile_bash,
	"include/slug.bash": include_slug_bash,
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
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"include": &_bintree_t{nil, map[string]*_bintree_t{
		"buildpack.bash": &_bintree_t{include_buildpack_bash, map[string]*_bintree_t{
		}},
		"buildpacks.txt": &_bintree_t{include_buildpacks_txt, map[string]*_bintree_t{
		}},
		"cedarish.txt": &_bintree_t{include_cedarish_txt, map[string]*_bintree_t{
		}},
		"cmd.bash": &_bintree_t{include_cmd_bash, map[string]*_bintree_t{
		}},
		"fn.bash": &_bintree_t{include_fn_bash, map[string]*_bintree_t{
		}},
		"herokuish.bash": &_bintree_t{include_herokuish_bash, map[string]*_bintree_t{
		}},
		"procfile.bash": &_bintree_t{include_procfile_bash, map[string]*_bintree_t{
		}},
		"slug.bash": &_bintree_t{include_slug_bash, map[string]*_bintree_t{
		}},
	}},
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

