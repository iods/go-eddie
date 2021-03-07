package util

// hasIndex Returns true if the index exists for the passed slice.
func hasIndex(slice []interface{}, index int) bool {
	return len(slice) > index
}

// Remove Returns a slice with an element removed at a specified index, if no index exists, the slice returns unchanged.
func Remove(slice []interface{}, index int) []interface{} {

	if ok := hasIndex(slice, index); ok {
		return append(slice[:index], slice[index + 1:]...)
	}

	return slice
}