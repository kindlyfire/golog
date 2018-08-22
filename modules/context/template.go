package context

import (
	"html/template"
)

func getTemplateFuncs() []template.FuncMap {
	return []template.FuncMap{map[string]interface{}{
		"ThAssetsUrl": func() string {
			return "/_/th-data"
		},
		"GlAssetsUrl": func() string {
			return "/_/gl-data"
		},
	}}
}
