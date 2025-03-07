package old

// TextArea represents a multi-line text input.
type TextArea struct {
	BaseField
}

func NewTextArea(name string, value string, required bool) *TextArea {
	return &TextArea{
		BaseField: BaseField{
			FieldName:     name,
			FieldType:     InputTypeTextArea,
			FieldValue:    value,
			FieldRequired: required,
		},
	}
}

func (s *TextArea) Validate() {
	if !s.BaseField.ValidateRequired() {
		return
	}
}
