package main

import (
	"flag"
	"fmt"
	"github.com/chr1st1ank/synergize/internal/fs"
	"github.com/chr1st1ank/synergize/internal/synergize"
)

var _ = fmt.Fprint
var _ = fs.NewFileInfo

func main() {
	flag.Parse()
	targetFolder := flag.Arg(0)
	foundFilesChannel := synergize.ScanTree(targetFolder, 10)
	hashedFilesChannel := synergize.TakeFileFingerprint(foundFilesChannel, 5)
	synergize.SkipOrSynergizeFile(hashedFilesChannel)
}
