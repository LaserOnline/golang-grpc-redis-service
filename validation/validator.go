package validation

import "errors"

var (
	ErrEmptyKey   = errors.New("key is required")
	ErrEmptyValue = errors.New("value is required")
)

type SetValidator interface {
	ValidateSet(key, value string) error
	ValidateGet(key string) error
}
