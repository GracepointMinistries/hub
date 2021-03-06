package sync

import (
	"fmt"
	"reflect"

	"github.com/GracepointMinistries/hub/modelext"
	"github.com/GracepointMinistries/hub/models"
	"github.com/gobuffalo/buffalo"
	"github.com/volatiletech/sqlboiler/v4/boil"
	sheets "google.golang.org/api/sheets/v4"
)

var (
	allGroupRange      string
	groupColumnBegin   int64
	groupColumnEnd     int64
	groupHeaders       []string
	groupHeadersLookup map[string]typeToCell
)

func init() {
	group := reflect.TypeOf(groupSlice{}).Elem().Elem()
	groupHeaders = make([]string, group.NumField())
	groupHeadersLookup = make(map[string]typeToCell, group.NumField())
	groupColumnBegin = 0
	for i := 0; i < group.NumField(); i++ {
		header := group.FieldByIndex([]int{i}).Name
		groupColumnEnd = int64(i)
		groupHeaders[i] = header
		groupHeadersLookup[header] = typeToCell{
			offset: groupColumnEnd,
			name:   string(rune('A') + rune(i)),
		}
	}
	groupRangeBegin := "A"
	groupRangeEnd := string(rune('A') + rune(group.NumField()))
	// 'SHEET TITLE'!X:Y
	allGroupRange = makeRange(groupsTitle, groupRangeBegin+":"+groupRangeEnd)
}

// generate these dynamically at some point
const (
	uniqueGroupNameValidation = "EQ(COUNTA('Groups'!$B$3:$B)-COUNTA(UNIQUE('Groups'!$B$3:$B)), 0)"
)

var groupSheetValidationRule = fmt.Sprintf("=IF(%s, \"VALID\", \"INVALID\")", uniqueGroupNameValidation)

type innerGroup struct {
	// order matters here, it needs to be the same as
	// the expected header ordering
	ID        *int    `json:"id"`
	Name      *string `json:"name"`
	ZoomLink  *string `json:"zoomLink"`
	Published *bool   `json:"published"`
	Archived  *bool   `json:"archived"`
}

type groupSlice []*innerGroup

func newGroupSlice() *groupSlice {
	return &groupSlice{}
}

func (g *groupSlice) Unmarshal(values *sheets.ValueRange) error {
	return unmarshal(g, values)
}

func (g *groupSlice) Marshal() *sheets.ValueRange {
	dataStart := int(headerStart) + 1
	values := make([][]interface{}, len(*g)+dataStart) // to skip past headers
	for i := 0; i < int(headerStart); i++ {
		value := []interface{}{}
		for j := 0; j < len(groupHeaders); j++ {
			if i == 0 && j == 0 {
				value = append(value, groupSheetValidationRule)
			} else {
				value = append(value, "")
			}
		}
		values[i] = value
	}
	values[headerStart] = stringSliceToInterfaceSlice(groupHeaders)
	for i, group := range *g {
		values[i+dataStart] = serializeStructToInterfaceSlice(group, groupHeaders)
	}
	return &sheets.ValueRange{
		Range:  allGroupRange,
		Values: values,
	}
}

func (g *groupSlice) Save(c buffalo.Context) error {
	for _, group := range *g {
		if group.ID != nil {
			// we have an update
			m := &models.Group{
				ID:        *group.ID,
				ZoomLink:  stringOrEmpty(group.ZoomLink),
				Published: boolOrFalse(group.Published),
				Archived:  boolOrFalse(group.Archived),
			}
			if !isStringEmpty(group.Name) {
				// we don't want to remove group names
				m.Name = *group.Name
			}
			if _, err := m.Update(c, modelext.GetTx(c), boil.Infer()); err != nil {
				return err
			}
		} else if !isStringEmpty(group.Name) {
			// nil ID + name == insert
			m := &models.Group{
				Name:      *group.Name,
				ZoomLink:  stringOrEmpty(group.ZoomLink),
				Published: boolOrFalse(group.Published),
				Archived:  boolOrFalse(group.Archived),
			}
			if err := m.Insert(c, modelext.GetTx(c), boil.Infer()); err != nil {
				return err
			}
		}
		// anything else we ignore
	}
	return nil
}

// constructs a group slice from the database
func dbGroupSlice(c buffalo.Context) (*groupSlice, error) {
	dbGroups, err := models.Groups().All(c, modelext.GetTx(c))
	if err != nil {
		return nil, err
	}
	groups := make(groupSlice, len(dbGroups))
	for i, group := range dbGroups {
		groups[i] = &innerGroup{
			ID:        &group.ID,
			Name:      &group.Name,
			ZoomLink:  &group.ZoomLink,
			Published: &group.Published,
			Archived:  &group.Archived,
		}
	}
	return &groups, nil
}
