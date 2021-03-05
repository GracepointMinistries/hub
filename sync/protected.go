package sync

import (
	sheets "google.golang.org/api/sheets/v4"
)

func updateUnprotectedRangesRequest(index, sheetID int64, ranges []*sheets.GridRange) *sheets.Request {
	return &sheets.Request{
		UpdateProtectedRange: &sheets.UpdateProtectedRangeRequest{
			ProtectedRange: &sheets.ProtectedRange{
				ProtectedRangeId: index,
				Editors: &sheets.Editors{
					Users: []string{
						// only this app
						syncEmail,
					},
				},
				Range: &sheets.GridRange{
					SheetId: sheetID,
				},
				UnprotectedRanges: ranges,
			},
			Fields: "*",
		},
	}
}

func addUnprotectedRangesRequest(index, sheetID int64, ranges []*sheets.GridRange) *sheets.Request {
	return &sheets.Request{
		AddProtectedRange: &sheets.AddProtectedRangeRequest{
			ProtectedRange: &sheets.ProtectedRange{
				ProtectedRangeId: index,
				Editors: &sheets.Editors{
					Users: []string{
						// only this app
						syncEmail,
					},
				},
				Range: &sheets.GridRange{
					SheetId: sheetID,
				},
				UnprotectedRanges: ranges,
			},
		},
	}
}

func unprotectedRangesRequest(update bool, index, sheetID int64, ranges []*sheets.GridRange) *sheets.Request {
	if update {
		return updateUnprotectedRangesRequest(index, sheetID, ranges)
	}
	return addUnprotectedRangesRequest(index, sheetID, ranges)
}
