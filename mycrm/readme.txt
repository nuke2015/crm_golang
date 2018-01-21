
启动mongo数据库,添加test库

拉代码
git clone https://github.com/nuke2015/crm_golang.git

安装第三方依赖
glide cc
glide up


使用:
bee run 或go run main.go
然后访问
http://127.0.0.1:8080/index/index
看到登陆首页

若无用户则注册一个用户名+密码都是admin
注意:
必须确保mongodb连接正确,并且其中test库存在。

CRM基础
作者:锋子
email:196971567@qq.com
