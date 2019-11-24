package util

func CheckIntInArray(i int64, a []int64) bool {
	for _, b := range a {
		if i == b {
			return true
		}
	}
	return false
}
