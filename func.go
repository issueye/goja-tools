package main

import (
	"go/ast"
	"unicode"
)

func (tools *GojaTools) FuncDecl(decl *ast.FuncDecl) {
	fnName := decl.Name.Name
	// 如果不是公共方法就不解析
	if !unicode.IsUpper(rune(fnName[0])) {
		return
	}

	// 判断方法是全局还是属于某个类型
	if tools.Name != "" {
		if decl.Recv != nil {
			name := tools.ParseRecv(decl.Recv)
			if tools.Name == name {
				fn := new(Func)
				fn.Name = fnName
				fn.Params = make([]*Param, 0)
				fn.Results = make([]*Result, 0)

				tools.FuncParams(decl.Type.Params, &fn.Params)
				tools.FuncResults(decl.Type.Results, &fn.Results)
				tools.Node.Func = append(tools.Node.Func, fn)
				return
			}
		}
	}
}

// 只获取类型名称
func (tools *GojaTools) ParseRecv(recv *ast.FieldList) string {
	// 有内容才往下走
	if len(recv.List) > 0 {
		switch t := recv.List[0].Type.(type) {
		case *ast.StarExpr:
			{
				i, ok := t.X.(*ast.Ident)
				if ok {
					return i.Name
				}
			}
		case *ast.Ident:
			{
				return t.Name
			}
		}
	}

	return ""
}

func (tools *GojaTools) FuncParams(fields *ast.FieldList, list *[]*Param) {
	for _, param := range fields.List {
		if len(param.Names) > 0 {
			rp := new(Param)
			// 参数名称
			rp.Name = param.Names[0].Name
			// 参数类型

			switch t := param.Type.(type) {
			case *ast.Ident:
				{
					rp.Type = t.Name
				}
			case *ast.StarExpr:
				{
					rp.Type = tools.StarExpr(t)
				}
			case *ast.SelectorExpr:
				{
					name, ok := t.X.(*ast.Ident)
					if ok {
						rp.Type = name.Name + "." + t.Sel.Name
					}
				}
			}
			ident, ok := param.Type.(*ast.Ident)
			if ok {
				rp.Type = ident.Name
			}

			*list = append(*list, rp)
		}
	}
}

func (tools *GojaTools) StarExpr(expr *ast.StarExpr) string {
	switch t := expr.X.(type) {
	case *ast.Ident:
		{
			return t.Name
		}
	case *ast.SelectorExpr:
		{
			ident, ok := t.X.(*ast.Ident)
			if ok {
				return ident.Name + "." + t.Sel.Name
			}
		}
	}
	return ""
}

func (tools *GojaTools) FuncResults(fields *ast.FieldList, list *[]*Result) {
	for _, result := range fields.List {
		rp := new(Result)
		// 参数类型
		switch t := result.Type.(type) {
		case *ast.Ident:
			{
				rp.Type = t.Name
			}
		case *ast.StarExpr:
			{
				rp.Type = tools.StarExpr(t)
			}
		case *ast.ArrayType:
			{
				expr, ok := t.Elt.(*ast.StarExpr)
				if ok {
					rp.Type = tools.StarExpr(expr)
				}
			}
		case *ast.SelectorExpr:
			{
				name, ok := t.X.(*ast.Ident)
				if ok {
					rp.Type = name.Name + "." + t.Sel.Name
				}
			}
		}
		*list = append(*list, rp)
	}
}
