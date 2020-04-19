# Docs For Developer

## 开发流程
1. 编写流程文档/UML 图
2. 数据库表结构、字段设计
3. 开发业务代码
4. 编写测试用例（单元测试、interface 测试、api 测试）
5. 编写 api 文档

注：如果不想写1、2，可以写伪代码，即写好大致逻辑函数、结构体

## Prepare
1. create MySQL server on your Development machine:

    docker run -d --rm -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=111111 -e MYSQL_DATABASE=mcp mysql:5.7 --character-set-server=utf8mb4 --collation-server=utf8mb4_general_ci

2. confirm config file configs/config-dev.yaml exist

3. confirm sync db if you need to dev:

    make run-syncdb


## How to start
1. config Makefile run-server (--port=8080 --config=configs/config-dev.yaml), start server:

    make run-server

## Swagger

make run-swagger

ref: https://github.com/emicklei/go-restful/blob/master/examples/restful-openapi.go

## Test
### Integration
docker-compose -f ./test/integration/docker-compose.yaml --project-directory . up --abort-on-container-exit --exit-code-from integration

### devtest
CONFIG_FILE=configs/config-dev.yaml go test -v -tags=integration ./test/integration/...

