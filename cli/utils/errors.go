package utils

import (
	"fmt"
	"net/http"
	"os"

	"github.com/GracepointMinistries/hub/cli/print"
)

// CheckError checks for an error and exits if there is one
func CheckError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, print.Criticalf("Error: %v", err))
		os.Exit(1)
	}
}

// CheckUnauthorized checks for an unauthorized status code and exits if there is one
func CheckUnauthorized(response *http.Response) {
	if response != nil && response.StatusCode == http.StatusUnauthorized {
		fmt.Fprintln(os.Stderr, print.Critical("You are not authorized to perform that action"))
		os.Exit(1)
	}
}
