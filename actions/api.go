package actions

import (
	"encoding/json"

	"github.com/gobuffalo/buffalo"
)

// APIErrorResponse returns an api error
// swagger:response apiErrorResponse
type APIErrorResponse struct {
	//in:body
	Body APIErrorPayload
}

// APIErrorPayload wraps errors in a well-defined api payload
type APIErrorPayload struct {
	Error string `json:"error"`
}

func apiErrorHandler(status int, err error, c buffalo.Context) error {
	response := c.Response()
	response.WriteHeader(status)
	return json.NewEncoder(response).Encode(&APIErrorPayload{
		Error: err.Error(),
	})
}

func setupErrorHandlers(api *buffalo.App) {
	api.ErrorHandlers[400] = apiErrorHandler
	api.ErrorHandlers[401] = apiErrorHandler
	api.ErrorHandlers[403] = apiErrorHandler
	api.ErrorHandlers[422] = apiErrorHandler
	api.ErrorHandlers[500] = apiErrorHandler
}
