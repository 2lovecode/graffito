package aop

import (
	"github.com/spf13/cobra"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{Use: "aop", Run: func(cmd *cobra.Command, args []string) {
		aa()
		//run()
	}}
}

func aa() {
	a := 23
	panic(&a)
}
func run() {
	src := `package main

type MyStruct struct {
	Field1 int
	Field2 string
}

func (m *MyStruct) Method1() {
	println(m.Field1)
}

func (m *MyStruct) Method2() {
	println(m.Field2)
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		log.Fatal(err)
	}

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			if x.Recv != nil && len(x.Recv.List) == 1 && x.Recv.List[0].Names[0].Name == "m" && x.Recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Name == "MyStruct" {

				logStmt := &ast.ExprStmt{
					X: &ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X:   ast.NewIdent("log"),
							Sel: ast.NewIdent("Println"),
						},
						Args: []ast.Expr{&ast.BasicLit{
							ValuePos: x.Pos(),
							Kind:     token.STRING,
							Value:    "\"Entering " + x.Name.Name + "\"",
						}},
					},
				}

				x.Body.List = append([]ast.Stmt{logStmt}, x.Body.List...)
				return false
			}

		}
		return true
	})

	printer.Fprint(os.Stdout, fset, f)
}
