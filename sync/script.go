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

function onEdit(e) {
	var sheet = e.range.getSheet().getSheetId();
	// columns are 1 v. 0 indexed via app script
	var column = e.range.getColumn() - 1;

	// double the validation
	if (sheet === %d && column === %d) {
		// e.value is undefined on a paste
		var value = e.range.getValue();
		var column = e.range.getColumn();
		var row = e.range.getRow();
		var spreadsheet = SpreadsheetApp.getActiveSpreadsheet();
		// ensure that there are no duplicates
		var range = spreadsheet.getRange("%s");
		var rows = range.getNumRows();
  	var columns = range.getNumColumns();
		for (var i = 1; i <= columns; i++) {
			for (var j = 1; j <= rows; j++) {
				var cell = range.getCell(j, i);
				var content = cell.getValue();
				if (i + 1 === column && j + 1 === row) {
					continue;
				}
				if (content === value) {
					// clear out what was just set and warn the user
					// that they can't duplicate values, because of
					// the validation rules, this is really only
					// possible if someone copy and pasted the data
					e.range.clearContent();
					var message = "Duplicate value found at cell ";
					message += cell.getA1Notation();
					message += ", all values must be unique";
					SpreadsheetApp.getUi().alert(message);
					return
				}
			}
		}

		// update all of the users teams that are linked to this
		range = spreadsheet.getRange("%s");
		rows = range.getNumRows();
  	columns = range.getNumColumns();

		for (var i = 1; i <= columns; i++) {
			for (var j = 1; j <= rows; j++) {
				var cell = range.getCell(j, i);
				var content = cell.getValue();
				if (content === e.oldValue) {
					cell.setValue(value);
				}
			}
		}
	}
}
`
	scriptManifestSource = `{
  "timeZone": "America/New_York",
	"exceptionLogging": "NONE",
	"urlFetchWhitelist": ["%s"]
}`
)

func getScriptSource(url string, sheet, column int64, nameRange, groupRange string) string {
	return fmt.Sprintf(scriptSource, url, sheet, column, nameRange, groupRange)
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

func updateScriptProject(id, url string, sheet, column int64, nameRange, groupRange string) error {
	_, err := scriptClient.Projects.UpdateContent(id, &script.Content{
		ScriptId: id,
		Files: []*script.File{{
			Name:   "script",
			Type:   "SERVER_JS",
			Source: getScriptSource(url, sheet, column, nameRange, groupRange),
		}, {
			Name:   "appsscript",
			Type:   "JSON",
			Source: getScriptManifestSource(url),
		}},
	}).Do()
	return err
}
