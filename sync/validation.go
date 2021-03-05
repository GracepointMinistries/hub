package sync

import (
	"fmt"

	sheets "google.golang.org/api/sheets/v4"
)

func booleanValidationRule(sheetID int64, column typeToCell) *sheets.Request {
	return &sheets.Request{
		SetDataValidation: &sheets.SetDataValidationRequest{
			Range: gridRangeForData(sheetID, column),
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

func rangeValidationRule(sheetID int64, column typeToCell, validationRange string) *sheets.Request {
	return &sheets.Request{
		SetDataValidation: &sheets.SetDataValidationRequest{
			Range: gridRangeForData(sheetID, column),
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

func uniqueValidationRule(sheetID int64, title string, column typeToCell) *sheets.Request {
	// =COUNTIF('SHEET NAME'!$A:$A,"="&'SHEET NAME'!A2)  < 2
	condition := fmt.Sprintf("=COUNTIF(%s,\"=\"&%s%d) < 2", makeRange(title, absoluteDataRangeFor(column)), makeRange(title, column.name), headerStart+2)
	return &sheets.Request{
		SetDataValidation: &sheets.SetDataValidationRequest{
			Range: gridRangeForData(sheetID, column),
			Rule: &sheets.DataValidationRule{
				Condition: &sheets.BooleanCondition{
					Type: "CUSTOM_FORMULA",
					Values: []*sheets.ConditionValue{
						&sheets.ConditionValue{
							UserEnteredValue: condition,
						},
					},
				},
				Strict: true,
			},
		},
	}
}
