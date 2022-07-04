package ts

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

func Run() {

	Trans()

	fmt.Println(strconv.Itoa(0))

	if strings.Contains("https://ppp.ttt.com/aaa/ccc", "ppp.ttt.com") {
		fmt.Println("good good")
	}

	fmt.Println(UrlAddAndReplaceQuery("http://hot.t.ziroom.com/2022/pano-vr/?inv_no=772908110&inv_no=772908110&zrVrId=858922680266300906ffff", map[string]string{
		"inv_no": "1234",
	}))

	fmt.Println(time.Now().Format("2006年01月02日 15:04:05"))
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

	var count int = 0
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for i := 0; i < 10000; i++ {
			count++
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 0; i < 10000; i++ {
			count++
		}
		wg.Done()
	}()

	wg.Wait()

	fmt.Println(count)

	decoded, err := base64.StdEncoding.DecodeString("aaa")

	result := ""

	if err == nil {
		reader := bytes.NewReader(decoded)
		if reader != nil {
			gzReader, err := gzip.NewReader(reader)
			if err == nil && gzReader != nil {
				res, err := ioutil.ReadAll(reader)
				if err == nil {
					result = string(res)
				}
			}
		}

	}
	fmt.Println(err, result)
}

func UrlAddAndReplaceQuery(link string, extraQuery map[string]string) string {
	if strings.Contains(link, "http") {
		if linkEntry, err := url.Parse(link); err == nil && linkEntry != nil {
			query := linkEntry.Query()
			for key, val := range extraQuery {
				query.Set(key, val)
			}
			linkEntry.RawQuery = query.Encode()
			return linkEntry.String()

		}
	}

	return link
}

func Trans() {
	original := "https://hot.ziroom.com/2022/pano-vr/?zrVrId=16559046272702861973"

	originalUrl, err := url.Parse(original)

	if err == nil && originalUrl != nil {
		originalUrl.Path = strings.TrimRight(originalUrl.Path, "/")
		fmt.Println(originalUrl.String())

	}

}
