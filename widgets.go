package wtf

type Widget func(f Field) string

func TextInput(f Field) string {
	return ""
}

func (Widget) String() string {
	return ""
}

func (Widget) Foo() {

}
