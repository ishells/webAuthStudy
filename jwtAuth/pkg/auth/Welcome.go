package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	// 获取用户发送的cookie
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// 如果没有cookie，返回未授权错误
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// 其它类型的错误，返回请求状态错误
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// 从cookie中获取token字符串并解码
	tokenString := cookie.Value
	// 实例化claims对象，后续存储用户传递来的数据
	claims := &Claims{}
	// 使用 jwt 的 ParseWithClaims 方法解析token，其中第二个参数为解析成功后需要存储解析结果的对象
	// 如果令牌无效（如果令牌已根据我们设置的登录到期时间过期了或者签名不匹配，此方法会返回错误）
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
}
