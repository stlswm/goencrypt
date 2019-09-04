package aes128

import (
	"github.com/stlswm/goencrypt/aes128"
	"strconv"
	"testing"
)

func TestEncrypt(t *testing.T) {
	for i := 1; i < 100000; i++ {
		encStr := aes128.Encrypt(strconv.Itoa(i), "e102397d02cddc19", "b7e4eba12397a82b")
		t.Log(encStr)
		decStr := aes128.Decrypt(encStr, "e102397d02cddc19", "b7e4eba12397a82b")
		t.Log(decStr)
		n, _ := strconv.Atoi(decStr)
		if i != n {
			t.Error("err:" + strconv.Itoa(i))
		}
	}
	// encStr := aes128.Encrypt("2&31&0&", "e102397d02cddc19", "b7e4eba12397a82b")
	// t.Log(encStr)
	// t.Log(url.QueryEscape(encStr))
	// t.Log(url.QueryUnescape(url.QueryEscape(encStr)))
}

func TestDecrypt(t *testing.T) {
	decStr := aes128.Decrypt("y93amfPgKhoBR6F90l+d6Q", "e102397d02cddc19", "b7e4eba12397a82b")
	t.Log(decStr)
}
