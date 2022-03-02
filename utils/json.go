package utils

import "github.com/gogf/gf/encoding/gjson"

//返回常用json
func Get_response_json(code int, body string) *gjson.Json {
	json := gjson.New(nil)
	json.Set("code", code)
	json.Set("body", body)
	return json
}
