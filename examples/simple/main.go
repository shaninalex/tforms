package main

import (
	"github.com/shaninalex/tforms"
	"log"
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
		tforms.NewInputField[string]("email", "email", ""),
		tforms.NewSelectField[string]("state", stateOptions, true, nil),
		tforms.NewSelectField[float64]("sizes", sizesOptions, true, nil),
	)

	form.SetValue(map[string]any{
		"email": "test@test.com",
		"state": []string{"GA", "OG"},
		"sizes": []float64{0.24},
	})

	log.Println(form)
}
