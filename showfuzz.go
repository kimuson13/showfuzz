package showfuzz

import (
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "showfuzz is the tool that analyze functions can do fuzz test"

// Analyzer is checking the function whether do fuzz test
var Analyzer = &analysis.Analyzer{
	Name: "showfuzz",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.Ident)(nil),
	}

	if strings.HasSuffix(pass.Pkg.Name(), "_test") {
		return nil, nil
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.Ident:
			if n.Obj != nil {
				if fd, ok := n.Obj.Decl.(*ast.FuncDecl); ok {
					if len(fd.Type.Params.List) != 0 {
						for _, l := range fd.Type.Params.List {
							switch t := l.Type.(type) {
							case *ast.ArrayType:
								if !isFuzzable(pass.TypesInfo.TypeOf(t.Elt).Underlying()) {
									return
								}
							case *ast.Ident:
								if !isFuzzable(pass.TypesInfo.TypeOf(l.Type).Underlying()) {
									return
								}
							}
						}

						pass.Reportf(n.Pos(), "%s can fuzz test", fd.Name)
					}
				}
			}
		}
	})

	return nil, nil
}

func isFuzzable(typ types.Type) bool {
	ns := []types.BasicKind{types.Int, types.Int8, types.Int16, types.Int32, types.Int64, types.Uint, types.Uint8, types.Uint16, types.Uint32, types.Uint64, types.Float32, types.Float64, types.Bool, types.String}
	fuzzableTypes := make([]*types.Basic, len(ns))
	for i, v := range ns {
		fuzzableTypes[i] = types.Typ[v]
	}

	for _, v := range fuzzableTypes {
		if types.Identical(typ, v) {
			return true
		}
	}

	return false
}
