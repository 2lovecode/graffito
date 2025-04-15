package top100

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	res := make([]string, 0)
	path := make([]byte, 0, len(digits))
	// 数字到字母的映射表
	letterMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	var backtrack func(index int)

	backtrack = func(index int) {
		if index == len(digits) {
			res = append(res, string(path))
			return
		}
		digit := digits[index]
		letters := letterMap[digit]
		for i := 0; i < len(letters); i++ {
			path = append(path, letters[i])
			backtrack(index + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return res
}
