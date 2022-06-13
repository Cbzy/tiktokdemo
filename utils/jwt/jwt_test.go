package JwtLib

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"testing"
	"time"
)

func tT(t *testing.T) {

}

var signedString string

func TestName(t *testing.T) {
	nowTime := time.Now()
	//Create a new Token.  Takes a signing method
	//	NotBefore int64  `json:"nbf,omitempty"`
	//	Subject   string `json:"sub,omitempty"`
	//nbf(Not Before)：如果当前时间在nbf里的时间之前
	//，则Tocken不被接受；一般都会留一些余地
	//，比如几分钟；是否使用是可选的
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserClaims{
		&jwt.StandardClaims{
			"douyin",
			nowTime.Add(time.Minute * 5).Unix(),
			"uini2123",
			nowTime.Unix(),
			"mind",
			5,
			"test",
		},
		"mind",
		123,
	})
	// Get the complete, signed token
	signedString, err := claims.SignedString([]byte("asd!@#.'.'p.@#ASDasd123"))
	if err != nil {

	}
	log.Println(signedString)
}
func TestJwtSign(t *testing.T) {
	str, err := JwtSignById(123)
	if err != nil {
		log.Println(err)
	}
	println(str)

	value, err := JwtParse(str)
	if err != nil {
		log.Println(err)
	}
	log.Println(value)

}
func TestValidJWT(t *testing.T) {

	jwtStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJkb3V5aW4iLCJleHAiOjE2NTUwMjAzMTcsImp0aSI6InVpbmkyMTIzIiwiaWF0IjoxNjU1MDIwMDE3LCJpc3MiOiJtaW5kIiwibmJmIjo1LCJzdWIiOiJ0ZXN0IiwiTmFtZSI6Im1pbmQiLCJVaWQiOjEyM30.fzC01yaSXqvyy_FsPEvwBg-hrN_QdVAA7Pk7rPE-TDY"

	// keyFunc was not provided.  short circuiting validation
	token, err := jwt.ParseWithClaims(jwtStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		//验证加密方式
		return []byte("asd!@#.'.'p.@#ASDasd123"), nil
	})

	if err != nil {
		t.Error("解析失败:", err.Error())
		return
	}
	// 断言类型
	// For a type to be a Claims object,
	// it must just have a Valid method that determines
	// if the token is invalid for any supported reason
	claim, ok := token.Claims.(*UserClaims)
	// 验证
	if !ok || !token.Valid {
		t.Error("解析失败,Token不合法!")
		return
	}
	fmt.Printf("解析结果: %+v\n", claim)
}
