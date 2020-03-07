package fs

import (
	"testing"
)

const knownFileMd5Sum = "31d2a6aa73dcfa42d62da2a8d2dcefbd" // MD5 checksum for test/file_with_known_hash.txt
const knownFileAdler32Sum = "e74a21c6"                     // Adler32 checksum for test/file_with_known_hash.txt

// testHashFunction tries to hash the "file_with_known_hash.txt" test file and validated
// the checksum. Afterwards it tries to hash a non existing file and expects an error.
func testHashFunction(function HashFunction, expectedHash string, t *testing.T) {
	var returnedHash, err = function("testdata/file_with_known_hash.txt")
	if err != nil {
		t.Errorf("Unexpected error when hashing test file %v", err)
	}
	if returnedHash != expectedHash {
		t.Errorf("Hash function returned wrong checksum for test file! Correct: %v Returned: %v",
			expectedHash, returnedHash)
	}

	_, err = function("invalid_file")
	if _, ok := err.(error); !ok {
		t.Errorf("Hash function should have failed for invalid file")
	}
}

// TestMd5Sum tests the Md5Sum checksum function
func TestMd5Sum(t *testing.T) {
	testHashFunction(Md5Sum, knownFileMd5Sum, t)
}

// TestMd5Sum tests the Adler32 checksum function
func TestAdler32Sum(t *testing.T) {
	testHashFunction(Adler32Sum, knownFileAdler32Sum, t)
}
