package render

import (
	"github.com/gobuffalo/buffalo/render"
	packr "github.com/gobuffalo/packr/v2"
)

var r *render.Engine

// Assets contain static file assets
var Assets = packr.New("app:assets", "../../public")

func init() {
	r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.plush.html",

		// Box containing all of the templates:
		TemplatesBox: packr.New("app:templates", "../../templates"),
		AssetsBox:    Assets,

		// Add template helpers here:
		Helpers: render.Helpers{
			// for non-bootstrap form helpers uncomment the lines
			// below and import "github.com/gobuffalo/helpers/forms"
			// forms.FormKey:     forms.Form,
			// forms.FormForKey:  forms.FormFor,
		},
	})
}

// JSON renders the given payload as json
func JSON(v interface{}) render.Renderer {
	return r.JSON(v)
}

// HTML renders the given templates
func HTML(names ...string) render.Renderer {
	return r.HTML(names...)
}
