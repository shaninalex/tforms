package tforms

import (
	"fmt"
)

type SelectField[T comparable] struct {
	*BaseInput
	Options []*SelectableOption[T]

	// Multiple in select should render something like "tags". But it required more JS
	Multiple bool
}

func NewSelectField[T comparable](name string, options []*SelectableOption[T], multiple bool, required bool) *SelectField[T] {
	s := &SelectField[T]{
		BaseInput: &BaseInput{
			name:      name,
			inputType: InputTypeSelect,
			required:  required,
		},
		Multiple: multiple,
	}
	s.Options = s.copy(options)
	s.makeOptions(s.Options)
	return s
}

func (s *SelectField[T]) Base() *BaseInput { return s.BaseInput }
func (s *SelectField[T]) SetValue(v any) IBaseFormControl {
	for _, option := range s.Options {
		option.Selected = false
	}
	for _, option := range s.Options {
		if fmt.Sprintf("%v", option.Value) == fmt.Sprintf("%v", v) {
			option.Selected = true
			break
		}
	}
	s.makeOptions(s.Options)
	return s
}
func (s *SelectField[T]) SetLabel(v string) IBaseFormControl {
	s.BaseInput.label = v
	return s
}
func (s *SelectField[T]) SetPlaceholder(v string) IBaseFormControl {
	s.BaseInput.placeholder = v
	return s
}
func (s *SelectField[T]) makeOptions(options []*SelectableOption[T]) {
	strOptions := make([]Option, len(options))
	for i, option := range options {
		strOptions[i] = option.ToOption()
	}
	s.BaseInput.options = strOptions
}
func (s *SelectField[T]) copy(options []*SelectableOption[T]) []*SelectableOption[T] {
	newOptions := make([]*SelectableOption[T], len(options))
	for i, o := range options {
		newOptions[i] = &SelectableOption[T]{
			Label:    o.Label,
			Selected: o.Selected,
			Value:    o.Value,
		}
	}
	return newOptions
}
func (s *SelectField[T]) SetError(e string) {
	s.BaseInput.inputError = e
}
