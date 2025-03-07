package tforms

// Checkbox represents a checkbox field.
type Checkbox struct {
	BaseField
	Checked bool
}

func NewCheckbox(name string, checked bool, required bool) *Checkbox {
	return &Checkbox{
		BaseField: BaseField{
			FieldName:     name,
			FieldType:     InputTypeCheckbox,
			FieldValue:    checked,
			FieldRequired: required,
		},
		Checked: checked,
	}
}

func (s *Checkbox) Validate() {
	if !s.BaseField.ValidateRequired() {
		return
	}
}
