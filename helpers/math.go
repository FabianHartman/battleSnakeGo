package helpers

func AbsValue(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
