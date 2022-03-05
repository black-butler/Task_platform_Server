package Test

import (
	"encoding/base64"
	"fmt"
	"platform/utils"
	"testing"
)

func Test_utils(t *testing.T) {

	zhi := "123123"
	jiami, err := utils.AesEncrypt([]byte(zhi), utils.Number_AES)
	if err != nil {
		panic(err)
	}
	fmt.Printf("加密后: %s\n", base64.StdEncoding.EncodeToString(jiami))

	jiema, err := base64.StdEncoding.DecodeString("AhhTnr5qMRYprFrV/hfpZw==")
	if err != nil {
		panic(err)
	}
	origin, err := utils.AesDecrypt(jiema, utils.Number_AES)
	if err != nil {
		panic(err)
	}
	fmt.Printf("解密后明文: %s\n", string(origin))
}
