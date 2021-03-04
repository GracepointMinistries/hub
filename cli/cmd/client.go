package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/GracepointMinistries/hub/client"
)

func checkUnauthorized(response *http.Response) {
	if response != nil && response.StatusCode == http.StatusUnauthorized {
		fmt.Fprintln(os.Stderr, "You are not authorized to perform that action")
		os.Exit(1)
	}
}

func newClient() *client.APIClient {
	return client.NewAPIClient(&client.Configuration{
		BasePath: fileConfig.Host + "/",
		DefaultHeader: map[string]string{
			"Authorization": "Bearer " + fileConfig.Token,
		},
		UserAgent: "Hub-CLI/1.0.0",
	})
}
