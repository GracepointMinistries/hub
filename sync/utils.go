package sync

import (
	"fmt"
	"strings"

	sheets "google.golang.org/api/sheets/v4"
)

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
