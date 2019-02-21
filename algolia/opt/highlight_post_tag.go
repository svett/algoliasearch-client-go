// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type HighlightPostTagOption struct {
	value string
}

func HighlightPostTag(v string) HighlightPostTagOption {
	return HighlightPostTagOption{v}
}

func (o HighlightPostTagOption) Get() string {
	return o.value
}

func (o HighlightPostTagOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *HighlightPostTagOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = ""
		return nil
	}
	return json.Unmarshal(data, &o.value)
}