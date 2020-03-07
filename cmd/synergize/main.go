package main

import (
	"flag"
	"fmt"
	"github.com/chr1st1ank/synergize/internal/synergize"
)

var _ = fmt.Fprint

func main() {
	flag.Parse()
	targetFolder := flag.Arg(0)
	synergize.RunPipeline(targetFolder)
}
