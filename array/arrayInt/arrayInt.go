package arrayint

func Contains(i []int, numSearch int) bool {
	for _, num := range i {
		if num == numSearch {
			return true
		}
	}

	return false
}
