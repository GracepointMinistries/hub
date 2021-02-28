package actions

import (
	"fmt"
	"os"

	"github.com/GracepointMinistries/hub/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo-pop/pop/popmw"
	"github.com/gobuffalo/envy"
	forcessl "github.com/gobuffalo/mw-forcessl"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
	"github.com/golang/gddo/httputil"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/rs/cors"
	"github.com/unrolled/secure"
)

var currentEnvironment = getEnvironment()

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
		app = buffalo.New(setSecureStore(buffalo.Options{
			Env:         string(currentEnvironment),
			SessionName: "_hub_session",
			PreWares:    []buffalo.PreWare{cors.Default().Handler},
		}))

		gothic.Store = app.SessionStore
		goth.UseProviders(
			google.New(os.Getenv("GOOGLE_OAUTH_KEY"), os.Getenv("GOOGLE_OAUTH_SECRET"), fmt.Sprintf(getHost(app)+"/auth/google/callback")),
		)

		// Automatically redirect to SSL
		app.Use(forceSSL())

		// Log request parameters (filters apply).
		app.Use(paramlogger.ParameterLogger)

		// TODO(dk): should probably add this, but makes API requests more complicated. If we enable it we need to also
		// uncomment the csrf meta tags in application.html.
		//
		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		// app.Use(csrf.New)

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.Connection)
		// Remove to disable this.
		app.Use(popmw.Transaction(models.DB))

		app.GET("/login", loginHandler)

		main := app.Group("")
		main.Use(requireLoggedInUser)
		main.GET("/", profileHandler)
		main.GET("/logout", logoutHandler)

		auth := app.Group("/auth")
		auth.GET("/{provider}", buffalo.WrapHandlerFunc(gothic.BeginAuthHandler))
		auth.GET("/{provider}/callback", authCallback)

		app.ServeFiles("/", assetsBox) // serve files from the public directory
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

type acceptType int

const (
	acceptJSON acceptType = iota
	acceptHTML
)

func acceptContentType(c buffalo.Context) acceptType {
	val := httputil.NegotiateContentType(c.Request(), []string{"application/json", "text/html"}, "text/html")
	switch val {
	case "application/json":
		return acceptJSON
	case "text/html":
		return acceptHTML
	default:
		panic(fmt.Sprintf("got negotiated Accept header that makes no sense: %s", val))
	}
}