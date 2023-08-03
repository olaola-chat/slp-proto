

# 项目结构

```

rbp-proto
    config 用于工具生成代码的配置
    model gf生成的model
        xianshi  数据库名
            xxx.go  表名
    dao gf生成的dao
        xianshi  数据库名
            xxx.go  表名
    gen_proto gf生成的proto
        db  数据库子目录
            xianshi  数据库名
                xxx.proto  表名
    proto 手写proto
        User 系统名
            User.Auth.proto 微服务名
    rpcclient proto生成的rpc客户端，各项目共用
        User  系统名
            User.Auth.go  微服务名
    gen_pb proto生成的go文件
        db
            xianshi 数据库表
        rpc
            User 系统名
    Makefile 生成脚本
    
```

rpc服务手写proto，用service定义函数方法，通过工具自动生成客户端代码

# 生成数据库表代码

在dev机器上执行

```
以xianshi库xs_user_profile为例

sh gen_db.sh xianshi xs_user_profile

```

# 生成rpc代码

```
以User.Profile rpc服务为例

sh gen_rpc.sh User Profile

```

# 安装gf-cli

git clone git@github.com:olaola-chat/rbp-gf-cli.git

go install

# 安装protoc-go-inject-tag

git clone git@github.com:favadi/protoc-go-inject-tag.git

go install
