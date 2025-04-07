package unused

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/ak4bento/code-function-watcher/pkg/scanner"
)

func FindUnusedFunctions(projectPath string, defined []scanner.FunctionInfo, ignore map[string]struct{}) ([]scanner.FunctionInfo, error) {
	calledSet := make(map[string]struct{})

	err := filepath.Walk(projectPath, func(filePath string, info os.FileInfo, err error) error {
		if err != nil || !strings.HasSuffix(filePath, ".go") {
			return nil
		}

		fs := token.NewFileSet()
		node, err := parser.ParseFile(fs, filePath, nil, 0)
		if err != nil {
			return err
		}

		// Walk AST untuk mengumpulkan nama-nama fungsi yang dipanggil
		ast.Inspect(node, func(n ast.Node) bool {
			callExpr, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}
			switch fun := callExpr.Fun.(type) {
			case *ast.Ident:
				calledSet[fun.Name] = struct{}{}
			case *ast.SelectorExpr:
				// untuk fungsi seperti: fmt.Println atau obj.Method()
				calledSet[fun.Sel.Name] = struct{}{}
			}
			return true
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Filter fungsi yang tidak pernah dipanggil
	var unused []scanner.FunctionInfo
	// for _, fn := range functions {
	// 	if _, found := calledSet[fn.Name]; !found && fn.Name != "main" && fn.Name != "init" {
	// 		unused = append(unused, fn)
	// 	}
	// }

	for _, fn := range defined {
		if _, ignored := ignore[fn.Name]; ignored {
			continue
		}
		if _, found := calledSet[fn.Name]; !found && fn.Name != "main" && fn.Name != "init" {
			unused = append(unused, fn)
		}
	}

	return unused, nil

}
