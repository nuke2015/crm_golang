
使用:
bee run
然后访问
http://127.0.0.1:8080/index/index
看到登陆首页
若无用户则注册一个用户名+密码都是admin
注意:
必须确保mongodb连接正确,并且其中test库存在。

CRM基础
作者:锋子
email:196971567@qq.com


开发环境是go1.6版本.

------>
设置gopath目录
把%gopath%\bin加入到path

安装beego
go get github.com/astaxie/beego
安装bee
go get github.com/beego/bee
在gopath/bin目录中会有bee.exe二进制产生

安装mongo中间件
这个用不了
go get gopkg.in/mgo.v2
改成
git clone https://github.com/go-mgo/mgo.git D:\advance\go1.6\gopath\src\gopkg.in\mgo.v2


启动mongo服务监听 127.0.0.1:27017
然后

问题修复:
问题一Build fails with: use of internal package not allowed 
go get github.com/prasmussen/google-api-go-client/googleapi/internal/uritemplates
go get gopkg.in/mgo.v2-unstable/internal/scram


有些源用go get是安装不上的.要换一种办法,用git.
git clone https://github.com/smartystreets/goconvey/
git clone https://github.com/jtolds/gls.git
git clone https://github.com/smartystreets/assertions.git

git clone https://github.com/prasmussen/google-api-go-client.git

可以直接在git目录直接运行,但注意要使用utf-8编码.
