package tforms

// CheckboxField represents a checkbox field.
type CheckboxField[T comparable] struct {
	*BaseInput
	Options []*SelectableOption[T]

	// If false it will render radiobuttons instead
	Multiple bool
}

func NewSelectableField[T comparable](name string, selectableType InputType, options []*SelectableOption[T], multiple bool, required bool) *CheckboxField[T] {
	s := &CheckboxField[T]{
		BaseInput: &BaseInput{
			name:      name,
			inputType: selectableType,
			required:  required,
		},
		Options:  options,
		Multiple: multiple,
	}
	s.makeOptions(options)
	return s
}

func (s *CheckboxField[T]) Base() *BaseInput { return s.BaseInput }
func (s *CheckboxField[T]) Validate()        {}
func (s *CheckboxField[T]) SetValue(v any) IBaseFormControl {
	if values, ok := v.([]T); ok {
		for _, value := range values {
			for _, option := range s.Options {
				if option.Value == value {
					option.Selected = true
					break
				}
			}
		}
		s.makeOptions(s.Options)
		return s
	}

	s.BaseInput.inputError = "Invalid value"
	return s
}
func (s *CheckboxField[T]) SetLabel(v string) IBaseFormControl {
	s.BaseInput.label = v
	return s
}
func (s *CheckboxField[T]) SetPlaceholder(v string) IBaseFormControl {
	s.BaseInput.placeholder = v
	return s
}

// private

func (s *CheckboxField[T]) makeOptions(options []*SelectableOption[T]) {
	strOptions := make([]Option, len(options))
	for i, option := range options {
		strOptions[i] = option.ToOption()
	}
	s.BaseInput.options = strOptions
}
