package fs

import (
	"crypto/sha1"
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

// Sha1Sum calculates and returns the sha1 checksum for the file in the given path
func Sha1Sum(filePath string) (checksum string, err error) {
	return hashFile(filePath, sha1.New())
}

// Adler32Sum calculates and returns the adler32 checksum for the file in the given path
func Adler32Sum(filePath string) (checksum string, err error) {
	return hashFile(filePath, adler32.New())
}
