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
1. 当post为null时会访问失败，或许可以在api处返回空列表
2. 需要在跳转到登录的时候记录一下是从哪里跳转过去的，然后登录之后再对应的跳转回去
3. 我现在应该再写一个统一的拦截器，在请求访问的时候使用自定义的axios请求，来让请求都带上authorization请求头，并在后端的部分也加上一个中间件，对每一个来访问的请求做一个拦截
4. 在登录之后要改一下footer，将login和register去掉，改为logout

bug
1. 登录成功之后的第一次删除会失败
2. 之后每次删除用登陆的accesstoken都会失败，都会重新取一个accesstoken才能成功