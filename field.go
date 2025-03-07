package tforms

type IBaseFormControl interface {
	Name() string
	Type() InputType
	Value() any
	SetValue(value any)
}

// BaseField provides common properties for all form fields.
type BaseField struct {
	FieldName  string
	FieldType  InputType
	FieldValue any
}

func (f *BaseField) Name() string       { return f.FieldName }
func (f *BaseField) Type() InputType    { return f.FieldType }
func (f *BaseField) Value() any         { return f.FieldValue }
func (f *BaseField) SetValue(value any) { f.FieldValue = value }
