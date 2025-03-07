package tforms

type SelectField[T any] struct {
	BaseField
	Options  []SelectableOption[T]
	Multiple bool
}

type SelectableOption[T any] struct {
	Label    string
	Value    T
	Selected bool
}

func NewSelectField[T any](name string, options []SelectableOption[T], multiple bool, value []T) *SelectField[T] {
	field := &SelectField[T]{
		BaseField: BaseField{
			FieldName:  name,
			FieldType:  InputTypeSelect,
			FieldValue: value,
		},
		Options:  options,
		Multiple: multiple,
	}

	return field
}
