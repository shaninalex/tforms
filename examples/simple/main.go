package main

import (
	"fmt"
	"github.com/shaninalex/tforms"
)

func main() {

	stateOptions := []tforms.SelectableOption[string]{
		{Label: "Alaska", Value: "AK"},
		{Label: "Georgia", Value: "GA"},
		{Label: "Oregon", Value: "OG"},
	}

	sizesOptions := []tforms.SelectableOption[float64]{
		{Label: "SM", Value: 0.43},
		{Label: "MD", Value: 0.24},
		{Label: "LG", Value: 0.83},
	}

	form := tforms.NewForm(
		"/post/action",
		false,
		tforms.NewInputField[string]("email", tforms.TextInputEmail, nil, true),
		tforms.NewSelectField[string]("state", stateOptions, false, nil, true),
		tforms.NewSelectField[float64]("sizes", sizesOptions, true, nil, true),
	)

	form.SetValue(map[string]any{
		"email": "testtest.com",
		"state": []string{"GA", "OG"},
		"sizes": []int{24},
	})

	form.Validate()
	if !form.IsValid() {
		fmt.Println("Form is not valid")
		for name, err := range form.GetErrors() {
			fmt.Printf("[%s]: %s\n", name, err)
		}
	}

	fmt.Println("Done")
}
