# polarisctl
The polaris-server command line tool is used to quickly initiate related OpenAPI requests

## 工具配置说明

**配置格式说明**

```josn
{
  "default_culster":"test",  // 指定当前操作的默认集群，如果工具参数没有指定，则使用该配置
  "clusters":[              // 配置多个集群
    {
      "name":"test",        // 	集群名，随意命名
      "host":"119.91.66.223:8090", // 集群控制面地址
      "token":""              // 集群控制面的 token，从控制台获取
    }
  ]
}
```



**配置使用**

1. 通过参数指定配置文件路径：./polarisctl -c "/xxx/xxx/xx.json"
2. 不指定，则查找默认配置：$$$HOME/.polarisctl/polarisctl_config.json



## 命令结构

## 全局参数

* --cluster=xx ，指定操作的集群
* --config=xxx.json，指定集群配置文件目录
* --debug   打开debug 日志

### 指定本地配置

```bash
./polarisctl namespaces list --config=test
```

### 配置多个集群，使用指定集群

```bash
./polarisctl namespaces list --cluster=test
```

