package clientext

import (
	"github.com/GracepointMinistries/hub/client"
)

// NewClient initializes an api client with the global
// configuration file
func NewClient() *client.APIClient {
	return client.NewAPIClient(&client.Configuration{
		BasePath: fileConfig.Host + "/",
		DefaultHeader: map[string]string{
			"Authorization": "Bearer " + fileConfig.Token,
		},
		UserAgent: "Hub-CLI/1.0.0",
	})
}
