package fs

import (
	"os"
	"syscall"
	"testing"
)

// TestFileInfo creates a FileInfo object from a test file and checks all known attributes
func TestFileInfo(t *testing.T) {
	stat, _ := os.Stat("testdata/file_with_known_hash.txt")
	fi := NewFileInfo(stat, "testdata/file_with_known_hash.txt")

	if fi.Name() != "file_with_known_hash.txt" {
		t.Errorf("Unexpected result of Name(): %v", fi.Name())
	}

	if fi.Size() != stat.Size() {
		t.Errorf("Unexpected result of Size(): %v", fi.Size())
	}

	if fi.Mode() != stat.Mode() {
		t.Errorf("Unexpected result of Mode(): %v", fi.Mode())
	}

	if fi.Path() != "testdata/file_with_known_hash.txt" {
		t.Errorf("Unexpected result of Path(): %v", fi.Path())
	}

	// if fi.FileSystem() != stat.Size() {
	// 	t.Errorf("Unexpected result of FileSystem(): %v", fi.FileSystem())
	// }

	sysStat, _ := stat.Sys().(*syscall.Stat_t)
	if inode, _ := fi.Inode(); inode != sysStat.Ino {
		t.Errorf("Unexpected result of Inode(): %v", inode)
	}

	if hash, _ := fi.Hash(); hash != knownFileMd5Sum {
		t.Errorf("Unexpected result of Hash(): %v", hash)
	}
}
