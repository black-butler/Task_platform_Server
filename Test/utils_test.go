package Test

import (
	"encoding/base64"
	"fmt"
	"platform/utils"
	"testing"
)

func Test_utils(t *testing.T) {

	zhi := "JASHGAHDJASBG"
	jiami, err := utils.AesEncrypt([]byte(zhi), utils.Number_AES)
	if err != nil {
		panic(err)
	}
	fmt.Printf("加密后: %s\n", base64.StdEncoding.EncodeToString(jiami))

	origin, err := utils.AesDecrypt([]byte("kP9ly8nvzhBtnPxbfM3rwA=="), utils.Number_AES)
	if err != nil {
		panic(err)
	}
	fmt.Printf("解密后明文: %s\n", string(origin))
}
