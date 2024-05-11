#### Golang Web接口认证的几种方式
- 1、基于Session和Cookie的认证
  - 这是一种有状态的认证方式。当用户登录成功后，服务器会生成一个唯一的Session ID，存储在服务器的Session存储中，并将其通过Cookie返回给客户端。客户端在后续的请求中携带这个Session ID，服务器通过ID查找对应的Session数据来验证用户身份。由于服务器需要维护Session数据，所以这种方式是有状态的
- 2、JSON Web Tokens (JWT)
  - JWT是一种无状态的认证方式。用户登录后，服务器生成一个包含用户信息的JWT，并发送给客户端。客户端在后续的请求中将此JWT放在Authorization头或Cookie中。服务器验证JWT的签名，无需存储额外的状态信息，因为所有必要的信息都在JWT内
- 3、OAuth 2.0 和 OpenID Connect
  - 这些是授权框架，可以用来进行无状态的认证。用户通过第三方提供商验证后，会收到一个访问令牌，这个令牌可以用于获取资源服务器上的受保护资源。服务器只验证令牌的有效性，不存储关于会话的特定状态
- 4、API密钥
  - 对于API接口，有时会使用API密钥进行认证，这是一种无状态的方式。每个客户端都有一个固定的密钥，每次请求时都将密钥放在请求头中。服务器验证密钥的有效性，但不需要存储会话信息。
- 5、预共享密钥 (PSK) 或者其他加密令
  - 类似于JWT，但可能使用不同的加密技术，如PASETO、Branca或nacl/secretbox，这些都是无状态的认证方法

这里学习练习下JWT的认证方式