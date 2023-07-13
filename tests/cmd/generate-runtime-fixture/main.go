package main

import (
	"flag"
	"path/filepath"

	"github.com/MrChaos1993/gen/tests/internal/runtimefixture"
)

func main() {
	outPath := flag.String("out", filepath.Join("fixture", "query"), "generated query output directory")
	flag.Parse()
	runtimefixture.Generate(*outPath)
}
