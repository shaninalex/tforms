package tforms

// Checkbox represents a checkbox field.
type Checkbox struct {
	BaseField
	Checked bool
}

func NewCheckbox(name string, checked bool) *Checkbox {
	return &Checkbox{
		BaseField: BaseField{
			FieldName:  name,
			FieldType:  InputTypeCheckbox,
			FieldValue: checked,
		},
		Checked: checked,
	}
}
