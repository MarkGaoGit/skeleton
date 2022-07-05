package stringMe

import "math/rand"

// RandomStr 生成指定长度的随机字符串
func RandomStr(strLength int) string {
	strSlice := [31]string{"a", "b", "c", "d", "e", "f", "g", "h", "j", "k", "m", "n", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "2", "3", "4", "5", "6", "7", "8", "9"}
	//切片最大容量索引
	for i := 30; i > 0; i-- {
		num := rand.Intn(i + 1)
		strSlice[i], strSlice[num] = strSlice[num], strSlice[i]
	}

	result := ""
	for i := 0; i < strLength; i++ {
		result += strSlice[i]
	}
	return result
}
