package tforms

type IForm interface {
	IsValid() bool
	SetValue(payload map[string]any)
	GetInputs() map[string]IBaseFormControl
	GetActionUrl() string
	Validate()
	GetErrors() map[string]string
	GetMethod() FormMethod
	GetType() FormType
}

type Form struct {
	actionUrl  string
	multipart  bool
	inputs     map[string]IBaseFormControl
	formMethod FormMethod
	formType   FormType
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

func (s *Form) IsValid() bool {
	for _, i := range s.inputs {
		if i.HasError() {
			return false
		}
	}
	return true
}
func (s *Form) GetInputs() map[string]IBaseFormControl { return s.inputs }
func (s *Form) GetActionUrl() string                   { return s.actionUrl }
func (s *Form) SetValue(payload map[string]any) {
	for k, v := range payload {
		if _, ok := s.inputs[k]; ok {
			s.inputs[k].SetValue(v)
		}
	}
}
func (s *Form) Validate() {
	for k, _ := range s.inputs {
		s.inputs[k].Validate()
	}
}
func (s *Form) GetErrors() map[string]string {
	errors := make(map[string]string, len(s.inputs))
	for _, input := range s.inputs {
		if input.GetError() != "" {
			errors[input.Name()] = input.GetError()
		}
	}

	return errors
}
func (s *Form) GetMethod() FormMethod {
	if s.formMethod == "" {
		return FormMethodPost
	}
	return s.formMethod
}
func (s *Form) GetType() FormType {
	if s.formType == "" {
		return FormTypeXWWWFormUrlEncoded
	}
	return s.formType
}
