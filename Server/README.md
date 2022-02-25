# 接口文档

/login (Post)

+ 携带参数:
  + body: `{Username:"", Password:""}`

/register (Post)

+ 携带参数:
  + body: `{Username:"", Password:""}`

http: /ws (Get)

+ 携带参数:
  + headers: `x-Token`(Jwt Token)
  + url: `?username=""&friend_id=""`

ws: /ws

+ 携带参数:
  + body: `{"sender":"", "recipient":"", "content":"", "date":""}`

/user (Delete)

+ 携带参数:
  + headers: `x-Token`
