package arraystring

func Contains(s []string, stringSearch string) bool {
	for _, str := range s {
		if str == stringSearch {
			return true
		}
	}

	return false
}

func IsEmpty(s []string) bool {
	return len(s) == 0
}

func IsNotEmpty(s []string) bool {
	return len(s) > 0
}
