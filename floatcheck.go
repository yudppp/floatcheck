package floatcheck

import (
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

const doc = "floatcheck is a static analysis tool that checks for potential floating-point precision issues in Go code"

// Analyzer is the main entry point for the floatcheck analysis tool.
// It inspects the AST for potential floating-point precision issues, such as
// floating-point formatting, division, and comparisons.
// It reports potential issues to the analysis pass.
// The analysis pass is expected to be run on Go source files, and it will
// report any potential issues found in the code.
var Analyzer = &analysis.Analyzer{
	Name: "floatcheck",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := func(node ast.Node) bool {
		// Detect fmt.Sprintf("%.nf", ...) pattern
		if call, ok := node.(*ast.CallExpr); ok {
			if fun, ok := call.Fun.(*ast.SelectorExpr); ok {
				if id, ok := fun.X.(*ast.Ident); ok && id.Name == "fmt" && fun.Sel.Name == "Sprintf" {
					if len(call.Args) >= 1 {
						if format, ok := call.Args[0].(*ast.BasicLit); ok && format.Kind.String() == "STRING" {
							if strings.Contains(format.Value, `%.`) && strings.Contains(format.Value, `f`) {
								pass.Reportf(format.Pos(), "potential floating-point precision issue in Sprintf format: %s", format.Value)
							}
						}
					}
				}
			}
		}

		// Detect floating-point division
		if binExpr, ok := node.(*ast.BinaryExpr); ok {
			if binExpr.Op == token.QUO {
				if isFloat(pass.TypesInfo.TypeOf(binExpr.X)) || isFloat(pass.TypesInfo.TypeOf(binExpr.Y)) {
					pass.Reportf(binExpr.Pos(), "potential floating-point division, consider potential precision issues")
				}
			}
		}

		// Detect floating-point equality and relational comparisons
		if binExpr, ok := node.(*ast.BinaryExpr); ok {
			switch binExpr.Op {
			case token.EQL, token.NEQ, token.LSS, token.LEQ, token.GTR, token.GEQ:
				if isFloat(pass.TypesInfo.TypeOf(binExpr.X)) || isFloat(pass.TypesInfo.TypeOf(binExpr.Y)) {
					pass.Reportf(binExpr.Pos(), "potential floating-point comparison, consider using a tolerance")
				}
			}
		}

		return true
	}
	for _, file := range pass.Files {
		ast.Inspect(file, inspect)
	}
	return nil, nil
}

func isFloat(t types.Type) bool {
	if t == nil {
		return false
	}
	_, ok := t.(*types.Basic)
	return ok && (t.String() == "float32" || t.String() == "float64")
}
