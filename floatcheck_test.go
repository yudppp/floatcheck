package floatcheck_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/yudppp/floatcheck"

	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAllAnalyzer is a test for AllAnalyzer.
func TestAllAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, floatcheck.AllAnalyzer, "all")
}

// TestFormatAnalyzer is a test for FormatAnalyzer.
func TestFormatAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, floatcheck.FormatAnalyzer, "format")
}

// TestDivisionAnalyzer is a test for DivisionAnalyzer.
func TestDivisionAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, floatcheck.DivisionAnalyzer, "division")
}

// TestComparisonAnalyzer is a test for ComparisonAnalyzer.
func TestComparisonAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, floatcheck.ComparisonAnalyzer, "comparison")
}
