# lmaker -  go项目代码生成
快速生成一个golang项目，扩展性强，加快接口开发。

## 安装依赖
 使用proto协议，定义接口,也是项目生成起点，需要本地安装 `[protoc](https://github.com/grpc-ecosystem/grpc-gateway)`
```$xslt
Then use go get -u to download the following packages:
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```
## 使用
```$xslt
git clone https://github.com/lmfuture-ma/lmaker.git
cd lmaker 
go install

// todolist 为项目名，切换到你的gopath下
lmaker create -name=todolist   

//测试
cd todolist
go build .   ==>生成二进制包 todolist
./todolist   ==>运行 端口默认8088
初始化共三个接口
// add
curl -X POST -H "Content-Type: application/json" -d '{"addTodoReq":{"id":2,"author":"lmaker3","description":"first+1 todo","done":true,"createdAt":"2019-10-11"}}' localhost:8088/addTodo  
curl -X POST -H "Content-Type: application/json" -d '{"addTodoReq":{"id":2,"author":"lmaker3","description":"first+1 todo","done":true,"createdAt":"2019-10-11"}}' localhost:8088/addTodo
// list
curl -X GET -H "Content-Type: application/json" localhost:8088/listTodos 
// get
curl -X GET -H "Content-Type: application/json" localhost:8088/getTodo/2
```
## API扩展

```$xslt
cd todolist
//1. 更改 pb/*.proto文件
lmaker gen 
// 2. 完成server目录下完成接口实现 
```
