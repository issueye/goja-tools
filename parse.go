package main

import (
	"go/ast"
	"go/token"
	"unicode"
)

type GojaTools struct {
	Name string `json:"name"`
	Node *TypeNode
}

func NewGojaTools(name string, modName string) *GojaTools {
	tools := &GojaTools{
		Name: name,
		Node: new(TypeNode),
	}

	tools.Node.Name = name
	tools.Node.ModName = modName
	tools.Node.Fields = make([]*Field, 0)
	tools.Node.Func = make([]*Func, 0)

	return tools
}

func (tools *GojaTools) Parse(pkg *ast.File) {
	ast.Inspect(pkg, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.GenDecl:
			{
				tools.GenDecl(t)
			}
		case *ast.FuncDecl:
			{
				tools.FuncDecl(t)
			}
		}

		return true
	})
}

func (tools *GojaTools) GenDecl(decl *ast.GenDecl) {
	switch decl.Tok {
	case token.TYPE: // 类型
		{
			for _, spec := range decl.Specs {
				switch t := spec.(type) {
				case *ast.TypeSpec:
					{
					}
					tools.ParseType(t)
				}
			}
		}
	}
}

func (tools *GojaTools) ParseType(ts *ast.TypeSpec) {
	switch t := ts.Type.(type) {
	case *ast.StructType:
		{
			// 如果不是要找的内容则退出
			if tools.Name != ts.Name.Name {
				return
			}

			for _, field := range t.Fields.List {
				if len(field.Names) > 0 {
					name := field.Names[0].Name
					if !unicode.IsUpper(rune(name[0])) {
						continue
					}

					rtField := new(Field)
					rtField.Name = field.Names[0].Name
					switch ft := field.Type.(type) {
					case *ast.Ident:
						{
							rtField.Type = ft.Name
						}
					}

					tools.Node.Fields = append(tools.Node.Fields, rtField)
				}
			}
		}
	}
}
