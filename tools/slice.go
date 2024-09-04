package tools

func RemoveInt32Slice(slice []int32, ele int32) []int32 {
	if len(slice) == 0 {
		return slice
	}
	for i, v := range slice {
		if v == ele {
			slice = append(slice[:i], slice[i+1:]...)
			return slice
		}
	}
	return slice
}
