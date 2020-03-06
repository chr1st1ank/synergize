package main

import (
	"flag"
	"fmt"
	"github.com/chr1st1ank/synergize/internal/fs"
)

var _ = fmt.Fprint
var _ = fs.NewFileInfo

func main() {
	flag.Parse()
	// targetFolder := flag.Arg(0)
	// foundFilesChannel := fs.ScanFolderTree(targetFolder, 10)
	// hashedFilesChannel := hashFileFromChannel(foundFilesChannel, 5)
	// storeOrSynergize(hashedFilesChannel)
}
