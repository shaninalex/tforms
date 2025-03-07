package tforms

type SelectField[T comparable] struct {
	BaseField
	Options  []SelectableOption[T]
	Multiple bool
}

// SelectableOption type for all selectable input types ( select, checkbox )
type SelectableOption[T comparable] struct {
	Label    string
	Value    T
	Selected bool
}

func NewSelectField[T comparable](name string, options []SelectableOption[T], multiple bool, value []T, required bool) *SelectField[T] {
	field := &SelectField[T]{
		BaseField: BaseField{
			FieldName:     name,
			FieldType:     InputTypeSelect,
			FieldValue:    value,
			FieldRequired: required,
		},
		Options:  options,
		Multiple: multiple,
	}
	return field
}

func (s *SelectField[T]) Validate() {
	if !s.BaseField.ValidateRequired() {
		return
	}

	// Ensure FieldValue is of the correct type
	selectedValues, ok := s.FieldValue.([]T)
	if !ok {
		s.FieldError = ErrorInvalidType
		return
	}

	// Validate that each selected value exists in the options
	for _, selected := range selectedValues {
		valid := false
		for _, option := range s.Options {
			if option.Value == selected {
				valid = true
				break
			}
		}
		if !valid {
			s.FieldError = ErrorInvalidSelection
			return
		}
	}

	// Validate multiple selections if Multiple is false
	if !s.Multiple && len(selectedValues) > 1 {
		s.FieldError = ErrorMultipleSelection
	}
}
