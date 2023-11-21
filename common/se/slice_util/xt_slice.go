package slice_util

// ContainsString checks whether a slice contains a string.
func ContainsString(slice []string, s string) bool {
	return FindString(slice, s) != -1
}

// FindString finds the index of a string in a slice.
func FindString(slice []string, s string) int {
	for index, item := range slice {
		if item == s {
			return index
		}
	}
	return -1
}

// RemoveString removes a string from a slice.
func RemoveString(slice []string, s string) []string {
	var newSlice []string
	copy(newSlice, slice)
	for i := 0; i < len(newSlice); i++ {
		if newSlice[i] == s {
			// 将后面的元素往前移一位，相当于删除指定元素
			copy(newSlice[i:], newSlice[i+1:])
			newSlice = newSlice[:len(newSlice)-1] // 切片长度减一
			i--
		}
	}
	return newSlice
}

// RemoveFirstString removes first string from a slice. O(n)
func RemoveFirstString(slice []string, s string) []string {
	var newSlice []string
	copy(newSlice, slice)
	for i := 0; i < len(newSlice); i++ {
		if newSlice[i] == s {
			// 将后面的元素往前移一位，相当于删除指定元素
			copy(newSlice[i:], newSlice[i+1:])
			newSlice = newSlice[:len(newSlice)-1] // 切片长度减一
			return newSlice
		}
	}
	return newSlice
}
