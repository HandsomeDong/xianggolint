package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"xianggolint/pkg/analyzers/loopgoroutinecheck"
)

func main() {
	singlechecker.Main(loopgoroutinecheck.Analyzer)
}
