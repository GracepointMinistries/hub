package sync

import sheets "google.golang.org/api/sheets/v4"

func validationRange(id, end int64) *sheets.GridRange {
	return &sheets.GridRange{
		SheetId:        id,
		EndColumnIndex: end + 1,
		EndRowIndex:    headerStart,
	}
}

func mergeValidationCell(id, end int64) *sheets.Request {
	return &sheets.Request{
		MergeCells: &sheets.MergeCellsRequest{
			MergeType: "MERGE_ALL",
			Range:     validationRange(id, end),
		},
	}
}

func formatValidationCell(id, end int64) *sheets.Request {
	return &sheets.Request{
		RepeatCell: &sheets.RepeatCellRequest{
			Fields: "userEnteredFormat(backgroundColor,textFormat,horizontalAlignment)",
			Range:  validationRange(id, end),
			Cell: &sheets.CellData{
				UserEnteredFormat: &sheets.CellFormat{
					BackgroundColor: &sheets.Color{
						// Forest green
						Red:   0.1643,
						Green: 0.6715,
						Blue:  0.1643,
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

func updateConditionalErrorFormattingRequest(index, id, end int64) *sheets.Request {
	return &sheets.Request{
		UpdateConditionalFormatRule: &sheets.UpdateConditionalFormatRuleRequest{
			Index: index,
			Rule: &sheets.ConditionalFormatRule{
				BooleanRule: &sheets.BooleanRule{
					Condition: &sheets.BooleanCondition{
						Type: "TEXT_EQ",
						Values: []*sheets.ConditionValue{
							&sheets.ConditionValue{
								UserEnteredValue: "INVALID",
							},
						},
					},
					Format: &sheets.CellFormat{
						BackgroundColor: &sheets.Color{
							Red:   1.0,
							Green: 0.39,
							Blue:  0.28,
						},
					},
				},
				Ranges: []*sheets.GridRange{
					validationRange(id, end),
				},
			},
		},
	}
}

func addConditionalErrorFormattingRequest(index, id, end int64) *sheets.Request {
	return &sheets.Request{
		AddConditionalFormatRule: &sheets.AddConditionalFormatRuleRequest{
			Index: index,
			Rule: &sheets.ConditionalFormatRule{
				BooleanRule: &sheets.BooleanRule{
					Condition: &sheets.BooleanCondition{
						Type: "TEXT_EQ",
						Values: []*sheets.ConditionValue{
							&sheets.ConditionValue{
								UserEnteredValue: "INVALID",
							},
						},
					},
					Format: &sheets.CellFormat{
						BackgroundColor: &sheets.Color{
							// Indian red
							Red:   1.0,
							Green: 0.39,
							Blue:  0.28,
						},
					},
				},
				Ranges: []*sheets.GridRange{
					validationRange(id, end),
				},
			},
		},
	}
}

func conditionalErrorFormattingRequest(update bool, index, id, end int64) *sheets.Request {
	if update {
		return updateConditionalErrorFormattingRequest(index, id, end)
	}
	return addConditionalErrorFormattingRequest(index, id, end)
}
