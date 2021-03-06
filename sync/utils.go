package sync

import (
	"fmt"
	"strings"

	sheets "google.golang.org/api/sheets/v4"
)

func stringOrEmpty(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

func isStringEmpty(v *string) bool {
	if v == nil {
		return true
	}
	return *v == ""
}

func boolOrFalse(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}

func urlToID(url string) string {
	id := strings.TrimPrefix(url, urlPrefix)
	id = strings.TrimSuffix(id, urlSuffix)
	return id
}

func makeRange(tab, dataRange string) string {
	return fmt.Sprintf("'%s'!%s", tab, dataRange)
}

func gridRangeForData(id int64, col typeToCell) *sheets.GridRange {
	return &sheets.GridRange{
		SheetId:          id,
		StartColumnIndex: col.offset,
		EndColumnIndex:   col.offset + 1,
		// skip header
		StartRowIndex: headerStart + 1,
	}
}

func gridRangeDataWithHeader(id, end int64) *sheets.GridRange {
	return &sheets.GridRange{
		SheetId:        id,
		EndColumnIndex: end + 1,
		StartRowIndex:  headerStart,
	}
}

func gridRangeFor(id int64, col typeToCell) *sheets.GridRange {
	return &sheets.GridRange{
		SheetId:          id,
		StartColumnIndex: col.offset,
		EndColumnIndex:   col.offset + 1,
	}
}

func gridRangeOverflow(id int64, lastColumn int64) *sheets.GridRange {
	return &sheets.GridRange{
		SheetId:          id,
		StartColumnIndex: lastColumn + 1,
	}
}
