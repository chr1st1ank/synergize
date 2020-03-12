package fs

import (
	"errors"
	"fmt"
	"os"
	"syscall"
	"time"
)

// FileInfo groups information about a certain file
type FileInfo interface {
	Name() string
	Size() int64
	Mode() os.FileMode
	Path() string
	FileSystem() string
	Inode() (uint64, error)
	Fingerprint() (string, error) // Fingerprint of the files contents. Files with the same fingerpint hold the same data.
}

type fileInfoImp struct {
	osFileInfo os.FileInfo
	path       string
	filesystem string
	inode      uint64
}

func (fi *fileInfoImp) Name() string       { return fi.osFileInfo.Name() }
func (fi *fileInfoImp) Size() int64        { return fi.osFileInfo.Size() }
func (fi *fileInfoImp) Mode() os.FileMode  { return fi.osFileInfo.Mode() }
func (fi *fileInfoImp) Path() string       { return fi.path }
func (fi *fileInfoImp) FileSystem() string { return fi.filesystem }
func (fi *fileInfoImp) ModTime() time.Time { return fi.osFileInfo.ModTime() }
func (fi *fileInfoImp) IsDir() bool        { return fi.Mode().IsDir() }

// NewFileInfo creates a new FileInfo object for the file at a given path
func NewFileInfo(fi os.FileInfo, path string) FileInfo {
	f := new(fileInfoImp)
	f.osFileInfo = fi
	f.path = path
	return f
}

func (fi *fileInfoImp) Inode() (uint64, error) {
	stat, ok := fi.osFileInfo.Sys().(*syscall.Stat_t)
	if !ok {
		return 0, errors.New("Unable to extract inode")
	}
	return stat.Ino, nil
}

func (fi *fileInfoImp) Fingerprint() (string, error) {
	sha1, err := Sha1Sum(fi.Path())
	if err != nil {
		return "", err
	}
	adler, err := Adler32Sum(fi.Path())
	if err != nil {
		return "", err
	}
	return fmt.Sprint(fi.Size(), "-", sha1, "-", adler), nil
}
