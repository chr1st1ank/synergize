package fs

import (
	"errors"
	"fmt"
	"os"
	"syscall"
	"time"
)

type FileInfo interface {
	Name() string
	Size() int64
	Mode() os.FileMode
	Path() string
	FileSystem() string
	Inode() (uint64, error)
	Hash() (string, error)
}

type fileInfoImp struct {
	osFileInfo os.FileInfo
	// sys        syscall.Stat_t
	path       string
	filesystem string
	inode      uint64
	hash       string
}

func (fs *fileInfoImp) Name() string       { return fs.osFileInfo.Name() }
func (fs *fileInfoImp) Size() int64        { return fs.osFileInfo.Size() }
func (fs *fileInfoImp) Mode() os.FileMode  { return fs.osFileInfo.Mode() }
func (fs *fileInfoImp) ModTime() time.Time { return fs.osFileInfo.ModTime() }

// func (fs *fileInfoImp) Sys() interface{}   { return &fs.osFileInfo.sys }

func NewFileInfo(fi os.FileInfo, path string) FileInfo {
	f := new(fileInfoImp)
	f.osFileInfo = fi
	f.path = path
	return f
}

func (fi *fileInfoImp) IsDir() bool {
	return fi.Mode().IsDir()
}

func (fi *fileInfoImp) Path() string {
	return fi.path
}

func (fi *fileInfoImp) FileSystem() string {
	return fi.filesystem
}

func (fi *fileInfoImp) Inode() (uint64, error) {
	stat, ok := fi.osFileInfo.Sys().(*syscall.Stat_t)
	if !ok {
		return 0, errors.New("Unable to extract inode")
	}
	return stat.Ino, nil
}

func (fi *fileInfoImp) Hash() (string, error) {
	if len(fi.hash) == 0 {
		hash, err := Adler32Sum(fi.Path())
		if err != nil {
			return "", err
		}
		fi.hash = fmt.Sprint(fi.Size(), "-", hash)
	}
	return fi.hash, nil
}
