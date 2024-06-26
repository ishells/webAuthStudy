#### 参考链接 
[从零实现 JWT 认证](https://learnku.com/go/t/52399)

#### 互联网认证一般访问流程
- 1、客户端向服务端发送用户名和密码
- 2、服务端验证用户名和密码是否正确
  - 如果不正确则返回未认证信息
  - 如果正确，生成jwt密钥，在当前对话（session）里面保存相关数据，比如用户角色、登录时间等等。
- 3、服务器向用户返回一个 session_id，写入用户的 Cookie
- 4、用户随后的每一次请求，都会通过 Cookie，将 session_id 传回服务器。
- 5、服务器收到 session_id，找到前期保存的数据，由此得知用户的身份。

```
这种模式的问题在于，扩展性（scaling）不好。
单机当然没有问题，如果是服务器集群，或者是跨域的服务导向架构，就要求 session 数据共享，每台服务器都能够读取 session。

举例来说，A 网站和 B 网站是同一家公司的关联服务。现在要求，用户只要在其中一个网站登录，再访问另一个网站就会自动登录，请问怎么实现？

一种解决方案是 session 数据持久化，写入数据库或别的持久层。
各种服务收到请求后，都向持久层请求数据。这种方案的优点是架构清晰，缺点是工程量比较大。另外，持久层万一挂了，就会单点失败。

另一种方案是服务器索性不保存 session 数据了，所有数据都保存在客户端，每次请求都发回服务器。JWT 就是这种方案的一个代表。
```

#### jwt加密流程
![jwtSignature](./imgs/jwtSignature.png)

#### signin代码编写流程(即jwt创建流程)
- 1、分别创建所需对应变量
  - 创建一个jwt加密时使用的密钥变量
      - 可以使用随机密钥或固定密钥，这里选择简单的固定密钥key
  - 创建固定的用户名密码的map变量 
  - 创建一个结构体变量用来存储用户请求中的用户名、密码
  - 创建生成jwt声明对象claims时所需的结构体变量（即payload部分）
- 2、编写处理函数
  - 首先从用户请求中拿到传递来的账户名、密码
  - 然后验证与服务端的账户名、密码是否匹配
    - 不匹配直接返回401状态码、结束本次访问流程
    - 匹配的话，首先设置下面要生成的token的过期时间，并设置好JWT声明（即payload的内容）
  - 然后调用 jwt.NewWithClaims() 方法生成一个jwt对象，该方法返回jwt加密所需的header、payload、signature三个部分
  - 最后调用 SignedString() 方法对jwt对象进行签名，并返回最终的jwt的token段
  - 将生成的token值返回到客户端的cookie中

- 3、后续经过signin jwt生成token的客户端都可以使用cookie存储用户信息

#### jwt认证流程
![authJwt](./imgs/authJwt.png)


#### 



