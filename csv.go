// Package csv provides `Marshal` and `UnMarshal` encoding functions for CSV(Comma Seperated Value) data.
// This package is built on the the standard library's encoding/csv.
package csv

import (
	"reflect"
)

func skipField(f reflect.StructField) bool {
	if f.Tag.Get("csv") == "-" {
		return true
	}

	return false
}

// fieldHeaderName returns the header name to use for the given StructField
// This can be a user defined name (via the Tag) or a default name.
func fieldHeaderName(f reflect.StructField) (string, bool) {
	h := f.Tag.Get("csv")

	if h == "-" {
		return "", false
	}

	// If there is no tag set, use a default name
	if h == "" {
		return f.Name, true
	}

	return h, true
}

func typeHeaders(t reflect.Type) (out []string) {
	l := t.NumField()

	for x := 0; x < l; x++ {
		f := t.Field(x)
		h, ok := fieldHeaderName(f)
		if ok {
			out = append(out, h)
		}
	}

	return
}
