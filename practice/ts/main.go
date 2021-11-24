package ts

import (
	"fmt"
	"net/url"
	"path/filepath"
	"strings"
)

func Run() {
	link := "http://aaa.ccc.com/cs/go/we.html?ddd=.png"
	u, e := url.Parse(link)

	if e == nil && u != nil {
		fmt.Println(u.Path)

		ext := strings.TrimLeft(strings.ToLower(filepath.Ext(u.Path)), ".")
		fmt.Println(ext)
	}

}
