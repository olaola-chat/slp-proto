

# 项目结构

```

rbp-proto
    config 用于工具生成代码的配置
    dao gf生成的dao
        xianshi  数据库名
            xxx.go  表名
    gen_proto gf生成的proto
        db  数据库子目录
            xianshi  数据库名
                xxx.proto  表名
    proto 手写proto
        rpc rpc服务接口
            user 系统名
                xxx.proto
        nsq nsq消息结构
        kafka kafka消息结构
    protoc-gen-rbp-rpc 生成RPC客户端的插件代码
    rpcclient proto生成的rpc客户端，各项目共用
        base RPC基类
        user  系统名
            Rbp.User.Auth.go  微服务名
    gen_pb proto生成的go文件
        db
            xianshi 数据库表
        rpc
            user 系统名
        nsq 
    Makefile 生成脚本
    
```

rpc服务手写proto，用service定义函数方法，通过工具自动生成客户端代码


# 注意！注意！注意！！！

需要将项目代码按照go规范放置目录，当前代码应放置于${GOPATH}/src/github.com/olaola-chat/rbp-proto

IDE 设置protobuf路径导入路径
${GOPATH}/src/github.com/olaola-chat/rbp-proto/proto
${GOPATH}/src/github.com/olaola-chat/rbp-proto/protoc-gen-rbp-rpc/proto
${GOPATH}/src/github.com/olaola-chat/rbp-proto/gen_proto


# 代码规范

* 数据库配置名需要与db名一致
* 参考make dao直接生成一张表的全套
* proto目录下手写RPC协议
* rpc的proto需要import "rbp/plugin/option.proto"
* rpc的service需要添加option (rbp.plugin.rbp_service).name = "User.Profile";
* 所有proto的go_package需要写全路径

# 生成数据库表代码

在dev机器上执行

```
以xianshi库xs_user_profile为例

sh gen_db.sh xianshi xs_user_profile

或

make dao

```

# 生成rpc代码

```
以User.Profile rpc服务为例

sh gen_rpc.sh User Profile

或 make rpc

```

# 安装protoc-gen-go
brew install protoc-gen-go

# 安装rpc生成工具
make cli

# 安装gf-cli

git clone git@github.com:olaola-chat/rbp-gf-cli.git

go install

# 安装protoc-go-inject-tag

go install github.com/favadi/protoc-go-inject-tag@latest
