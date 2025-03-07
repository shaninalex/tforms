package tforms

import "fmt"

type IBaseFormControl interface {
	Name() string
	Type() InputType
	Value() any
	SetValue(value any)
	Validate()
	ValidateRequired() bool
	SetError(e string)
	HasError() bool
	GetError() string
}

// BaseField provides common properties for all form fields.
type BaseField struct {
	FieldName     string
	FieldType     InputType
	FieldValue    any
	FieldError    string
	FieldRequired bool
	// TODO: validation rules like this: https://github.com/go-ozzo/ozzo-validation
}

func (s *BaseField) Name() string       { return s.FieldName }
func (s *BaseField) Type() InputType    { return s.FieldType }
func (s *BaseField) Value() any         { return s.FieldValue }
func (s *BaseField) SetValue(value any) { s.FieldValue = value }
func (s *BaseField) Validate() {
	// NOTE: concrete implementations SHOULD override this method!
	panic(fmt.Sprintf("field \"%s\" not implemented correctly", s.FieldName))
}
func (s *BaseField) SetError(e string) { s.FieldError = e }
func (s *BaseField) HasError() bool    { return s.FieldError != "" }
func (s *BaseField) ValidateRequired() bool {
	if s.FieldRequired && s.FieldValue == nil {
		s.FieldError = ErrorRequired
		return false
	}
	return true
}
func (s *BaseField) GetError() string { return s.FieldError }
