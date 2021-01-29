// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

// LanguageOption is a wrapper for an Language option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type LanguageOption struct {
	value string
}

// Language wraps the given value into a LanguageOption.
func Language(v string) *LanguageOption {
	return &LanguageOption{v}
}

// Get retrieves the actual value of the option parameter.
func (o *LanguageOption) Get() string {
	if o == nil {
		return ""
	}
	return o.value
}

// MarshalJSON implements the json.Marshaler interface for
// LanguageOption.
func (o LanguageOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// LanguageOption.
func (o *LanguageOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = ""
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *LanguageOption) Equal(o2 *LanguageOption) bool {
	if o == nil {
		return o2 == nil || o2.value == ""
	}
	if o2 == nil {
		return o == nil || o.value == ""
	}
	return o.value == o2.value
}

// LanguageEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func LanguageEqual(o1, o2 *LanguageOption) bool {
	return o1.Equal(o2)
}
