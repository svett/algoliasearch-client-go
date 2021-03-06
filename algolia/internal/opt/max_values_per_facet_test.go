// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestMaxValuesPerFacet(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.MaxValuesPerFacetOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.MaxValuesPerFacet(100),
		},
		{
			opts:     []interface{}{opt.MaxValuesPerFacet(0)},
			expected: opt.MaxValuesPerFacet(0),
		},
		{
			opts:     []interface{}{opt.MaxValuesPerFacet(1)},
			expected: opt.MaxValuesPerFacet(1),
		},
		{
			opts:     []interface{}{opt.MaxValuesPerFacet(-42)},
			expected: opt.MaxValuesPerFacet(-42),
		},
	} {
		var (
			in  = ExtractMaxValuesPerFacet(c.opts...)
			out opt.MaxValuesPerFacetOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, *c.expected, out)
	}
}
