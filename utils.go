package main

import (
	"fmt"
	"math/big"
	"net/url"
	"regexp"
	"strings"
)

const (
	PAUL_NUM = 58
)

var (
	CHARTS   = "fZodR9XQDSUm21yCkr6zBqiveYah8bt4xsWpHnJE7jL5VG3guMTKNPAwcF"
	MAX_CODE = big.NewInt((1 << 51) - 1) // 0x1FFFFFFFFFFFF
	XOR_CODE = big.NewInt(23442827791579)
)

// swapString 交换字符串中两个下标的字符
func swapString(s string, i, j int) string {
	if i >= len(s) || j >= len(s) {
		return s
	}
	runes := []rune(s)
	runes[i], runes[j] = runes[j], runes[i]
	return string(runes)
}

// bvidToAvid 转换 BV 号为 AV 号
func bvidToAvid(bvid string) *big.Int {
	// 两次 swap
	s := swapString(swapString(bvid, 3, 9), 4, 7)
	bv1 := strings.Split(s[3:], "")

	temp := big.NewInt(0)
	for _, c := range bv1 {
		idx := strings.Index(CHARTS, c)
		if idx < 0 {
			panic(fmt.Sprintf("Invalid character: %s", c))
		}
		temp.Mul(temp, big.NewInt(PAUL_NUM))
		temp.Add(temp, big.NewInt(int64(idx)))
	}

	// (temp & MAX_CODE) ^ XOR_CODE
	temp.And(temp, MAX_CODE)
	temp.Xor(temp, XOR_CODE)
	return temp
}

// convertBvUrlToAv 把包含 BV 号的链接转成 AV 链接
func ConvertBvUrlToAv(url string) string {
	re := regexp.MustCompile(`BV[0-9A-Za-z]+`)
	match := re.FindString(url)
	if match == "" {
		return url
	}
	bvid := match
	avid := bvidToAvid(bvid)

	baseUrl := strings.Split(url, "?")[0]
	baseUrl = strings.Replace(baseUrl, bvid, fmt.Sprintf("av%s", avid.String()), 1)
	return baseUrl
}

func CleanBilibiliURL(raw string) string {
	u, err := url.Parse(raw)
	if err != nil {
		return raw // 解析失败就返回原始的
	}

	// 保留路径部分（/video/BVxxxxxx/ 或 /video/avxxxxxx/）
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
