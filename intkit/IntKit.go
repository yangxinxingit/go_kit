package intkit

// @Title 判断多个整形是否为0
// @Description
// @param strs
// usage:
//	IntIsZero(0) return false
func IntIsZero(ints ...int) bool {
	if len(ints) == 0 {
		return false
	}
	for _, v := range ints {
		if v != 0 {
			return false
		}
	}
	return true
}