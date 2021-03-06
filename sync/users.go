package sync

import (
	"fmt"
	"reflect"

	"github.com/GracepointMinistries/hub/modelext"
	"github.com/GracepointMinistries/hub/models"
	"github.com/gobuffalo/buffalo"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	sheets "google.golang.org/api/sheets/v4"
)

var (
	allUserRange      string
	userColumnBegin   int64
	userColumnEnd     int64
	userHeaders       []string
	userHeadersLookup map[string]typeToCell
)

func init() {
	user := reflect.TypeOf(userSlice{}).Elem().Elem()
	userHeaders = make([]string, user.NumField())
	userHeadersLookup = make(map[string]typeToCell, user.NumField())
	userColumnBegin = 0
	for i := 0; i < user.NumField(); i++ {
		header := user.FieldByIndex([]int{i}).Name
		userColumnEnd = int64(i)
		userHeaders[i] = header
		userHeadersLookup[header] = typeToCell{
			offset: userColumnEnd,
			name:   string(rune('A') + rune(i)),
		}
	}
	userRangeBegin := "A"
	userRangeEnd := string(rune('A') + rune(user.NumField()))
	// 'SHEET TITLE'!X:Y
	allUserRange = makeRange(usersTitle, userRangeBegin+":"+userRangeEnd)
}

type innerUser struct {
	// order matters here, it needs to be the same as
	// the expected header ordering
	ID      *int    `json:"id"`
	Email   *string `json:"email"`
	Name    *string `json:"name"`
	Group   *string `json:"group"`
	Blocked *bool   `json:"blocked"`
}

type userSlice []*innerUser

func newUserSlice() *userSlice {
	return &userSlice{}
}

func (u *userSlice) Unmarshal(values *sheets.ValueRange) error {
	return unmarshal(u, values)
}

// generate these dynamically at some point
const (
	usersGroupNameValidation = "EQ(SUM(ARRAYFORMULA(COUNTIF('Groups'!$B$3:$B, 'Users'!$E$3:$E)))-COUNTA('Users'!$E$3:$E), 0)"
)

var userSheetValidationRule = fmt.Sprintf("=IF(%s, \"VALID\", \"INVALID\")", usersGroupNameValidation)

func (u *userSlice) Marshal() *sheets.ValueRange {
	dataStart := int(headerStart) + 1
	values := make([][]interface{}, len(*u)+dataStart) // to skip past the haders
	for i := 0; i < int(headerStart); i++ {
		value := []interface{}{}
		for j := 0; j < len(groupHeaders); j++ {
			if i == 0 && j == 0 {
				value = append(value, userSheetValidationRule)
			} else {
				value = append(value, "")
			}
		}
		values[i] = value
	}
	values[headerStart] = stringSliceToInterfaceSlice(userHeaders)
	for i, user := range *u {
		values[i+dataStart] = serializeStructToInterfaceSlice(user, userHeaders)
	}
	return &sheets.ValueRange{
		Range:  allUserRange,
		Values: values,
	}
}

func (u *userSlice) ToDB() []*models.User {
	return nil
}

// constructs a user slice from the database
func dbUserSlice(c buffalo.Context) (*userSlice, error) {
	dbUsers, err := models.Users(
		qm.Load(models.UserRels.Groups, models.GroupWhere.Archived.EQ(false)),
	).All(c, modelext.GetTx(c))
	if err != nil {
		return nil, err
	}
	users := make(userSlice, len(dbUsers))
	for i, user := range dbUsers {
		group := modelext.GroupForUser(user)
		inner := &innerUser{
			ID:      &user.ID,
			Name:    &user.Name,
			Email:   &user.Email,
			Blocked: &user.Blocked,
		}
		if group != nil {
			inner.Group = &group.Name
		}
		users[i] = inner
	}
	return &users, nil
}
