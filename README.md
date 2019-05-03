# HuihuCrawller
并发版爬虫

1.先在本地运行elasticsearch容器

`docker run -d -p 9200:9200 elasticsearch`

2.启动前端展示模块

`go run frontend/start.go`

3.启动爬虫

`go run main.go`

运行后本地9200端口即可看到搜索界面，关键字用空格分隔
