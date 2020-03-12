package synergize

import (
	"fmt"
	"github.com/chr1st1ank/synergize/internal/fs"
)

func TakeFingerprints(fi <-chan fs.FileInfo, bufferSize int) <-chan fs.FileOnDiskInfo {
	knownFingerprints := make(map[string]string, 100)
	out := make(chan fs.FileOnDiskInfo, bufferSize)
	go func() {
		for fileInfo := range fi {
			fileOnDiskInfo, err := fs.GenerateFileOnDiskInfo(fileInfo, knownFingerprints)
			if err != nil {
				fmt.Println("Skipping ", fileInfo.Path(), " because ", err.Error())
			} else {
				out <- *fileOnDiskInfo
			}
		}
		close(out)
	}()
	return out
}
