package tforms

import "log"

// InputField represents text-based input fields.
type InputField[T any] struct {
	BaseField
	HtmlType TextInputType
}

// NewInputField creates a new text-based input field.
func NewInputField[T any](name string, textInputType TextInputType, value T) *InputField[T] {
	return &InputField[T]{
		BaseField: BaseField{
			FieldName:  name,
			FieldType:  InputTypeText,
			FieldValue: value,
		},
		HtmlType: textInputType,
	}
}

func (s *InputField[T]) Validate() {
	log.Println("validation in InputField using override method", s.Name())

	if s.HtmlType == TextInputEmail {
		v, ok := s.FieldValue.(string)
		if ok {
			if ok := rxEmail.MatchString(v); !ok {
				s.SetError("incorrect email")
			}
		}
	}
}
