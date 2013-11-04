package main

import "fmt"
import "crast.us/wtf"
import v "crast.us/wtf/validators"
import "crast.us/wtf/fields"

func main() {
	MyForm := wtf.Form{
		"username": wtf.StringField("Username", v.Required{}),
		"age": wtf.IntegerField("Username", chain(v.Required{}, v.NumberRange{Min: 18, Max: 100})),
		"fooflaf": wtf.Field{
			Type: fields.IntegerField{},
			Label: "Foo",
			Validators: chain(v.NumberRange{18, 42}),
			Widget: wtf.TextInput,
		},
	}
	form := MyForm.Bind()
	fmt.Printf("%#v\n\n%#v", MyForm, form)
}

func chain(args ...wtf.Validator) []wtf.Validator {
	/// TODO make this a builtin
	return args
}
