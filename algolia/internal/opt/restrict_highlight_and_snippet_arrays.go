package opt

import "github.com/algolia/algoliasearch-client-go/algolia/opt"

func ExtractRestrictHighlightAndSnippetArrays(opts ...interface{}) *opt.RestrictHighlightAndSnippetArraysOption {
	for _, o := range opts {
		if v, ok := o.(opt.RestrictHighlightAndSnippetArraysOption); ok {
			return &v
		}
	}
	return nil
}