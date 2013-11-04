package wtf

import "net/url"

var coerceFuncs []func(interface{}) url.Values

func init() {
	coerceFuncs = append(coerceFuncs, defaultCoerceFormData)
}

func coerceFormData(formdata interface{}) url.Values {
	for _, f := range coerceFuncs {
		val := f(formdata)
		if val != nil {
			return val
		}
	}
	return nil
}

func defaultCoerceFormData(formdata interface{}) url.Values {
	switch f := formdata.(type) {
	case url.Values:
		return f
	case map[string][]string:
		return f
	case map[string]string:
		return formdataFromFlatMap(f)
	}
	return nil
}

func formdataFromFlatMap(m map[string]string) url.Values {
	formdata := url.Values{}
	for k, v := range m {
		formdata[k] = []string{v}
	}
	return formdata
}
