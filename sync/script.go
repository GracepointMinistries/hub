package sync

import (
	script "google.golang.org/api/script/v1"
)

const (
	scriptURLPrefix = "https://script.google.com/d/"

	scriptSource = `
function helloWorld() {
	console.log('Hello, world!');
}
`
	scriptManifestSource = `{
  "timeZone": "America/New_York",
	"exceptionLogging": "NONE"
}`
)

func getScriptSource(url string) string {
	return scriptSource
}

func getScriptManifestSource(url string) string {
	return scriptManifestSource
}

func createScriptProject(parent string) (string, error) {
	response, err := scriptClient.Projects.Create(&script.CreateProjectRequest{
		ParentId: parent,
		Title:    "Hub Script",
	}).Do()
	if err != nil {
		return "", err
	}
	url := scriptURLPrefix + response.ScriptId + urlSuffix
	return url, nil
}

func updateScriptProject(id, url string) error {
	_, err := scriptClient.Projects.UpdateContent(id, &script.Content{
		ScriptId: id,
		Files: []*script.File{{
			Name:   "script",
			Type:   "SERVER_JS",
			Source: getScriptSource(url),
		}, {
			Name:   "appsscript",
			Type:   "JSON",
			Source: getScriptManifestSource(url),
		}},
	}).Do()
	return err
}
