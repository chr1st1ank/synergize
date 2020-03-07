package synergize

import (
	"fmt"
	"github.com/chr1st1ank/synergize/internal/fs"
)

func SkipOrSynergizeFile(fi <-chan fs.FileOnDiskInfo) {
	filesByHash := make(map[string]fs.FileOnDiskInfo, 100)
	for nextFile := range fi {
		hash, err := nextFile.FileInfo.Hash()
		if err != nil {
			fmt.Printf("Skipping file %v because of error %v", nextFile.FileInfo.Path(), err)
		}
		otherFile, known := filesByHash[hash]
		if !known {
			fmt.Println(nextFile.FileInfo.Path(), hash, " (new)")
			filesByHash[hash] = nextFile
		} else {
			fmt.Println(nextFile.FileInfo.Path(), hash, " (duplicate of ", otherFile.FileInfo.Path(), ")")
		}
	}
}
