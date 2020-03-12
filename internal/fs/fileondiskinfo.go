package fs

import (
	"fmt"
)

// FileOnDiskInfo combines the information about a file, its filesystem storage place and its data
type FileOnDiskInfo struct {
	FileInfo    FileInfo
	InodeInfo   InodeInfo
	Fingerprint string // Fingerprint of the Inodes contents. Inodes with the same fingerpint hold the same data.
}

// GenerateFileOnDiskInfo calculates the fingerprint of the given FileInfo object and also collects inode information on it
func GenerateFileOnDiskInfo(fileInfo FileInfo, knownFingerprints map[string]string) (*FileOnDiskInfo, error) {
	inodeInfo, err := InodeInfoFromFileInfo(fileInfo)
	if err != nil {
		return nil, fmt.Errorf("Inode inaccessible")
	}
	fingerprint, err := lookupOrCalculateFingerprint(inodeInfo.Address(), fileInfo, knownFingerprints)
	if err != nil {
		return nil, err
	}
	return &FileOnDiskInfo{
		FileInfo:    fileInfo,
		InodeInfo:   inodeInfo,
		Fingerprint: fingerprint,
	}, nil
}

func lookupOrCalculateFingerprint(address string, fileInfo FileInfo, knownFingerprints map[string]string) (string, error) {
	fingerprint, known := knownFingerprints[address]
	if !known {
		fingerprint, err := fileInfo.Fingerprint()
		if err != nil {
			return "", fmt.Errorf("Fingerprint could not be generated")
		}
		knownFingerprints[address] = fingerprint
	}
	return fingerprint, nil
}
