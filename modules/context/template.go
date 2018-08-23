package context

import (
	"html/template"

	"github.com/kindlyfire/golog/modules/config"
)

func getTemplateFuncs() []template.FuncMap {
	return []template.FuncMap{map[string]interface{}{
		"ThAssetsUrl": func() string {
			return "/_/th-data"
		},
		"GlAssetsUrl": func() string {
			return "/_/gl-data"
		},
		"Debug": func() bool {
			return config.Debug
		},
		"BaseUrl": func() string {
			return config.BaseUrl
		},
	}}
}
