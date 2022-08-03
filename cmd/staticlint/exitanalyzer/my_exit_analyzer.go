package exitanalyzer

import (
	"fmt"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/analysis"
)

// ExitCheckAnalyzer notifies if we have an os.Exit() call in main.
var ExitCheckAnalyzer = &analysis.Analyzer{
	Name: "exitcheck",
	Doc:  "check for os.exit call in main",
	Run:  run,
}

// run inspects the AST to find the os.Exit() call.
func run(pass *analysis.Pass) (interface{}, error) {
	// создаём token.FileSet
	fset := token.NewFileSet()

	for _, file := range pass.Files {
		// функцией ast.Inspect проходим по всем узлам AST
		// запускаем инспектор, который рекурсивно обходит ветви AST
		// передаём инспектирующую функцию анонимно
		ast.Inspect(file, func(n ast.Node) bool {
			// проверяем, какой конкретный тип лежит в узле
			switch x := n.(type) {
			case *ast.File:
				if x.Name.Name == `main` { // package main check
					fmt.Printf("package: %v; pos: %v ", x.Name, fset.Position(x.Pos()))
					fmt.Println()
					for _, dec := range x.Decls {
						switch y := dec.(type) {
						case *ast.FuncDecl:
							// ast.FuncDecl представляет декларацию функции - ищем func main()
							se, ok := ExitCallCheck(pass, y, fset)
							if ok {
								pass.Reportf(se.Pos(), "os.Exit() call in main")
							}
						}
					}
				}
			}
			return true
		})
	}
	return nil, nil
}

func ExitCallCheck(pass *analysis.Pass, call *ast.FuncDecl, fset *token.FileSet) (ast.Expr, bool) {
	if call.Name.Name == `main` {
		fmt.Printf("func decl: %s, pos: %v", call.Name.Name, fset.Position(call.Pos()))
		//printer.Fprint(os.Stdout, fset, x)
		fmt.Println()
		for _, l := range call.Body.List {
			switch es := l.(type) {
			case *ast.ExprStmt:
				switch ec := es.X.(type) {
				case *ast.CallExpr:
					// ast.CallExpr вызов функции - ищем os.Exit()
					switch se := ec.Fun.(type) {
					case *ast.SelectorExpr:
						packageCall := fmt.Sprintf("%v", se.X)
						if se.Sel.Name == "Exit" && packageCall == "os" {
							//fmt.Printf("os.Exit() call in main, pos: %v", fset.Position(se.Pos()))
							return se, true
						}
					}
				}
			}
		}
	}
	return nil, false
}
