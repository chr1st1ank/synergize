package fs

import (
	"crypto/md5"
	"fmt"
	"hash"
	"hash/adler32"
	"io"
	"strings"
)

func hashFile(path string, hasher hash.Hash) string {
	input := strings.NewReader(path)
	if _, err := io.Copy(hasher, input); err != nil {
		return "Error"
	}
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func Md5Sum(path string) string {
	return hashFile(path, md5.New())
}

func Adler32Sum(path string) string {
	return hashFile(path, adler32.New())
}
