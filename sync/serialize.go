package sync

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/GracepointMinistries/hub/modelext"
	"github.com/GracepointMinistries/hub/models"
	"github.com/gobuffalo/buffalo"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	sheets "google.golang.org/api/sheets/v4"
)

type typeToCell struct {
	offset int64
	name   string
}

var (
	allGroupRange      string
	groupColumnBegin   int64
	groupColumnEnd     int64
	allUserRange       string
	userColumnBegin    int64
	userColumnEnd      int64
	groupHeaders       []string
	userHeaders        []string
	groupHeadersLookup map[string]typeToCell
	userHeadersLookup  map[string]typeToCell
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
	userRangeEnd := string(rune('A') + rune(group.NumField()))
	// 'SHEET TITLE'!X:Y
	allUserRange = makeRange(usersTitle, userRangeBegin+":"+userRangeEnd)
}

func dataRangeFor(col typeToCell) string {
	column := col.name
	// skip the header
	return column + "2:" + column
}

func stringSliceToInterfaceSlice(a []string) []interface{} {
	b := make([]interface{}, len(a))
	for i := range a {
		b[i] = a[i]
	}
	return b
}

func pointerToInterface(v interface{}) interface{} {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr {
		// we can't work with non-pointers
		return nil
	}
	if val.IsNil() {
		return nil
	}
	return val.Elem().Interface()
}

func serializeStructToInterfaceSlice(v interface{}, headers []string) []interface{} {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr {
		// we can't work with non-pointers
		return nil
	}
	// we're just going to assume proper bounds
	ind := reflect.Indirect(val)
	values := make([]interface{}, ind.NumField())
	for i := 0; i < ind.NumField(); i++ {
		field := ind.FieldByName(headers[i])
		values[i] = pointerToInterface(field.Interface())
	}
	return values
}

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
	values := make([][]interface{}, len(*g)+1) // +1 for the header
	values[0] = stringSliceToInterfaceSlice(groupHeaders)
	for i, group := range *g {
		values[i+1] = serializeStructToInterfaceSlice(group, groupHeaders)
	}
	return &sheets.ValueRange{
		Range:  allGroupRange,
		Values: values,
	}
}

func (g *groupSlice) ToDB() []*models.Group {
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

type innerUser struct {
	// order matters here, it needs to be the same as
	// the expected header ordering
	ID      *int    `json:"id"`
	Name    *string `json:"name"`
	Email   *string `json:"email"`
	Blocked *bool   `json:"blocked"`
	Group   *string `json:"group"`
}

type userSlice []*innerUser

func newUserSlice() *userSlice {
	return &userSlice{}
}

func (u *userSlice) Unmarshal(values *sheets.ValueRange) error {
	return unmarshal(u, values)
}

func (u *userSlice) Marshal() *sheets.ValueRange {
	values := make([][]interface{}, len(*u)+1) // +1 for the header
	values[0] = stringSliceToInterfaceSlice(userHeaders)
	for i, user := range *u {
		values[i+1] = serializeStructToInterfaceSlice(user, userHeaders)
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

func setString(v reflect.Value, s string) error {
	if v.Kind() != reflect.Ptr {
		return errors.New("expected field struct to be nil")
	}
	switch v.Type().Elem().Kind() {
	case reflect.Int:
		value, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(&value))
	case reflect.String:
		v.Set(reflect.ValueOf(&s))
	case reflect.Bool:
		// google spreadsheets uses TRUE and FALSE
		var value bool
		switch s {
		case "TRUE":
			value = true
		case "FALSE":
			value = false
		default:
			return fmt.Errorf("invalid boolean string '%s'", s)
		}
		v.Set(reflect.ValueOf(&value))
	default:
		return errors.New("unhandled conversion kind")
	}
	return nil
}

func setInt(v reflect.Value, i int) error {
	if v.Kind() != reflect.Ptr {
		return errors.New("expected field struct to be nil")
	}
	switch v.Type().Elem().Kind() {
	case reflect.Int:
		v.Set(reflect.ValueOf(&i))
	case reflect.String:
		value := strconv.Itoa(i)
		v.Set(reflect.ValueOf(&value))
	case reflect.Bool:
		value := false
		if i > 0 {
			value = true
		}
		v.Set(reflect.ValueOf(&value))
	default:
		return errors.New("unhandled conversion kind")
	}
	return nil
}

func setBool(v reflect.Value, b bool) error {
	if v.Kind() != reflect.Ptr {
		return errors.New("expected field struct to be nil")
	}
	switch v.Type().Elem().Kind() {
	case reflect.Int:
		value := 0
		if b {
			value = 1
		}
		v.Set(reflect.ValueOf(&value))
	case reflect.String:
		value := strconv.FormatBool(b)
		v.Set(reflect.ValueOf(&value))
	case reflect.Bool:
		v.Set(reflect.ValueOf(&b))
	default:
		return errors.New("unhandled conversion kind")
	}
	return nil
}

func unmarshalWithHeaders(v interface{}, headers []interface{}, data [][]interface{}) error {
	// sanity check the arguments passed in
	if reflect.TypeOf(v).Kind() != reflect.Ptr ||
		reflect.TypeOf(v).Elem().Kind() != reflect.Slice ||
		reflect.TypeOf(v).Elem().Elem().Kind() != reflect.Ptr {
		return errors.New("unmarshaled value must be a pointer to a slice of pointers")
	}

	// double check the headers match the field names
	typ := reflect.TypeOf(v).Elem().Elem()
	fieldStruct := reflect.Indirect(reflect.New(typ.Elem()))
	if fieldStruct.NumField() != len(headers) {
		return fmt.Errorf("invalid headers, unmarshaling structure has %d elements, headers have %d elements", fieldStruct.NumField(), len(headers))
	}
	for i := 0; i < fieldStruct.NumField(); i++ {
		field := fieldStruct.Type().Field(i)
		header := field.Name
		dataHeader, ok := headers[i].(string)
		if !ok {
			return fmt.Errorf("expected string in header, found %T", headers[i])
		}
		if dataHeader != header {
			return fmt.Errorf("expected %s in header, found %s", header, dataHeader)
		}
	}

	// proceed with the rest of the unmarshaling
	val := reflect.Indirect(reflect.ValueOf(v))
	for _, row := range data {
		element := reflect.New(typ.Elem())
		for i, column := range row {
			if column == nil {
				continue
			}
			name := reflect.Indirect(element).Type().Field(i).Name
			field := reflect.Indirect(element).FieldByIndex([]int{i})
			switch c := column.(type) {
			// via the docs, this can only be a string, int, bool, or float64 (which we have none of)
			case string:
				setString(field, c)
			case int:
				setInt(field, c)
			case bool:
				setBool(field, c)
			default:
				return fmt.Errorf("cannot convert field '%s' from '%T' to '%v'", name, column, field.Type())
			}
		}
		val.Set(reflect.Append(val, element))
	}
	return nil
}

func unmarshal(v interface{}, values *sheets.ValueRange) error {
	data := values.Values
	if len(data) == 0 {
		return errors.New("invalid group data")
	}
	if err := unmarshalWithHeaders(v, data[0], data[1:]); err != nil {
		return err
	}
	return nil
}
