package tforms

// TextArea represents a multi-line text input.
type TextArea struct {
	BaseField
}

func NewTextArea(name string, value string) *TextArea {
	return &TextArea{
		BaseField: BaseField{
			FieldName:  name,
			FieldType:  InputTypeTextArea,
			FieldValue: value,
		},
	}
}
