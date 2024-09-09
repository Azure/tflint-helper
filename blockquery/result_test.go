// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package blockquery

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func TestNewStringResult(t *testing.T) {
	source := []byte(`
{
		"key": "value"
}`)
	results := NewStringResults("value")
	got := gjson.GetBytes(source, "key")
	assert.Equal(t, results[0].Value(), got.Value())
}

func TestNewIntResult(t *testing.T) {
	source := []byte(`
{
		"key": 1.0
}`)
	results := NewNumberResults(1.0)
	got := gjson.GetBytes(source, "key")
	assert.Equal(t, results[0].Value(), got.Value())
}

func TestNewTrueResult(t *testing.T) {
	source := []byte(`
{
		"key": true
}`)
	result := NewTrueResult()
	got := gjson.GetBytes(source, "key")
	assert.Equal(t, result.Value(), got.Value())
}

func TestNewFalseResult(t *testing.T) {
	source := []byte(`
{
		"key": false
}`)
	result := NewFalseResult()
	got := gjson.GetBytes(source, "key")
	assert.Equal(t, result.Value(), got.Value())
}

func TestNewJsonResult(t *testing.T) {
	source := []byte(`
{
		"key": {
			"nested": "value"
		}
}`)
	results := NewComplexResults(map[string]any{
		"nested": "value",
	})
	got := gjson.GetBytes(source, "key")
	assert.Equal(t, results[0].Value(), got.Value())
}

func TestNewJsonResultArray(t *testing.T) {
	source := []byte(`
{
		"key": [1, 2, 3]
}`)
	results := NewComplexResults([]any{1, 2, 3})
	got := gjson.GetBytes(source, "key")
	assert.Equal(t, results[0].Value(), got.Value())
}
