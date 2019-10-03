package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/plush"
)

var r *render.Engine

func init() {
	r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "base.plush.html",

		// Box containing all of the templates:
		TemplatesBox: packr.New("../templates", "../templates"),

		// Add template helpers here:
		Helpers: render.Helpers{
			"isCurrentPath": func(name string, ctx plush.HelperContext) bool {
				if cr, ok := ctx.Value("current_route").(buffalo.RouteInfo); ok {
					return cr.PathName == name
				}
				return false
			},
		},
	})
}
