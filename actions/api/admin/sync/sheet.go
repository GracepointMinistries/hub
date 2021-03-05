package sync

import (
	"net/http"

	"github.com/GracepointMinistries/hub/actions/render"
	"github.com/GracepointMinistries/hub/settings"
	"github.com/GracepointMinistries/hub/sync"
	"github.com/gobuffalo/buffalo"
)

// SheetResponse returns a synchronized Google sheet
// swagger:response adminSyncSheetResponse
type SheetResponse struct {
	//in:body
	Body SheetResponsePayload
}

// SheetResponsePayload contains the synchronized Google sheet
type SheetResponsePayload struct {
	Sheet string `json:"sheet,omitempty"`
}

// Initialize creates a new Google sheet used for synchronizing data
//
// swagger:route POST /api/v1/admin/sync admin initializeSync
// Returns a google sheet reference.
// responses:
//   200: adminSyncSheetResponse
//	 400: errorResponse
//	 401: errorResponse
//	 403: errorResponse
//	 422: errorResponse
//	 500: errorResponse
func Initialize(c buffalo.Context) error {
	sheet, err := sync.CreateSpreadsheet()
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	if err := settings.UpdateSheet(c, sheet); err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, render.JSON(&SheetResponsePayload{
		Sheet: sheet,
	}))
}

// Export dumps the user and group state to the stored data sheet
//
// swagger:route POST /api/v1/admin/sync/dump admin dumpSync
// Returns a google sheet reference.
// responses:
//   200: adminSyncSheetResponse
//	 400: errorResponse
//	 401: errorResponse
//	 403: errorResponse
//	 422: errorResponse
//	 500: errorResponse
func Export(c buffalo.Context) error {
	sheet := settings.Sheet()
	err := sync.DumpToSpreadsheet(sheet)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, render.JSON(&SheetResponsePayload{
		Sheet: sheet,
	}))
}
