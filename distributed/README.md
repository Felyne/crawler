
#### 爬取珍爱网数据的分布式爬虫

**准备工作**

配置 config/config.go 后启动以下服务
```shell
# 数据存储在elastic里面
docker run -d --name=elastic -p 9200:9200 -v /home/chen/work/elastic:/usr/share/elasticsearch/data elasticsearch:5.6.6
# redis用来去重
docker run -d --name=myredis -p 6379:6379 -v /home/chen/work/redis:/data redis:5.0 redis-server --appendonly yes
```

**启动rpc服务，用到的是go的jsonrpc**

```shell
# 启用保存数据到ES的rpc服务:
go run service/persist/server/itemsaver.go --port=1234

#启用两个爬取工作的rpc服务:
go run service/worker/server/worker.go --port 9000
go run service/worker/server/worker.go --port 9001
```

**入口程序**
```
# 启动主程序, 工作过程中若30s内没有接收到解析结果则超时停止
go run cmd/distributed/main.go --itemsaver_host=":1234" --worker_hosts=":9000,:9001"
```
