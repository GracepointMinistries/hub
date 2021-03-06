package sync

import (
	"fmt"

	script "google.golang.org/api/script/v1"
)

const (
	scriptURLPrefix = "https://script.google.com/d/"

	scriptSource = `
function sync() {
	UrlFetchApp.fetch("%s");
}

function onOpen(e) {
  SpreadsheetApp.getUi()
      .createMenu('Hub')
      .addItem('Sync', 'sync')
      .addToUi();
}
`
	scriptManifestSource = `{
  "timeZone": "America/New_York",
	"exceptionLogging": "NONE",
	"urlFetchWhitelist": ["%s"]
}`
)

func getScriptSource(url string) string {
	return fmt.Sprintf(scriptSource, url)
}

func getScriptManifestSource(url string) string {
	return fmt.Sprintf(scriptManifestSource, url)
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
