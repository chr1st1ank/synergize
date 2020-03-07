package synergize

import (
	"fmt"
	"github.com/chr1st1ank/synergize/internal/fs"
)

func TakeFileFingerprint(fi <-chan fs.FileInfo, bufferSize int) <-chan fs.FileOnDiskInfo {
	// inodeTable := make(map[string]fs.InodeInfo, 100)
	out := make(chan fs.FileOnDiskInfo, bufferSize)
	go func() {
		for fileInfo := range fi {
			inodeInfo, err := fs.InodeInfoFromFileInfo(fileInfo)
			if err != nil {
				fmt.Println("Skipping because inode inaccessible: ", fileInfo.Path())
			}
			out <- fs.FileOnDiskInfo{
				FileInfo:  fileInfo,
				InodeInfo: inodeInfo,
			}
		}
		close(out)
	}()
	return out
}
