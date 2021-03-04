package sync

import (
	drive "google.golang.org/api/drive/v3"
	sheets "google.golang.org/api/sheets/v4"
)

// GetSpreadsheetData returns data from the given spreadsheet
func GetSpreadsheetData(id string) (*sheets.Spreadsheet, error) {
	return syncClient.Spreadsheets.Get(id).Do()
}

// CreateSpreadsheet creates a new Google spreadsheet for synchronization
func CreateSpreadsheet() (string, error) {
	sheet, err := syncClient.Spreadsheets.Create(&sheets.Spreadsheet{
		Properties: &sheets.SpreadsheetProperties{
			Title: "Hub Synchronized Groupings",
		},
		Sheets: []*sheets.Sheet{
			&sheets.Sheet{
				Properties: &sheets.SheetProperties{
					Title: "Groups - DO NOT RENAME",
				},
			},
			&sheets.Sheet{
				Properties: &sheets.SheetProperties{
					Title: "Users - DO NOT RENAME",
				},
			},
		},
	}).Do()
	if err != nil {
		return "", err
	}
	_, err = driveClient.Permissions.Create(sheet.SpreadsheetId, &drive.Permission{
		EmailAddress: "andrew.stucki@gmail.com",
		Role:         "writer",
		Type:         "user",
	}).SendNotificationEmail(false).Do()
	if err != nil {
		return "", err
	}
	return sheet.SpreadsheetUrl, nil
}
