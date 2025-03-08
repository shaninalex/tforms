package tforms

import "fmt"

// InputField represents text-based input fields.
type InputField[T any] struct {
	*BaseInput
}

// NewInputField creates a new text-based input field.
func NewInputField[T any](name string, textInputType TextInputType, required bool) *InputField[T] {
	return &InputField[T]{
		BaseInput: &BaseInput{
			name:      name,
			inputType: InputTypeText,
			htmlType:  textInputType,
			required:  required,
		},
	}
}

// public

func (s *InputField[T]) Validate()        {}
func (s *InputField[T]) Base() *BaseInput { return s.BaseInput }
func (s *InputField[T]) SetValue(v any) IBaseFormControl {
	s.BaseInput.value = fmt.Sprintf("%v", v)
	return s
}
func (s *InputField[T]) SetLabel(v string) IBaseFormControl {
	s.BaseInput.label = v
	return s
}
func (s *InputField[T]) SetPlaceholder(v string) IBaseFormControl {
	s.BaseInput.placeholder = v
	return s
}
