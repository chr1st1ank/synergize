package fs

import (
	"os"
	"syscall"
	"time"
)

type FileInfo interface {
	os.FileInfo
	Path() string
	FileSystem() string
	Inode() uint64
	Hash() string
}

type fileInfoImp struct {
	name       string
	size       int64
	mode       os.FileMode
	modTime    time.Time
	sys        syscall.Stat_t
	path       string
	filesystem string
	inode      uint64
	hash       string
}

func (fs *fileInfoImp) Name() string       { return fs.name }
func (fs *fileInfoImp) Size() int64        { return fs.size }
func (fs *fileInfoImp) Mode() os.FileMode  { return fs.mode }
func (fs *fileInfoImp) ModTime() time.Time { return fs.modTime }
func (fs *fileInfoImp) Sys() interface{}   { return &fs.sys }

func NewFileInfo(fi os.FileInfo, path string) FileInfo {
	f := new(fileInfoImp)
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

func (fi *fileInfoImp) Inode() uint64 {
	return fi.inode
}

func (fi *fileInfoImp) Hash() string {
	if len(fi.hash) == 0 {
		fi.hash, _ = Md5Sum(fi.Path())
	}
	return fi.hash
}
