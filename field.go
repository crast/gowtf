package wtf

type FieldImpl interface {
	Render() string
}

type Validator interface {
	Validate()
}

type Field struct {
	Type       FieldImpl
	Label      string
	Validators []Validator
	Widget     Widget
	Name       string
}
