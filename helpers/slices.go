package helpers

func RemoveStringFromSlice(slice []string, value string) []string {
	for i, s := range slice {
		if s == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}

	return slice
}
