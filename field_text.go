package tforms

// InputField represents text-based input fields.
type InputField[T any] struct {
	BaseField
	HtmlType TextInputType
}

// NewInputField creates a new text-based input field.
func NewInputField[T any](name string, textInputType TextInputType, value any, required bool) *InputField[T] {
	return &InputField[T]{
		BaseField: BaseField{
			FieldName:     name,
			FieldType:     InputTypeText,
			FieldValue:    value,
			FieldRequired: required,
		},
		HtmlType: textInputType,
	}
}

func (s *InputField[T]) Validate() {
	if !s.BaseField.ValidateRequired() {
		return
	}

	if s.HtmlType == TextInputEmail {
		v, ok := s.FieldValue.(string)
		if ok {
			if ok := rxEmail.MatchString(v); !ok {
				s.SetError("incorrect email")
			}
		}
	}
}
