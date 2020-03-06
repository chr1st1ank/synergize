package fs

import (
	"fmt"
	"os"
	"time"
)

type InodeInfo interface {
	Address() string
	Fingerprint() string
}

type inodeInfoImp struct {
	fileSystem string    // Identifier of the file system
	inodeID    uint64    // address of the inode on the file system
	size       int64     // length in bytes
	modTime    time.Time // modification time
	hash       string    // some good hash digest as checksum
}

func FromFileInfo(os.FileInfo) InodeInfo {
	return new(inodeInfoImp)
}

func (inode *inodeInfoImp) String() string {
	return fmt.Sprint("Inode ", inode.Address(), inode.modTime, inode.size, " ", inode.hash)
}

func (inode *inodeInfoImp) Address() string {
	return fmt.Sprint(inode.fileSystem, ":", inode.inodeID)
}

func (inode *inodeInfoImp) Fingerprint() string {
	return fmt.Sprint(inode.size, "-", inode.hash)
}
