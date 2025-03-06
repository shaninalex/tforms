package tforms

type InputType string

const (
	InputTypeText     InputType = "text"
	InputTypeTextArea InputType = "textarea"
	InputTypeSelect   InputType = "select"
	InputTypeCheckbox InputType = "checkbox"
)

type IFormControl interface {
	HasError() bool
	GetError() string
	SetError(v string)
	GetName() string
	SetName(v string)
	GetValue() string
	SetValue(v string)
	GetPlaceholder() string
	SetPlaceholder(v string)
	GetLabel() string
	SetLabel(v string)
	GetType() InputType
	GetHTMLType() string
	SetHTMLType(v string)
	GetID() string
	IsValid() (bool, string)
	Validate(value any)
}

type ISelectableFormControl interface {
	IFormControl
	SetOptions(options map[string]string)
	GetOptions() map[string]any
}
