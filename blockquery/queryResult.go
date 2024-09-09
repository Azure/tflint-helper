// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package blockquery

import (
	"fmt"

	"github.com/zclconf/go-cty/cty"
)

type ResultCompareFunc func(cty.Value, ...cty.Value) (bool, string, error)

// IsNotNull is a compare function that checks if the result exists.
func IsNotNull(got cty.Value, _ ...cty.Value) (bool, string, error) {
	if got.IsNull() {
		return false, "returned value is null but not expected to be", nil
	}
	return true, "", nil
}

// IsNull is a compare function that checks if the result does not exist.
func IsNull(got cty.Value, _ ...cty.Value) (bool, string, error) {
	if !got.IsNull() {
		return false, "returned value is not null but expected to be", nil
	}
	return true, "", nil
}

// EachIsOneOf is a compare function that checks if the result is one of the expected values.
// This is useful when the result is an array and you want to check if each element is one of the expected values.
func EachIsOneOf(got cty.Value, expected ...cty.Value) (bool, string, error) {
	// go through them all and check if they are one of the expected values
}

// IsOneOf is a compare function that checks if the result is one of the expected values.
func IsOneOf(got cty.Value, expected ...cty.Value) (bool, string, error) {
	ok := compareResults(got, expected)
	if !ok {
		return false, fmt.Sprintf("returned value `%s` not in expected values `%v`", got, expected), nil
	}
	return ok, "", nil
}

// compareResults compares the result with the expected values.
func compareResults(got cty.Value, want []cty.Value) bool {
	for _, w := range want {
		ok := w.Equals(got).True()
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
