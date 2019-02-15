package opt

import "github.com/algolia/algoliasearch-client-go/algolia/opt"

func ExtractAroundRadius(opts ...interface{}) *opt.AroundRadiusOption {
	for _, o := range opts {
		if v, ok := o.(opt.AroundRadiusOption); ok {
			return &v
		}
	}
	return nil
}