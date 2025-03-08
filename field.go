package tforms

type IBaseFormControl interface {
	Base() *BaseInput
	Validate()
	SetValue(v any) IBaseFormControl
	SetLabel(v string) IBaseFormControl
	SetPlaceholder(v string) IBaseFormControl
}
