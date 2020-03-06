package fs

import (
	"testing"
)

const knownFileMd5Sum = "31d2a6aa73dcfa42d62da2a8d2dcefbd" // MD5 checksum for test/file_with_known_hash.txt
const knownFileAdler32Sum = "e74a21c6"                     // Adler32 checksum for test/file_with_known_hash.txt

func TestMd5Sum(t *testing.T) {
	var returnedHash = Md5Sum("testdata/file_with_known_hash.txt")
	if returnedHash != knownFileMd5Sum {
		t.Errorf("Md5Sum returned wrong checksum for test file! Correct: %v Returned: %v",
			knownFileMd5Sum, returnedHash)
	}
}

func TestAdler32Sum(t *testing.T) {
	var returnedHash = Adler32Sum("testdata/file_with_known_hash.txt")
	if returnedHash != knownFileAdler32Sum {
		t.Errorf("Adler32Sum returned wrong checksum for test file! Correct: %v Returned: %v",
			knownFileAdler32Sum, returnedHash)
	}
}
