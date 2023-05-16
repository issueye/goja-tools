package main

import (
	"fmt"
	"go/parser"
	"go/token"

	"github.com/flosch/pongo2/v6"
)

func main() {
	ParseDemo()
}

func ParseDemo() {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseFile(fset, "repository.go_", nil, 0)
	// pkgs, err := parser.ParseFile(fset, "demo.go_", nil, 0)
	if err != nil {
		panic(err)
	}

	// ast.Print(fset, pkgs)
	tools := NewGojaTools("Repository", "git")
	tools.Parse(pkgs)

	// data, _ := json.Marshal(tools.Node)
	// fmt.Println(string(data))

	code, err := Generater(tl, pongo2.Context{
		"data": tools.Node,
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("code", code)
}
