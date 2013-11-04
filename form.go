package wtf

import "net/url"

type Form map[string]Field

func (f Form) Bind(args ...interface{}) BoundForm {
	var formdata url.Values
	if len(args) >= 1 {
		formdata = coerceFormData(args[0])
	}
	return BoundForm{
		Form:     f,
		FormData: formdata,
	}
}

type BoundForm struct {
	Form
	FormData url.Values
}
