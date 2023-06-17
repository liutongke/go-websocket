package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 指定加密密钥
var jwtSecret = []byte("D65GQbemixhAcu4gKsMzZRdHl8pO0BU2fPVEI7wFJknCyWLrY9ovjST3atXq1NfscETSuEl5YUzMjkb0HBnsFbInneMVyj0JEHF")

// Claim是一些实体（通常指的用户）的状态和额外的元数据
type Claims struct {
	AppName string `json:"appName"`
	Appid   string `json:"appid"`
	Module  string `json:"module"`
	Ver     string `json:"ver"`
	Openid  string `json:"openid"`
	Uid     int    `json:"uid"`
	jwt.StandardClaims
}

// 根据用户的信息生成token
func GenerateToken(AppName, appid, module, ver, openid string, uid int) (string, error) {
	//设置token有效时间
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
		AppName: AppName,
		Appid:   appid,
		Module:  module,
		Ver:     ver,
		Openid:  openid,
		Uid:     uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 过期时间
			Issuer:    "nat-x",           // 指定token发行人
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret) //该方法内部生成签名字符串，再用于获取完整、已签名的token
	return token, err
}

// 根据传入的token值获取到Claims对象信息，（进而获取其中的用户名和密码）
func ParseToken(token string) (*Claims, error) {
	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空间。
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err

}

func Verify(token string) (*Claims, bool) {
	if token == "" {
		return nil, false
	}
	claims, err := ParseToken(token)
	if err != nil || time.Now().Unix() > claims.ExpiresAt {
		return claims, false
	}
	return claims, true
}
