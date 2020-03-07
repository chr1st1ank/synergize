package fs

import (
	"fmt"
)

type InodeInfo interface {
	Address() string
	Fingerprint() string
}

type inodeInfoImp struct {
	fileSystem string // Identifier of the file system
	inodeID    uint64 // address of the inode on the file system
	size       int64  // length in bytes
	hash       string // some good hash digest as checksum
}

func InodeInfoFromFileInfo(fi FileInfo) (InodeInfo, error) {
	in := new(inodeInfoImp)
	in.fileSystem = fi.FileSystem()
	in.size = fi.Size()

	var err error
	in.inodeID, err = fi.Inode()
	if err != nil {
		return nil, err
	}

	in.hash, err = fi.Hash()
	if err != nil {
		return nil, err
	}

	return in, nil
}

func (inode *inodeInfoImp) String() string {
	return fmt.Sprint("Inode ", inode.Address(), inode.size, " ", inode.hash)
}

func (inode *inodeInfoImp) Address() string {
	return fmt.Sprint(inode.fileSystem, ":", inode.inodeID)
}

func (inode *inodeInfoImp) Fingerprint() string {
	return fmt.Sprint(inode.size, "-", inode.hash)
}
