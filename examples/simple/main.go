package main

import (
	"fmt"
	"github.com/shaninalex/tforms"
	"net/http"
)

var (
	stateOptions = []*tforms.SelectableOption{
		{Label: "Alaska", Value: "AK"},
		{Label: "Georgia", Value: "GA"},
		{Label: "Oregon", Value: "OG"},
	}
	sizesOptions = []*tforms.SelectableOption{
		{Label: "SM", Value: 0.43},
		{Label: "MD", Value: 0.24},
		{Label: "LG", Value: 0.83},
	}
	departmentsOptions = []*tforms.SelectableOption{
		{Label: "Workshop", Value: "workshop"},
		{Label: "Inventory", Value: "inventory"},
		{Label: "Sales", Value: "sales"},
		{Label: "Shop", Value: "shop"},
		{Label: "E-Shop", Value: "eshop"},
		{Label: "Support", Value: "support"},
	}
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/post/action", handlePostAction)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

/*
Use to add validation tags
*/
type Employee struct {
	Email       string    `form:"email" validate:"required,email"`
	Address     string    `form:"address" validate:"required"`
	Salary      float64   `form:"salary" validate:"required,min=1000,max=10000"`
	State       string    `form:"state" validate:"required"`
	Sizes       []float64 `form:"sizes" validate:"required"`
	Departments []string  `form:"departments" validate:"required,dive,alphanum"`
	Description string    `form:"description"`
}

func handlePostAction(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	form := MakeForm()
	form.SetValue(r.Form)

	// NOTE: this library is not for validating. It's only for rendering.
	// 		 For validating use:
	//		 - github.com/go-playground/validator/v10
	//		 - github.com/go-ozzo/ozzo-validation
	// 		 or other. Form takes map[string]string as an error and match errors for fields.
	form.SetErrors(map[string]string{
		"email": "invalid email",
	})
	IndexPage(form, true).Render(r.Context(), w)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	form := MakeForm()
	IndexPage(form, false).Render(r.Context(), w)
}

func MakeForm() tforms.IForm {
	form := tforms.NewForm(
		"/post/action",
		false,
		tforms.NewHiddenField("_csrf", "csrf-token-string", true),
		tforms.NewInputField("email", tforms.TextInputEmail, true).
			SetPlaceholder("example@mail.com").
			SetLabel("Employee email"),
		tforms.NewInputField("address", tforms.TextInputText, true).
			SetPlaceholder("st.84, New York, USA").
			SetLabel("Employee address"),
		tforms.NewInputField("salary", tforms.TextInputNumber, true).
			SetPlaceholder("$10000").
			SetLabel("Salary"),
		tforms.NewSelectField("state", stateOptions, false, true).
			SetPlaceholder("Select state...").
			SetLabel("State where employee leaves"),
		tforms.NewSelectField("sizes", sizesOptions, false, true).
			SetPlaceholder("Select sizes...").
			SetLabel("Employee size for work clothes"),
		tforms.NewSelectableField("departments", tforms.InputTypeCheckbox, departmentsOptions, true, true).
			SetLabel("Departments"),
		tforms.NewTextArea("description", true).
			SetLabel("Employee description"),
	)
	form.SetMethod(tforms.FormMethodPost)
	return form
}
