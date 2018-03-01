package controller

import (
	"github.com/unrolled/render"
)

var r *render.Render

func init() {
	r = render.New(render.Options{
		Directory:     "views",
		IndentJSON:    true,
		Layout:        "container",
		IsDevelopment: true,
		Extensions:    []string{".html"},
	})
}
