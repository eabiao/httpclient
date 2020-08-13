package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

//go:generate go test -run TestGenAssets assets_test.go
func TestGenAssets(t *testing.T) {
	genAssets("main", "assets.go", "page")
}

// 生成内嵌资源
func genAssets(packageName, file string, dirs ...string) error {
	w, err := os.Create(file)
	if err != nil {
		return err
	}
	defer w.Close()
	fmt.Fprintf(w, `package %s

import (
	"bytes"
	"compress/zlib"
	"io"
	"os"
)

var Assets = &fs{}
var assets = map[string][]byte{}

type fs struct{}

// Get Data
func (fs *fs) Get(name string) ([]byte, error) {
	data, ok := assets[name]
	if !ok {
		return nil, os.ErrNotExist
	}
	return uncompress(data), nil
}

// uncompress
func uncompress(data []byte) []byte {
	var in bytes.Buffer
	in.Write(data)
	r, _ := zlib.NewReader(&in)
	defer r.Close()

	var out bytes.Buffer
	io.Copy(&out, r)
	return out.Bytes()
}

func init() {
`, packageName)

	defer fmt.Fprintln(w, "}")

	for _, dir := range dirs {
		filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			b := compress(data)

			path = filepath.ToSlash(path)
			fmt.Fprintf(w, "assets[%q] = []byte{\n", path)
			for i := 0; i < len(b); i++ {
				if i > 0 {
					if i%20 == 0 {
						fmt.Fprintf(w, ",\n")
					} else {
						fmt.Fprintf(w, ", ")
					}
				}
				fmt.Fprintf(w, "0x%02x", b[i])
			}
			fmt.Fprintln(w, "}")
			return nil
		})
	}
	return nil
}

// 压缩
func compress(data []byte) []byte {
	var buff bytes.Buffer
	w := zlib.NewWriter(&buff)
	defer w.Close()

	w.Write(data)
	w.Flush()
	return buff.Bytes()
}
