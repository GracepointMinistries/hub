package sync

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gobuffalo/buffalo"
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

func booleanValidationRule(sheetID, dataColumnOffset int64) *sheets.Request {
	return &sheets.Request{
		SetDataValidation: &sheets.SetDataValidationRequest{
			Range: &sheets.GridRange{
				SheetId:          sheetID,
				StartColumnIndex: dataColumnOffset,
				EndColumnIndex:   dataColumnOffset + 1,
				// Skip the header
				StartRowIndex: 1,
			},
			Rule: &sheets.DataValidationRule{
				Condition: &sheets.BooleanCondition{
					Type: "BOOLEAN",
					Values: []*sheets.ConditionValue{
						&sheets.ConditionValue{
							UserEnteredValue: "TRUE",
						},
						&sheets.ConditionValue{
							UserEnteredValue: "FALSE",
						},
					},
				},
				ShowCustomUi: true,
				Strict:       true,
			},
		},
	}
}

func rangeValidationRule(sheetID, dataColumnOffset int64, validationRange string) *sheets.Request {
	return &sheets.Request{
		SetDataValidation: &sheets.SetDataValidationRequest{
			Range: &sheets.GridRange{
				SheetId:          sheetID,
				StartColumnIndex: dataColumnOffset,
				EndColumnIndex:   dataColumnOffset + 1,
				// Skip the header
				StartRowIndex: 1,
			},
			Rule: &sheets.DataValidationRule{
				Condition: &sheets.BooleanCondition{
					Type: "ONE_OF_RANGE",
					Values: []*sheets.ConditionValue{
						&sheets.ConditionValue{
							UserEnteredValue: validationRange,
						},
					},
				},
				ShowCustomUi: true,
				Strict:       true,
			},
		},
	}
}

func createValidationRules(id string, groupID, userID int64) error {
	userGroupColumnOffset := userHeadersLookup["Group"].offset
	userBlockedColumnOffset := userHeadersLookup["Blocked"].offset
	groupPublishedColumnOffset := groupHeadersLookup["Published"].offset
	groupArchivedColumnOffset := groupHeadersLookup["Archived"].offset

	groupNameColumn := groupHeadersLookup["Name"]
	groupNameRangeFormula := "=" + makeRange(groupsTitle, dataRangeFor(groupNameColumn))

	request := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			// user.Group mapping
			rangeValidationRule(userID, userGroupColumnOffset, groupNameRangeFormula),
			// Group boolean mappings
			booleanValidationRule(groupID, groupPublishedColumnOffset),
			booleanValidationRule(groupID, groupArchivedColumnOffset),
			// User boolean mappings
			booleanValidationRule(userID, userBlockedColumnOffset),
		},
	}
	_, err := syncClient.Spreadsheets.BatchUpdate(id, request).Do()
	return err
}

func autoSizeRequest(sheetID, begin, end int64, orientation string) *sheets.Request {
	return &sheets.Request{
		AutoResizeDimensions: &sheets.AutoResizeDimensionsRequest{
			Dimensions: &sheets.DimensionRange{
				Dimension:  orientation,
				SheetId:    sheetID,
				StartIndex: begin,
				EndIndex:   end + 1,
			},
		},
	}
}

func headerFormattingRequest(sheetID, begin, end int64) *sheets.Request {
	return &sheets.Request{
		RepeatCell: &sheets.RepeatCellRequest{
			Fields: "userEnteredFormat(backgroundColor,textFormat,horizontalAlignment)",
			Range: &sheets.GridRange{
				SheetId:          sheetID,
				StartColumnIndex: begin,
				EndColumnIndex:   end + 1,
				StartRowIndex:    0,
				EndRowIndex:      1,
			},
			Cell: &sheets.CellData{
				UserEnteredFormat: &sheets.CellFormat{
					BackgroundColor: &sheets.Color{
						Red:   0.0,
						Green: 0.0,
						Blue:  0.0,
					},
					HorizontalAlignment: "CENTER",
					TextFormat: &sheets.TextFormat{
						ForegroundColor: &sheets.Color{
							Red:   1.0,
							Green: 1.0,
							Blue:  1.0,
						},
						Bold:     true,
						FontSize: 12,
					},
				},
			},
		},
	}
}

func formatAlignLeftRequest(sheetID, column int64) *sheets.Request {
	return &sheets.Request{
		RepeatCell: &sheets.RepeatCellRequest{
			Fields: "userEnteredFormat(horizontalAlignment)",
			Range: &sheets.GridRange{
				SheetId:          sheetID,
				StartColumnIndex: column,
				EndColumnIndex:   column + 1,
				// skip the header
				StartRowIndex: 1,
			},
			Cell: &sheets.CellData{
				UserEnteredFormat: &sheets.CellFormat{
					HorizontalAlignment: "LEFT",
				},
			},
		},
	}
}

func formatAlignRightRequest(sheetID, column int64) *sheets.Request {
	return &sheets.Request{
		RepeatCell: &sheets.RepeatCellRequest{
			Fields: "userEnteredFormat(horizontalAlignment)",
			Range: &sheets.GridRange{
				SheetId:          sheetID,
				StartColumnIndex: column,
				EndColumnIndex:   column + 1,
				// skip the header
				StartRowIndex: 1,
			},
			Cell: &sheets.CellData{
				UserEnteredFormat: &sheets.CellFormat{
					HorizontalAlignment: "RIGHT",
				},
			},
		},
	}
}

func formatData(id string, groupID, userID int64) error {
	groupIDColumnOffset := groupHeadersLookup["ID"].offset
	groupNameColumnOffset := groupHeadersLookup["Name"].offset
	groupZoomLinkColumnOffset := groupHeadersLookup["ZoomLink"].offset
	userIDColumnOffset := userHeadersLookup["ID"].offset
	userNameColumnOffset := userHeadersLookup["Name"].offset
	userEmailColumnOffset := userHeadersLookup["Email"].offset
	userGroupColumnOffset := userHeadersLookup["Group"].offset

	request := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			headerFormattingRequest(groupID, groupColumnBegin, groupColumnEnd),
			headerFormattingRequest(userID, userColumnBegin, userColumnEnd),
			formatAlignRightRequest(groupID, groupIDColumnOffset),
			formatAlignRightRequest(userID, userIDColumnOffset),
			formatAlignLeftRequest(groupID, groupNameColumnOffset),
			formatAlignLeftRequest(groupID, groupZoomLinkColumnOffset),
			formatAlignLeftRequest(userID, userNameColumnOffset),
			formatAlignLeftRequest(userID, userEmailColumnOffset),
			formatAlignLeftRequest(userID, userGroupColumnOffset),
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

// DumpToSpreadsheet exports the group and user space to the given spreadsheet
func DumpToSpreadsheet(c buffalo.Context, url string) error {
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
	// ensure validation rules are set
	if err := createValidationRules(id, groups.Properties.SheetId, users.Properties.SheetId); err != nil {
		return err
	}
	// dump data
	if err := exportGroups(c, id); err != nil {
		return err
	}
	// format spreadsheet
	if err := formatData(id, groups.Properties.SheetId, users.Properties.SheetId); err != nil {
		return err
	}

	// slurp up some data
	// groupRange := makeRange(groupsTitle, "A:E")
	// userRange := makeRange(usersTitle, "A:E")
	// batch, err := syncClient.Spreadsheets.Values.BatchGet(id).Ranges(groupRange, userRange).Do()
	// if err != nil {
	// 	return err
	// }
	// grps := newGroupSlice()
	// usrs := newUserSlice()
	// if err := grps.Unmarshal(batch.ValueRanges[0]); err != nil {
	// 	return err
	// }
	// if err := usrs.Unmarshal(batch.ValueRanges[1]); err != nil {
	// 	return err
	// }
	// fmt.Println(grps.String())
	// fmt.Println(usrs.String())
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
