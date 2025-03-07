package main

import (
	"fmt"
	"github.com/shaninalex/tforms"
	"net/http"
)

var (
	stateOptions = []tforms.SelectableOption[string]{
		{Label: "Alaska", Value: "AK"},
		{Label: "Georgia", Value: "GA"},
		{Label: "Oregon", Value: "OG"},
	}
	sizesOptions = []tforms.SelectableOption[float64]{
		{Label: "SM", Value: 0.43},
		{Label: "MD", Value: 0.24},
		{Label: "LG", Value: 0.83},
	}
)

func main() {
	http.HandleFunc("/", handleIndex)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	//form.Validate()
	//if !form.IsValid() {
	//	fmt.Println("Form is not valid")
	//	for name, err := range form.GetErrors() {
	//		fmt.Printf("[%s]: %s\n", name, err)
	//	}
	//}
	//
	//fmt.Println("Done")
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	form := tforms.NewForm(
		"/post/action",
		false,
		tforms.NewHiddenField("_csrf", "csrf-token-string", true),
		tforms.NewInputField[string]("email", tforms.TextInputEmail, true).
			SetPlaceholder("example@mail.com").
			SetLabel("Employee email"),
		tforms.NewInputField[string]("address", tforms.TextInputText, true).
			SetPlaceholder("st.84, New York, USA").
			SetLabel("Employee address"),
		//tforms.NewSelectField[string]("state", stateOptions, false, nil, true),
		//tforms.NewSelectField[float64]("sizes", sizesOptions, true, nil, true),
	)

	//form.SetValue(map[string]any{
	//	"email": "testtest.com",
	//	"state": []string{"GA", "OG"},
	//	"sizes": []int{24},
	//})
	IndexPage(form).Render(r.Context(), w)
}
