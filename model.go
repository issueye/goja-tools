package main

type TypeNode struct {
	Name    string   `json:"name"`    // 名称
	Fields  []*Field `json:"fields"`  // 字段
	Func    []*Func  `json:"fns"`     // 方法
	ModName string   `json:"modName"` // 包名
}

type Field struct {
	Name string `json:"name"` // 名称
	Type string `json:"type"` // 类型
}

type Func struct {
	Name    string    `json:"name"`   // 方法名称
	Params  []*Param  `json:"params"` // 参数
	Results []*Result `json:"result"` // 返回值
}

type Result struct {
	Type string `json:"type"` // 类型
}

type Param struct {
	Name string `json:"name"` // 名称
	Type string `json:"type"` // 类型
}
