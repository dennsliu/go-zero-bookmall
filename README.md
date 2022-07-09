创建mall服务
mkdir go-zero-bookmall
cd go-zero-bookmall 
go mod init go-zero-bookmall



user
    ├── api //  http访问服务，业务需求实现
    ├── cronjob // 定时任务，定时数据更新业务
    ├── rmq // 消息处理系统：mq和dq，处理一些高并发和延时消息业务
    ├── rpc // rpc服务，给其他子系统提供基础数据访问
    └── script // 脚本，处理一些临时运营需求，临时数据修复




model生成

准备工作#
找到user/model下的user.sql文件，将其在你自己的数据库中执行建表。

代码生成(带缓存)#
方式一(ddl)#
进入service/user/model目录，执行命令

$ cd service/user/model
$ goctl model mysql ddl -src user.sql -dir . -c
Done.
方式二(datasource)#
$ goctl model mysql datasource -url="$datasource" -table="user" -c -dir .
Done.

api文件编写
编写user.api文件
$ vim service/user/api/user.api

生成api服务
$ cd service/user/api
$ goctl api go -api user.api -dir . 

添加Mysql配置#
$ vim service/user/api/internal/config/config.go

	Mysql struct {
		DataSource string
	}
	CacheRedis cache.CacheConf

完善yaml配置#
$ vim service/user/api/etc/user-api.yaml

完善yaml配置#
$ vim service/user/api/etc/user-api.yaml
Name: user-api
Host: 0.0.0.0
Port: 8888
Mysql:
  DataSource: $user:$password@tcp($url)/$db?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai
CacheRedis:
  - Host: $host
    Pass: $pass
    Type: node

TIP
$user: mysql数据库user
$password: mysql数据库密码
$url: mysql数据库连接地址
$db: mysql数据库db名称，即user表所在database
$host: redis连接地址 格式：ip:port，如:127.0.0.1:6379
$pass: redis密码


完善服务依赖#
$ vim service/user/api/internal/svc/servicecontext.go
type ServiceContext struct {
    Config    config.Config
    UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
    conn:=sqlx.NewMysql(c.Mysql.DataSource)
    return &ServiceContext{
        Config: c,
        UserModel: model.NewUserModel(conn,c.CacheRedis),
    }
}


填充登录逻辑#
$ vim service/user/api/internal/logic/loginlogic.go



go-zero中怎么使用jwt

添加配置定义和yaml配置项#
$ vim service/user/api/internal/config/config.go
Auth       struct {
		AccessSecret string
		AccessExpire int64
	}

$ vim service/user/api/etc/user-api.yaml     加jwt密码和过期时间 
Auth:
  AccessSecret: ad879037-c7a4-4063-9236-6bfc35d44b6d
  AccessEpire: 86400



$ vim service/user/api/internal/logic/loginlogic.go



$ curl -i -X POST \
  http://127.0.0.1:8888/user/login \
  -H 'Content-Type: application/json' \
  -d '{
    "username":"666",
    "password":"123456"
}'

$ curl -i -X POST \
  http://127.0.0.1:8888/user/signup \
  -H 'Content-Type: application/json' \
  -d '{
    "username":"6666",
    "password":"123456",
    "email":"6666@testing.com",
    "gender":"male"
}'


curl -i -X POST \
  http://127.0.0.1:8888/book/search \
  -H 'Content-Type: application/json' \
  -d '{
    "keyword":"te",
    "page":1,
    "pagesize":20,
    "orderby":"id"
}'


curl -i -X POST \
  http://127.0.0.1:8888/book/add \
  -H 'Content-Type: application/json' \
  -d '{
    "plu":"plu2",
    "sku":"sku2",
    "name":"testing2",
    "image":"image2",
    "description":"image2",
    "in_stocked":"1"
}'