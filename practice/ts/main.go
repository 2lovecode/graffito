package ts

import (
	"fmt"
	"net/url"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

func Run() {
	link := "http://aaa.ccc.com/cs/go/we.html?ddd=.png"
	u, e := url.Parse(link)

	if e == nil && u != nil {
		fmt.Println(u.Path)

		ext := strings.TrimLeft(strings.ToLower(filepath.Ext(u.Path)), ".")
		fmt.Println(ext)
	}

	originalNickname := "我sdfdfsdfs是谁"
	nickname := ""
	nicknameLen := utf8.RuneCountInString(originalNickname)

	if nicknameLen <= 0 {
		nickname = "***"
	} else if nicknameLen == 1 {
		nickname = originalNickname
	} else {
		f, _ := utf8.DecodeRuneInString(originalNickname)
		l, _ := utf8.DecodeLastRuneInString(originalNickname)
		nickname = string(f) + "***" + string(l)
	}

	fmt.Println(nickname)
}
