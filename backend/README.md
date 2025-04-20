api

/api/post : 

    创建文章
    POST /create

    获取全部文章
    GET /getposts

    删除文章
    DELETE /delete/:uid

/api/user :

    登录
    POST /login

    注册
    POST /register

    获取全部用户
    GET /getalluser

    刷新令牌
    POST /refreshaccesstoken

初始化
GET /init  

想法：
之后可以做一个回收箱的功能，每次删除之后先是软删除，然后当超过设定的时间还没有恢复的，就从数据库永久删除