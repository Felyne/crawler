
### 爬取珍爱网数据的分布式爬虫

1. 准备工作
```shell
# 数据存储在elasticsearch里面
docker run -d --name=elastic -p 9200:9200 -v /home/chen/work/elastic:/usr/share/elasticsearch/data elasticsearch:5.6
# redis用来去重
docker run -d --name=myredis -p 6379:6379 -v /home/chen/work/redis:/data redis:5.0 redis-server --appendonly yes
```

2. 启动rpc服务：
```shell
# 启用保存数据到ES的rpc服务:
go run crawler_distributed/persist/server/itemsaver.go --port=1234

#启用两个爬取工作的rpc服务:
go run crawler_distributed/worker/server/worker.go --port=9000
go run crawler_distributed/worker/server/worker.go --port=9001

# 启动主程序, 工作过程中若30s内没有接收到解析结果则超时停止
go run crawler_distributed/main.go --itemsaver_host=":1234" --worker_hosts=":9000,:9001"
```

3. 前端展示
```shell
go run crawler/frontend/starter.go
```
4. 浏览器访问localhost 进入搜索首页
