package sync

import (
	script "google.golang.org/api/script/v1"
)

const scriptSource = `
function helloWorld() {
	console.log('Hello, world!');
}
`

func getScriptSource(url string) string {
	return scriptSource
}

func createScriptProject(parent string) (string, error) {
	response, err := scriptClient.Projects.Create(&script.CreateProjectRequest{
		ParentId: parent,
		Title:    "Hub Script",
	}).Do()
	if err != nil {
		return "", err
	}
	return response.ScriptId, nil
}

func updateScriptProject(id, url string) error {
	_, err := scriptClient.Projects.UpdateContent(id, &script.Content{
		ScriptId: id,
		Files: []*script.File{{
			Name:   "script",
			Type:   "SERVER_JS",
			Source: getScriptSource(url),
		}},
	}).Do()
	return err
}
