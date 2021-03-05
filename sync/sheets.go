package sync

import (
	"errors"
	"fmt"
	"strings"

	drive "google.golang.org/api/drive/v3"
	sheets "google.golang.org/api/sheets/v4"
)

const (
	urlPrefix = "https://docs.google.com/spreadsheets/d/"
	urlSuffix = "/edit"

	groupsTitle = "Groups - DO NOT RENAME"
	usersTitle  = "Users - DO NOT RENAME"
)

func urlToID(url string) string {
	id := strings.TrimPrefix(url, urlPrefix)
	id = strings.TrimSuffix(id, urlSuffix)
	return id
}

func makeRange(tab, dataRange string) string {
	return fmt.Sprintf("'%s'!%s", tab, dataRange)
}

// DumpToSpreadsheet exports the group and user space to the given spreadsheet
func DumpToSpreadsheet(url string) error {
	id := urlToID(url)
	sheet, err := syncClient.Spreadsheets.Get(id).Do()
	// validate the sheet metadata
	if len(sheet.Sheets) < 2 {
		return errors.New("missing required tabs on spreadsheet")
	}
	groups := sheet.Sheets[0]
	users := sheet.Sheets[1]
	if groups.Properties.Title != groupsTitle || users.Properties.Title != usersTitle {
		return errors.New("unrecognized sheet name")
	}
	// slurp up some data
	groupRange := makeRange(groupsTitle, "A:E")
	userRange := makeRange(usersTitle, "A:E")
	batch, err := syncClient.Spreadsheets.Values.BatchGet(id).Ranges(groupRange, userRange).Do()
	if err != nil {
		return err
	}
	grps := newGroupSlice()
	usrs := newUserSlice()
	if err := grps.Unmarshal(batch.ValueRanges[0]); err != nil {
		return err
	}
	if err := usrs.Unmarshal(batch.ValueRanges[1]); err != nil {
		return err
	}
	fmt.Println(grps.String())
	fmt.Println(usrs.String())
	return err
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
					Title: groupsTitle,
				},
			},
			&sheets.Sheet{
				Properties: &sheets.SheetProperties{
					Title: usersTitle,
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
