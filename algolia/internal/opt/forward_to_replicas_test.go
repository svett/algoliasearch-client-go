// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestForwardToReplicas(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.ForwardToReplicasOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.ForwardToReplicas(false),
		},
		{
			opts:     []interface{}{opt.ForwardToReplicas(true)},
			expected: opt.ForwardToReplicas(true),
		},
		{
			opts:     []interface{}{opt.ForwardToReplicas(false)},
			expected: opt.ForwardToReplicas(false),
		},
	} {
		var (
			in  = ExtractForwardToReplicas(c.opts...)
			out opt.ForwardToReplicasOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, *c.expected, out)
	}
}
