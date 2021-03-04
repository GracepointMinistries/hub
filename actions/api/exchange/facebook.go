package exchange

import (
	"net/http"

	"github.com/GracepointMinistries/hub/actions/api/common"
	"github.com/GracepointMinistries/hub/actions/render"
	"github.com/GracepointMinistries/hub/actions/utils"
	"github.com/GracepointMinistries/hub/modelext"
	"github.com/gobuffalo/buffalo"
)

const facebookEndpointProfile = "https://graph.facebook.com/me?fields=email,id,name&access_token="

func exchangeFacebookUserKey(c buffalo.Context, accessToken string) (string, error) {
	fbUser, err := getThirdPartyUser(c, "facebook", facebookEndpointProfile, accessToken)
	if err != nil {
		return "", err
	}
	user, err := modelext.EnsureUserWithOauth(c, "facebook", fbUser.ID, fbUser.Name, fbUser.Email)
	if err != nil {
		return "", err
	}
	session, err := modelext.CreateUserSession(c, user, utils.UserIP(c))
	if err != nil {
		return "", err
	}
	return utils.GenerateUserToken(user.ID, session.ID)
}

// FacebookToken exchanges a facebook auth token with a hub api token.
//
// swagger:route POST /api/v1/exchange/facebook auth exchangeFacebook
// Exchanges a facebook authentication token with an api token.
// responses:
//   200: tokenResponse
//	 400: errorResponse
//	 401: errorResponse
//	 403: errorResponse
//	 422: errorResponse
//	 500: errorResponse
func FacebookToken(c buffalo.Context) error {
	payload := &common.TokenPayload{}
	if err := c.Bind(payload); err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	apiKey, err := exchangeFacebookUserKey(c, payload.Token)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, render.JSON(&common.TokenPayload{
		Token: apiKey,
	}))
}
