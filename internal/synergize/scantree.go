package synergize

import (
	"fmt"
	"github.com/chr1st1ank/synergize/internal/fs"
	"io/ioutil"
	"path/filepath"
)

// ScanTree traverses a folder recursively
func ScanTree(folder string, bufferSize int) <-chan fs.FileInfo {
	files, ok := ioutil.ReadDir(folder)
	if ok != nil {
		fmt.Println("Can't read folder")
		return nil
	}
	out := make(chan fs.FileInfo, bufferSize)
	go func() {
		for _, f := range files {
			fi := fs.NewFileInfo(f, filepath.Join(folder, f.Name()))
			if f.IsDir() {
				for subfolderItem := range ScanTree(fi.Path(), bufferSize) {
					out <- subfolderItem
				}
			} else {
				out <- fi
			}
		}
		close(out)
	}()
	return out
}
