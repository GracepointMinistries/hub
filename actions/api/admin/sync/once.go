package sync

import (
	"errors"
	"net/http"

	"github.com/GracepointMinistries/hub/actions/render"
	"github.com/GracepointMinistries/hub/settings"
	"github.com/GracepointMinistries/hub/sync"
	"github.com/gobuffalo/buffalo"
)

func urlForOnce(slug string) string {
	return settings.Host() + "/api/v1/admin/sync/once/" + slug
}

// OnceParameters documents the inbound parameters used
// for the Once endpoint
// swagger:parameters runSyncOnce
type OnceParameters struct {
	// in:path
	// required:true
	ID int `json:"id"`
}

// Once synchronizes the user and group state to the stored data sheet
// via a generated slug
//
// swagger:route POST /api/v1/admin/sync/once/{id} admin runSyncOnce
// Returns a google sheet reference.
// responses:
//   200: adminSyncOnceResponse
//	 400: errorResponse
//	 401: errorResponse
//	 403: errorResponse
//	 422: errorResponse
//	 500: errorResponse
func Once(c buffalo.Context) error {
	id := c.Param("slug")
	slug := settings.OneTimeSyncSlug()
	if id == "" || id != slug {
		return c.Error(http.StatusUnauthorized, errors.New("unauthorized"))
	}
	sheet := settings.Sheet()
	script := settings.Script()
	slug, err := settings.RegenerateOneTimeSyncSlug(c)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	err = sync.ExportToSpreadsheet(c, sheet, script, urlForOnce(slug))
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, render.JSON(&SheetResponsePayload{
		Sheet: sheet,
	}))
}
