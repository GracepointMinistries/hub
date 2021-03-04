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
	"github.com/GracepointMinistries/hub/settings"
	"github.com/GracepointMinistries/hub/sync"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo-pop/pop/popmw"
	"github.com/gobuffalo/envy"
	csrf "github.com/gobuffalo/mw-csrf"
	forcessl "github.com/gobuffalo/mw-forcessl"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
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

var app *buffalo.App

// App initializes the buffalo application
func App() *buffalo.App {
	if app == nil {
		if err := settings.Initialize(); err != nil {
			fmt.Fprintf(os.Stderr, "Error initializing application settings: %v", err)
			os.Exit(1)
		}
		if err := sync.SetupClient(); err != nil {
			fmt.Fprintf(os.Stderr, "Error setting up google sheets client: %v", err)
			os.Exit(1)
		}

		app = buffalo.New(utils.SetSecureStore(buffalo.Options{
			Env:         string(currentEnvironment),
			SessionName: "_hub_session",
			PreWares:    []buffalo.PreWare{cors.Default().Handler},
		}))
		host := envy.Get("HOST", app.Host)
		gothic.Store = app.SessionStore
		adminProvider := google.New(os.Getenv("GOOGLE_OAUTH_KEY"), os.Getenv("GOOGLE_OAUTH_SECRET"), fmt.Sprintf(host+"/admin/auth/callback"))
		adminProvider.SetName("admin")
		goth.UseProviders(
			adminProvider,
			google.New(os.Getenv("GOOGLE_OAUTH_KEY"), os.Getenv("GOOGLE_OAUTH_SECRET"), fmt.Sprintf(host+"/auth/google/callback")),
		)

		// Automatically redirect to SSL
		app.Use(forcessl.Middleware(secure.Options{
			SSLRedirect:     currentEnvironment.IsDeployed(),
			SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
		}))
		app.Use(paramlogger.ParameterLogger)
		app.Use(popmw.Transaction(modelext.DB))
		app.Use(addHelpers)

		// Have API endpoints register themselves
		api.Register(app.Group("/api/v1"))
		// Register the rest of the "page" endpoints
		{
			pages := app.Group("")
			pages.Use(csrf.New)

			mainPages := app.Group("")
			mainPages.Use(middleware.RequireLoggedInUser)
			mainPages.GET("/", Profile)

			authPages := pages.Group("/auth")
			authPages.GET("/login", auth.Login)
			authPages.GET("/logout", middleware.RequireLoggedInUser(auth.Logout))
			authPages.GET("/{provider}", buffalo.WrapHandlerFunc(gothic.BeginAuthHandler))
			authPages.GET("/{provider}/callback", auth.Callback)

			adminPages := pages.Group("/admin")
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
