package api

import (
	"encoding/json"
	"net/http"

	"github.com/GracepointMinistries/hub/actions/api/admin"
	"github.com/GracepointMinistries/hub/actions/api/admin/sync"
	"github.com/GracepointMinistries/hub/actions/api/exchange"
	"github.com/GracepointMinistries/hub/actions/api/middleware"
	"github.com/gobuffalo/buffalo"
)

// ErrorResponse returns an api error
// swagger:response apiErrorResponse
type ErrorResponse struct {
	//in:body
	Body ErrorPayload
}

// ErrorPayload wraps errors in a well-defined api payload
type ErrorPayload struct {
	Error string `json:"error"`
}

func errorHandler(status int, err error, c buffalo.Context) error {
	if status == http.StatusInternalServerError {
		// log all unexpected errors
		c.Logger().Error(err)
	}
	response := c.Response()
	response.WriteHeader(status)
	return json.NewEncoder(response).Encode(&ErrorPayload{
		Error: err.Error(),
	})
}

func setupErrorHandlers(api *buffalo.App) {
	api.ErrorHandlers[400] = errorHandler
	api.ErrorHandlers[401] = errorHandler
	api.ErrorHandlers[403] = errorHandler
	api.ErrorHandlers[422] = errorHandler
	api.ErrorHandlers[500] = errorHandler
}

// Register registers the api endpoints with the app
func Register(app *buffalo.App) {
	setupErrorHandlers(app)

	userAPI := app.Group("")
	userAPI.Use(middleware.RequireUser)
	userAPI.GET("/profile", Profile)
	userAPI.DELETE("/logout", Logout)

	exchangeAPI := app.Group("/exchange")
	exchangeAPI.POST("/admin", exchange.AdminToken)
	exchangeAPI.POST("/google", exchange.GoogleToken)
	exchangeAPI.POST("/facebook", exchange.FacebookToken)

	adminAPI := app.Group("/admin")
	adminAPI.Use(middleware.RequireAdmin)
	adminAPI.GET("/users", admin.Users)
	adminAPI.GET("/users/{id}", admin.User)
	adminAPI.GET("/groups", admin.Groups)
	adminAPI.GET("/groups/{id}", admin.Group)
	adminAPI.GET("/settings", admin.CurrentSettings)
	adminAPI.POST("/sync", sync.Initialize)
	adminAPI.GET("/impersonate/{id}", admin.Impersonate)
}
