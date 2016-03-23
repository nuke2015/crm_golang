
go1.6版本还在调试

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


