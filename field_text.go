package tforms

// InputField represents text-based input fields.
type InputField[T any] struct {
	BaseField
	HtmlType string
}

// NewInputField creates a new text-based input field.
func NewInputField[T any](name, textInputType string, value T) *InputField[T] {
	return &InputField[T]{
		BaseField: BaseField{
			FieldName:  name,
			FieldType:  InputTypeText,
			FieldValue: value,
		},
		HtmlType: textInputType,
	}
}

func (s *InputField[T]) SetValue(value any) {
	s.BaseField.SetValue(value)
}
