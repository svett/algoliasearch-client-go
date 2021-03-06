// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestSnippetEllipsisText(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.SnippetEllipsisTextOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.SnippetEllipsisText("…"),
		},
		{
			opts:     []interface{}{opt.SnippetEllipsisText("")},
			expected: opt.SnippetEllipsisText(""),
		},
		{
			opts:     []interface{}{opt.SnippetEllipsisText("content of the string value")},
			expected: opt.SnippetEllipsisText("content of the string value"),
		},
	} {
		var (
			in  = ExtractSnippetEllipsisText(c.opts...)
			out opt.SnippetEllipsisTextOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, *c.expected, out)
	}
}
