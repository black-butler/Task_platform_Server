package Route

import (
	"encoding/base64"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"platform/Config"
	"platform/Data"
	"platform/log"
	"platform/utils"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func init() {
	g := g.Server()
	group := g.Group("/public")
	group.POST("/sign_in", ZhuCe) //注册接口
	group.POST("/Login", Login)   //登录
	group.POST("/Token_login", Token_login)
	group.GET("/test", test)
}

//注册账号
func ZhuCe(r *ghttp.Request) {

	number := r.GetString("number")
	Aes_password := r.GetString("password")
	if utf8.RuneCountInString(number) > 20 || utf8.RuneCountInString(Aes_password) > 70 || number == "" || Aes_password == "" {
		r.Response.WriteJson(utils.Get_response_json(1, "字段不合法"))
		return
	}

	jiema, err := base64.StdEncoding.DecodeString(Aes_password)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "安全校验失败"))
		return
	}

	password, err := utils.AesDecrypt(jiema, utils.Number_AES)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "安全校验失败"))
		return
	}

	err = Data.Data_Add_user(number, string(password))
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, err.Error()))
		return
	}

	r.Response.WriteJson(utils.Get_response_json(0, "注册成功"))
}

//登录账号
func Login(r *ghttp.Request) {

	number := r.GetString("number")
	Aes_password := r.GetString("password")

	if utf8.RuneCountInString(number) > 20 || utf8.RuneCountInString(Aes_password) > 70 || number == "" || Aes_password == "" {
		r.Response.WriteJson(utils.Get_response_json(1, "字段不合法"))
		return
	}

	jiema, err := base64.StdEncoding.DecodeString(Aes_password)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "安全校验失败"))
		return
	}

	password, err := utils.AesDecrypt(jiema, utils.Number_AES)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "安全校验失败"))
		return
	}

	user, err := Data.Data_Get_user(number, string(password))
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, err.Error()))
		return
	}

	sessionid := r.Session.Id()
	r.Session.Set(Config.Session_user, user)

	token := strconv.Itoa(user.Id) + "`" + strconv.Itoa(int(time.Now().AddDate(0, 0, 10).Unix()))
	Encrypt_token, _ := utils.AesEncrypt([]byte(token), utils.Token_AES)
	Base_64_Encrypt_token := base64.StdEncoding.EncodeToString(Encrypt_token)
	log.Login_log().Line().Println("用户", user.Number, "登录成功", sessionid+"----"+Base_64_Encrypt_token)

	r.Response.WriteJson(utils.Get_response_json(0, sessionid+"----"+Base_64_Encrypt_token))
}

//token登录
func Token_login(r *ghttp.Request) {

	Aes_Token := r.GetString("token")
	if Aes_Token == "" || utf8.RuneCountInString(Aes_Token) > 50 {
		r.Response.WriteJson(utils.Get_response_json(1, "字段不合法"))
		return
	}

	jiema, err := base64.StdEncoding.DecodeString(Aes_Token)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "安全校验失败"))
		return
	}

	token, err := utils.AesDecrypt(jiema, utils.Token_AES)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "安全校验失败"))
		return
	}

	s := strings.Split(string(token), "`")
	if len(s) != 2 {
		r.Response.WriteJson(utils.Get_response_json(1, "安全校验失败"))
		return
	}

	userid := s[0]
	token_time, err := strconv.Atoi(s[1])
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "安全校验失败"))
		return
	}

	//检测是否过期
	if time.Now().Unix() >= int64(token_time) {
		//过期
		r.Response.WriteJson(utils.Get_response_json(1, "token过期"))
		return
	}

	user, err := Data.Data_Get_userid(userid)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "账号不存在"))
		return
	}

	sessionid := r.Session.Id()
	r.Session.Set(Config.Session_user, user)

	r.Response.WriteJson(utils.Get_response_json(0, sessionid))
}

func test(r *ghttp.Request) {
	r.Session.Set("123", "1231")
	r.Response.Write("asd")
}
