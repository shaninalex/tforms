package tforms

import "fmt"

type Option struct {
	Value    string
	Label    string
	Selected bool
}

type SelectableOption[T comparable] struct {
	Label    string
	Value    T
	Selected bool
}

func (s *SelectableOption[T]) ToOption() Option {
	return Option{
		Label:    s.Label,
		Selected: s.Selected,
		Value:    fmt.Sprintf("%v", s.Value),
	}
}
