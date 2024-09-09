// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package blockquery

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zclconf/go-cty/cty"
)

func TestQuery(t *testing.T) {
	ctyVal := cty.ObjectVal(map[string]cty.Value{
		"key": cty.StringVal("value"),
	})
	query := "key"
	result, err := Query(ctyVal, cty.DynamicPseudoType, query)
	require.NoError(t, err)
	require.Equal(t, "value", result.String())
}
