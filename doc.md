# Polarisclt

> 命令行工具使用说明



# 配置说明

工具需要一份 json 格式的配置文件，配置集群控制台的地址和token，配置示例：

```json
{
  "default_culster":"test",  // 指定当前操作的默认集群，如果工具参数没有指定，则使用该配置
  "clusters":[              // 配置多个集群
    {
      "name":"test",               // 用于区分不同的配置，工具可以使用改命名使用不同的集群配置
      "host":"119.91.66.223:8090", // 集群控制面地址
      "token":""                  // 集群控制面的 token，从控制台获取
    }
  ]
}
```

* 默认的配置路径：$HOME/.polarisctl/polarisctl.json

* 可以使用全局参数指定配置路径 

  ```shell
  ./polarisctl -f /xx/xx/xx/xx.json
  ```

# 命令结构

使用树形结构组织所有的资源操作命令，使用 -h 获取帮忙，查看支持的所有资源

```shell
polarisctl polaris command line tool is used to quickly initiate related OpenAPI requests
 Find OpenApi doc at:
 https://polarismesh.cn/docs/%E5%8F%82%E8%80%83%E6%96%87%E6%A1%A3/%E6%8E%A5%E5%8F%A3%E6%96%87%E6%A1%A3/open_api

Usage:
  polarisctl [flags]
  polarisctl [command]

Available Commands:
  circuitbreaker circuitbreaker [list|create|delete|enable|update]
  completion     Generate the autocompletion script for the specified shell
  config         config []
  faultdetectors faultdetectors [list|create|delete|update]
  help           Help about any command
  instances      instances [list|labels|create|delete|count|update]
  maintain       maintain [loglevel|setloglevel|leaders|cmdb|clients]
  namespaces     namespaces [list|create|delete|update]
  ratelimits     ratelimits [list|create|update|enable|delete]
  routings       routings [list|create|update|enable|delete]
  services       services [list|all|create|update|delete]
  Flags:
      --cluster string   current cluster
      --config string    cluster config path (default "")
      --debug            debug log
  -h, --help             help for polarisctl
```

# 全局参数

所有命令都生效的可选全局命令：

* --cluster : 指定当前使用的集群，用于有多个集群时，基于该值去查找不同的集群配置
* --config： 指定集群配置文件路径
* --print：控制打印参数，使用示例：--print=name,id

# 支持的命令

1. 命名空间
2. 注册发现
3. 服务治理
4. 配置中心
5. 用户鉴权

## 命名空间操作

namespaces 支持的命令：

```
namespaces [list|create|delete|update]

Usage:
  polarisctl namespaces [list|create|delete|update] [flags]
  polarisctl namespaces [command]

Available Commands:
  create      create (-f create_namespaces.json)
  delete      delete (-f delete_namespaces.json)
  list        list namespaces
  update      update (-f update_namespaces.json)

Flags:
  -f, --file string    json file for create/delete/update namespaces
  -h, --help           help for namespaces
      --print string   namespaces print field,eg:"jsontag1,jsontag2"

Global Flags:
      --cluster string   current cluster
      --config string    cluster config path (default "/Users/zhangyong/.polarisctl/polarisctl_config.json")
      --debug            debug log
```

### 获取命令空间列表

example:

```go
./polarisctl namespaces list -l 10
```

参数说明：-l 10 ：分页查询参数，限制一页 10 个

指输出 naemspace 的部分字段：

```g0
./polarisctl namespaces list  -l 10 --print=name,comment,owners
```

参数说明： --print=name,comment,owners ，设置输出 namespace 时只输出 3 个字段

### 批量创建命令空间

```go
./polarisctl namespaces create -f example/namespace/create.json
```

1. 写操作相关的命令，使用 json 编辑参数，参数参考 openapi 文档
2. 可以使用 example 目录下面的 json 文件示例

