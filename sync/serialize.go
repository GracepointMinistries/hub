package sync

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	sheets "google.golang.org/api/sheets/v4"
)

type typeToCell struct {
	offset int64
	name   string
}

const headerStart int64 = 1

func dataRangeFor(col typeToCell) string {
	column := col.name
	// skip the header
	return column + strconv.Itoa(int(headerStart)+1) + ":" + column
}

func absoluteDataRangeFor(col typeToCell) string {
	column := col.name
	// skip the header
	return column + "$" + strconv.Itoa(int(headerStart)+1) + ":" + column
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
	if len(data) <= int(headerStart)+1 { // we don't even have the header data
		return errors.New("invalid spreadsheet data")
	}
	if err := unmarshalWithHeaders(v, data[headerStart], data[headerStart+1:]); err != nil {
		return err
	}
	return nil
}
