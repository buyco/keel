package transform

import (
	"database/sql"
	"github.com/buyco/keel/pkg/helper"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/jinzhu/gorm/dialects/postgres"
	"strconv"
)

// SQLNullStringToString converts an SQL NullString to a string
func SQLNullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

// SQLNullStringToString converts an SQL NullString to a Protobuf StringValue
func SQLNullStringToWrapperString(ns sql.NullString) *wrappers.StringValue {
	if ns.Valid {
		return &wrappers.StringValue{
			Value: ns.String,
		}
	}
	return nil
}

// StringToSQLNullString builds an SQL NullString from a string
func StringToSQLNullString(s string) sql.NullString {
	ns := sql.NullString{}
	if len(s) > 0 {
		ns.String = s
		ns.Valid = true
	}
	return ns
}

// StringToSQLNullStringList builds a list of SQL NullString from a list of string
func StringToSQLNullStringList(s []string) []sql.NullString {
	var nullStringList []sql.NullString
	for _, val := range s {
		nullStringList = append(nullStringList, StringToSQLNullString(val))
	}
	return nullStringList
}

// WrapperStringToNullString converts a Protobuf StringValue to an SQL NullString
func WrapperStringToNullString(w *wrappers.StringValue) sql.NullString {
	ns := sql.NullString{}
	if w != nil && len(w.GetValue()) > 0 {
		ns.String = w.GetValue()
		ns.Valid = true
	}
	return ns
}

// WrapperFloatToSQLNullFloat64 converts a Protobuf FloatValue to an SQL NullFloat64
func WrapperFloatToSQLNullFloat64(f *wrappers.FloatValue) sql.NullFloat64 {
	nf := sql.NullFloat64{}
	if f != nil {
		nf.Float64 = float64(f.Value)
		nf.Valid = true
	}
	return nf
}

// SQLNullFloat64ToWrapperFloat converts an SQL NullFloat64 to a Protobuf NullFloat
func SQLNullFloat64ToWrapperFloat(nf sql.NullFloat64) *wrappers.FloatValue {
	if nf.Valid {
		f := wrappers.FloatValue{}
		f.Value = float32(nf.Float64)
		return &f
	}
	return nil
}

// WrapperInt32ToSQLNullString converts a Protobuf Int32Value to an SQL NullString
func WrapperInt32ToSQLNullString(f *wrappers.Int32Value) sql.NullString {
	nf := sql.NullString{}
	if f != nil {
		nf.String = strconv.Itoa(int(f.Value))
		nf.Valid = true
	}
	return nf
}

// WrapperInt32ToSQLNullInt32 converts a Protobuf Int32Value to an SQL NullInt32
func WrapperInt32ToSQLNullInt32(f *wrappers.Int32Value) sql.NullInt32 {
	nf := sql.NullInt32{}
	if f != nil {
		nf.Int32 = f.Value
		nf.Valid = true
	}
	return nf
}

// SQLNullStringToWrapperInt32 converts an SQL NullString to a Protobuf Int32Value
func SQLNullStringToWrapperInt32(nf sql.NullString) *wrappers.Int32Value {
	if nf.Valid {
		f := wrappers.Int32Value{}
		val, err := strconv.ParseInt(nf.String, 10, 32)
		if err == nil {
			f.Value = int32(val)
		}
		return &f
	}
	return nil
}

// SQLNullInt32ToWrapperInt32 converts an SQL NullInt32 to a Protobuf Int32Value
func SQLNullInt32ToWrapperInt32(nf sql.NullInt32) *wrappers.Int32Value {
	if nf.Valid {
		f := wrappers.Int32Value{}
		f.Value = nf.Int32
		return &f
	}
	return nil
}

// WrapperInt64ToSQLNullInt64  converts a Protobuf Int64Value to an SQL NullInt64
func WrapperInt64ToSQLNullInt64(f *wrappers.Int64Value) sql.NullInt64 {
	nf := sql.NullInt64{}
	if f != nil {
		nf.Int64 = f.Value
		nf.Valid = true
	}
	return nf
}

// SQLNullInt64ToWrapperInt64 converts an SQL NullInt64 to a Protobuf Int64Value
func SQLNullInt64ToWrapperInt64(nf sql.NullInt64) *wrappers.Int64Value {
	if nf.Valid {
		f := wrappers.Int64Value{}
		f.Value = nf.Int64
		return &f
	}
	return nil
}

// HstoreToMapString converts an PostgreSQL HSTORE field into a map of strings
func HstoreToMapString(hs postgres.Hstore) map[string]string {
	if hs != nil {
		m := make(map[string]string)
		for key, val := range hs {
			m[key] = *val
		}
		return m
	}
	return nil
}

// MapStringToHStore converts a map of strings into an PostgreSQL HSTORE field
func MapStringToHStore(m map[string]string) (hs postgres.Hstore) {
	if m != nil {
		var hs = make(map[string]*string, 0)
		for key, val := range m {
			tmp := val
			hs[key] = &tmp
		}
		return hs
	}
	return nil
}

// Int32ToSQLNullInt32 converts an int32 to an SQL Int32
func Int32ToSQLNullInt32(i int32) sql.NullInt32 {
	return sql.NullInt32{
		Int32: i,
		Valid: true,
	}
}

// Int32sToUints converts a list of int32 to a list of uint
func Int32sToUints(input []int32) []uint {
	var output []uint
	for _, inputVal := range input {
		output = append(output, uint(inputVal))
	}
	return output
}

// StringToInt32 converts a string to an int32
func StringToInt32(s string) (int32, error) {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, helper.ErrorPrint("Unable to convert string to int")
	}
	return int32(i), nil
}
