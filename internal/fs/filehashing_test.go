package fs

import (
	"testing"
)

const knownFileSha1Sum = "44bc03ad6d33f3ea06e37cb92ac9b1ee86961a3d" // Sha1 checksum for test/file_with_known_hash.txt
const knownFileAdler32Sum = "e74a21c6"                              // Adler32 checksum for test/file_with_known_hash.txt

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

// TestSha1Sum tests the Sha1Sum checksum function
func TestSha1Sum(t *testing.T) {
	testHashFunction(Sha1Sum, knownFileSha1Sum, t)
}

// TestSha1Sum tests the Adler32 checksum function
func TestAdler32Sum(t *testing.T) {
	testHashFunction(Adler32Sum, knownFileAdler32Sum, t)
}
