[toc]

# 新手任务

## 功能点：

1. 使用golang，实现员工基本信息(姓名、工号、岗位、电话…...) CRUD 的 grpc 服务，信息存储到mysql；

2. 使用golang，实现api网关，对接1中的grpc服务，对外提供 http json接口；

3. 使用docker打包grpc服务和api网关，并以docker-compose编排部署(包含mysql)。

## 目录说明：

### /api

主要存放所有的接口代码，文件目录接口如下：

```
api/
├── grpc_api
│   └── grpc_api.go
├── http_api
│   └── http_api.go
├── mysql_api
│   └── mysql_api.go
└── proto
    ├── google
    │   └── api
    │       ├── annotations.pb.go
    │       ├── annotations.proto
    │       ├── httpbody.pb.go
    │       ├── httpbody.proto
    │       ├── http.pb.go
    │       └── http.proto
    ├── staff_service.pb.go
    ├── staff_service.pb.gw.go
    └── staff_service.proto
    
6 directories, 12 files
```

#### /api/grpc_api/grpc_api.go：

grpc 服务的接口

#### /api/http_api/http_api.go

http 服务的接口

#### /api/mysql_api/mysql_api.go

mysql 数据库 CRUD 操作接口

#### /api/proto/google/api

直接从包`github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google` 将该文件夹拷贝过来，然后利用 protoc 工具编译 .pb.go 文件：

```
# windows系统
$ protoc -ID:\Download\protoc-3.12.3-win64\include -I. --go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor:. google/api/*.proto

# Linux系统
待补充（相对简单）
```

#### /api/proto/staff_service.proto

定义 gprc 服务的接口文件，利用该文件生成 `staff_service.pb.go`，`staff_service.pb.gw.go` 文件：

```
# windows系统
# 编译staff_service.pb.go
$ protoc -ID:\Download\protoc-3.12.3-win64\include -I. --go_out=plugins=grpc,Mgoogle/api/annotations.proto=http_grpc/api/proto/google/api:. .\staff_service.proto

# 编译staff_service.pb.gw.go gateway
$ protoc -ID:\Download\protoc-3.12.3-win64\include -I. --grpc-gateway_out=logtostderr=true:. .\staff_service.proto
```

---

### /build

存放 Dockerfile 文件和编译生成的可执行文件，并作为编译docker 镜像的上下文目录。

```
build/
└── Dockerfile

0 directories, 1 file
```

---

### /deployments

存放 docker-compose 的配置文件以及相应的依赖文件：

```
deployments/
├── docker-compose.yaml
└── mysql
    └── init
        └── init.sql

2 directories, 2 files
```

#### /deployments/docker-compose.yaml

docker-compose 的启动配置文件，定义服务。

#### /deployments/mysql/init/init.sql

mysql 服务的初始化文件

---

### /go.mod、go.sum

go module 相关文件

---

### /main.go

主文件，调用相应接口启动服务。

---

### /Makefile

用于编译目标文件：http_gateway 可执行文件与 http_gateway 的docker 容器镜像。

---

## 运行步骤

1. 进入 http_grpc/ 目录

   ```
   cd http_grpc
   ```

2. 编译目标文件

   ```
   make
   ```

3. 运行 docker-compose 编排部署服务

   ```
   docker-compose -f deployments/docker-compose.yaml up -d
   ```

4. 结束服务

   ```
   docker-compose -f deployments/docker-compose.yaml stop
   ```

5. 删除相关的容器

   ```
   docker-compose -f deployments/docker-compose.yaml rm
   ```

6. 删除文件

   ```
   make clean
   ```

---