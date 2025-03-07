package tforms

// FileUpload represents a file input field.
type FileUpload struct {
	BaseField
	Accept []string // Allowed file types (e.g., [".jpg", ".png"])
}

func NewFileUpload(name string, accept []string) *FileUpload {
	return &FileUpload{
		BaseField: BaseField{
			FieldName: name,
			FieldType: InputTypeFile,
		},
		Accept: accept,
	}
}
