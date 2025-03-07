package tforms

// RadioButton represents a radio button field.
type RadioButton struct {
	BaseField
	Group string
}

func NewRadioButton(name, group string, value string) *RadioButton {
	return &RadioButton{
		BaseField: BaseField{
			FieldName:  name,
			FieldType:  InputTypeRadio,
			FieldValue: value,
		},
		Group: group,
	}
}
