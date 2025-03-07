package tforms

type IForm interface {
	IsValid() bool
	SetValue(payload map[string]any)
	GetInputs() map[string]IBaseFormControl
	GetActionUrl() string
}

type Form struct {
	actionUrl string
	multipart bool
	inputs    map[string]IBaseFormControl
	errors    map[string]string // NOTE: multiple errors per field ?
}

func NewForm(action string, multipart bool, controls ...IBaseFormControl) IForm {
	form := &Form{
		actionUrl: action,
		multipart: multipart,
		inputs:    make(map[string]IBaseFormControl, len(controls)),
	}

	for _, input := range controls {
		form.inputs[input.Name()] = input
	}

	return form
}

func (s *Form) IsValid() bool                          { return len(s.errors) == 0 }
func (s *Form) GetInputs() map[string]IBaseFormControl { return s.inputs }
func (s *Form) GetActionUrl() string                   { return s.actionUrl }
func (s *Form) SetValue(payload map[string]any) {
	for k, v := range payload {
		s.inputs[k].SetValue(v)
	}
}
