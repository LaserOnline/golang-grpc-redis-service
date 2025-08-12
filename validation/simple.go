package validation

import "strings"

type SimpleSetValidator struct{}

func NewSimpleSetValidator() *SimpleSetValidator { return &SimpleSetValidator{} }

func (v *SimpleSetValidator) ValidateSet(key, value string) error {
	key = strings.TrimSpace(key)
	if key == "" {
		return ErrEmptyKey
	}
	if strings.TrimSpace(value) == "" {
		return ErrEmptyValue
	}
	return nil
}

func (v *SimpleSetValidator) ValidateGet(key string) error {
	if strings.TrimSpace(key) == "" {
		return ErrEmptyKey
	}
	return nil
}
