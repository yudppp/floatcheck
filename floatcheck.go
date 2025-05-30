package floatcheck

import (
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
)

// Analyzer for format check only
var FormatAnalyzer = &analysis.Analyzer{
	Name: "floatcheck_format",
	Doc:  "checks for potential floating-point precision issues in fmt.Sprintf formatting",
	Run: func(pass *analysis.Pass) (any, error) {
		if !strings.HasPrefix(pass.Pkg.Path(), "ng") {
			return nil, nil
		}
		inspect := func(node ast.Node) bool {
			checkFormat(pass, node)
			return true
		}
		for _, file := range pass.Files {
			ast.Inspect(file, inspect)
		}
		return nil, nil
	},
	// No analyzer requirements to avoid running on dependencies
	Requires: nil,
}

// Analyzer for division check only
var DivisionAnalyzer = &analysis.Analyzer{
	Name: "floatcheck_division",
	Doc:  "checks for potential floating-point division precision issues",
	Run: func(pass *analysis.Pass) (any, error) {
		if !strings.HasPrefix(pass.Pkg.Path(), "ng") {
			return nil, nil
		}
		inspect := func(node ast.Node) bool {
			checkDivision(pass, node)
			return true
		}
		for _, file := range pass.Files {
			ast.Inspect(file, inspect)
		}
		return nil, nil
	},
	// No analyzer requirements to avoid running on dependencies
	Requires: nil,
}

// Analyzer for comparison check only
var ComparisonAnalyzer = &analysis.Analyzer{
	Name: "floatcheck_comparison",
	Doc:  "checks for potential floating-point comparison issues",
	Run: func(pass *analysis.Pass) (any, error) {
		if !strings.HasPrefix(pass.Pkg.Path(), "ng") {
			return nil, nil
		}
		inspect := func(node ast.Node) bool {
			checkComparison(pass, node)
			return true
		}
		for _, file := range pass.Files {
			ast.Inspect(file, inspect)
		}
		return nil, nil
	},
	// No analyzer requirements to avoid running on dependencies
	Requires: nil,
}

// Analyzer for all checks
var AllAnalyzer = &analysis.Analyzer{
	Name: "floatcheck_all",
	Doc:  "checks for potential floating-point precision issues in fmt.Sprintf formatting, division, and comparison",
	Run: func(pass *analysis.Pass) (any, error) {
		if !strings.HasPrefix(pass.Pkg.Path(), "ng") {
			return nil, nil
		}
		inspect := func(node ast.Node) bool {
			checkFormat(pass, node)
			checkDivision(pass, node)
			checkComparison(pass, node)
			return true
		}
		for _, file := range pass.Files {
			ast.Inspect(file, inspect)
		}
		return nil, nil
	},
	// No analyzer requirements to avoid running on dependencies
	Requires: nil,
}

func checkFormat(pass *analysis.Pass, node ast.Node) {
	if call, ok := node.(*ast.CallExpr); ok {
		if fun, ok := call.Fun.(*ast.SelectorExpr); ok {
			if id, ok := fun.X.(*ast.Ident); ok && id.Name == "fmt" && fun.Sel.Name == "Sprintf" {
				if len(call.Args) >= 1 {
					if format, ok := call.Args[0].(*ast.BasicLit); ok && format.Kind.String() == "STRING" {
						if containsFloatSpecifier(format.Value) {
							pass.Reportf(format.Pos(), "potential floating-point precision issue in Sprintf format: %s", format.Value)
						}
					}
				}
			}
		}
	}
}

func checkDivision(pass *analysis.Pass, node ast.Node) {
	if binExpr, ok := node.(*ast.BinaryExpr); ok {
		if binExpr.Op == token.QUO {
			if isFloat(pass.TypesInfo.TypeOf(binExpr.X)) || isFloat(pass.TypesInfo.TypeOf(binExpr.Y)) {
				pass.Reportf(binExpr.Pos(), "potential floating-point division, consider potential precision issues")
			}
		}
	}
}

func checkComparison(pass *analysis.Pass, node ast.Node) {
	if binExpr, ok := node.(*ast.BinaryExpr); ok {
		switch binExpr.Op {
		case token.EQL, token.NEQ, token.LSS, token.LEQ, token.GTR, token.GEQ:
			if isFloat(pass.TypesInfo.TypeOf(binExpr.X)) || isFloat(pass.TypesInfo.TypeOf(binExpr.Y)) {
				pass.Reportf(binExpr.Pos(), "potential floating-point comparison, consider using a tolerance")
			}
		}
	}
}

func isFloat(t types.Type) bool {
	if t == nil {
		return false
	}
	_, ok := t.(*types.Basic)
	return ok && (t.String() == "float32" || t.String() == "float64")
}

func containsFloatSpecifier(s string) bool {
	s = strings.Trim(s, "`\"")
	for i := 0; i < len(s); i++ {
		if s[i] == '%' {
			j := i + 1
			for j < len(s) && strings.ContainsRune("#0- +0123456789.", rune(s[j])) {
				j++
			}
			if j < len(s) {
				switch s[j] {
				case 'f', 'F', 'g', 'G', 'e', 'E':
					return true
				}
			}
		}
	}
	return false
}
