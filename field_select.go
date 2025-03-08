package tforms

import "fmt"

type SelectField[T comparable] struct {
	*BaseInput
	Options  []SelectableOption[T]
	Multiple bool
}

func NewSelectField[T comparable](name string, options []SelectableOption[T], multiple bool, required bool) *SelectField[T] {
	strOptions := make([]Option, len(options))
	for i, option := range options {
		strOptions[i] = option.ToOption()
	}
	return &SelectField[T]{
		BaseInput: &BaseInput{
			name:      name,
			inputType: InputTypeSelect,
			required:  required,
			options:   strOptions,
		},
		Options:  options,
		Multiple: multiple,
	}
}

func (s *SelectField[T]) HTML() *BaseInput { return s.BaseInput }
func (s *SelectField[T]) HasError() bool   { return len(s.BaseInput.inputError) != 0 }
func (s *SelectField[T]) Validate()        {}
func (s *SelectField[T]) Name() string     { return s.BaseInput.Name() }
func (s *SelectField[T]) Error() string    { return s.BaseInput.inputError }
func (s *SelectField[T]) SetValue(v any) IBaseFormControl {
	fmt.Println("Setting the value")
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
