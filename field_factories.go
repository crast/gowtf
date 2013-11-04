package wtf

import "crast.us/wtf/fields"

func coerceValidator(input interface{}) []Validator {
	if v, ok := input.(Validator); ok {
		return []Validator{v}
	}
	if vlist, ok := input.([]Validator); ok {
		return vlist
	}
	return []Validator{}
}

func makeField(t FieldImpl, defaultWidget Widget, args []interface{}) Field {
	label := ""
	validators := []Validator{}
	l := len(args)
	if l >= 1 {
		if _v, ok := args[0].(string); ok {
			label = _v
		}
	}
	if l >= 2 {
		validators = coerceValidator(args[1])
	}
	return Field{
		Type:       t,
		Widget:     defaultWidget,
		Label:      label,
		Validators: validators,
	}
}

func StringField(args ...interface{}) Field {
	return makeField(fields.TextField{}, TextInput, args)
}

func IntegerField(args ...interface{}) Field {
	return makeField(fields.IntegerField{}, TextInput, args)
}
