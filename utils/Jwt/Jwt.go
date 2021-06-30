package Jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

//指定加密密钥
var jwtSecret = []byte("N<7orG't|i0j4H6pxJP^_Q.&$(:a[lf2vS]cCOEAZn%3V#TdB1!IhW)*M@k0egL8-U,z+w?/`}y5uXq>F~=;RbY9sm{DKad-zAg")

//Claim是一些实体（通常指的用户）的状态和额外的元数据
type Claims struct {
	//appName string `json:"appName"`
	//appid   string `json:"appid"`
	//module  string `json:"module"`
	//ver     string `json:"ver"`
	Openid string `json:"openid"`
	UserId int    `json:"userId"`
	jwt.StandardClaims
}

// 根据用户的信息生成token
func GenerateToken(userId int, openId string) (string, error) {
	expireTime := time.Now().Add(168 * time.Hour) //设置token有效时间
	claims := Claims{
		UserId: userId,
		Openid: openId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 过期时间
			Issuer:    "go-websocket",    // 指定token发行人
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

func Verify(token string) (bool, *Claims) {
	var (
		code int64
		c    *Claims
		err  error
	)
	if token == "" {
		code = 400
	} else {
		c, err = ParseToken(token)
		if err != nil {
			code = 400
		} else if time.Now().Unix() > c.ExpiresAt {
			code = 400
		}
	}
	//如果token验证不通过，直接终止程序，c.Abort()
	if code != 0 {
		// 返回错误信息
		//fmt.Println("token error")
		//终止程序
		return false, c
	}
	//fmt.Println("token succ")
	return true, c
}
