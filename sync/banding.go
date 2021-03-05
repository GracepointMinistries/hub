package sync

import sheets "google.golang.org/api/sheets/v4"

func bandedRange(index, sheetID, start, end int64) *sheets.BandedRange {
	return &sheets.BandedRange{
		BandedRangeId: index,
		Range: &sheets.GridRange{
			SheetId:          sheetID,
			StartColumnIndex: start,
			EndColumnIndex:   end + 1,
			// skip header
			StartRowIndex: headerStart + 1,
		},
		RowProperties: &sheets.BandingProperties{
			FirstBandColor: &sheets.Color{
				Red:   1.0,
				Green: 1.0,
				Blue:  1.0,
			},
			SecondBandColor: &sheets.Color{
				Red:   0.9,
				Green: 0.9,
				Blue:  0.9,
			},
		},
	}
}

func addBandingRequest(index, sheetID, start, end int64) *sheets.Request {
	return &sheets.Request{
		AddBanding: &sheets.AddBandingRequest{
			BandedRange: bandedRange(index, sheetID, start, end),
		},
	}
}

func updateBandingRequest(index, sheetID, start, end int64) *sheets.Request {
	return &sheets.Request{
		UpdateBanding: &sheets.UpdateBandingRequest{
			BandedRange: bandedRange(index, sheetID, start, end),
			Fields:      "*",
		},
	}
}

func bandingRequest(update bool, index, sheetID, start, end int64) *sheets.Request {
	if update {
		return updateBandingRequest(index, sheetID, start, end)
	}
	return addBandingRequest(index, sheetID, start, end)
}
