package auth

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
)

func Refresh(w http.ResponseWriter, r *http.Request) {
	//  (BEGIN) 此处的代码与`Welcome`路由的第一部分相同，作用就是检验客户端的token合法性
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
	// (END)  此处的代码与`Welcome`路由的第一部分相同

	// 我们确保在足够的时间之前不会发行新令牌。
	// 在这种情况下，仅当旧令牌过期时间小于等于30秒时才发行新令牌。
	// 否则，返回错误的请求状态。
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("token过期时间大于30s")
		return
	}

	// 现在，为当前用户创建一个新令牌，并延长其到期时间
	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newTokenString, err := newToken.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 查看用户新的`token` cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   newTokenString,
		Expires: expirationTime,
	})
}
