package tforms

type SelectField[T comparable] struct {
	*BaseInput
	Options []SelectableOption[T]

	// Multiple in select should render something like "tags". But it required more JS
	Multiple bool
}

func NewSelectField[T comparable](name string, options []SelectableOption[T], multiple bool, required bool) *SelectField[T] {
	strOptions := make([]Option, len(options))
	for i, option := range options {
		strOptions[i] = option.ToOption()
	}
	return &SelectField[T]{
		BaseInput: &BaseInput{
			name:      name,
			inputType: InputTypeSelect,
			required:  required,
			options:   strOptions,
		},
		Options:  options,
		Multiple: multiple,
	}
}

func (s *SelectField[T]) Base() *BaseInput { return s.BaseInput }
func (s *SelectField[T]) Validate()        {}
func (s *SelectField[T]) SetValue(v any) IBaseFormControl {
	return s
}
func (s *SelectField[T]) SetLabel(v string) IBaseFormControl {
	s.BaseInput.label = v
	return s
}
func (s *SelectField[T]) SetPlaceholder(v string) IBaseFormControl {
	s.BaseInput.placeholder = v
	return s
}
