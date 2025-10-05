package service

import (
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
