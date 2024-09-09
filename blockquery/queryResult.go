// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package blockquery

import (
	"fmt"
	"reflect"

	"github.com/tidwall/gjson"
)

type ResultCompareFunc func(gjson.Result, ...gjson.Result) (bool, string, error)

// Exists is a compare function that checks if the result exists.
func Exists(got gjson.Result, _ ...gjson.Result) (bool, string, error) {
	if !got.Exists() {
		return false, "returned value does not exist but expected", nil
	}
	return true, "", nil
}

// NotExists is a compare function that checks if the result does not exist.
func NotExists(got gjson.Result, _ ...gjson.Result) (bool, string, error) {
	if got.Exists() {
		return false, "returned value exists but not expected", nil
	}
	return true, "", nil
}

// EachIsOneOfAndMustExist is a compare function that checks if the result is one of the expected values and must exist.
// This is useful when the result is an array and you want to check if each element is one of the expected values.
func EachIsOneOfAndMustExist(got gjson.Result, expected ...gjson.Result) (bool, string, error) {
	if !got.Exists() {
		return false, "returned value does not exist but expected", nil
	}
	return EachIsOneOf(got, expected...)
}

// EachIsOneOf is a compare function that checks if the result is one of the expected values.
// This is useful when the result is an array and you want to check if each element is one of the expected values.
func EachIsOneOf(got gjson.Result, expected ...gjson.Result) (bool, string, error) {
	if !got.Exists() {
		return true, "", nil
	}
	if len(got.Array()) == 1 {
		var message string
		ok := compareResults(got, expected)
		if !ok {
			message = fmt.Sprintf("returned value `%s` not in expected values `%s`", got, expected)
		}
		return ok, message, nil
	}
	results := make([]bool, len(got.Array()))
	for i, qr := range got.Array() {
		results[i] = compareResults(qr, expected)
	}
	if !allTrue(results...) {
		return false, fmt.Sprintf("returned value `%s` not in expected values `%v`", got, expected), nil
	}
	return true, "", nil
}

// IsOneOfAndMustExist is a compare function that checks if the result is one of the expected values and must exist.
func IsOneOfAndMustExist(got gjson.Result, expected ...gjson.Result) (bool, string, error) {
	if !got.Exists() {
		return false, "returned value does not exist but expected", nil
	}
	return IsOneOf(got, expected...)
}

// IsOneOf is a compare function that checks if the result is one of the expected values.
func IsOneOf(got gjson.Result, expected ...gjson.Result) (bool, string, error) {
	if !got.Exists() {
		return true, "", nil
	}
	ok := compareResults(got, expected)
	if !ok {
		return false, fmt.Sprintf("returned value `%s` not in expected values `%v`", got, expected), nil
	}
	return ok, "", nil
}

// compareResults compares the result with the expected values.
func compareResults(got gjson.Result, want []gjson.Result) bool {
	for _, w := range want {
		ok := reflect.DeepEqual(got.Value(), w.Value())
		if ok {
			return true
		}
	}
	return false
}

// allTrue checks if all the values are true.
func allTrue(in ...bool) bool {
	for _, b := range in {
		if !b {
			return false
		}
	}
	return true
}
