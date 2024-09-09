// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package blockquery

import (
	"fmt"

	"github.com/tidwall/gjson"
	"github.com/zclconf/go-cty/cty"
	ctyjson "github.com/zclconf/go-cty/cty/json"
)

// Query takes a cty value and a gjson query string and returns the result.
// The cty.Value is marshalled to JSON and the query is run against the resulting JSON.
func Query(val cty.Value, ty cty.Type, query string) (gjson.Result, error) {
	jsonbytes, err := ctyjson.Marshal(val, ty)
	if err != nil {
		return gjson.Result{}, fmt.Errorf("could not marshal cty value: %s", err)
	}
	return gjson.GetBytes(jsonbytes, "value."+query), nil
}
