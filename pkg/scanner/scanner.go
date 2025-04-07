package scanner

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

type FunctionInfo struct {
	Name     string `json:"name"`
	File     string `json:"file"`
	Position int    `json:"position"`
}

// Scan walks through all Go files in a directory and returns function info.
func Scan(path string) ([]FunctionInfo, error) {
	var functions []FunctionInfo
	fset := token.NewFileSet()

	err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(p) == ".go" {
			node, err := parser.ParseFile(fset, p, nil, parser.ParseComments)
			if err != nil {
				return err
			}

			for _, decl := range node.Decls {
				if fn, isFn := decl.(*ast.FuncDecl); isFn {
					pos := fset.Position(fn.Pos())
					functions = append(functions, FunctionInfo{
						Name:     fn.Name.Name,
						File:     p,
						Position: pos.Line,
					})
				}
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return functions, nil
}
