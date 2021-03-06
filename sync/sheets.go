package sync

import (
	"errors"

	"github.com/gobuffalo/buffalo"
	drive "google.golang.org/api/drive/v3"
	sheets "google.golang.org/api/sheets/v4"
)

const (
	urlPrefix = "https://docs.google.com/spreadsheets/d/"
	urlSuffix = "/edit"

	groupsTitle = "Groups"
	usersTitle  = "Users"
)

func createProtectionRules(update bool, id string, groupID, userID int64) error {
	userUnprotectedRanges := []*sheets.GridRange{
		gridRangeForData(userID, userHeadersLookup["Name"]),
		gridRangeForData(userID, userHeadersLookup["Blocked"]),
		gridRangeForData(userID, userHeadersLookup["Group"]),
		gridRangeOverflow(userID, userColumnEnd),
	}
	groupUnprotectedRanges := []*sheets.GridRange{
		gridRangeForData(groupID, groupHeadersLookup["Name"]),
		gridRangeForData(groupID, groupHeadersLookup["ZoomLink"]),
		gridRangeForData(groupID, groupHeadersLookup["Published"]),
		gridRangeForData(groupID, groupHeadersLookup["Archived"]),
		gridRangeOverflow(groupID, groupColumnEnd),
	}
	request := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			unprotectedRangesRequest(update, 3, groupID, groupUnprotectedRanges),
			unprotectedRangesRequest(update, 4, userID, userUnprotectedRanges),
		},
	}
	_, err := syncClient.Spreadsheets.BatchUpdate(id, request).Do()
	return err
}

func createValidationRules(id string, groupID, userID int64) error {
	groupNameColumn := groupHeadersLookup["Name"]
	groupNameRangeFormula := "=" + makeRange(groupsTitle, dataRangeFor(groupNameColumn))

	request := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			// user.Group mapping
			rangeValidationRule(userID, userHeadersLookup["Group"], groupNameRangeFormula),
			// Group boolean mappings
			booleanValidationRule(groupID, groupHeadersLookup["Published"]),
			booleanValidationRule(groupID, groupHeadersLookup["Archived"]),
			// User boolean mappings
			booleanValidationRule(userID, userHeadersLookup["Blocked"]),
			// Uniqueness on group names
			uniqueValidationRule(groupID, groupsTitle, groupNameColumn),
			// Uniqueness on user emails
			uniqueValidationRule(userID, usersTitle, userHeadersLookup["Email"]),
		},
	}
	_, err := syncClient.Spreadsheets.BatchUpdate(id, request).Do()
	return err
}

func formatData(update bool, id string, groupID, userID int64) error {
	groupIDColumn := groupHeadersLookup["ID"]
	groupNameColumn := groupHeadersLookup["Name"]
	groupZoomLinkColumn := groupHeadersLookup["ZoomLink"]
	userIDColumn := userHeadersLookup["ID"]
	userNameColumn := userHeadersLookup["Name"]
	userEmailColumn := userHeadersLookup["Email"]
	userGroupColumn := userHeadersLookup["Group"]
	request := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			formatValidationCell(groupID, groupColumnEnd),
			formatValidationCell(userID, userColumnEnd),
			conditionalErrorFormattingRequest(update, 0, groupID, groupColumnEnd),
			conditionalErrorFormattingRequest(update, 0, userID, userColumnEnd),
			mergeValidationCell(groupID, groupColumnEnd),
			mergeValidationCell(userID, userColumnEnd),
			headerFormattingRequest(groupID, groupColumnBegin, groupColumnEnd),
			headerFormattingRequest(userID, userColumnBegin, userColumnEnd),
			freezeHeaderRequest(groupID, groupColumnEnd),
			freezeHeaderRequest(userID, userColumnEnd),
			formatAlignRightRequest(groupID, groupIDColumn),
			formatAlignRightRequest(userID, userIDColumn),
			formatAlignLeftRequest(groupID, groupNameColumn),
			formatAlignLeftRequest(groupID, groupZoomLinkColumn),
			formatAlignLeftRequest(userID, userNameColumn),
			formatAlignLeftRequest(userID, userEmailColumn),
			formatAlignLeftRequest(userID, userGroupColumn),
			bandingRequest(update, 1, groupID, groupColumnBegin, groupColumnEnd),
			bandingRequest(update, 2, userID, userColumnBegin, userColumnEnd),
			formatLockedRequest(groupID, groupIDColumn),
			formatLockedRequest(userID, userIDColumn),
			formatLockedRequest(userID, userEmailColumn),
			autoSizeRequest(groupID, groupColumnBegin, groupColumnEnd, "COLUMNS"),
			autoSizeRequest(groupID, 0, -1, "ROWS"),
			autoSizeRequest(userID, userColumnBegin, userColumnEnd, "COLUMNS"),
			autoSizeRequest(userID, 0, -1, "ROWS"),
		},
	}
	_, err := syncClient.Spreadsheets.BatchUpdate(id, request).Do()
	return err
}

func exportGroups(c buffalo.Context, id string) error {
	groups, err := dbGroupSlice(c)
	if err != nil {
		return err
	}
	users, err := dbUserSlice(c)
	if err != nil {
		return err
	}

	_, err = syncClient.Spreadsheets.Values.BatchUpdate(id, &sheets.BatchUpdateValuesRequest{
		Data: []*sheets.ValueRange{
			groups.Marshal(),
			users.Marshal(),
		},
		ValueInputOption: "USER_ENTERED",
	}).Do()
	return err
}

// ExportToSpreadsheet exports the group and user space to the given spreadsheet
func ExportToSpreadsheet(c buffalo.Context, url string) error {
	return exportToSpreadsheet(c, true, url)
}

func exportToSpreadsheet(c buffalo.Context, update bool, url string) error {
	id := urlToID(url)
	sheet, err := syncClient.Spreadsheets.Get(id).Do()
	if err != nil {
		return err
	}
	// validate the sheet metadata
	if len(sheet.Sheets) < 2 {
		return errors.New("missing required tabs on spreadsheet")
	}
	groups := sheet.Sheets[0]
	users := sheet.Sheets[1]
	if groups.Properties.Title != groupsTitle || users.Properties.Title != usersTitle {
		return errors.New("unrecognized sheet name")
	}

	if update {
		// slurp up data
		batch, err := syncClient.Spreadsheets.Values.BatchGet(id).Ranges(allGroupRange, allUserRange).Do()
		if err != nil {
			return err
		}
		syncGroups := newGroupSlice()
		syncUsers := newUserSlice()
		if err := syncGroups.Unmarshal(batch.ValueRanges[0]); err != nil {
			return err
		}
		if err := syncUsers.Unmarshal(batch.ValueRanges[1]); err != nil {
			return err
		}
		// save the groups first since some users might reference the new ones
		if err := syncGroups.Save(c); err != nil {
			return err
		}
		if err := syncUsers.Save(c); err != nil {
			return err
		}
	}

	// sync it back to the sheet

	// ensure sheet protection rules are set
	if err := createProtectionRules(update, id, groups.Properties.SheetId, users.Properties.SheetId); err != nil {
		return err
	}
	// ensure validation rules are set
	if err := createValidationRules(id, groups.Properties.SheetId, users.Properties.SheetId); err != nil {
		return err
	}
	// dump data
	if err := exportGroups(c, id); err != nil {
		return err
	}
	// format spreadsheet
	return formatData(update, id, groups.Properties.SheetId, users.Properties.SheetId)
}

// CreateSpreadsheet creates a new Google spreadsheet for synchronization
func CreateSpreadsheet(c buffalo.Context) (string, error) {
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
	return sheet.SpreadsheetUrl, exportToSpreadsheet(c, false, sheet.SpreadsheetUrl)
}
