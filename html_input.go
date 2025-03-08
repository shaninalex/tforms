// IHTMLInput represent only html element. There are only getters.

package tforms

import (
	"fmt"
)

type IHTMLInput interface {
	Name() string
	Type() InputType
	HTMLInputType() string // TextInputType
	Error() string
	HasError() bool
	HTMLValue() string
	Options() []Option
	ID() string
	Placeholder() string
	Required() bool
	Label() string
}

type BaseInput struct {
	name        string
	inputType   InputType
	htmlType    TextInputType
	inputError  string
	value       string
	options     []Option
	placeholder string
	required    bool
	label       string
}

func (s *BaseInput) Name() string          { return s.name }
func (s *BaseInput) Type() InputType       { return s.inputType }
func (s *BaseInput) HTMLInputType() string { return string(s.htmlType) }
func (s *BaseInput) Error() string         { return s.inputError }
func (s *BaseInput) HasError() bool        { return len(s.inputError) != 0 }
func (s *BaseInput) HTMLValue() string     { return s.value }
func (s *BaseInput) Options() []Option     { return s.options }
func (s *BaseInput) ID() string            { return fmt.Sprintf("form_%s", s.Name()) }
func (s *BaseInput) Placeholder() string   { return s.placeholder }
func (s *BaseInput) Required() bool        { return s.required }
func (s *BaseInput) Label() string         { return s.label }
