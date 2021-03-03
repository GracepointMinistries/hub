package actions

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/GracepointMinistries/hub/modelext"
	"github.com/gobuffalo/buffalo"
)

const (
	googleEndpointProfile   = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
	facebookEndpointProfile = "https://graph.facebook.com/me?fields=email,id,name&access_token="
)

type thirdPartyUser struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func getThirdPartyUser(c buffalo.Context, provider, endpoint, accessToken string) (*thirdPartyUser, error) {
	profileURL := endpoint + url.QueryEscape(accessToken)
	response, err := http.DefaultClient.Get(profileURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s responded with a %d trying to fetch user information", provider, response.StatusCode)
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	tpUser := &thirdPartyUser{}
	err = json.NewDecoder(bytes.NewReader(data)).Decode(tpUser)
	if err != nil {
		return nil, err
	}
	return tpUser, nil
}

func exchangeGoogleAdminKey(c buffalo.Context, accessToken string) (string, error) {
	user, err := getThirdPartyUser(c, "google", googleEndpointProfile, accessToken)
	if err != nil {
		return "", err
	}
	for _, admin := range admins {
		if admin == user.Email {
			session, err := modelext.CreateAdminSession(c, admin, userIP(c))
			if err != nil {
				return "", err
			}
			return generateAdminToken(user.Email, session.ID)
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
	session, err := modelext.CreateUserSession(c, user, userIP(c))
	if err != nil {
		return "", err
	}
	return generateUserToken(user.ID, session.ID)
}

func exchangeFacebookUserKey(c buffalo.Context, accessToken string) (string, error) {
	fbUser, err := getThirdPartyUser(c, "facebook", facebookEndpointProfile, accessToken)
	if err != nil {
		return "", err
	}
	user, err := modelext.EnsureUserWithOauth(c, "facebook", fbUser.ID, fbUser.Name, fbUser.Email)
	if err != nil {
		return "", err
	}
	session, err := modelext.CreateUserSession(c, user, userIP(c))
	if err != nil {
		return "", err
	}
	return generateUserToken(user.ID, session.ID)
}

// TokenRequest is a request for a token exchange
// swagger:parameters exchangeFacebook exchangeGoogle exchangeAdmin
type TokenRequest struct {
	// in: body
	// required: true
	Body TokenPayload
}

// TokenResponse is a response after a token exchange
// swagger:response tokenResponse
type TokenResponse struct {
	// in: body
	Body TokenPayload
}

// TokenPayload contain the body parameters for a token exchange
type TokenPayload struct {
	Token string `json:"token"`
}

// swagger:route POST /api/v1/exchange/admin auth exchangeAdmin
// Exchanges a google authentication token with an admin api token.
// responses:
//   200: tokenResponse
//	 400: apiErrorResponse
//	 401: apiErrorResponse
//	 403: apiErrorResponse
//	 422: apiErrorResponse
//	 500: apiErrorResponse
func apiExchangeAdminToken(c buffalo.Context) error {
	payload := &TokenPayload{}
	if err := c.Bind(payload); err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	apiKey, err := exchangeGoogleAdminKey(c, payload.Token)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, r.JSON(&TokenPayload{
		Token: apiKey,
	}))
}

// swagger:route POST /api/v1/exchange/google auth exchangeGoogle
// Exchanges a google authentication token with an api token.
// responses:
//   200: tokenResponse
//	 400: apiErrorResponse
//	 401: apiErrorResponse
//	 403: apiErrorResponse
//	 422: apiErrorResponse
//	 500: apiErrorResponse
func apiExchangeGoogleToken(c buffalo.Context) error {
	payload := &TokenPayload{}
	if err := c.Bind(payload); err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	apiKey, err := exchangeGoogleUserKey(c, payload.Token)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, r.JSON(&TokenPayload{
		Token: apiKey,
	}))
}

// swagger:route POST /api/v1/exchange/facebook auth exchangeFacebook
// Exchanges a facebook authentication token with an api token.
// responses:
//   200: tokenResponse
//	 400: apiErrorResponse
//	 401: apiErrorResponse
//	 403: apiErrorResponse
//	 422: apiErrorResponse
//	 500: apiErrorResponse
func apiExchangeFacebookToken(c buffalo.Context) error {
	payload := &TokenPayload{}
	if err := c.Bind(payload); err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	apiKey, err := exchangeFacebookUserKey(c, payload.Token)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, r.JSON(&TokenPayload{
		Token: apiKey,
	}))
}
