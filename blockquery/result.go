// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package blockquery

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
)

func NewResult(ty gjson.Type, val any) gjson.Result {
	raw, err := json.Marshal(val)
	if err != nil {
		panic(fmt.Sprintf("could not marshal value: %s", err))
	}
	res := gjson.Result{
		Raw:  string(raw),
		Type: ty,
	}
	switch ty {
	case gjson.Number:
		res.Num = val.(float64)
	case gjson.String:
		res.Str = val.(string)
	case gjson.JSON:
		res.Raw = string(raw)
	}
	return res
}

// NewNumberResults creates results of type Number,
// E.g. If you expect vales to be 1, 2, or 3 then use
// NewNumberResults(1, 2, 3).
func NewNumberResults(vals ...float64) []gjson.Result {
	results := make([]gjson.Result, len(vals))
	for i, val := range vals {
		results[i] = NewResult(gjson.Number, val)
	}
	return results
}

// NewStringResults creates results of type String,
// E.g. If you expect vales to be "a", "b", or "c" then use
// NewStringResults("a", "b", "c").
func NewStringResults(vals ...string) []gjson.Result {
	results := make([]gjson.Result, len(vals))
	for i, val := range vals {
		results[i] = NewResult(gjson.String, val)
	}
	return results
}

// NewTrueResult creates a result of type True.
func NewTrueResult(vals ...bool) gjson.Result {
	return NewResult(gjson.True, true)
}

// NewFalseResult creates a result of type False.
func NewFalseResult(vals ...bool) gjson.Result {
	return NewResult(gjson.False, false)
}

// NewComplexResults creates results for complex types.
//
// E.g. If you expect vales to be an object: {"nested": "value"},
// then use NewComplexResults(map[string]any{"nested": "value"}).
//
// If you want to compare a list, then use NewComplexResults([]any{1, 2, 3}).
func NewComplexResults(vals ...any) []gjson.Result {
	results := make([]gjson.Result, len(vals))
	for i, val := range vals {
		results[i] = NewResult(gjson.JSON, val)
	}
	return results
}
