package main

import (
	pongo2 "github.com/flosch/pongo2/v6"
)

var tl = `

package main

import (
	"github.com/dop251/goja"
)

func New{{ data.Name }}(rt *goja.Runtime, option *{{data.ModName}}.{{ data.Name }}) *goja.Object {
	o := rt.NewObject()
	
	{% for field in data.Fields -%}
	# {{ field.Name }} Setter
	o.Set("set{{ field.Name }}", func() {

	})

	# {{ field.Name }} Getter
	o.Set("get{{ field.Name }}", func() {

	})
	{% endfor %}

	{% for func in data.Func -%}
	# {{ func.Name }} method
	o.Set("{{ func.Name }}", func() {

	})

	{% endfor %}
	return o
}
`

func Generater(tl string, ctx pongo2.Context) (string, error) {
	t, err := pongo2.FromString(tl)
	if err != nil {
		return "", err
	}

	s, err := t.Execute(ctx)
	if err != nil {
		return "", err
	}

	return s, nil
}
