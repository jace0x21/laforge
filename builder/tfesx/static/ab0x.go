// Code generated by fileb0x at "2019-04-06 20:36:59.22708159 +0000 UTC m=+0.008699485" from config file "assets.toml" DO NOT EDIT.
// modification hash(d20a740e55ac8e6345055047c9212e21.e5b9c5ef4c0b7aef8593382d0449dfd6)

package static

import (
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"net/http"
	"os"
	"path"

	"golang.org/x/net/webdav"
)

var (
	// CTX is a context for webdav vfs
	CTX = context.Background()

	// FS is a virtual memory file system
	FS = webdav.NewMemFS()

	// Handler is used to server files through a http handler
	Handler *webdav.Handler

	// HTTP is the http file system
	HTTP http.FileSystem = new(HTTPFS)
)

// HTTPFS implements http.FileSystem
type HTTPFS struct {
	// Prefix allows to limit the path of all requests. F.e. a prefix "css" would allow only calls to /css/*
	Prefix string
}

// FileCommandTfTmpl is "command.tf.tmpl"
var FileCommandTfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xb4\x93\x31\x8f\xd4\x30\x10\x85\x7b\xff\x8a\x91\x45\x01\x12\x44\x14\x88\xee\x0a\xb8\x02\xe8\x10\x14\x14\xe8\x64\xe5\xe2\xd9\xdb\x11\xf1\x4c\xe4\x99\xec\xde\x29\xca\x7f\x47\x93\x90\xb0\xc5\x22\x41\x81\x9b\x64\xf4\x66\x9e\x9f\x3f\xd9\xd3\x04\x19\x0f\xc4\x08\xb1\x93\x52\x5a\xce\x11\xe6\x39\x54\x54\x19\x6b\x87\x10\x79\xec\xfb\xb4\x95\x11\xe2\x50\xe5\x44\x4a\xc2\x69\x9a\xa0\xf9\x80\x06\x71\x53\x13\xb7\x05\x7d\x3c\xa9\xe1\xb0\xcb\x5e\x24\x1e\xcb\x3d\x56\x17\x23\x4c\x01\x20\xe3\x80\x9c\x35\x09\xc3\x0d\x7c\x0f\x00\x00\x91\xee\x4b\xea\xa4\x0c\xa3\x61\x3a\x95\x44\xac\xd6\x72\x87\xcd\x9f\x37\x8a\x01\xe0\x2e\x04\x80\x3d\x15\x56\x6f\x2b\x62\xf8\x0a\x1f\xb1\x5b\x37\x03\x98\x26\xa0\x03\x34\x1f\x45\xad\xf9\xa4\xdf\x88\xb3\x9c\xd5\x0f\x0a\xcb\xea\x84\x19\x3b\x23\xe1\x5f\xfd\xbe\x8e\xa2\xb6\xfc\xdc\x40\x7c\x36\xfd\x7b\xb8\x86\x86\xd3\x9b\xd4\xe6\x5c\x51\x75\x89\xba\x2e\x7b\x1a\x70\xf3\x3d\x13\xd7\xf2\x5b\x1a\x15\xeb\x26\xbd\xcb\x85\x98\xd4\x6a\x6b\x52\x2f\xa6\xa9\xa0\x8c\xb6\xb4\xbc\x7d\x7d\x31\x3b\xb4\xaa\x67\xa9\xd9\x05\x0f\x75\x2b\x65\x40\x23\x3f\x54\xf3\x45\xc4\x3e\x6f\xfa\xbc\x67\x99\x37\x36\xd8\x2b\xfe\x2d\x8d\xff\x08\x64\xb5\x56\x3d\x5e\x21\xb2\x6a\x55\xc4\xae\xb0\x80\x2b\x38\x2a\x9d\x5a\xc3\xf4\x03\x9f\xd6\xbc\x07\xea\xf1\xb9\x93\x21\xce\xf8\x08\xcd\xfb\x91\xfa\xdc\xdc\x0a\x1f\xe8\xc1\xc3\xf6\x49\xf5\x98\x2e\xc6\x92\x4f\x2c\xb7\xec\xc5\x15\x62\xec\x20\xc3\x52\x12\xf7\xfe\x80\xb6\x7b\x0c\x3b\x7f\x7f\x4e\xdb\xf7\xab\x55\xe2\x07\x77\x7b\xb9\x74\xdd\x05\x77\x9b\xc3\xee\xf5\x33\x00\x00\xff\xff\x0d\x32\x43\xcf\x8a\x03\x00\x00")

// FileDNSRecordTfTmpl is "dns_record.tf.tmpl"
var FileDNSRecordTfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xac\x92\x41\x6b\xe3\x30\x10\x85\xef\xfa\x15\x0f\xb1\xc7\x20\x02\x7b\xce\x6d\x61\xe9\x25\x87\x16\x7a\x68\x29\xc2\xb5\x26\xae\xa8\x33\x32\x1a\xb9\x85\x0a\xfd\xf7\x22\x13\xdb\x09\xa1\x97\xb6\x47\xcd\x7b\xbc\x99\xf9\x46\x39\xc3\xd1\xc1\x33\x41\x3b\x16\x1b\xa9\x0d\xd1\x69\x94\xa2\x54\x24\x09\x63\x6c\x4f\x4a\x73\xd2\xac\x50\xd2\xd0\x43\x0c\x6f\x5e\x7c\x60\x9b\x33\xcc\x7f\x4a\xd0\xb3\xdf\x72\x73\xa4\x1a\x61\x25\xd1\xb0\xc8\xf5\x61\x79\x3c\x3e\x53\xac\xa2\x46\x56\xc0\x47\x60\xc2\x0e\xba\xba\xfe\xed\xef\x6e\xa7\x16\xe6\xa1\x56\x4b\x31\x5a\x01\x35\xec\xda\xb1\xaf\xd5\x52\xaa\xa1\x71\x2e\x92\x08\x09\x76\x78\x54\x00\x90\x33\xfc\xe1\xdc\x7d\xc3\x2f\x14\x7d\x22\x57\xf7\xaa\x0e\xfd\x27\x77\x21\x74\x3d\xd9\x36\x1c\x87\x31\x91\xf5\x2c\xa9\xe1\x96\xcc\xd7\xeb\x18\xa6\xf4\x1e\xe2\xab\xf5\x9c\x28\x1e\x9a\x96\xcc\x76\xad\x0d\x45\x6f\xe6\xee\xd4\x0b\x2d\xad\x2e\x07\xbf\x6f\xfa\x91\xa4\x8e\xbe\xba\xf9\x34\xd7\x93\x02\x52\xea\xb1\xc3\xdf\xed\x56\x5d\x5c\x80\xc7\xbe\xb7\xf3\xf3\xf7\xf0\xa7\xe8\xbb\x8e\xa2\x4c\x0f\x60\x86\x60\xbd\xab\xc8\xbf\x05\xe9\x2c\x63\x3a\x4f\x51\x0a\x70\x34\x10\x3b\xb1\x81\x97\x23\xe9\x25\x63\x15\xf5\xca\xe5\xea\xcf\x99\x9f\xaf\xbc\x99\x18\x17\xa5\x56\xe8\x9f\x01\x00\x00\xff\xff\xac\x82\xb9\x56\xff\x02\x00\x00")

// FileInfraTfTmpl is "infra.tf.tmpl"
var FileInfraTfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xcc\x56\x5d\x6f\xdb\x36\x14\x7d\x96\x7e\xc5\x85\xec\x60\x2f\x15\x67\xf7\xa1\x28\x06\xe4\x21\x6d\xdc\x2c\x40\xe7\x14\x8e\xb3\x0e\x18\x06\x81\x16\xaf\x6d\x2e\x12\x29\xf0\x43\x5e\x66\xf8\xbf\x0f\xa4\x28\x7f\xbb\xf3\x92\x00\x9d\x1f\x1c\x46\xf7\xde\x73\x0f\x0f\x2f\x8f\xd5\x79\xfe\x27\xee\xc0\xe7\xab\x4f\x77\xa3\x9b\x01\xdc\x0c\x86\x83\xd1\xd5\x78\x70\x0d\xe3\xc1\x68\xe4\x1e\xfe\x02\x1f\xef\x86\x9f\x6e\x6f\x1e\x46\x57\xe3\xdb\xbb\x61\xdc\x81\x34\x85\xaf\x57\xa3\xe1\xed\xf0\x06\xd2\x34\xee\xc0\x78\xce\x35\x4c\x79\x81\xc0\x35\x50\x6b\x64\x49\x0d\xcf\x69\x51\x3c\xc1\x0c\x05\x2a\x6a\x90\x11\xb8\x96\x20\xa4\x01\x64\xdc\x00\x37\x3f\xe8\xb8\x03\xb9\x14\x06\x85\xd1\xc0\xb8\xc2\xdc\x14\x4f\x04\x1e\x34\xc2\x67\x3a\x95\x6a\x86\x40\x05\x03\x85\x30\xb1\xbc\x60\x60\xda\x26\x24\x7e\xc9\x4e\xe3\x4a\xc9\x9a\x33\x54\x90\xd4\xba\x9a\xa3\xc2\x04\x96\x71\x14\xd6\x99\x46\x55\xa3\x82\x4b\x48\x96\x4b\xe0\x82\xe1\x5f\xd0\x25\x1f\x1c\x01\xf2\x51\x8a\x29\x9f\xad\xcb\x42\x6a\x02\xab\x55\x12\x47\x56\x9f\x59\xe5\x12\x9b\x1a\x00\x80\x8a\x6a\xbd\x90\x8a\x9d\x55\xda\x26\x87\x96\xb4\x28\xe4\x22\xb3\xa2\x46\xc5\xa7\x1c\x59\xa6\x75\x01\x97\x60\x94\xc5\x78\x15\xc7\x1d\x98\x53\xc5\x20\x97\x8a\x21\x03\x96\xc7\x8c\x1a\xba\x01\x73\xff\xe5\x28\x8c\x63\x93\xb0\xdc\x8b\x20\x68\x89\x8e\xc9\xe0\xfe\xb7\xdb\x64\x07\xc3\x43\x50\x43\xb5\x91\x0a\x8f\x20\xf9\xe7\x0e\x68\xb3\xde\xc2\x5b\x3f\xed\x27\x71\xb4\x69\x9c\x71\xbf\xef\xee\xd2\x3d\x22\x87\xc4\x08\xcb\x09\x67\xab\xe4\x70\x37\x0a\xb5\xb4\x2a\x47\xa8\xa4\x2c\xf6\xe8\xb4\xb1\xcc\xc5\x12\x48\x9a\x3f\x5b\x6c\x46\x21\x41\xbf\x9c\x8c\x97\x85\x6b\xa3\xf8\xc4\x1a\x64\xa0\x17\xdc\xe4\xf3\x7d\x7d\x36\x09\x59\xcd\x95\xb1\xb4\xc8\x9a\x44\x27\x58\xad\x77\xc8\x5d\xdf\x37\x91\x57\xa1\x66\xb0\xac\x0a\x6a\x10\x14\x4e\x51\xa1\xc8\xf7\x8f\xae\xa5\x53\xd2\x7c\xce\x85\x3b\x40\x3b\xb1\xc2\xd8\xfe\xfb\x1d\x52\xe3\x80\x93\x3e\xf8\x28\xf4\xdf\x7f\x37\x7e\x0b\x2e\xfa\xbd\xe3\xe4\xbe\x72\xc1\xe4\x42\x43\xbf\xf7\xfd\xd8\xbd\x7d\xec\xbf\xfb\x36\xbb\xfb\xc6\x60\xde\xf6\xfa\xef\x5e\x40\xb3\x90\xb2\x02\x23\x21\x57\xe8\x18\x56\x52\x19\x98\x29\x69\x2b\x1d\x2f\x97\xa0\xa8\x98\x21\x74\x2b\x81\x86\xb3\x37\xcd\x02\x7e\xba\x84\x2e\x19\x23\x2d\xc9\x17\xe7\x80\x9a\x4b\x81\x6c\x88\x66\x21\xd5\xa3\x86\xd5\xca\x15\x76\xdb\x44\x57\x41\x42\x70\x2b\xe6\x77\xe5\xe2\x2e\xfc\x85\x9a\xb9\x8b\xad\x2f\xe3\xd1\x89\x77\xd4\x32\x4f\xcd\xdd\xc5\x59\x6a\x1c\x54\xc3\xc3\x7d\x0d\x6d\x39\x41\x05\xab\x55\x1a\x5a\x90\x0f\x54\xa3\x33\x38\x58\xc6\xb0\x96\xf1\xec\xaa\x18\xa2\xd3\xf7\x2d\xb3\xf6\xb8\xca\x27\x2b\x08\xab\x75\x23\x3b\x44\x75\x41\x45\x73\x4a\x27\xc8\x6c\x0c\xdc\x31\xfa\x95\x2a\x0d\x49\x28\x72\x96\x1d\x7b\x19\x51\x30\xb7\x8e\x3b\x50\xd2\x47\x84\x8a\x2a\x14\x06\xa6\xb2\x60\xa8\x8e\x48\xd9\x04\x12\x48\x0c\xd2\x32\x3d\xd1\xd8\x4f\x5c\xe5\x8e\xc3\x4f\x1c\x2d\x4f\x31\x4c\xe2\xc8\x3c\x55\x5e\xd1\xba\x7c\xbd\xf1\xd3\x76\x12\x76\xf0\xbf\x18\xbf\xb3\x34\x3b\x3e\x6f\xad\x88\xdd\xe5\x2e\x18\xf9\x16\x14\x71\x55\xab\x1f\x0f\x00\x5f\x43\xec\x9d\x89\xf1\xb2\x87\x01\x85\xe0\x3a\x2f\x50\x1c\x20\x50\x96\x93\x3f\x8f\xca\xee\x13\x5a\xec\xb9\xd4\x01\xdc\xad\x36\xf9\x5b\xe0\x3f\x4b\x6d\x02\x72\x83\xbd\xc9\x74\x2b\xe2\xe2\xdb\xe1\xf5\x4f\x75\x7b\xb0\x95\xe2\xc2\x4c\x21\xb9\xd0\xa9\xb9\x60\xe9\x85\x4e\x2f\x74\x02\x5d\x32\x10\x35\x57\x52\x94\x28\x82\xbc\x87\xc7\x10\xf6\x11\xc2\xbe\x5d\x38\x88\x3d\x5b\xaf\x4b\xd0\x15\xe6\x7c\xca\x73\x6a\xb8\x14\xfa\x0d\x58\x8d\x50\xd2\xca\xf3\x3a\x1c\xa7\x43\x9b\x3f\xe4\x1e\xe6\xc7\x7f\x5a\xcf\x3a\x9a\x15\x47\x9b\x16\xfe\x0d\xe5\xf8\x2c\xec\xa4\x10\xff\xe5\x07\x22\x82\xcd\x6b\xd8\xe9\x29\xf2\x61\xb2\x59\xb5\xa5\xc2\x96\x59\x5e\x59\xdd\x58\x58\xf0\x2a\x2f\x55\x30\xab\x32\xaf\xac\xb7\xaa\x08\xa0\xc4\x52\xaa\xa7\xd3\xa9\x25\x96\x6d\xea\xcc\xa2\x36\xc7\xe9\xec\x89\x47\xda\xa9\x20\x77\xf7\xee\xe6\xb4\x95\x0d\x3f\x9d\x6b\x9e\xb5\x77\xe6\xbf\x01\xad\x4b\x1b\xa4\x29\x57\xe5\x82\xaa\x67\x00\xb5\x95\x01\xc7\xdf\xfe\x17\x3a\xc2\x7a\x2c\xc3\xf9\x8b\xe6\x86\x65\xdc\xdd\xf5\x29\xcd\xd1\xd9\x77\xb4\x7e\xca\x76\xdb\x1d\xff\x31\x25\x67\xfc\x94\x6e\xf5\x5d\x8f\x00\x00\x65\xb4\x72\x1e\xf4\x3c\x9d\x0f\xb8\x7b\x1c\xfd\x7b\xef\x8f\x06\xdf\xcf\x03\xe3\xfa\xd1\xef\x49\xf3\xbf\x9f\xd1\xc3\x95\x6b\xd2\x23\xae\xda\xa1\x46\x05\x9d\x60\x71\xce\x6b\xc0\xd6\x8e\xd3\x3d\xd0\xba\x64\x8f\x49\xb8\xa1\x9e\x64\x5e\x48\xd1\x28\xdf\xbe\xe8\x9d\x7a\x3d\xf8\x17\xba\x8d\xb4\x91\xe1\x25\x4a\x6b\xe0\x12\xde\xf5\x1a\x25\x9a\x56\xde\xea\x82\x87\x6f\xb9\xf9\x3f\x01\x00\x00\xff\xff\x9f\x89\xf2\x96\xae\x0f\x00\x00")

// FileProvisionedHostTfTmpl is "provisioned_host.tf.tmpl"
var FileProvisionedHostTfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x94\x90\x31\x6b\xf3\x30\x10\x86\x77\xfd\x8a\x17\x93\x39\x81\x40\xe0\xfb\x86\x0c\x2d\x1d\x9a\xa5\x53\xa1\xa3\x11\xd6\x99\x88\xd8\xba\xa0\xbb\xc4\x04\xa3\xff\x5e\x2c\x25\x6e\x0c\x5d\xea\xed\x9e\xf7\xfc\xde\x83\x1a\x0e\x81\x1a\xf5\x1c\x50\x8d\x23\x56\xeb\x4f\xb2\xfd\xfa\xf0\x86\x94\x36\x81\x74\xe0\x78\x92\x4d\x0e\x3e\xca\xb4\x7e\xb5\x42\x53\x7a\x64\xd1\x7b\xf4\xce\xa2\x33\x9f\x1a\x2b\x8c\x06\xb0\x8d\xfa\x2b\xe1\xf1\xed\x51\xad\xc6\xe9\xaf\xba\x04\xa9\x32\x40\xa4\x9e\x95\x6a\xeb\x5c\x9c\x77\x9e\x58\xde\xe9\xb8\xb1\xdd\xbc\x52\x76\x7e\xd8\xbd\x46\xf8\x12\x1b\xaa\x83\xed\x69\xae\x79\x62\xa9\x32\xc6\x00\xe3\x08\xdf\x3e\x8c\x0f\xf2\xe5\x83\xe3\x41\x90\x92\x01\x06\x1f\x62\x9f\xc5\x97\x5a\xbf\x3a\x01\x67\x8e\x8a\x3d\x76\xff\xff\xed\xf2\x7c\x54\x3d\x0b\xf6\x68\x6d\x27\x94\x89\x9c\xfc\xb9\xbe\x52\xf4\xed\x6d\xc1\x2f\x42\xb9\xf5\xc5\xf5\x3e\x78\xd1\x68\x95\xe3\xbd\xd4\x8a\x0c\x1c\x5d\x39\xfa\x98\xf2\xc5\x54\xf4\xa9\xcb\xcf\x6c\x00\x91\xe3\x9f\x6d\xb7\xdb\x85\x41\x64\xd6\x92\x7b\x47\x41\xbd\xde\xea\xd6\x77\x54\x4a\x16\x68\xa1\x10\xdc\x64\x90\xbe\x03\x00\x00\xff\xff\x91\x4b\xc8\x07\x3c\x02\x00\x00")

// FileRemoteFileTfTmpl is "remote_file.tf.tmpl"
var FileRemoteFileTfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xcc\x94\xcf\x6e\xd4\x30\x10\xc6\xef\x79\x8a\x91\xd5\x03\x48\x95\xcb\x89\x03\xd2\x1e\x4a\x11\x7f\x2e\x08\xf5\xc2\x01\x21\xcb\xac\x67\xb3\x23\x12\x4f\xe4\x99\x74\xa9\xa2\xbc\x3b\xb2\xb3\x49\x53\x11\xe0\x42\x25\x7c\x8a\xf5\xf9\x9b\xcc\xfc\xc6\x9e\x61\x80\x80\x07\x8a\x08\x26\x61\xcb\x8a\xee\x40\x0d\x1a\x18\xc7\x2a\xa1\x70\x9f\xf6\x08\x26\xf6\x4d\xe3\xe6\xad\x01\xd3\x25\xbe\x23\x21\x8e\x6e\x18\xc0\xbe\x43\xcd\xe6\x49\x75\xd1\xb7\xc5\xee\x44\xb1\x5b\xe4\xbc\x71\xb1\x6f\xbf\x61\x2a\x62\xdf\x35\xec\x83\x81\xa1\x02\xd0\x44\x75\x8d\x49\xca\x06\x80\xa2\xa8\x8f\x7b\x74\x14\x60\x07\xe6\x62\xa8\x99\xeb\x06\xdd\x9e\xdb\xae\x57\x74\xb3\x6e\x7f\xff\x6f\xbb\x8a\x31\x9a\x0a\x60\xac\x2a\x80\x80\x1d\xc6\x20\x8e\x23\xec\xe0\x4b\xf9\x97\x59\x62\x3c\x88\x39\x40\xf6\x7c\xcd\x9e\x61\x80\x8b\x0c\x24\x47\x86\x57\x3b\xb0\xb7\x05\xd2\x5b\x6a\xd0\x5e\x8b\xa0\x7e\xcc\xc2\x58\xe2\x2f\x54\x30\x81\x99\x28\x4e\x15\x0d\x03\xd0\x01\xec\x7b\x16\xb5\x1f\xe4\x33\xc5\xc0\x27\xc9\x26\x28\x6b\xcf\x31\xe2\x5e\x89\xe3\xf9\x7c\x5e\x47\x16\x2d\x1f\x5b\x08\x7c\x08\x09\x45\xfe\x44\xe0\x7c\xa4\x54\x32\x2d\xbd\xef\x70\x8e\x78\xa2\x98\xda\x07\xa9\x17\x4c\xb3\x74\x1d\x5a\x8a\x24\x9a\xbc\x72\x5a\xb9\xa9\x45\xee\xb5\x1c\x79\xf9\x62\xe5\xed\xbc\xc8\x89\x53\xe9\x55\xce\xe7\x86\xdb\x0e\x95\x72\x39\xf6\x96\x59\x3f\xcd\xfa\xb8\xe4\x32\xce\x54\xb0\x11\xfc\x0b\x07\x5f\x63\xd4\xf3\xf7\x0e\xcc\xc1\x37\x82\xe6\x57\x4a\x4f\x02\x6a\x0a\x2a\x72\xdc\x20\x35\x69\x89\x59\x37\x18\xc1\x06\xa6\x44\x77\x5e\xd1\x7d\xc7\xfb\x29\xd3\x7c\x41\x9e\x65\x62\x14\x03\xfe\x00\xfb\xba\xa7\x26\xd8\x1b\x8e\x07\xaa\x73\x9e\x8d\x13\x39\xba\x95\x6d\x79\x98\xe6\xf9\x06\xc9\x18\x66\x90\xe7\x47\x3b\x67\x61\xed\x95\xb5\x57\xc1\xab\xbf\x7a\x74\x99\xe7\x76\x04\x14\xa5\xe8\x0b\xf6\x73\x0b\x57\x77\xfc\xcd\x4a\xcd\x8e\xf2\x96\xc6\xea\x29\x67\xc3\xff\x3b\x14\x2e\x27\xf9\x51\xc5\xf6\x9f\xcd\xc2\xcb\x32\x73\xc6\xaa\x5a\xda\xf9\x33\x00\x00\xff\xff\xbf\xe6\x1b\x66\x9c\x05\x00\x00")

// FileRootModuleTfTmpl is "root_module.tf.tmpl"
var FileRootModuleTfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x4c\xca\x3b\x0a\x43\x21\x10\x05\xd0\x7e\x56\x71\x11\xcb\x44\xfb\x80\x4d\xd6\x90\x0d\x48\x1c\x82\xe0\x07\xfc\x54\xc3\xec\x3d\xf8\xaa\x57\x1e\x38\x22\x18\xb1\xfd\x18\x76\xe5\xf4\x80\x5d\x1c\x2b\x5e\x01\xd6\xbd\x77\x2e\xc9\x7d\x38\xd6\x09\x55\xaa\x3d\xed\xc2\x30\x27\x3c\x45\xae\x0f\x55\x03\x21\x60\xf6\x3d\xbe\x8c\x00\xe3\xfc\x09\xd3\xdf\x06\x29\x91\x08\xb8\x1d\xfd\x03\x00\x00\xff\xff\x66\x90\xa0\x30\x70\x00\x00\x00")

// FileScriptTfTmpl is "script.tf.tmpl"
var FileScriptTfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xec\x57\x4b\x8f\xdb\x36\x10\xbe\xeb\x57\x0c\x88\x14\x48\x50\x5b\xdb\xa0\x68\x0f\x01\xf6\x90\x5d\xf4\x05\x04\xc5\xa2\x3e\xf4\x50\x07\x04\x23\x8d\x65\xa2\x14\xa9\x72\xa8\xd8\x0b\x45\xff\xbd\x18\xca\x92\xe5\x67\xd2\xee\x23\x41\xb0\x3c\x18\xa6\x87\xc3\x99\xf9\xf8\xf1\xe3\xb8\x69\x20\xc7\x85\xb6\x08\x82\x32\xaf\xab\x20\xa0\x6d\x13\x8f\xe4\x6a\x9f\x21\x08\x5b\x1b\x23\xfb\xa9\x00\x51\x79\xf7\x5e\x93\x76\x56\x36\x0d\xa4\xbf\x60\x00\xd1\x5b\xa5\x55\x25\xb2\xbb\xa4\x80\xd5\x60\xe6\x89\xb4\x75\xf9\x0e\x7d\x34\xd6\x95\x71\x2a\x17\xd0\x24\x00\xc1\xeb\xa2\x40\x4f\x71\x02\xa0\x2d\x05\x65\x33\x94\x3a\x87\x4b\x10\xcf\x9a\xc2\xb9\xc2\xa0\xcc\x5c\x59\xd5\x01\x65\x6f\x4f\x4f\xc7\x4e\x47\x7b\xb4\x22\x01\x68\x93\x04\x20\xc7\x0a\x6d\x4e\xd2\x59\xb8\x84\xbf\x62\x2c\x31\xec\xb1\x35\xf2\x06\xec\xf3\x96\x7d\x86\x4a\xd1\x83\x58\x68\x83\x62\x93\x65\xd3\x80\x5e\x40\xfa\xab\xa3\x90\xfe\x46\x7f\x6a\x9b\xbb\x15\x31\x6a\x10\x47\xe6\xac\xc5\x2c\x68\x67\x37\xeb\x79\x2c\x1d\x85\xf8\xe5\x58\x59\x2a\xcf\x3d\x12\x9d\xab\x6a\xb3\x24\x66\xd7\x8d\x70\x5b\x61\xbf\xe3\x4a\x5b\x5f\x6e\x4d\x35\xa1\xef\x4d\xaf\xf3\x52\x5b\x4d\xc1\xab\xe0\xfc\xc8\x5b\x97\xe8\xea\x10\x97\xfc\xf8\xdd\xc8\xb7\x52\x44\x2b\xe7\x23\xfe\x9c\xcf\xb5\x2b\x2b\x0c\x9a\xcb\x49\xff\x70\x2e\xdc\xf4\xf6\x76\xc8\x25\x22\xcc\x63\xc3\x99\x3e\x34\xbb\x47\x90\xf8\x83\x2b\x81\xb6\xbd\x50\x44\x18\xe8\x82\x6d\xb3\xc8\xb7\xf4\x4a\x11\x8e\x76\xeb\xd0\xc5\x7f\x06\xfb\x1b\x65\x8b\x5a\x15\x4c\xd0\x25\x1a\x23\xb6\x48\xf3\xb9\x52\xd0\x56\x45\xb0\xbb\x88\x95\xd7\x36\x2c\x40\x5c\xbf\x9a\xcf\xe7\x73\xa3\x16\xce\x17\x28\x3b\x6a\xcb\x6f\x28\x7d\xa7\x82\x80\xe7\x87\xcc\x7c\xb1\x9b\x02\x9a\x98\xd4\x1d\x02\x55\xf4\xf2\x93\x02\xd9\xbc\x8f\x73\x10\xb6\x69\xe0\x59\x4e\x81\xb9\x07\xaf\x2e\x87\x88\x17\xde\xb9\x70\x71\x18\x91\x96\x27\x03\x9e\xa3\xa6\x2a\xd0\x06\x18\x4e\x6d\xa1\x0c\xa1\x38\x24\xee\x83\x70\xb7\xdb\x94\x68\x79\x84\xbc\x9d\x8d\x6b\x3d\x42\x5b\x38\xc2\x5c\xaf\xdf\xab\x80\xf2\x6f\xbc\xed\x32\x65\xdc\x9e\xf3\x51\x69\x9b\xe3\x1a\xd2\xab\x5a\x9b\x3c\xbd\x76\x76\xa1\x0b\xce\xd3\x48\xa2\xa5\x1c\xb9\xc9\xee\x96\xb7\xad\x78\x31\x22\xf7\x3d\x72\xfb\x90\x45\xc3\xf9\xf6\x8b\xc6\x94\x68\x93\x36\x79\x50\x29\xc6\x35\x66\x9f\x5f\x88\x77\xea\x4a\xef\xed\x81\x99\x1c\x15\x72\x8f\xa5\x0b\x38\xdd\x56\xfe\xa4\xe7\x27\xf5\x5c\x5b\xc3\x8d\x41\x7f\x50\xff\x59\x9d\xef\xac\xc7\x93\x71\xdc\x7d\x49\x16\x95\x5b\xa1\x8f\x51\x61\xfa\xbb\xbb\xf1\x2e\x5e\xa4\xe9\x4f\x6b\xcc\x6a\xae\xef\xc6\x19\x9d\xdd\xc2\xd5\x2d\x63\x00\xd3\x9f\xd9\x3a\xbf\xab\x74\xcf\xc5\x5e\x56\x5b\xfd\x66\xba\x3d\xe9\xf8\x17\xab\xe3\xa7\x49\x2d\xb2\x65\xe9\x72\xf8\x76\x0d\x7b\x8a\x3c\xd9\xa1\xf2\xc8\x32\x92\xf8\xd7\xbe\x98\x05\xaf\x6d\x31\x76\x78\xfb\xf8\x5a\xbe\x52\x3a\x7c\x7d\x5a\x1e\x75\x7a\x50\x72\x3e\x04\xa3\x6d\xbd\x96\x64\x10\x2b\xc9\x34\xf2\x7c\xa5\xbe\xef\x40\x66\xfb\xaa\xd3\xef\xfd\x15\x2f\x7f\xe0\x25\xc9\xa0\xf6\x6f\x5c\xa6\xcc\xa1\xdc\xef\xbc\x15\x86\xd7\xec\x3c\x15\x1f\x7d\x2c\xf8\x6e\x96\xa5\xb2\x11\xea\x59\x50\x3e\x4c\x67\x9c\xc8\xc9\xd4\xce\x76\x9d\x67\xf6\x3a\x84\xe1\x54\x5b\xc9\x0c\x08\xe8\x2b\x8f\x01\x3d\x9f\x94\xb8\x61\xe1\x9c\x45\xb9\x9e\x80\x98\x5e\x77\x51\x44\x47\xda\x0d\x8e\xe3\x54\xee\x15\x14\xba\x17\x38\xe8\x7f\x00\xd1\x57\xd6\xfd\xf2\xb0\xd7\xf1\x71\x6e\xe2\xe4\x31\xaf\x62\x14\x98\xa7\xa6\xea\x73\x37\x55\x00\x22\x47\x03\x0f\xda\x5a\x7d\x7a\x8c\xb3\xff\x72\x9f\x1a\xa5\xaf\xad\x51\xf2\x25\x4c\x17\x7b\x5d\x12\x7c\xf8\x00\xc1\xd7\xf8\x91\xe6\x67\x3b\x4f\xfe\x0d\x00\x00\xff\xff\x3e\x2f\xbd\x02\x74\x14\x00\x00")

func init() {
	err := CTX.Err()
	if err != nil {
		panic(err)
	}

	var f webdav.File

	var rb *bytes.Reader
	var r *gzip.Reader

	rb = bytes.NewReader(FileCommandTfTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "command.tf.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileDNSRecordTfTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "dns_record.tf.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileInfraTfTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "infra.tf.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileProvisionedHostTfTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "provisioned_host.tf.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileRemoteFileTfTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "remote_file.tf.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileRootModuleTfTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "root_module.tf.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileScriptTfTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "script.tf.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	Handler = &webdav.Handler{
		FileSystem: FS,
		LockSystem: webdav.NewMemLS(),
	}

}

// Open a file
func (hfs *HTTPFS) Open(path string) (http.File, error) {
	path = hfs.Prefix + path

	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// ReadFile is adapTed from ioutil
func ReadFile(path string) ([]byte, error) {
	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))

	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(f)
	return buf.Bytes(), err
}

// WriteFile is adapTed from ioutil
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := FS.OpenFile(CTX, filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// WalkDirs looks for files in the given dir and returns a list of files in it
// usage for all files in the b0x: WalkDirs("", false)
func WalkDirs(name string, includeDirsInList bool, files ...string) ([]string, error) {
	f, err := FS.OpenFile(CTX, name, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	fileInfos, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}

	err = f.Close()
	if err != nil {
		return nil, err
	}

	for _, info := range fileInfos {
		filename := path.Join(name, info.Name())

		if includeDirsInList || !info.IsDir() {
			files = append(files, filename)
		}

		if info.IsDir() {
			files, err = WalkDirs(filename, includeDirsInList, files...)
			if err != nil {
				return nil, err
			}
		}
	}

	return files, nil
}
