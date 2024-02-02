package tools

func FindCommonElements[T comparable](slice, otherSlice []T) []T {
	var commonElements []T

	for _, elementOfSlice := range slice {
		for _, elementOfOtherSlice := range otherSlice {
			if elementOfSlice == elementOfOtherSlice {
				commonElements = append(commonElements, elementOfOtherSlice)
				break
			}
		}
	}

	return commonElements
}
