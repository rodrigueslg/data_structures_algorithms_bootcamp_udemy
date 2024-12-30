package slices

func MergeSortedSlices(slice1, slice2 []int) []int {
	merged := make([]int, 0, len(slice1)+len(slice2))
	i, j := 0, 0

	for i < len(slice1) && j < len(slice2) {
		if slice1[i] < slice2[j] {
			merged = append(merged, slice1[i])
			i++
		} else {
			merged = append(merged, slice2[j])
			j++
		}
	}

	for i < len(slice1) {
		merged = append(merged, slice1[i])
		i++
	}

	for j < len(slice2) {
		merged = append(merged, slice2[j])
		j++
	}

	return merged
}
