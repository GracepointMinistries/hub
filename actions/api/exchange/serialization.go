package exchange

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gobuffalo/buffalo"
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
