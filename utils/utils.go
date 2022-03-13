package utils

const (
	Time_Format = "2006-01-02 15:04:05"
)

//判断某个float64是不是整数
func Check_float64_zheng(zhi float64) bool {
	if zhi == float64(int(zhi)) {
		return true
	} else {
		return false
	}
}
