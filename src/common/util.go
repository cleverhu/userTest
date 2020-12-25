package common

import (
	"crypto/md5"
	"fmt"
	"hash/crc32"
)

func Hashcode(s string) int64 {
	v := int64(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}

	return 0
}

func MD5(str string) string {
	bytes := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", bytes)
}
