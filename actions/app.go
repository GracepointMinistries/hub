package actions

import (
	"fmt"
	"os"

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
	currentEnvironment = getEnvironment()
	admins             []string
)

func init() {
	currentEnvironment.Load()
	admins = getAdmins()
}

func getHost(server *buffalo.App) string {
	return envy.Get("HOST", server.Host)
}

func userIP(c buffalo.Context) string {
	request := c.Request()
	if ip := request.Header.Get("X-Real-Ip"); ip != "" {
		return ip
	}
	if ip := request.Header.Get("X-Forwarded-For"); ip != "" {
		return ip
	}
	return request.RemoteAddr
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
		app = buffalo.New(setSecureStore(buffalo.Options{
			Env:         string(currentEnvironment),
			SessionName: "_hub_session",
			PreWares:    []buffalo.PreWare{cors.Default().Handler},
		}))

		gothic.Store = app.SessionStore
		adminProvider := google.New(os.Getenv("GOOGLE_OAUTH_KEY"), os.Getenv("GOOGLE_OAUTH_SECRET"), fmt.Sprintf(getHost(app)+"/auth/admin/callback"))
		adminProvider.SetName("admin")
		goth.UseProviders(
			adminProvider,
			google.New(os.Getenv("GOOGLE_OAUTH_KEY"), os.Getenv("GOOGLE_OAUTH_SECRET"), fmt.Sprintf(getHost(app)+"/auth/google/callback")),
		)

		// Automatically redirect to SSL
		app.Use(forceSSL())

		// Log request parameters (filters apply).
		// app.Use(paramlogger.ParameterLogger)

		// TODO(dk): should probably add this, but makes API requests more complicated. If we enable it we need to also
		// uncomment the csrf meta tags in application.html.
		//
		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		// app.Use(csrf.New)

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.Connection)
		// Remove to disable this.
		app.Use(popmw.Transaction(modelext.DB))
		app.Use(addHelpers)

		{
			// api resources
			api := app.Group("/api/v1")
			setupErrorHandlers(api)

			api.POST("/exchange/admin", apiExchangeAdminToken)
			api.POST("/exchange/google", apiExchangeGoogleToken)
			api.POST("/exchange/facebook", apiExchangeFacebookToken)

			main := api.Group("")
			main.Use(requireAPIUser)
			main.GET("/profile", apiProfile)

			admin := api.Group("/admin")
			admin.Use(requireAPIAdmin)
			admin.GET("/users", apiAdminUsers)
			admin.GET("/zgroups", apiAdminZgroups)
			admin.GET("/zgroups/{id}", apiAdminZgroup)
		}

		{
			// page resources
			app.GET("/login", loginPage)
			app.GET("/admin/login", adminLoginPage)

			main := app.Group("")
			main.Use(requireLoggedInUser)
			main.GET("/", profilePage)
			main.GET("/logout", logoutPage)

			auth := app.Group("/auth")
			auth.GET("/{provider}", buffalo.WrapHandlerFunc(gothic.BeginAuthHandler))
			auth.GET("/admin/callback", adminCallback)
			auth.GET("/{provider}/callback", authCallback)

			admin := app.Group("/admin")
			admin.Use(requireAdmin)
			admin.GET("/", adminUsersPage)
			admin.GET("/zgroups", adminZgroupsPage)
			admin.GET("/zgroups/{id}", adminZgroupPage)
			admin.GET("/impersonate/{id}", adminImpersonatePage)
			admin.GET("/logout", adminLogoutPage)

			app.ServeFiles("/", assetsBox) // serve files from the public directory
		}
	}

	return app
}

// forceSSL will return a middleware that will redirect an incoming request
// if it is not HTTPS. "http://example.com" => "https://example.com".
// This middleware does **not** enable SSL. for your application. To do that
// we recommend using a proxy: https://gobuffalo.io/en/docs/proxy
// for more information: https://github.com/unrolled/secure/
func forceSSL() buffalo.MiddlewareFunc {
	return forcessl.Middleware(secure.Options{
		SSLRedirect:     currentEnvironment.IsDeployed(),
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}
