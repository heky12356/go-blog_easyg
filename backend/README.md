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

初始化
GET /init  

问题：
删除一个标签之后不能再创建，或许可以在创建的时候再检查一下是否存在（软删除）然后再将其deleteat设置为null，在建立连接