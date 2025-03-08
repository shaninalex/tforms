package tforms

import "fmt"

// CheckboxField represents a checkbox field.
type CheckboxField[T comparable] struct {
	*BaseInput
	Options []SelectableOption[T]

	// If false it will render radiobuttons instead
	Multiple bool
}

func NewSelectableField[T comparable](name string, selectableType InputType, options []SelectableOption[T], multiple bool, required bool) *CheckboxField[T] {
	// TODO: check for InputTypeCheckbox and InputTypeSelect
	strOptions := make([]Option, len(options))
	for i, option := range options {
		strOptions[i] = option.ToOption()
	}
	return &CheckboxField[T]{
		BaseInput: &BaseInput{
			name:      name,
			inputType: selectableType,
			required:  required,
			options:   strOptions,
		},
		Options:  options,
		Multiple: multiple,
	}
}

func (s *CheckboxField[T]) HTML() *BaseInput { return s.BaseInput }
func (s *CheckboxField[T]) HasError() bool   { return len(s.BaseInput.inputError) != 0 }
func (s *CheckboxField[T]) Validate()        {}
func (s *CheckboxField[T]) Name() string     { return s.BaseInput.Name() }
func (s *CheckboxField[T]) Error() string    { return s.BaseInput.inputError }
func (s *CheckboxField[T]) SetValue(v any) IBaseFormControl {
	fmt.Println("Setting the value")
	return s
}
func (s *CheckboxField[T]) SetLabel(v string) IBaseFormControl {
	s.BaseInput.label = v
	return s
}
func (s *CheckboxField[T]) SetPlaceholder(v string) IBaseFormControl {
	s.BaseInput.placeholder = v
	return s
}
