package Test

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"testing"
)

func Test_base_byte(t *testing.T) {
	f, err := os.OpenFile("2.jpg", os.O_RDONLY, 0600)
	defer f.Close()
	if err != nil {
		panic(err)
		return
	}

	contentByte, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
		return
	}

	base64_string := base64.StdEncoding.EncodeToString(contentByte)
	println(base64_string)

	jiema, err := base64.StdEncoding.DecodeString(base64_string)
	if err != nil {
		panic(err)
	}

	f, err = os.Create("3.jpg")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	_, err = f.Write(jiema)
}
