package types

func ErrDetails(err interface{}) *AnyValue {
	anyValue := AnyValue(err)
	return &anyValue
}
