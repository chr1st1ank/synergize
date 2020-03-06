package fs

import (
	"crypto/md5"
	"fmt"
	"hash"
	"hash/adler32"
	"io"
	"os"
)

func hashFile(path string, hasher hash.Hash) string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if _, err := io.Copy(hasher, file); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func Md5Sum(path string) string {
	return hashFile(path, md5.New())
}

func Adler32Sum(path string) string {
	return hashFile(path, adler32.New())
}
