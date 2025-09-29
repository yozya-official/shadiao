package main

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

var (
	XOR_CODE = int64(23442827791579)
	MAX_CODE = int64(2251799813685247)
	CHARTS   = "FcwAPNKTMug3GV5Lj7EJnHpWsx4tb8haYeviqBz6rkCy12mUSDQX9RdoZf"
	PAUL_NUM = int64(58)
)

func swapString(s string, x, y int) string {
	chars := []rune(s)
	chars[x], chars[y] = chars[y], chars[x]
	return string(chars)
}

func bvid2Avid(bvid string) (avid int64) {
	s := swapString(swapString(bvid, 3, 9), 4, 7)
	bv1 := string([]rune(s)[3:])
	temp := int64(0)
	for _, c := range bv1 {
		idx := strings.IndexRune(CHARTS, c)
		temp = temp*PAUL_NUM + int64(idx)
	}
	avid = (temp & MAX_CODE) ^ XOR_CODE
	return
}

// convertBvUrlToAv 把包含 BV 号的链接转成 AV 链接
func ConvertBvUrlToAv(url string) string {
	re := regexp.MustCompile(`BV[0-9A-Za-z]+`)
	match := re.FindString(url)
	if match == "" {
		return url
	}
	bvid := match
	avid := bvid2Avid(bvid)

	baseUrl := strings.Split(url, "?")[0]
	baseUrl = strings.Replace(baseUrl, bvid, fmt.Sprintf("av%d", avid), 1)
	return baseUrl
}

func CleanBilibiliURL(raw string) string {
	u, err := url.Parse(raw)
	if err != nil {
		return raw // 解析失败就返回原始的
	}

	cleanURL := u.Scheme + "://" + u.Host + u.Path

	// 如果有必要，可以只保留 "p" 参数（分P时有用）
	query := u.Query()
	if p := query.Get("p"); p != "" {
		cleanURL = cleanURL + "?p=" + p
	}

	// 去掉结尾的 "/"
	cleanURL = strings.TrimSuffix(cleanURL, "/")

	return cleanURL
}
