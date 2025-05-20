//go:build go1.12
// +build go1.12

package main

import (
	"flag"

	"github.com/yudppp/floatcheck"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/unitchecker"
)

// analyzers returns analyzers for floatcheck-comparison only.
func analyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		floatcheck.ComparisonAnalyzer,
	}
}

func main() {
	unitchecker.Main(analyzers()...)
}

func init() {
	flag.String("unsafeptr", "", "")
}
