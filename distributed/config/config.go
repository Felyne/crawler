package config

const (
	Qps           = 1  //访问网站的速率
	Timeout       = 30 //该时间内没有任何一个worker返回结果则退出程序
	ElasticURL    = "http://127.0.0.1:9200"
	ElasticIndex  = "db_test" //相当于数据库名
	RedisAddr     = "localhost:6379"
	RedisPassword = ""
	RedisDB       = 0

	//Parser names
	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ParseProfile"
	NilParser     = "NilParser"

	//要注册和本地调用的rpc服务
	ItemSaverRpc    = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"
)
