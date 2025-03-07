package tforms

type IBaseFormControl interface {
	HTML() *BaseInput
	HasError() bool
	Validate()
	Name() string
	GetError() string
	SetValue(v any) IBaseFormControl
	SetLabel(v string) IBaseFormControl
	SetPlaceholder(v string) IBaseFormControl
}
