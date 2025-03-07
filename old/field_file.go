package old

// FileUpload represents a file input field.
type FileUpload struct {
	BaseField
	Accept []string // Allowed file types (e.g., [".jpg", ".png"])
}

func NewFileUpload(name string, accept []string, required bool) *FileUpload {
	return &FileUpload{
		BaseField: BaseField{
			FieldName:     name,
			FieldType:     InputTypeFile,
			FieldRequired: required,
		},
		Accept: accept,
	}
}

func (s *FileUpload) Validate() {
	if !s.BaseField.ValidateRequired() {
		return
	}
}
