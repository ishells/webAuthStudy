package auth

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
)

// 创建一个jwt使用的密钥
var jwtKey = []byte("secret_key")

// 创建一个用户列表，用户名与密码
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

// Credentials 创建一个结构体保存请求正文中的用户名与密码
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// JWT对象包含三部分(由.进行分隔)，
// 形如eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIxIiwiZXhwIjoxNTQ3OTc0MDgyfQ.2Ye5_w1z3zpD4dSGdRp3s98ZipCNQqmsHRB9vioOx54
// 第一部分是标题header，表头指定加密算法和请求类型，这部分是固定的（即使用相同算法的任何JWT都是相同的）
// 第二部分是有效载荷payload，其中包含特定于应用程序的信息（在我们的示例中，是用户名和有效时间）
// 第三部分是签名，这一部分是通过加密算法与指定密钥生成的

// 创建将被编码为JWT对象的结构体
// jwt.StandardClaims结构体中会包含一些默认字段，如过期时间，下面会将值存入
// claims 结构体相当于 payload 载荷内容
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Signin(w http.ResponseWriter, r *http.Request) {
	// 实例化 请求体中账户密码 的结构体对象
	var creds Credentials
	// 获取请求的json正文并解析
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// 如果请求体获取错误，返回http错误
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 获取服务端map中保存的密码与请求体中保存的密码进行校验，一致则通过，不一致则返回认证失败
	expectedPasswd, ok := users[creds.Username]
	if !ok || creds.Password != expectedPasswd {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// 声明令牌的到期时间
	expirationTime := time.Now().Add(5 * time.Minute)
	// 实例化JWT声明，其中包括用户名和有效时间，存入具体值
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			// 设置过期时间，按照jwt的StandardClaims结构体字段来设置
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// 使用指定的签名方法创建签名对象
	// jwt.NewWithClaims()方法创建一个jwt对象
	/*
		return &Token{
			Header: map[string]interface{}{
				"typ": "JWT",
				"alg": method.Alg(),
			},
			Claims: claims,
			Method: method,
		}
	*/
	// 根据以上return可知，jwt.NewWithClaims()方法返回一个jwt对象，第一个参数是Header（基本固定格式），第二个参数是Claims，第三个参数是签名方法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 创建JWT字符串
	// 这部分可以点进去看下 SignedString 方法的具体实现
	// 其实就是将 NewWithClaims() 方法返回的token中的值进行处理，分别序列化为json字节切片，然后使用 EncodeSegment 函数对这些序列化后的数据进行编码，并用句点分隔符连接这些编码过的部分，从而生成签名字符串
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// 如果创建JWT时出现错误，返回服务器内部错误
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("JWT生成错误...")
		return
	}
	// 最后将客户端cookie token设置为刚刚生成的JWT
	// 我们还设置了与令牌本身相同的cookie到期时间
	log.Printf(tokenString)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
