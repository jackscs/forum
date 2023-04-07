package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// TokenExpireDuration 统一过期时间
const TokenExpireDuration = time.Hour * 24

// Secret 加密密钥
var Secret = []byte("夏天夏天悄悄过去")

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成Token
func GenToken(userID int64, username string) (string, error) {
	claims := MyClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),                          //生效时间
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), //失效时间
			Issuer:    "destiny",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //生成token
	//fmt.Println(token)
	return token.SignedString(Secret) //加密token  传输过程传送加密后的token
}

// ParseToken 解析Token
func ParseToken(keyToken string) (*MyClaims, error) {
	//需要使用new关键字初始化 否则会报空指针
	var claim = new(MyClaims)
	token, err := jwt.ParseWithClaims(keyToken, claim, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { //校验token
		return claim, nil
	}
	return nil, errors.New("invalid token")
}
