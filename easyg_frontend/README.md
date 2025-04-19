页面

home /  
about /about  
post /post  
    article /post/:uid  
admin /admin  
    create /admin/create  
    delete /admin/delete  

login /login  
register /register


问题：
1. 需要在跳转到登录的时候记录一下是从哪里跳转过去的，然后登录之后再对应的跳转回去
2. 图片或许需要做一下懒加载
3. 当创建完文章之后就跳转到管理页面（管理页面还没写，乐）

bug
1. 登录成功之后的第一次删除会失败
2. 之后每次删除用登陆的accesstoken都会失败，都会重新取一个accesstoken才能成功

想做：
1. 需要做一下管理页面，可以管理页面的状态，比如修改内容，修改分类，删除，这样。