# 给 OPQ 打一个小补丁，加入鉴权和发送队列功能

# 配置

此补丁和 OPQ 使用同一个配置文件，所以需要放入同一个文件夹中，否则将按默认配置运行

只有两个配置项:
`PatchKey`: 用于 api 调用鉴权
`PatchServerPort`: 该补丁程序运行的端口

# 使用

该补丁与 OPQ 是否运行无关,配置好后直接运行即可, 如果配置没有改动(包括 OPQ 的端口)，就无需重新启动

与 webapi 相关的操作可直接调用此补丁运行的端口,而无需开放 OPQ 运行的端口裸奔(当然可以用其他反代工具设置鉴权), 如果不想开启鉴权需设置`PatchKey=""`

## 鉴权

调用 api 时在路径加入`_key=你设置的PatchKey`，否则会拒绝访问

## 发送队列

(仅当)调用 `funcname=SendMsg` 时加入额外参数`_queue=1` 即可使用队列发送，该接口会立即返回，你无法知道处理结果

# 演示

示例配置:
`CoreConf.conf`

```ini
Port = "0.0.0.0:8888"
WorkerThread = 50
OPQVer = "v3.1.8"
Token = ""

PatchKey = "xiyaowong"
PatchServerPort = 8899
```

则:
普通发送消息 : `http://127.0.0.0:8899/v1/LuaApiCaller?qq=123456&funcname=SendMsg&timeout=10&_key=xiyaowong`

使用队列延时发送消息: `http://127.0.0.0:8899/v1/LuaApiCaller?qq=123456&funcname=SendMsg&timeout=10&_key=xiyaowong&_queue=1`

与控制面板(WebUI) 相关的 api 无需鉴权
