package admin

import (
	"net/http"

	"github.com/GracepointMinistries/hub/actions/render"
	"github.com/GracepointMinistries/hub/models"
	"github.com/GracepointMinistries/hub/settings"
	"github.com/gobuffalo/buffalo"
)

// CurrentSettingsResponse returns a synchronized Google sheet
// swagger:response adminCurrentSettingsResponse
type CurrentSettingsResponse struct {
	//in:body
	Body *models.Setting
}

// CurrentSettings returns the current application settings
//
// swagger:route GET /api/v1/admin/settings admin currentSettings
// Returns the current application settings.
// responses:
//   200: adminCurrentSettingsResponse
//	 400: errorResponse
//	 401: errorResponse
//	 403: errorResponse
//	 422: errorResponse
//	 500: errorResponse
func CurrentSettings(c buffalo.Context) error {
	return c.Render(http.StatusOK, render.JSON(settings.SerializableSettings()))
}
