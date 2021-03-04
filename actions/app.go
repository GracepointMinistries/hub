package actions

import (
	"fmt"
	"os"

	"github.com/GracepointMinistries/hub/actions/admin"
	"github.com/GracepointMinistries/hub/actions/api"
	"github.com/GracepointMinistries/hub/actions/auth"
	"github.com/GracepointMinistries/hub/actions/middleware"
	"github.com/GracepointMinistries/hub/actions/render"
	"github.com/GracepointMinistries/hub/actions/utils"
	"github.com/GracepointMinistries/hub/modelext"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo-pop/pop/popmw"
	"github.com/gobuffalo/envy"
	forcessl "github.com/gobuffalo/mw-forcessl"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/rs/cors"
	"github.com/unrolled/secure"
)

var (
	currentEnvironment = utils.GetEnvironment()
)

func init() {
	currentEnvironment.Load()
}

func getHost(server *buffalo.App) string {
	return envy.Get("HOST", server.Host)
}

var app *buffalo.App

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
//
// Routing, middleware, groups, etc... are declared TOP -> DOWN.
// This means if you add a middleware to `app` *after* declaring a
// group, that group will NOT have that new middleware. The same
// is true of resource declarations as well.
//
// It also means that routes are checked in the order they are declared.
// `ServeFiles` is a CATCH-ALL route, so it should always be
// placed last in the route declarations, as it will prevent routes
// declared after it to never be called.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(utils.SetSecureStore(buffalo.Options{
			Env:         string(currentEnvironment),
			SessionName: "_hub_session",
			PreWares:    []buffalo.PreWare{cors.Default().Handler},
		}))

		gothic.Store = app.SessionStore
		adminProvider := google.New(os.Getenv("GOOGLE_OAUTH_KEY"), os.Getenv("GOOGLE_OAUTH_SECRET"), fmt.Sprintf(getHost(app)+"/admin/auth/callback"))
		adminProvider.SetName("admin")
		goth.UseProviders(
			adminProvider,
			google.New(os.Getenv("GOOGLE_OAUTH_KEY"), os.Getenv("GOOGLE_OAUTH_SECRET"), fmt.Sprintf(getHost(app)+"/auth/google/callback")),
		)

		// Automatically redirect to SSL
		app.Use(forcessl.Middleware(secure.Options{
			SSLRedirect:     currentEnvironment.IsDeployed(),
			SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
		}))

		// Log request parameters (filters apply).
		// app.Use(paramlogger.ParameterLogger)

		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		// app.Use(csrf.New)

		app.Use(popmw.Transaction(modelext.DB))
		app.Use(addHelpers)
		{
			api.Register(app.Group("/api/v1"))
		}

		{
			mainPages := app.Group("")
			mainPages.Use(middleware.RequireLoggedInUser)
			mainPages.GET("/", Profile)

			authPages := app.Group("/auth")
			authPages.GET("/login", auth.Login)
			authPages.GET("/logout", middleware.RequireLoggedInUser(auth.Logout))
			authPages.GET("/{provider}", buffalo.WrapHandlerFunc(gothic.BeginAuthHandler))
			authPages.GET("/{provider}/callback", auth.Callback)

			adminPages := app.Group("/admin")
			adminAuthPages := adminPages.Group("/auth")
			adminAuthPages.GET("/login", admin.Login)
			adminAuthPages.GET("/callback", admin.Callback)
			adminAuthPages.GET("/logout", middleware.RequireAdmin(admin.Logout))
			adminProtectedPages := adminPages.Group("")
			adminProtectedPages.Use(middleware.RequireAdmin)
			adminProtectedPages.GET("/", admin.Users)
			adminProtectedPages.GET("/groups", admin.Groups)
			adminProtectedPages.GET("/groups/{id}", admin.Group)
			adminProtectedPages.GET("/impersonate/{id}", admin.Impersonate)

			app.ServeFiles("/", render.Assets) // serve files from the public directory
		}
	}

	return app
}
