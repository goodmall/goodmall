JWT
===========

跟传统浏览器session机制对比

|     session     |           VS  |            JWT      |
|  :-             | -:           |  :-:              | 
|     引用        |              |               值     |
|     服务端存储实现session机制    |              |     无需存储     |
| （session_id）  |              |       payload      |
|   cookie中传递session_id  |       |   用户非敏感信息编码进去  |
|  sess_encode session特有编解码  |   |       base64编解码      |

 


很不错哦！
https://medium.com/@raul_11817/securing-golang-api-using-json-web-token-jwt-2dc363792a48

八幅漫画理解JWT
https://blog.csdn.net/ztguang/article/details/53155877


https://stackoverflow.com/questions/35749646/jwt-authentication-in-slim-v3-and-android

蛮详细的：
https://blog.csdn.net/a82793510/article/details/53509427

client端集成jwt
https://www.thepolyglotdeveloper.com/2017/03/authenticate-jwt-nativescript-angular-mobile-application/

## refresh-token

https://segmentfault.com/q/1010000012407467
~~~
	
1.    jwt有两种tokenrefreshToken和accessToken,两个都可以在AuthServiceProvider设置过期时间
2.    一种存在cookie里或localstorage里,一种存在session或sessionstorage
3.    laravel 把refreshToken放在cookie里,把accessToken和过期时间放在body里返回
    前端按照返回数据保存两种token
    refreshToken过期则重新登录,accessToken过期则用refreshToken换取accessToken
4.    可以用命令php artisan route:list查看所有auth相关api的路由

~~~


