package sync

import sheets "google.golang.org/api/sheets/v4"

func mergeValidationCell(id, end int64) *sheets.Request {
	return &sheets.Request{
		MergeCells: &sheets.MergeCellsRequest{
			MergeType: "MERGE_ALL",
			Range: &sheets.GridRange{
				SheetId:        id,
				EndColumnIndex: end + 1,
				EndRowIndex:    headerStart,
			},
		},
	}
}

func formatValidationCell(id, end int64) *sheets.Request {
	return &sheets.Request{
		RepeatCell: &sheets.RepeatCellRequest{
			Fields: "userEnteredFormat(backgroundColor,textFormat,horizontalAlignment)",
			Range: &sheets.GridRange{
				SheetId:        id,
				EndColumnIndex: end + 1,
				EndRowIndex:    headerStart,
			},
			Cell: &sheets.CellData{
				UserEnteredFormat: &sheets.CellFormat{
					BackgroundColor: &sheets.Color{
						Red:   0.5,
						Green: 0.5,
						Blue:  0.5,
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

func headerFormattingRequest(sheetID, begin, end int64) *sheets.Request {
	return &sheets.Request{
		RepeatCell: &sheets.RepeatCellRequest{
			Fields: "userEnteredFormat(backgroundColor,textFormat,horizontalAlignment)",
			Range: &sheets.GridRange{
				SheetId:          sheetID,
				StartColumnIndex: begin,
				EndColumnIndex:   end + 1,
				StartRowIndex:    headerStart,
				EndRowIndex:      headerStart + 1,
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

func formatAlignLeftRequest(id int64, column typeToCell) *sheets.Request {
	return &sheets.Request{
		RepeatCell: &sheets.RepeatCellRequest{
			Fields: "userEnteredFormat(horizontalAlignment)",
			Range:  gridRangeForData(id, column),
			Cell: &sheets.CellData{
				UserEnteredFormat: &sheets.CellFormat{
					HorizontalAlignment: "LEFT",
				},
			},
		},
	}
}

func formatAlignRightRequest(id int64, column typeToCell) *sheets.Request {
	return &sheets.Request{
		RepeatCell: &sheets.RepeatCellRequest{
			Fields: "userEnteredFormat(horizontalAlignment)",
			Range:  gridRangeForData(id, column),
			Cell: &sheets.CellData{
				UserEnteredFormat: &sheets.CellFormat{
					HorizontalAlignment: "RIGHT",
				},
			},
		},
	}
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

func freezeHeaderRequest(id, end int64) *sheets.Request {
	return &sheets.Request{
		UpdateSheetProperties: &sheets.UpdateSheetPropertiesRequest{
			Fields: "gridProperties(frozenRowCount,frozenColumnCount)",
			Properties: &sheets.SheetProperties{
				SheetId: id,
				GridProperties: &sheets.GridProperties{
					FrozenColumnCount: end + 1,
					FrozenRowCount:    headerStart + 1,
				},
			},
		},
	}
}

func formatLockedRequest(id int64, column typeToCell) *sheets.Request {
	return &sheets.Request{
		RepeatCell: &sheets.RepeatCellRequest{
			Fields: "userEnteredFormat(backgroundColor,textFormat)",
			Range:  gridRangeForData(id, column),
			Cell: &sheets.CellData{
				UserEnteredFormat: &sheets.CellFormat{
					BackgroundColor: &sheets.Color{
						Red:   0.8,
						Green: 0.8,
						Blue:  0.8,
					},
					TextFormat: &sheets.TextFormat{
						ForegroundColor: &sheets.Color{
							Red:   0.2,
							Green: 0.2,
							Blue:  0.2,
						},
					},
				},
			},
		},
	}
}
