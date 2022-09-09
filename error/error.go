package error

func IsNil(e error) bool {
	return e == nil
}

func IsNotNil(e error) bool {
	return e != nil
}
