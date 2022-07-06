model生成

准备工作#
book/book.sql文件，将其在你自己的数据库中执行建表。

代码生成(带缓存)#
进入service/user/model目录，执行命令

$ cd service/book/model
$ goctl model mysql ddl -src book.sql -dir . -c



api文件编写
book.api文件
$ vim service/book/api/book.api

生成api服务
$ cd service/book/api
$ goctl api go -api book.api -dir . 


go run book.go -f etc/book.yaml