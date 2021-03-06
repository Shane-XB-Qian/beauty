- [English](README.md)
- [中文](README_ZH.md)

[![GoDoc](https://godoc.org/github.com/yang-f/beauty?status.svg)](https://godoc.org/github.com/yang-f/beauty)

# 这是一个 Golang 实现的简易框架

你通过它可以实现一个简单的 restful 工程或者是一个 web 应用。
如果你不想使用 mysql 数据库，可以自己实现 Auth 中间件与 session。你可以用自己的 DAO，或者你喜欢的。

## 快速开始:

- 执行命令
  ```
  mkdir demo && cd demo
  go get gopkg.in/alecthomas/kingpin.v2
  go get github.com/yang-f/beauty
  ```
- 将$GOPATH/bin 加入到$PATH 环境变量
- 执行命令
  ```
  beauty
  ```
- 这时显示如下
  ```
      usage: beauty [<flags>] <command> [<args> ...]

      A command-line tools of beauty.

      Flags:
        --help  Show context-sensitive help (also try --help-long and --help-man).

      Commands:
        help [<command>...]
          Show help.

        demo
          Demo of web server.

        generate <name>
          Generate a new app.
  ```
- 测试 beauty
  ```
  beauty demo
  ```
- 这时 terminal 显示
  ```golang
  2017/05/04 16:21:05 start server on port :8080
  ```
- 通过浏览器访问 127.0.0.1:8080

  ```golang
  {"description":"this is json"}
  ```

- 访问 127.0.0.1:8080/demo1

  ```golang
  {"status":403,"description":"token not found.","code":"AUTH_FAILED"}
  ```

- 访问 127.0.0.1:8080/demo2
  ```golang
  {"description":"this is json"}
  ```
- 访问 127.0.0.1:8080/demo3
  ```golang
  {"status":403,"description":"token not found.","code":"AUTH_FAILED"}
  ```
- 恭喜你，运行成功。

## 如何使用:

- 生成 app
  ```
  beauty generate app的名字
  ```
- 生成的 app 目录列表
  ```
  GOPATH/src/yourAppName
      ├── controllers
      │   ├── adminController.go
      │   └── controller_test.go
      ├── decorates
  ├       └── http.go
      ├── main.go
      ├── models
      ├── tpl
      └── utils
  ```
- 关于路由

  - 例子

  ```golang
      r := router.New()

      r.GET("/", controllers.Config().ContentJSON())

      r.GET("/demo1", controllers.Config().ContentJSON().Auth())

      r.GET("/demo2", controllers.Config().ContentJSON().Verify())

      r.GET("/demo3", controllers.Config().ContentJSON().Auth().Verify())
  ```

- token 生成

  ```golang
  tokenString, err := token.Generate(fmt.Sprintf("%v|%v", user_id, user_pass))

  ```

- 小例子

  ```golang
  package main

  import (
      "net/http"
      "github.com/yang-f/beauty/consts/contenttype"
      "github.com/yang-f/beauty/utils/log"
      "github.com/yang-f/beauty/router"
      "github.com/yang-f/beauty/settings"
      "github.com/yang-f/beauty/controllers"
      "github.com/yang-f/beauty/decorates"

  )

  func main() {

      log.Printf("start server on port %s", settings.Listen)

      settings.Listen = ":8080"//服务运行端口

      settings.Domain = "yourdomain.com"//部署服务的域名

      settings.LogFile = "/your/path/yourname.log"//日志所在文件

      settings.DefaultOrigin = "http://defaultorigin.com"//默认的请求来源

      settings.HmacSampleSecret = "whatever"//令牌生产需要的字符串

      r := router.New()

      r.GET("/", controllers.Config().ContentJSON())

      r.GET("/demo1", controllers.Config().ContentJSON().Auth())

      r.GET("/demo2", controllers.Config().ContentJSON().Verify())

      r.GET("/demo3", controllers.Config().ContentJSON().Auth().Verify())

      log.Fatal(http.ListenAndServe(settings.Listen, r))

  }
  ```

## 支持特性:

- 参数校验
  - 这是一个统一的参数校验，主要是针对 SQL 注入，有了它，或许就不用一个一个参数做校验了。
  - 使用方式
  ```golang
  r.GET("/", controllers.Config().ContentJSON().Verify())//需要验证用户信息
  ```
- 令牌

  ```golang
  settings.HmacSampleSecret = "whatever"

  token, err := token.Generate(origin)

  origin, err := token.Valid(token)
  ```

- 数据库操作（基于 mymysql）
  ```golang
  db.Query(sql, params...)
  ```
- cors 跨域

  - 静态文件

  ```golang
  router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", decorates.CorsHeader2(http.FileServer(http.Dir("/your/static/path")))))
  ```

  - 其他:
    - 默认是支持跨域操作的

- 日志

  - 使用

  ```golang
  settings.LogFile = "/you/log/path/beauty.log"

  log.Printf(msg, params...)
  ```

  - 每隔 12 小时且日志大于 50M 自动归档

- 会话
  ```golang
  currentUser := sessions.CurrentUser(r *http.Request)
  ```
- 错误处理以及 http 状态管理

  ```golang
  func XxxxController() decorates.Handler{
      return func (w http.ResponseWriter, r *http.Request) *models.APPError {
          xxx,err := someOperation()
          if err != nil{
              return &models.APPError {err, Message, Code, Status}
          }
          ...
          return nil
      }
  }
  ```

- 工具

  - Response
  - Rand
  - MD5
  - Post

- 测试
  - go test -v -bench=".\*"
  - go test -v -short \$(go list ./... | grep -v /vendor/)
  - ...

## 其他:

- sql

  ```golang
  create database yourdatabase;
  use yourdatabase;
  create table if not exists user
  (
      user_id int primary key not null  auto_increment,
      user_name varchar(64),
      user_pass varchar(64),
      user_mobile varchar(32),
      user_type enum('user', 'admin', 'test') not null,
      add_time timestamp not null default CURRENT_TIMESTAMP
  );

  insert into user (user_name, user_pass) values('admin', 'admin');
  ```

- 你需要编辑一个配置文件'/srv/filestore/settings/latest.json' 格式如下：
  ```golang
  {
      "mysql_host":"127.0.0.1:3306",
      "mysql_user":"root",
      "mysql_pass":"root",
      "mysql_database":"yourdatabase"
  }
  ```

## 贡献代码:

1. Fork 代码!
2. 创建新的分支: `git checkout -b my-new-feature`
3. 提交更改: `git commit -m 'Add some feature'`
4. push 到分支: `git push origin my-new-feature`
5. 发起一个 pull request :D

## 将要实现:

- [x] 命令行工具
- [ ] 继续提升文档质量
- [ ] 权限控制
- [ ] 增加测试覆盖率
- [ ] 错误处理 
