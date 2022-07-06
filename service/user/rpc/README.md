rpc服务编写#
编译proto文件
$ vim service/user/rpc/user.proto


生成rpc服务代码
$ cd service/user/rpc
$ goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
