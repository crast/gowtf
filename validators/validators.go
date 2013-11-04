package validators

type Required struct{}

func (Required) Validate() {

}

type Email struct{}

func (Email) Validate() {

}

type NumberRange struct{
	Min int
	Max int
}

func (NumberRange) Validate() {}
