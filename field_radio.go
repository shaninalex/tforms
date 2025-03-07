package tforms

// TODO: Basically radio button type field is the selectable field of subtype checkbox.
//		 Without ability to select multiple options. This type look's like redundant.

// RadioButton represents a radio button field.
type RadioButton struct {
	BaseField
	Group string
}

func NewRadioButton(name, group string, value string, required bool) *RadioButton {
	return &RadioButton{
		BaseField: BaseField{
			FieldName:     name,
			FieldType:     InputTypeRadio,
			FieldValue:    value,
			FieldRequired: required,
		},
		Group: group,
	}
}

func (s *RadioButton) Validate() {
	if !s.BaseField.ValidateRequired() {
		return
	}
}
