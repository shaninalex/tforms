package tforms

type IForm interface {
	AddInput(f IFormControl)
	IsValid() bool
	Validate(payload any)
	GetInputs() []IFormControl
	GetActionUrl() string
}
