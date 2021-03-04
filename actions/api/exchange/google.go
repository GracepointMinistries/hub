package exchange

import (
	"errors"
	"net/http"

	"github.com/GracepointMinistries/hub/actions/api/common"
	"github.com/GracepointMinistries/hub/actions/render"
	"github.com/GracepointMinistries/hub/actions/utils"
	"github.com/GracepointMinistries/hub/modelext"
	"github.com/gobuffalo/buffalo"
)

const googleEndpointProfile = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func exchangeGoogleAdminKey(c buffalo.Context, accessToken string) (string, error) {
	user, err := getThirdPartyUser(c, "google", googleEndpointProfile, accessToken)
	if err != nil {
		return "", err
	}
	for _, admin := range utils.GetAdmins() {
		if admin == user.Email {
			session, err := modelext.CreateAdminSession(c, admin, utils.UserIP(c))
			if err != nil {
				return "", err
			}
			return utils.GenerateAdminToken(user.Email, session.ID)
		}
	}
	return "", errors.New("user is not an admin")
}

func exchangeGoogleUserKey(c buffalo.Context, accessToken string) (string, error) {
	googleUser, err := getThirdPartyUser(c, "google", googleEndpointProfile, accessToken)
	if err != nil {
		return "", err
	}
	user, err := modelext.EnsureUserWithOauth(c, "google", googleUser.ID, googleUser.Name, googleUser.Email)
	if err != nil {
		return "", err
	}
	session, err := modelext.CreateUserSession(c, user, utils.UserIP(c))
	if err != nil {
		return "", err
	}
	return utils.GenerateUserToken(user.ID, session.ID)
}

// AdminToken exchanges a google auth token with a hub api token.
//
// swagger:route POST /api/v1/exchange/admin auth exchangeAdmin
// Exchanges a google authentication token with an admin api token.
// responses:
//   200: tokenResponse
//	 400: apiErrorResponse
//	 401: apiErrorResponse
//	 403: apiErrorResponse
//	 422: apiErrorResponse
//	 500: apiErrorResponse
func AdminToken(c buffalo.Context) error {
	payload := &common.TokenPayload{}
	if err := c.Bind(payload); err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	apiKey, err := exchangeGoogleAdminKey(c, payload.Token)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, render.JSON(&common.TokenPayload{
		Token: apiKey,
	}))
}

// GoogleToken exchanges a google auth token with a hub api token.
//
// swagger:route POST /api/v1/exchange/google auth exchangeGoogle
// Exchanges a google authentication token with an api token.
// responses:
//   200: tokenResponse
//	 400: errorResponse
//	 401: errorResponse
//	 403: errorResponse
//	 422: errorResponse
//	 500: errorResponse
func GoogleToken(c buffalo.Context) error {
	payload := &common.TokenPayload{}
	if err := c.Bind(payload); err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	apiKey, err := exchangeGoogleUserKey(c, payload.Token)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, render.JSON(&common.TokenPayload{
		Token: apiKey,
	}))
}
