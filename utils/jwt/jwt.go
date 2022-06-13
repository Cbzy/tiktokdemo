package JwtLib

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"time"
)

type UserClaims struct {
	*jwt.StandardClaims
	Name string
	Uid  uint
}

//但我觉得如果能把每一个用户单独一个加密方式，有问题就能定位哪个用户了,因为uid可能会被伪造
func JwtSign() (string, error) {
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
			nowTime.Add(time.Minute * 15).Unix(),
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
	return claims.SignedString([]byte(viper.GetString("Jwt.key")))
}
func JwtSignById(id uint) (string, error) {
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
		id,
	})
	// Get the complete, signed token
	return claims.SignedString([]byte(viper.GetString("Jwt.key")))
}
func JwtSignByIdName(id uint, name string) (string, error) {
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
		name,
		id,
	})
	// Get the complete, signed token
	return claims.SignedString([]byte(viper.GetString("Jwt.key")))
}
func JwtParse(jwtStr string) (*UserClaims, error) {

	// keyFunc was not provided.  short circuiting validation
	token, err := jwt.ParseWithClaims(jwtStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		//验证加密方式

		return []byte(viper.GetString("Jwt.key")), nil

	})
	if err != nil {
		return nil, fmt.Errorf("解析失败:%s", err.Error())
		//return nil, errors.New("解析失败,Token不合法")
	}
	// 断言类型
	// For a type to be a Claims object,
	// it must just have a Valid method that determines
	// if the token is invalid for any supported reason
	claim, ok := token.Claims.(*UserClaims)
	// 验证
	if !ok || !token.Valid {
		return nil, errors.New("解析失败,Token不合法")

	}
	return claim, nil
	//fmt.Printf("解析结果: %+v\n", claim)

}
