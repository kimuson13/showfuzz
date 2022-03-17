package showfuzz

import (
	"go/ast"
	"go/types"
	"reflect"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "showfuzz is the tool that analyze functions can do fuzz test"

type Results struct {
	events []event
}

type event struct {
	Name string
	Args []tp
}

type tp struct {
	TypName        string
	UnderlyingName string
	IsArr          bool
}

// Analyzer is checking the function whether do fuzz test
var Analyzer = &analysis.Analyzer{
	Name: "showfuzz",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
	ResultType: reflect.TypeOf(new(Results)),
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.Ident)(nil),
	}

	if strings.HasSuffix(pass.Pkg.Name(), "_test") {
		return nil, nil
	}

	results := &Results{}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		e := event{}
		switch n := n.(type) {
		case *ast.Ident:
			if n.Obj != nil {
				if fd, ok := n.Obj.Decl.(*ast.FuncDecl); ok {
					if len(fd.Type.Params.List) != 0 {
						e.Name = fd.Name.Name
						for _, l := range fd.Type.Params.List {
							switch t := l.Type.(type) {
							case *ast.ArrayType:
								typ := pass.TypesInfo.TypeOf(t.Elt)
								if !isFuzzable(typ.Underlying()) {
									return
								}

								e.Args = append(e.Args, tp{typ.String(), typ.Underlying().String(), true})
							case *ast.Ident:
								typ := pass.TypesInfo.TypeOf(l.Type)
								if !isFuzzable(typ.Underlying()) {
									return
								}

								e.Args = append(e.Args, tp{typ.String(), typ.Underlying().String(), false})
							}
						}

						pass.Reportf(n.Pos(), "%s can fuzz test", fd.Name)
						results.events = append(results.events, e)
					}
				}
			}
		}
	})

	return results, nil
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
