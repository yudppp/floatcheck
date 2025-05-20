//go:build !go1.12
// +build !go1.12

package main

import (
	"github.com/yudppp/floatcheck"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(floatcheck.AllAnalyzer) }
