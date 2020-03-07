package fs

import (
	"crypto/md5"
	"fmt"
	"hash"
	"hash/adler32"
	"io"
	"os"
)

// HashFunction is the type of a function which reads a file and calculates a checksum on it
type HashFunction func(filePath string) (checksum string, err error)

func hashFile(filePath string, hasher hash.Hash) (checksum string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}

// Md5Sum calculates and returns the md5 checksum for the file in the given path
func Md5Sum(filePath string) (checksum string, err error) {
	return hashFile(filePath, md5.New())
}

// Adler32Sum calculates and returns the adler32 checksum for the file in the given path
func Adler32Sum(filePath string) (checksum string, err error) {
	return hashFile(filePath, adler32.New())
}
