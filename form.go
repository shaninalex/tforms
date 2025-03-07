package tforms

type IForm interface {
	IsValid() bool
	SetValue(payload map[string]any)
	GetInputs() []IBaseFormControl
	GetActionUrl() string
	Validate()
	GetErrors() map[string]string
	GetMethod() FormMethod
	GetType() FormType
}

type Form struct {
	actionUrl  string
	multipart  bool
	inputs     []IBaseFormControl
	formMethod FormMethod
	formType   FormType
}

func NewForm(action string, multipart bool, controls ...IBaseFormControl) IForm {
	form := &Form{
		actionUrl: action,
		multipart: multipart,
		inputs:    make([]IBaseFormControl, len(controls)),
	}

	for idx, input := range controls {
		form.inputs[idx] = input
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
func (s *Form) GetInputs() []IBaseFormControl { return s.inputs }
func (s *Form) GetActionUrl() string          { return s.actionUrl }
func (s *Form) SetValue(payload map[string]any) {
	//for k, v := range payload {
	//	if _, ok := s.inputs[k]; ok {
	//		s.inputs[k].SetValue(v)
	//	}
	//}
}
func (s *Form) Validate() {
	for k, _ := range s.inputs {
		s.inputs[k].Validate()
	}
}
func (s *Form) GetErrors() map[string]string {
	errors := make(map[string]string, len(s.inputs))
	for _, input := range s.inputs {
		if input.Error() != "" {
			errors[input.Name()] = input.Error()
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
