package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

const UserIDKey = "session"

var SampleSecretKey = []byte("GoLinuxCloudKey")

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// cookie, err := c.Cookie(UserIDKey)
		claims, err := validateToken(c.Writer, c.Request)
		if err != nil {
			data := map[string]interface{}{
				"Code":    http.StatusUnauthorized,
				"Message": "未登录啊",
			}
			c.JSON(http.StatusOK, data)
			c.Abort()
			return
		}
		c.Set("username", claims["username"])
		c.Next()
	}
}

// 登录的时候生产并返回token
// 发送请求的时候带上Token header 进行认证
// https://juejin.cn/post/7110044736848658445
// validateToken()参考 一下
// https://blog.csdn.net/weixin_51507240/article/details/123878839

func validateToken(_ http.ResponseWriter, r *http.Request) (claims jwt.MapClaims, err error) {
	if r.Header["Authorization"] == nil {
		// _, _ = fmt.Fprintf(w, "can not find token in header")
		// w.WriteHeader(http.StatusUnauthorized)
		return nil, errors.New("token error")
	}

	Token := r.Header["Authorization"][0]
	if len(Token) <= 7 {
		// w.WriteHeader(http.StatusUnauthorized)
		return nil, errors.New("token error")
	}

	token, err := jwt.Parse(Token[7:], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// w.WriteHeader(http.StatusUnauthorized)
			return nil, fmt.Errorf("there was an error in parsing")
		}
		return SampleSecretKey, nil
	})

	if token == nil {
		// _, _ = fmt.Fprintf(w, "invalid token")
		// w.WriteHeader(http.StatusUnauthorized)
		return nil, errors.New("token error")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		// _, _ = fmt.Fprintf(w, "couldn't parse claims")
		// w.WriteHeader(http.StatusUnauthorized)
		return nil, errors.New("token error")
	}
	exp := claims["exp"].(float64)
	if int64(exp) < time.Now().Local().Unix() {
		// _, _ = fmt.Fprintf(w, "token expired")
		// w.WriteHeader(http.StatusUnauthorized)
		return nil, errors.New("token error")
	}

	return claims, nil
}
