package tforms

type InputType string

const (
	InputTypeText     InputType = "text"
	InputTypeTextArea InputType = "textarea"
	InputTypeSelect   InputType = "select"
	InputTypeCheckbox InputType = "checkbox"
	InputTypeRadio    InputType = "radio"
	InputTypeFile     InputType = "file"
)
